# Arbitrum integration for sedge

Status: design approved 2026-05-25 (internal branch, off `core`).
Out of scope: README/CHANGELOG updates, Docusaurus site, the `run-a-single-node-from-branch.yaml` workflow dropdown.

## Goal

`sedge generate arb-full-node` produces a working Arbitrum node setup against either Arbitrum One (chain id 42161, L1 = Ethereum mainnet) or Arbitrum Sepolia (chain id 421614, L1 = Ethereum sepolia). The L2 side uses Offchain Labs' `nitro-node` paired with Nethermind's `nethermind-arbitrum` fork as the execution client. By default the compose also brings up a full L1 stack (EL + CL); URL flags let users point at external L1 endpoints and skip the bundled L1.

## Reference

`https://github.com/NethermindEth/nethermind-arbitrum/blob/main/docker-compose.yml` — two services: `nethermind-arbitrum` (forked EL) and `nitro` (rollup node). Healthcheck-gated `depends_on`, shared JWT via mounted volume, parent-chain RPC + beacon URLs as required env. This design rewires that into sedge's existing L2 templating model (Optimism/Taiko/Surge).

## Architecture

Four services in bundled mode, two in external-L1 mode:

```
   sedge generate arb-full-node -n <l1> [--chain <arb-chain>]
            │
   ┌────────┴────────┬──────────────────┬──────────────────┬─────────────────┐
   │ execution       │ consensus        │ arbexecution     │ arbitrum-init   │ arbitrum
   │ (geth|neth|     │ (lighthouse|     │ (nethermind-     │ (nitro w/       │ (nitro
   │  besu|erigon)   │  prysm|teku|     │  arbitrum fork)  │  --init flags   │  follower)
   │                 │  lodestar|       │                  │  + then-quit)   │
   │                 │  nimbus)         │                  │                 │
   │  8545/8551      │  5052            │  8547/8552       │  one-shot       │
   │  30303          │                  │  30313           │                 │
   └─────── jwtsecret ──────┘           └── shared jwt ────┴────── shared volume + jwt
```

Service roles added: `arbexecution` (parallel to `opexecution`) and `arbitrum` (parallel to `optimism`). Init companion is part of the `arbitrum` role template, gated by template var.

## Chain selection

CLI: `-n <l1>` and `--chain <arb-chain>` are both flags but each can be inferred from the other.

| user provides                          | resolved L1 | resolved arb chain |
|----------------------------------------|-------------|--------------------|
| `-n sepolia`                           | sepolia     | arbitrum-sepolia   |
| `-n mainnet`                           | mainnet     | arbitrum-one       |
| `--chain arbitrum-sepolia`             | sepolia     | arbitrum-sepolia   |
| `--chain arbitrum-one`                 | mainnet     | arbitrum-one       |
| `-n mainnet --chain arbitrum-one`      | mainnet     | arbitrum-one       |
| `-n mainnet --chain arbitrum-sepolia`  | **error**   | **mismatch**       |
| (neither)                              | mainnet     | arbitrum-one       |

`mainnet` and `sepolia` remain L1 identifiers; Arbitrum chains are *not* registered as sedge networks in `configs/networks.go`. A new `configs/arbitrum_chains.go` maps `--chain` to the per-chain knobs:

```go
var arbitrumChains = map[string]ArbitrumChainConfig{
    "arbitrum-one":     {NethermindChainspec: "arbitrum-mainnet", ChainID: 42161,  L1Network: "mainnet"},
    "arbitrum-sepolia": {NethermindChainspec: "arbitrum-sepolia", ChainID: 421614, L1Network: "sepolia"},
}
```

## Flag surface (`arb-full-node`)

Reused from `op-full-node`:
```
-e, --execution <name[:image]>          # L1 EL client
-c, --consensus <name[:image]>          # L1 CL client
    --checkpoint-sync-url <url>
    --el-extra-flag, --cl-extra-flag, --fallback-execution-urls
    --execution-api-url, --consensus-url  # external L1 mode
    --fee-recipient, --jwt-secret-path, --map-all, --latest
```

New for Arbitrum:
```
    --chain <arbitrum-one|arbitrum-sepolia>     # inferred from -n if omitted
    --arb-image <image>                         # override nitro image
    --arb-execution <name[:image]>              # override nethermind-arbitrum image
    --arb-extra-flag <flag>                     # passed to nitro
    --arb-el-extra-flag <flag>                  # passed to nethermind-arbitrum
    --parent-chain-rpc-url <url>                # external L1 RPC for nitro
    --parent-chain-beacon-url <url>             # external L1 beacon for nitro
    --snapshot-url <url>                        # nitro --init.url; defaults to Offchain Labs published snapshot per chain
```

Setting any of `--execution-api-url`, `--consensus-url`, `--parent-chain-rpc-url`, `--parent-chain-beacon-url` switches the compose to external-L1 mode (L1 EL+CL services omitted). Validation: all four URL surfaces must be provided when in external mode (RPC + beacon for both the L1 client wiring nitro talks to, and the consensus client wiring sedge would otherwise have generated).

## Init pattern (mirrors Nimbus's `consensus-sync`)

```
  arbitrum-init:
    image: ${ARB_IMAGE_VERSION}
    entrypoint: /usr/local/bin/nitro
    volumes:
      - ${ARB_DATA_DIR}:/app/nitro-data
      - ${EC_L2_JWT_SECRET_PATH}:/tmp/jwt/jwtsecret:ro
    command:
      - --init.then-quit=true
      - --init.url=${ARB_SNAPSHOT_URL}              # or --init.empty=true when ARB_SNAPSHOT_URL is empty
      - --persistent.global-config=/app/nitro-data
      - --chain.id={{.ArbChainID}}
      - --parent-chain.connection.url=${EC_API_URL}
      - --parent-chain.blob-client.beacon-url={{.ConsensusApiURL}}
      ...

  arbitrum:
    depends_on:
      arbitrum-init:
        condition: service_completed_successfully
    # same image + jwt + parent-chain wiring, no --init flags
```

The `--init.then-quit=true` flag on the init service guarantees a clean exit after snapshot extraction (or empty-init) so docker compose's `service_completed_successfully` gate fires. The main service then runs in follower mode (`--node.staker.enable=false --node.sequencer=false --execution.forwarding-target=null`).

## JWT lifecycle

Single `jwtsecret` file generated by sedge's existing flow at `JWT_SECRET_PATH`, mounted into:
- `arbexecution` (rw, at `/tmp/jwt/jwtsecret`)
- `arbitrum-init` (ro)
- `arbitrum` (ro)

L1 EL+CL keep their own jwt (the engine API one), separate from the L2 jwt. Nitro speaks plain JSON-RPC to L1, no jwt needed for the parent-chain connection.

## Port plan

Reuses sedge's existing L2 port constants (`configs/ports.go`):

| service        | host port(s)               | role                |
|----------------|----------------------------|---------------------|
| execution      | 8545, 8551, 30303, 8008    | L1 EL               |
| consensus      | 5052, 9000, 8009           | L1 CL               |
| arbexecution   | 8547, 8552, 30313, 8018    | L2 EL               |
| arbitrum-init  | —                          | one-shot            |
| arbitrum       | —                          | rollup-node         |

`arbitrum` exposes nothing externally by default; users hit the L2 RPC via `arbexecution:8547`. `--map-all` opens host bindings for the L2 EL's `8547` / `8552` (mirrors opexecution behavior).

## File layout

New = ✦, modified = ◇.

```
configs/
  ◇ networks.go               # no new constants — Arbitrum chains aren't sedge "networks"
  ◇ init.go                   # untouched
  ✦ arbitrum_chains.go        # ArbitrumChainConfig + lookup + L1↔chain inference
  ◇ client_images.yaml        # add arbexecution.nethermind-arbitrum + arbitrum.nitro pins
  ◇ images.go                 # plumbing for two new roles
cli/
  ◇ generate.go               # const list adds "arbitrum", "arbExecution"
  ◇ sub_gen.go                # new ArbFullNodeSubCmd; chain↔network inference; URL validation
  ◇ factory/node_factory.go   # initializers for arbitrum + arbexecution
internal/pkg/
  ◇ clients/init.go           # registers nethermind-arbitrum + nitro
  ◇ clients/types.go          # SetArbitrumImage / SetArbExecutionImage
  ◇ generate/types.go         # ArbitrumClient + ArbExecutionClient on GenData;
                              # ArbChainID, ArbChainspec, ArbSnapshotUrl, ParentChainRPC, ParentChainBeacon in env data
  ◇ generate/generate_scripts.go  # mapClients() handles the new keys; data flow for new fields
templates/
  ✦ services/merge/arbexecution/nethermind-arbitrum.tmpl
  ✦ services/merge/arbexecution/empty.tmpl
  ✦ services/merge/arbitrum/nitro.tmpl    # contains both arbitrum-init and arbitrum services
  ✦ services/merge/arbitrum/empty.tmpl
  ✦ envs/mainnet/arbexecution/nethermind-arbitrum.tmpl
  ✦ envs/mainnet/arbitrum/nitro.tmpl
  ✦ envs/sepolia/arbexecution/nethermind-arbitrum.tmpl
  ✦ envs/sepolia/arbitrum/nitro.tmpl
```

The env templates carry chain-id, chainspec name, default snapshot URL, and the parent-chain URLs (defaulted to internal docker DNS in bundled mode, overridable by the new flags in external mode).

## Image pins

Both upstream artifacts are pre-release; pin them in `client_images.yaml` like every other client:

```yaml
arbexecution:
  nethermind-arbitrum:
    name: nethermind/nethermind-arbitrum
    version: 0.1.0-alpha
arbitrum:
  nitro:
    name: offchainlabs/nitro-node
    version: v3.10.0-rc.2-746bda2
```

Versions bumped in follow-up PRs as stable tags land. `--latest` flag is *not* wired for these two — the upstreams don't ship a `:latest` tag we trust.

## Validation logic in `arb-full-node` `PreRunE`

1. If `--chain` not provided, infer from `-n`.
2. If `--chain` provided and `-n` not provided, infer L1 from `arbitrumChains[--chain].L1Network`.
3. If both provided, error if `arbitrumChains[--chain].L1Network != -n`.
4. If any of `--execution-api-url`, `--consensus-url`, `--parent-chain-rpc-url`, `--parent-chain-beacon-url` are set, require all four (external-L1 mode is all-or-nothing) and skip L1 service generation.
5. Run sedge's existing network/logging/URL validators.

## End-to-end verification plan

Same shape as the XDC verification we just ran:

1. `make compile-sedge`
2. `sedge generate arb-full-node -n sepolia --execution-api-url <L1 RPC> --consensus-url <L1 beacon> --parent-chain-rpc-url <L1 RPC> --parent-chain-beacon-url <L1 beacon> -p /tmp/sedge-arb-sepolia` (L1 endpoints supplied externally per maintainer's note)
3. `sedge run -p /tmp/sedge-arb-sepolia`
4. Watch `docker logs sedge-arb-init` for snapshot extraction (or empty-init) completing
5. Watch `docker logs sedge-arbitrum` for `Created block` / chain head advancing
6. Watch `docker logs sedge-execution-arb-l2-client` (nethermind-arbitrum) for engine API requests landing

Mainnet variant deferred until the testnet variant boots cleanly.

## Risks / known unknowns

- **`--init.then-quit=true` semantics** must be confirmed against the nitro version we pin. If the flag doesn't exist or behaves differently, fall back to a single-service variant where nitro does init+run in one process (matching the reference compose). Loses the Nimbus-style separation but works.
- **Snapshot URLs for arbitrum-one are large** (hundreds of GB) and the init service will dominate first-run time. Document this in the eventual user-facing readme.
- **Healthcheck portability**: the reference compose's `bash -c '</dev/tcp/...'` assumes bash in the EL image. Switch to a portable form (`nc -z localhost <port>` or a small dedicated check binary) when authoring the template.
- **External-L1 mode JWT**: confirmed nitro→L1 needs no jwt (plain JSON-RPC); nitro→nethermind-arbitrum uses the engine-API jwt that sedge already generates. No extra jwt plumbing.
