# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What Sedge is

Sedge is a Go CLI that generates `docker-compose.yml` + `.env` files to run PoS validator/node stacks (execution + consensus + validator clients, optionally MEV-Boost, Lido CSM, Charon DV, and L2s like Optimism/Base/Taiko/Surge). It does not run the node itself — it shells out to `docker compose` against templates rendered from data the user supplies (interactively or via flags).

Two binaries are produced:
- `cmd/sedge` — the main CLI (`build/sedge`).
- `cmd/lido-exporter` — a Prometheus exporter for Lido CSM metrics (`build/lido-exporter`).

Go module: `github.com/NethermindEth/sedge`, Go 1.22.

## Common commands

All day-to-day commands flow through the Makefile.

```
make install-deps     # one-time: installs gofumpt, mockgen@v1.6.0, abigen
make generate         # regenerates abigen contract bindings + `go generate ./...` (MUST run before tests)
make compile          # builds both binaries into ./build (calls generate)
make compile-sedge    # builds only sedge
make test             # full test suite incl. e2e (timeout 25m)
make test-no-e2e      # unit tests only, skips TestE2E*
make codecov-test     # what CI runs on ubuntu — unit tests with coverage, no e2e
make e2e-test         # e2e suite under ./e2e/sedge/...
make e2e-test-windows # e2e on Windows (skips monitoring stack tests)
make format           # gofumpt -l -w .
make format-check     # CI parity: fails if gofumpt would change anything
```

Run a single test:
```
go test ./cli/... -run TestGenerateCmd -count=1
go test ./internal/pkg/generate -run TestEnvs -v
```

Run the built CLI: `./build/sedge <subcommand>`, or `make run-cli` to compile and launch with `./config.yaml`.

### Code-generation gate (CI will fail without this)

`.github/workflows/code-generation-checks.yml` runs `make generate` and fails the PR if any tracked file changes. Always run `make generate` after editing:
- ABIs/BINs under `internal/lido/contracts/*` (the `make generate` target re-runs `abigen` for csmodule, csfeedistributor, csaccounting, mevboostrelaylist, vebo).
- Anything driving `//go:generate` directives (mockgen sources across `internal/`).

## Architecture

### Entry → cobra → actions

`cmd/sedge/main.go` wires concrete implementations (docker client, filesystem via `afero`, `flock`-based locker, compose+command runners) into:
- `actions.SedgeActions` — high-level operations (generate, run, importKeys, setup, jwt secrets, slashing import/export, getContainers).
- A tree of `cobra.Command`s defined in `cli/` (`cli.go`, `generate.go`, `run.go`, `keys.go`, `importKeys.go`, `monitoring.go`, `slashingExport.go`, `slashingImport.go`, `down.go`, `logs.go`, `listClients.go`, `listNetworks.go`, `lidoStatus.go`, `show.go`, etc.).

`cli/cli.go` is the interactive flow (`sedge cli`); `cli/generate.go` is the non-interactive `sedge generate` and is the workhorse most other commands lean on.

### Templates are the source of truth for docker-compose output

`templates/templates.go` embeds four trees with `//go:embed`:
- `templates/services/merge/{execution,consensus,validator,distributedValidator,optimism,opexecution,taiko,texecution,surge,sexecution}/<client>.tmpl` — per-client compose service fragments.
- `templates/envs/<network>/{execution,consensus,validator,optimism,opexecution,taiko,texecution}/<client>.tmpl` plus an `env_base.tmpl` — per-network `.env` fragments.
- `templates/setup/linux/...` — setup helper scripts.
- `templates/config/` — packaged config defaults.
- `templates/deposit-cli/` — bundled deposit assets.

`internal/pkg/generate/generate_scripts.go` walks `GenData` (clients chosen, network, ports, flags, mev settings, L2 wiring, etc. — see `internal/pkg/generate/types.go`), assembles the right template fragments per network/client combo, executes them with `text/template`, and writes the `docker-compose.yml` + `.env` pair. **Adding a new network or client almost always means: add templates under both `services/merge/<role>/` and `envs/<network>/<role>/`, register the network in `configs/networks.go` + `configs/init.go`, and extend `internal/pkg/clients` + `cli/factory/node_factory.go`.**

### Configs layer

`configs/` holds package-level data, not user config:
- `networks.go` — `NetworkSupported()`, `NetworkCheck()`, the `Network*` constants. Mirror constants also exist in `cli/cli.go` — both must be updated together.
- `init.go` — `networksConfigs` map of `NetworkConfig` (genesis fork version, MEV support, checkpoint sync URL, MEV relay list, chain id, randomization weight).
- `client_images.yaml` + `images.go` — pinned default docker images per client/role; **versions are hardcoded on purpose, do not switch back to `:latest`** (see `cf1ffb2 chore: hardcode version on taiko, not allowing latest`).
- `public_rpcs.go`, `sequencer_urls.go`, `ports.go`, `paths.go`, `superchain*.go` — companion lookup tables.

### Other internal packages

- `internal/pkg/clients` — typed client registry (`Client`, `OrderedClients`, validation, per-role image resolution). `cli/factory/node_factory.go` plugs role-specific initializers (execution / consensus / validator / optimism / taiko / surge / dv) over a shared `NodeInitializer` interface.
- `internal/pkg/options` — `SedgeOptions` polymorphism for ethereum vs lido setups.
- `internal/pkg/keystores`, `internal/crypto` — keystore generation (own implementation, not staking-deposit-cli; the README calls out this is unaudited).
- `internal/pkg/services`, `internal/compose`, `internal/pkg/commands` — docker / `docker compose` / shell wrappers used by actions.
- `internal/pkg/dependencies` — checks and (Linux) installs docker.
- `internal/monitoring` + `internal/monitoring/services/{grafana,prometheus,node_exporter,lido_exporter}` — the `sedge monitoring` stack; the manager composes these `ServiceAPI` implementations into its own compose project under `~/.sedge/monitoring`.
- `internal/lido/contracts/{csmodule,csfeedistributor,csaccounting,mevboostrelaylist,vebo}` — abigen-generated bindings; do not hand-edit the `*.go` files, edit the `.abi`/`.bin` and rerun `make generate`.
- `internal/ui` — survey-based prompts and table renderers used by interactive commands.

### E2E

`e2e/sedge/...` — drives the built binary against real docker. `make e2e-test` runs everything; Windows skips `TestE2E_MonitoringStack`. Tests are gated by the `TestE2E` prefix, which is why `test-no-e2e` uses `-skip TestE2E`.

## Conventions worth knowing

- **Commits**: Conventional Commits — `<type>(<scope>): <subject>` (see CONTRIBUTING.md). Branches are `feature/<name>` or `fix/<name>` in kebab-case.
- **CHANGELOG.md**: every PR updates the `Unreleased` section using Keep-a-Changelog categories (Added/Changed/Deprecated/Removed/Fixed/Security).
- **Formatter is gofumpt, not gofmt** — `make format-check` is enforced in CI.
- **Version stamping**: `make compile` injects `internal/utils.Version` from the latest git tag via `-ldflags`. `sedge version` reads it.
- The README's "Supported networks and clients" table is the user-facing source of truth for which client × network × role triples are expected to work; if you wire a new combo, update it.
