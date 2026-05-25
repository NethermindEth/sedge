# Arbitrum integration implementation plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Add a `sedge generate arb-full-node` subcommand that emits a working Arbitrum L2 stack (`nethermind-arbitrum` + `nitro`) sitting on top of an Ethereum L1 stack, mirroring sedge's existing `op-full-node` pattern.

**Architecture:** Two new service roles (`arbexecution`, `arbitrum`), a new chain-config lookup in `configs/arbitrum_chains.go`, a Nimbus-style init companion service for nitro snapshot extraction, and a subcommand with bidirectional `-n` ↔ `--chain` inference. Bundled L1 templates are built but verification targets external-L1 mode using maintainer-supplied endpoints.

**Tech Stack:** Go 1.22, cobra, embedded `text/template`, docker compose, `nethermind/nethermind-arbitrum:0.1.0-alpha`, `offchainlabs/nitro-node:v3.10.0-rc.2-746bda2`.

**Reference design:** @docs/plans/2026-05-25-arbitrum-integration-design.md

**Branch:** `add-arbitrum-chain` (off `core`). Internal branch — no README/CHANGELOG/docs-site work.

---

## Pre-flight

Before Task 1, on `add-arbitrum-chain`:

```bash
git rev-parse --abbrev-ref HEAD   # expect: add-arbitrum-chain
git log -1 --oneline              # expect: 9d87649 docs(arbitrum): design ...
make install-deps                 # gofumpt, mockgen, abigen
make generate                     # abigen + go generate ./...
make test-no-e2e                  # baseline must be green before any code change
```

If any of those fail, stop and fix the environment first.

---

## Task 1: `configs/arbitrum_chains.go` — chain config table + inference

**Files:**
- Create: `configs/arbitrum_chains.go`
- Create: `configs/arbitrum_chains_test.go`

**Step 1: Write the failing test**

```go
// configs/arbitrum_chains_test.go
package configs

import (
	"errors"
	"testing"
)

func TestArbitrumChainConfig(t *testing.T) {
	cfg, err := ArbitrumChain("arbitrum-one")
	if err != nil {
		t.Fatalf("arbitrum-one: unexpected error %v", err)
	}
	if cfg.ChainID != 42161 {
		t.Errorf("arbitrum-one ChainID: want 42161, got %d", cfg.ChainID)
	}
	if cfg.NethermindChainspec != "arbitrum-mainnet" {
		t.Errorf("arbitrum-one NethermindChainspec: want arbitrum-mainnet, got %q", cfg.NethermindChainspec)
	}
	if cfg.L1Network != NetworkMainnet {
		t.Errorf("arbitrum-one L1Network: want %q, got %q", NetworkMainnet, cfg.L1Network)
	}

	cfg, err = ArbitrumChain("arbitrum-sepolia")
	if err != nil {
		t.Fatalf("arbitrum-sepolia: unexpected error %v", err)
	}
	if cfg.ChainID != 421614 {
		t.Errorf("arbitrum-sepolia ChainID: want 421614, got %d", cfg.ChainID)
	}
	if cfg.L1Network != NetworkSepolia {
		t.Errorf("arbitrum-sepolia L1Network: want %q, got %q", NetworkSepolia, cfg.L1Network)
	}

	if _, err := ArbitrumChain("does-not-exist"); !errors.Is(err, ErrInvalidArbitrumChain) {
		t.Errorf("unknown chain: want ErrInvalidArbitrumChain, got %v", err)
	}
}

func TestArbitrumSupportedChains(t *testing.T) {
	chains := ArbitrumSupportedChains()
	if len(chains) != 2 {
		t.Fatalf("want 2 chains, got %d (%v)", len(chains), chains)
	}
	wantSet := map[string]bool{"arbitrum-one": false, "arbitrum-sepolia": false}
	for _, c := range chains {
		if _, ok := wantSet[c]; !ok {
			t.Errorf("unexpected chain %q", c)
		}
		wantSet[c] = true
	}
	for c, seen := range wantSet {
		if !seen {
			t.Errorf("missing chain %q in supported list", c)
		}
	}
}

func TestResolveArbitrumChainAndL1(t *testing.T) {
	tests := []struct {
		name      string
		network   string
		chain     string
		wantNet   string
		wantChain string
		wantErr   bool
	}{
		{"L1 mainnet only -> default arbitrum-one", "mainnet", "", "mainnet", "arbitrum-one", false},
		{"L1 sepolia only -> default arbitrum-sepolia", "sepolia", "", "sepolia", "arbitrum-sepolia", false},
		{"chain only -> derive L1", "", "arbitrum-sepolia", "sepolia", "arbitrum-sepolia", false},
		{"both consistent", "mainnet", "arbitrum-one", "mainnet", "arbitrum-one", false},
		{"both inconsistent", "mainnet", "arbitrum-sepolia", "", "", true},
		{"unknown chain", "mainnet", "arbitrum-nova", "", "", true},
		{"unsupported L1 holesky", "holesky", "", "", "", true},
		{"neither given", "", "", "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			net, chain, err := ResolveArbitrumChainAndL1(tt.network, tt.chain)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if net != tt.wantNet || chain != tt.wantChain {
				t.Errorf("got (%q,%q), want (%q,%q)", net, chain, tt.wantNet, tt.wantChain)
			}
		})
	}
}
```

**Step 2: Run test to verify it fails**

```bash
go test ./configs/ -run TestArbitrum -v
```

Expected: `undefined: ArbitrumChain` / `undefined: ResolveArbitrumChainAndL1` / `undefined: ErrInvalidArbitrumChain` — FAIL.

**Step 3: Write minimal implementation**

```go
// configs/arbitrum_chains.go
/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package configs

import (
	"errors"
	"fmt"
	"sort"
)

const (
	ArbitrumChainOne     = "arbitrum-one"
	ArbitrumChainSepolia = "arbitrum-sepolia"
)

var (
	ErrInvalidArbitrumChain    = errors.New("invalid arbitrum chain")
	ErrArbitrumChainL1Mismatch = errors.New("arbitrum chain does not match selected L1 network")
	ErrArbitrumNoSelector      = errors.New("must provide --network and/or --chain for arbitrum")
)

type ArbitrumChainConfig struct {
	NethermindChainspec string // --config=<value> for nethermind-arbitrum
	ChainID             uint64 // --chain.id=<value> for nitro
	L1Network           string // sedge network name of the parent chain
	DefaultSnapshotURL  string // optional --init.url=<value>; empty -> --init.empty=true
}

var arbitrumChains = map[string]ArbitrumChainConfig{
	ArbitrumChainOne: {
		NethermindChainspec: "arbitrum-mainnet",
		ChainID:             42161,
		L1Network:           NetworkMainnet,
		DefaultSnapshotURL:  "", // operator must supply for production timelines; sedge stays neutral
	},
	ArbitrumChainSepolia: {
		NethermindChainspec: "arbitrum-sepolia",
		ChainID:             421614,
		L1Network:           NetworkSepolia,
		DefaultSnapshotURL:  "",
	},
}

func ArbitrumChain(name string) (ArbitrumChainConfig, error) {
	cfg, ok := arbitrumChains[name]
	if !ok {
		return ArbitrumChainConfig{}, fmt.Errorf("%w: %s", ErrInvalidArbitrumChain, name)
	}
	return cfg, nil
}

func ArbitrumSupportedChains() []string {
	out := make([]string, 0, len(arbitrumChains))
	for k := range arbitrumChains {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

// ResolveArbitrumChainAndL1 applies the bidirectional inference rules from the design.
// At least one of network/chain must be non-empty.
func ResolveArbitrumChainAndL1(network, chain string) (resolvedNetwork, resolvedChain string, err error) {
	switch {
	case network == "" && chain == "":
		return "", "", ErrArbitrumNoSelector
	case chain == "":
		// Network -> default chain
		for name, cfg := range arbitrumChains {
			if cfg.L1Network == network {
				return network, name, nil
			}
		}
		return "", "", fmt.Errorf("%w: no arbitrum chain registered for L1 %q", ErrInvalidArbitrumChain, network)
	case network == "":
		cfg, err := ArbitrumChain(chain)
		if err != nil {
			return "", "", err
		}
		return cfg.L1Network, chain, nil
	default:
		cfg, err := ArbitrumChain(chain)
		if err != nil {
			return "", "", err
		}
		if cfg.L1Network != network {
			return "", "", fmt.Errorf("%w: chain %q expects L1 %q, got %q",
				ErrArbitrumChainL1Mismatch, chain, cfg.L1Network, network)
		}
		return network, chain, nil
	}
}
```

**Step 4: Run test to verify it passes**

```bash
go test ./configs/ -run TestArbitrum -v
```

Expected: all 3 tests PASS, no `notest` warnings.

**Step 5: Commit**

```bash
git add configs/arbitrum_chains.go configs/arbitrum_chains_test.go
git commit -S -m "feat(arbitrum): chain config table and L1 inference"
```

---

## Task 2: `client_images.yaml` + `images.go` — pin nethermind-arbitrum and nitro

**Files:**
- Modify: `configs/client_images.yaml` (append two new top-level keys)
- Modify: `configs/images.go` (add the two new role keys to whatever struct/map is used)

**Step 1: Find the structure being amended**

```bash
cat configs/client_images.yaml
grep -n "opexecution\|optimism\|taiko\|surge" configs/images.go
```

Expected: `client_images.yaml` has `execution`/`consensus`/`validator`/`distributed`/`optimism`/`opexecution`/`taiko`/`texecution`/`surge`/`sexecution` keys. `configs/images.go` has an embed of the yaml plus a Go struct or map mirroring it.

**Step 2: Write the failing test**

Edit (or create) `configs/images_test.go`:

```go
package configs

import "testing"

func TestArbitrumClientImagesPresent(t *testing.T) {
	imgs := ClientImages() // existing accessor — confirm name with grep before editing
	cases := []struct {
		role, client, wantImage, wantVersion string
	}{
		{"arbexecution", "nethermind-arbitrum", "nethermind/nethermind-arbitrum", "0.1.0-alpha"},
		{"arbitrum", "nitro", "offchainlabs/nitro-node", "v3.10.0-rc.2-746bda2"},
	}
	for _, tc := range cases {
		img, ok := imgs.Get(tc.role, tc.client) // also confirm Get(role, client) signature
		if !ok {
			t.Fatalf("%s/%s missing from client_images.yaml", tc.role, tc.client)
		}
		if img.Name != tc.wantImage || img.Version != tc.wantVersion {
			t.Errorf("%s/%s: got %s:%s, want %s:%s", tc.role, tc.client,
				img.Name, img.Version, tc.wantImage, tc.wantVersion)
		}
	}
}
```

**Note for executor:** the accessor name (`ClientImages`, `Get`) may differ; before writing the test, `grep -n "func.*ClientImages\|Image{\|images\[" configs/images.go` and adjust the test to use the actual API. The test's intent is "these entries exist and have these values."

**Step 3: Run test to verify it fails**

```bash
go test ./configs/ -run TestArbitrumClientImagesPresent -v
```

Expected: keys missing — FAIL.

**Step 4: Add YAML entries**

Edit `configs/client_images.yaml`, append:

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

If `configs/images.go` declares an explicit Go struct mirroring the YAML schema, add the two corresponding fields. If it uses `map[string]map[string]Image` or similar generic lookup, no Go change needed beyond the YAML edit.

**Step 5: Run test to verify it passes**

```bash
go test ./configs/ -run TestArbitrumClientImagesPresent -v
```

Expected: PASS.

**Step 6: Commit**

```bash
git add configs/client_images.yaml configs/images.go configs/images_test.go
git commit -S -m "feat(arbitrum): pin nethermind-arbitrum and nitro images"
```

---

## Task 3: `internal/pkg/clients/init.go` + `types.go` — register the two new clients

**Files:**
- Modify: `internal/pkg/clients/init.go`
- Modify: `internal/pkg/clients/types.go`

**Step 1: Read the existing pattern**

```bash
sed -n '1,80p' internal/pkg/clients/init.go
grep -n "setOptimismImage\|setOpExecutionImage\|SetOpExecutionImage\|SetTaikoImage" internal/pkg/clients/types.go
```

Expected: `init.go` populates `clientsImages` and supported-client lists keyed by role. `types.go` has a `SetImageOrDefault` switch on `c.Type` calling a per-role `setXImage` private helper.

**Step 2: Write the failing test**

Append to `internal/pkg/clients/types_test.go` (or create if absent):

```go
func TestSetArbitrumImages(t *testing.T) {
	t.Run("arbexecution defaults", func(t *testing.T) {
		c := &Client{Name: "nethermind-arbitrum", Type: "arbexecution"}
		c.SetImageOrDefault("")
		if c.Image == "" {
			t.Fatalf("default image not set")
		}
	})
	t.Run("arbitrum defaults", func(t *testing.T) {
		c := &Client{Name: "nitro", Type: "arbitrum"}
		c.SetImageOrDefault("")
		if c.Image == "" {
			t.Fatalf("default image not set")
		}
	})
	t.Run("arbexecution override", func(t *testing.T) {
		c := &Client{Name: "nethermind-arbitrum", Type: "arbexecution"}
		c.SetImageOrDefault("my.registry/custom:tag")
		if c.Image != "my.registry/custom:tag" {
			t.Errorf("override not honored, got %q", c.Image)
		}
	})
}

func TestArbitrumSupportedClients(t *testing.T) {
	c := ClientInfo{Network: "sepolia"}
	for _, role := range []string{"arbexecution", "arbitrum"} {
		got, err := c.SupportedClients(role)
		if err != nil {
			t.Fatalf("role %s: %v", role, err)
		}
		if len(got) == 0 {
			t.Errorf("role %s: empty supported list", role)
		}
	}
}
```

**Step 3: Run test to verify it fails**

```bash
go test ./internal/pkg/clients/ -run "TestSetArbitrumImages|TestArbitrumSupportedClients" -v
```

Expected: missing case in switch / unknown client type — FAIL.

**Step 4: Wire the clients**

In `internal/pkg/clients/init.go`, register `nethermind-arbitrum` under `arbexecution` and `nitro` under `arbitrum`. Follow the exact pattern used for `opexecution` and `optimism` — copy those two clauses verbatim, then rename.

In `internal/pkg/clients/types.go`, add to the `SetImageOrDefault` switch:

```go
case "arbexecution":
    c.setArbExecutionImage(image)
case "arbitrum":
    c.setArbitrumImage(image)
```

Add the two helpers below the existing `setOpExecutionImage` and `setOptimismImage` definitions — they should be exact duplicates with names changed and the role string changed.

**Step 5: Run test to verify it passes**

```bash
go test ./internal/pkg/clients/ -v
```

Expected: all tests PASS including the two new ones, no regressions in the existing types_test.go.

**Step 6: Commit**

```bash
git add internal/pkg/clients/
git commit -S -m "feat(arbitrum): register nethermind-arbitrum and nitro clients"
```

---

## Task 4: `cli/factory/node_factory.go` — add `ArbExecutionNodeInitializer` and `ArbitrumNodeInitializer`

**Files:**
- Modify: `cli/factory/node_factory.go`

**Step 1: Read the existing pattern**

```bash
grep -n "OptimismNodeInitializer\|OpExecutionNodeInitializer\|TaikoNodeInitializer" cli/factory/node_factory.go
```

Find the two definitions for Optimism — `OptimismNodeInitializer` and `OpExecutionNodeInitializer` — and read the surrounding 40 lines for each.

**Step 2: Add the two new initializers**

After the Optimism block (whichever appears last in the file), add:

```go
// ArbExecutionNodeInitializer handles nethermind-arbitrum initialization
type ArbExecutionNodeInitializer struct {
	BaseNodeInitializer
}

func NewArbExecutionNodeInitializer() *ArbExecutionNodeInitializer {
	return &ArbExecutionNodeInitializer{
		BaseNodeInitializer{
			serviceType: "arbexecution",
			config: clientConfig{
				clientType: "arbexecution",
				flagName:   "", // wired by InitializeClients from flags.arbExecutionName at call time
			},
		},
	}
}

func (i *ArbExecutionNodeInitializer) UpdateResult(result *clients.Clients, client *clients.Client) {
	result.ArbExecution = client
}

// ArbitrumNodeInitializer handles nitro initialization
type ArbitrumNodeInitializer struct {
	BaseNodeInitializer
}

func NewArbitrumNodeInitializer() *ArbitrumNodeInitializer {
	return &ArbitrumNodeInitializer{
		BaseNodeInitializer{
			serviceType: "arbitrum",
			config: clientConfig{
				clientType: "arbitrum",
				flagName:   "",
			},
		},
	}
}

func (i *ArbitrumNodeInitializer) UpdateResult(result *clients.Clients, client *clients.Client) {
	result.Arbitrum = client
}
```

**Note:** `clients.Clients` needs `ArbExecution` and `Arbitrum` fields. Add them in Task 5 below; this task is allowed to leave a compile error temporarily *only because Task 5 immediately follows*. The cleaner sequence is to add the fields first — flip the order if you prefer.

**Step 3: Register the initializers**

Find the function that constructs the initializer slice (likely `NewNodeFactory` or `InitializeClients`). Add:

```go
&ArbExecutionNodeInitializer{...},
&ArbitrumNodeInitializer{...},
```

…using the same `NewX` constructors as the Optimism pattern. Wire `flags.arbExecutionName`/`flags.arbitrumName` into their `clientConfig.flagName` in the same place the existing factory wires `flags.optimismName`.

**Step 4: Build (still expected to fail)**

```bash
go build ./cli/factory/...
```

Expected: errors about `Clients.ArbExecution` / `Clients.Arbitrum` undefined — these get added in Task 5.

**Step 5: Stash and proceed to Task 5**

Don't commit yet. Task 5 will land the `Clients` struct fields and a single passing build, then we commit factory + clients together.

---

## Task 5: Extend `clients.Clients` + `clients.OrderedClients` with the new roles

**Files:**
- Modify: `internal/pkg/clients/clients.go` (or wherever `Clients` and `OrderedClients` are defined — confirm with `grep -n "type Clients\|type OrderedClients" internal/pkg/clients/`)

**Step 1: Read the existing pattern**

```bash
grep -n "Optimism\s*\*Client\|OpExecution\s*\*Client" internal/pkg/clients/clients.go
```

Expected: `Clients` struct has fields `Execution`, `Consensus`, `Validator`, `Optimism`, `OpExecution`, etc.

**Step 2: Add new fields**

```go
type Clients struct {
    // ... existing ...
    Arbitrum     *Client
    ArbExecution *Client
}
```

If `OrderedClients` is a `map[string][]*Client` keyed by role, ensure the init function in `init.go` populates `"arbitrum"` and `"arbexecution"` keys (Task 3 should already have done this).

**Step 3: Run the build**

```bash
go build ./...
```

Expected: PASS, no compile errors.

**Step 4: Run the test suite**

```bash
make test-no-e2e
```

Expected: All existing tests still pass; new tests from Tasks 1, 2, 3 also pass.

**Step 5: Commit Tasks 4 + 5 together**

```bash
git add cli/factory/node_factory.go internal/pkg/clients/clients.go
git commit -S -m "feat(arbitrum): node factory initializers and Clients fields"
```

---

## Task 6: `cli/generate.go` — add role consts

**Files:**
- Modify: `cli/generate.go`

**Step 1: Add to the const block**

Edit `cli/generate.go:48-51`:

```go
const (
    execution, consensus, validator, distributedValidator, mevBoost, optimism, opExecution, taiko, tExecution, surge, sExecution, arbitrum, arbExecution = "execution", "consensus", "validator", "distributedValidator", "mev-boost", "optimism", "opexecution", "taiko", "texecution", "surge", "sexecution", "arbitrum", "arbexecution"
    jwtPathName                                                                                                                                          = "jwtsecret"
)
```

**Step 2: Verify build**

```bash
go build ./cli/...
```

Expected: PASS.

**Step 3: Commit**

```bash
git add cli/generate.go
git commit -S -m "feat(arbitrum): generate.go role consts"
```

---

## Task 7: `internal/pkg/generate/types.go` — `ArbitrumClient`, `ArbExecutionClient`, env fields

**Files:**
- Modify: `internal/pkg/generate/types.go`

**Step 1: Add to `GenData`**

After `SurgeClient *clients.Client` (or wherever the L2 client fields cluster):

```go
type GenData struct {
    // ... existing ...
    ArbitrumClient     *clients.Client
    ArbExecutionClient *clients.Client
    ArbChain           string // e.g. "arbitrum-one"
    ArbSnapshotURL     string // optional override; empty => use chain default or --init.empty=true
    ParentChainRPCURL  string // external L1 RPC; empty => bundled L1 via internal DNS
    ParentChainBeacon  string // external L1 beacon; empty => bundled L1 via internal DNS
}
```

**Step 2: Add to `EnvData`** (the struct rendered into the .env file)

```go
type EnvData struct {
    // ... existing ...
    ArbImage             string
    ArbExecutionImage    string
    ArbDataDir           string
    ArbExecutionDataDir  string
    ArbChainID           uint64
    ArbChainspec         string
    ArbSnapshotURL       string
    ParentChainRPCURL    string
    ParentChainBeacon    string
}
```

**Step 3: Add to `DockerComposeData`** (the struct rendered into docker-compose.yml)

```go
type DockerComposeData struct {
    // ... existing ...
    ArbChainID    uint64
    ArbChainspec  string
    ArbInitMode   string // "snapshot" | "empty" — selects the template branch
}
```

**Step 4: Verify build**

```bash
go build ./...
```

Expected: PASS — these are pure type additions, nothing else references them yet.

**Step 5: Commit**

```bash
git add internal/pkg/generate/types.go
git commit -S -m "feat(arbitrum): GenData/EnvData/DockerComposeData fields"
```

---

## Task 8: `internal/pkg/generate/generate_scripts.go` — wire the new clients through

**Files:**
- Modify: `internal/pkg/generate/generate_scripts.go`

**Step 1: Update `mapClients`**

Locate `mapClients(gd *GenData)` (around line 135). Add to the returned map:

```go
cls := map[string]*clients.Client{
    // ... existing ...
    arbitrum:     gd.ArbitrumClient,
    arbExecution: gd.ArbExecutionClient,
}
```

Use the const names from `cli/generate.go` if they're re-exported in this file, or define local equivalents (the existing file defines `optimism = "optimism"`, etc. — follow that pattern).

**Step 2: Wire the env-data fields**

In `EnvFile`, where `EnvData{...}` is constructed (search for `EnvData{`), populate the new fields:

```go
ArbImage:            imageOrEmpty(cls[arbitrum], gd.LatestVersion),
ArbExecutionImage:   imageOrEmpty(cls[arbExecution], gd.LatestVersion),
ArbDataDir:          "./arbitrum-data",
ArbExecutionDataDir: "./arbexecution-data",
ArbSnapshotURL:      gd.ArbSnapshotURL,
ParentChainRPCURL:   gd.ParentChainRPCURL,
ParentChainBeacon:   gd.ParentChainBeacon,
```

And resolve the chain config:

```go
if gd.ArbChain != "" {
    chainCfg, err := configs.ArbitrumChain(gd.ArbChain)
    if err != nil {
        return err
    }
    envData.ArbChainID = chainCfg.ChainID
    envData.ArbChainspec = chainCfg.NethermindChainspec
    if envData.ArbSnapshotURL == "" {
        envData.ArbSnapshotURL = chainCfg.DefaultSnapshotURL
    }
}
```

…and analogously in `ComposeFile` for `DockerComposeData.ArbChainID`, `ArbChainspec`, and `ArbInitMode` (set to `"snapshot"` when `gd.ArbSnapshotURL != ""`, else `"empty"`).

**Step 3: Build + run existing tests**

```bash
make test-no-e2e
```

Expected: all tests PASS (the new fields are not yet exercised, but nothing should regress).

**Step 4: Commit**

```bash
git add internal/pkg/generate/generate_scripts.go
git commit -S -m "feat(arbitrum): wire client + env + compose data"
```

---

## Task 9: `templates/services/merge/arbexecution/*` — service templates for nethermind-arbitrum

**Files:**
- Create: `templates/services/merge/arbexecution/empty.tmpl`
- Create: `templates/services/merge/arbexecution/nethermind-arbitrum.tmpl`

**Step 1: Add the empty template**

```
{{/* empty.tmpl */}}
{{ define "arbexecution" }}{{ end }}
```

**Step 2: Add the nethermind-arbitrum template**

Use `templates/services/merge/opexecution/opnethermind.tmpl` as the reference. Key differences:
- service name: `arbexecution`
- container name: `sedge-arbexecution-client`
- image var: `EC_L2_IMAGE_VERSION`
- command: `-c=${ARB_CHAINSPEC}` (passed via env from `EnvData.ArbChainspec`)
- no `--Optimism.SequencerUrl` (Arbitrum's EL doesn't need a sequencer URL)
- mount `${EC_L2_JWT_SECRET_PATH}` rw
- ports: `{{.ElL2DiscoveryPort}}`, `{{.ElL2MetricsPort}}` (and `{{.ElL2ApiPort}}`/`{{.ElL2AuthPort}}` when `.MapAllPorts`)
- expose `{{.ElL2ApiPort}}`, `{{.ElL2AuthPort}}`
- healthcheck: `nc -z localhost {{.ElL2AuthPort}}` (portable; design Risk #3)

```
{{/* nethermind-arbitrum.tmpl */}}
{{ define "arbexecution" }}
  arbexecution:
    tty: true
    environment:
      - TERM=xterm-256color
      - COLORTERM=truecolor
    stop_grace_period: 30s
    container_name: sedge-arbexecution-client{{if .ContainerTag}}-{{.ContainerTag}}{{end}}
    restart: unless-stopped
    image: ${EC_L2_IMAGE_VERSION}
    networks:
      - sedge
    volumes:
      - ${EC_L2_DATA_DIR}:/app/nethermind_db
      - ${EC_L2_JWT_SECRET_PATH}:/tmp/jwt/jwtsecret
    ports:
      - "{{.ElL2DiscoveryPort}}:{{.ElL2DiscoveryPort}}/tcp"
      - "{{.ElL2DiscoveryPort}}:{{.ElL2DiscoveryPort}}/udp"
      - "{{.ElL2MetricsPort}}:{{.ElL2MetricsPort}}"{{if .MapAllPorts}}
      - "{{.ElL2ApiPort}}:{{.ElL2ApiPort}}"
      - "{{.ElL2AuthPort}}:{{.ElL2AuthPort}}"{{end}}
    expose:
      - {{.ElL2ApiPort}}
      - {{.ElL2AuthPort}}
    command:
      - -c=${ARB_CHAINSPEC}
      - --data-dir=/app/nethermind_db
      - --JsonRpc.Enabled=true
      - --JsonRpc.Host=0.0.0.0
      - --JsonRpc.Port={{.ElL2ApiPort}}
      - --JsonRpc.JwtSecretFile=/tmp/jwt/jwtsecret
      - --JsonRpc.EngineHost=0.0.0.0
      - --JsonRpc.EnginePort={{.ElL2AuthPort}}
      - --Network.DiscoveryPort={{.ElL2DiscoveryPort}}
      - --HealthChecks.Enabled=true
      - --Metrics.Enabled=true
      - --Metrics.ExposePort={{.ElL2MetricsPort}}{{range $flag := .ElL2ExtraFlags}}
      - --{{$flag}}{{end}}
    healthcheck:
      test: ["CMD-SHELL", "nc -z localhost {{.ElL2AuthPort}} || exit 1"]
      interval: 10s
      timeout: 5s
      retries: 10
      start_period: 30s{{if .LoggingDriver}}
    logging:
      driver: "{{.LoggingDriver}}"{{if eq .LoggingDriver "json-file"}}
      options:
        max-size: "10m"
        max-file: "10"{{end}}{{end}}
{{ end }}
```

**Step 3: Verify embeds**

```bash
go build ./templates/...
go test ./templates/... 2>&1 | head -5
```

Expected: PASS. The `//go:embed services` directive picks up the new files automatically.

**Step 4: Commit**

```bash
git add templates/services/merge/arbexecution/
git commit -S -m "feat(arbitrum): arbexecution service templates"
```

---

## Task 10: `templates/services/merge/arbitrum/*` — service templates for nitro (init + main)

**Files:**
- Create: `templates/services/merge/arbitrum/empty.tmpl`
- Create: `templates/services/merge/arbitrum/nitro.tmpl`

**Step 1: Empty template**

```
{{/* empty.tmpl */}}
{{ define "arbitrum" }}{{ end }}
```

**Step 2: nitro template (two services: arbitrum-init + arbitrum)**

Pattern reference: `templates/services/merge/consensus/nimbus.tmpl` (split-service pattern, init companion with `depends_on: condition: service_completed_successfully`).

```
{{/* nitro.tmpl */}}
{{ define "arbitrum" }}
  arbitrum-init:
    tty: true
    environment:
      - TERM=xterm-256color
      - COLORTERM=truecolor
    container_name: sedge-arbitrum-init{{if .ContainerTag}}-{{.ContainerTag}}{{end}}
    image: ${ARB_IMAGE_VERSION}
    entrypoint: /usr/local/bin/nitro
    networks:
      - sedge
    volumes:
      - ${ARB_DATA_DIR}:/app/nitro-data
      - ${EC_L2_JWT_SECRET_PATH}:/tmp/jwt/jwtsecret:ro
    depends_on:
      arbexecution:
        condition: service_healthy
    command:
      - --init.then-quit=true{{if eq .ArbInitMode "snapshot"}}
      - --init.url=${ARB_SNAPSHOT_URL}{{else}}
      - --init.empty=true{{end}}
      - --persistent.global-config=/app/nitro-data
      - --chain.id={{.ArbChainID}}
      - --parent-chain.connection.url=${PARENT_CHAIN_RPC_URL}
      - --parent-chain.blob-client.beacon-url=${PARENT_CHAIN_BEACON_URL}
      - --node.execution-rpc-client.url=http://arbexecution:{{.ElL2AuthPort}}
      - --node.execution-rpc-client.jwtsecret=/tmp/jwt/jwtsecret
      - --init.validate-genesis-assertion=false{{if .LoggingDriver}}
    logging:
      driver: "{{.LoggingDriver}}"{{if eq .LoggingDriver "json-file"}}
      options:
        max-size: "10m"
        max-file: "10"{{end}}{{end}}

  arbitrum:
    tty: true
    environment:
      - TERM=xterm-256color
      - COLORTERM=truecolor
    stop_grace_period: 5m
    container_name: sedge-arbitrum{{if .ContainerTag}}-{{.ContainerTag}}{{end}}
    restart: unless-stopped
    image: ${ARB_IMAGE_VERSION}
    entrypoint: /usr/local/bin/nitro
    networks:
      - sedge
    volumes:
      - ${ARB_DATA_DIR}:/app/nitro-data
      - ${EC_L2_JWT_SECRET_PATH}:/tmp/jwt/jwtsecret:ro
    depends_on:
      arbitrum-init:
        condition: service_completed_successfully
      arbexecution:
        condition: service_healthy
    command:
      - --persistent.global-config=/app/nitro-data
      - --chain.id={{.ArbChainID}}
      - --parent-chain.connection.url=${PARENT_CHAIN_RPC_URL}
      - --parent-chain.blob-client.beacon-url=${PARENT_CHAIN_BEACON_URL}
      - --node.execution-rpc-client.url=http://arbexecution:{{.ElL2AuthPort}}
      - --node.execution-rpc-client.jwtsecret=/tmp/jwt/jwtsecret
      - --node.staker.enable=false
      - --node.sequencer=false
      - --execution.forwarding-target=null
      - --init.validate-genesis-assertion=false{{range $flag := .OPExtraFlags}}
      - --{{$flag}}{{end}}{{if .LoggingDriver}}
    logging:
      driver: "{{.LoggingDriver}}"{{if eq .LoggingDriver "json-file"}}
      options:
        max-size: "10m"
        max-file: "10"{{end}}{{end}}
{{ end }}
```

**Note:** `OPExtraFlags` is reused here for nitro extra flags; if you prefer a separate `ArbExtraFlags` field on `DockerComposeData`, add it in Task 7 and use it here instead.

**Step 3: Build**

```bash
go build ./templates/...
```

Expected: PASS.

**Step 4: Commit**

```bash
git add templates/services/merge/arbitrum/
git commit -S -m "feat(arbitrum): nitro service templates with init companion"
```

---

## Task 11: `templates/envs/{mainnet,sepolia}/arbexecution/nethermind-arbitrum.tmpl`

**Files:**
- Create: `templates/envs/mainnet/arbexecution/nethermind-arbitrum.tmpl`
- Create: `templates/envs/sepolia/arbexecution/nethermind-arbitrum.tmpl`

**Step 1: Mainnet env template**

```
{{/* nethermind-arbitrum.tmpl */}}
{{ define "arbexecution" }}
# --- Arbitrum L2 Execution - Nethermind-Arbitrum - configuration ---
EC_L2_IMAGE_VERSION={{.ArbExecutionImage}}
EC_L2_DATA_DIR={{.ArbExecutionDataDir}}
EC_L2_JWT_SECRET_PATH={{.JWTSecretPath}}
ARB_CHAINSPEC={{.ArbChainspec}}
{{ end }}
```

**Step 2: Sepolia env template**

Identical content. The chainspec value differs because `EnvData.ArbChainspec` is populated from `configs.ArbitrumChain(gd.ArbChain).NethermindChainspec` at generation time, not hardcoded per env folder. The two files only exist because sedge resolves env templates by folder name = `gd.Network`.

```
{{/* nethermind-arbitrum.tmpl */}}
{{ define "arbexecution" }}
# --- Arbitrum L2 Execution - Nethermind-Arbitrum - configuration ---
EC_L2_IMAGE_VERSION={{.ArbExecutionImage}}
EC_L2_DATA_DIR={{.ArbExecutionDataDir}}
EC_L2_JWT_SECRET_PATH={{.JWTSecretPath}}
ARB_CHAINSPEC={{.ArbChainspec}}
{{ end }}
```

**Step 3: Build**

```bash
go build ./templates/...
```

**Step 4: Commit**

```bash
git add templates/envs/mainnet/arbexecution/ templates/envs/sepolia/arbexecution/
git commit -S -m "feat(arbitrum): per-L1 env templates for arbexecution"
```

---

## Task 12: `templates/envs/{mainnet,sepolia}/arbitrum/nitro.tmpl`

**Files:**
- Create: `templates/envs/mainnet/arbitrum/nitro.tmpl`
- Create: `templates/envs/sepolia/arbitrum/nitro.tmpl`

**Step 1: Both env templates (identical structure, content driven by GenData)**

```
{{/* nitro.tmpl */}}
{{ define "arbitrum" }}
# --- Arbitrum L2 Rollup-Node - Nitro - configuration ---
ARB_IMAGE_VERSION={{.ArbImage}}
ARB_DATA_DIR={{.ArbDataDir}}
ARB_SNAPSHOT_URL={{.ArbSnapshotURL}}
PARENT_CHAIN_RPC_URL={{.ParentChainRPCURL}}
PARENT_CHAIN_BEACON_URL={{.ParentChainBeacon}}
{{ end }}
```

**Step 2: Commit**

```bash
git add templates/envs/mainnet/arbitrum/ templates/envs/sepolia/arbitrum/
git commit -S -m "feat(arbitrum): per-L1 env templates for nitro"
```

---

## Task 13: `cli/sub_gen.go` — `ArbFullNodeSubCmd`

**Files:**
- Modify: `cli/sub_gen.go`

**Step 1: Add the subcommand**

Reference: `OpFullNodeSubCmd` (around line 135 of `cli/sub_gen.go`). Copy-rename-adapt.

Insert after `OpFullNodeSubCmd` (or alongside the other L2 subcommands):

```go
func ArbFullNodeSubCmd(sedgeAction actions.SedgeActions) *cobra.Command {
    var flags GenCmdFlags
    var arbChain, arbSnapshotURL, parentRPC, parentBeacon string

    cmd := &cobra.Command{
        Use:   "arb-full-node [flags]",
        Short: "Generate a full node config for Arbitrum",
        Long: `Generate a docker-compose and an environment file for an Arbitrum full node.

By default this bundles an Ethereum L1 stack (execution + consensus) plus the L2 stack
(nethermind-arbitrum + nitro). Provide all four of --execution-api-url, --consensus-url,
--parent-chain-rpc-url, --parent-chain-beacon-url to skip the L1 services and target
external L1 endpoints instead.

You can select the Arbitrum chain via --chain (arbitrum-one, arbitrum-sepolia). If
omitted, the chain is inferred from -n (mainnet -> arbitrum-one, sepolia -> arbitrum-sepolia).`,
        Args: cobra.NoArgs,
        PreRunE: func(cmd *cobra.Command, args []string) error {
            if err := validateCustomNetwork(&flags.CustomFlags, network); err != nil {
                return err
            }
            resolvedNet, resolvedChain, err := configs.ResolveArbitrumChainAndL1(network, arbChain)
            if err != nil {
                return err
            }
            network = resolvedNet
            arbChain = resolvedChain
            // External-L1 mode validation: all four URLs must be set, or none.
            externalCount := 0
            for _, u := range []string{flags.executionApiUrl, flags.consensusApiUrl, parentRPC, parentBeacon} {
                if u != "" {
                    externalCount++
                }
            }
            if externalCount != 0 && externalCount != 4 {
                return fmt.Errorf("external-L1 mode requires all of --execution-api-url, --consensus-url, --parent-chain-rpc-url, --parent-chain-beacon-url (got %d/4 set)", externalCount)
            }
            return preValidationGenerateCmd(network, logging, &flags)
        },
        RunE: func(cmd *cobra.Command, args []string) error {
            services := []string{arbitrum, arbExecution}
            if flags.executionApiUrl == "" {
                services = append([]string{execution, consensus}, services...)
            }
            // Stash arb-specific values onto flags for runGenCmd to pick up.
            // runGenCmd is shared with op/taiko/surge — extend GenCmdFlags with the new fields
            // and have runGenCmd populate GenData.ArbChain / ArbSnapshotURL / ParentChainRPCURL /
            // ParentChainBeacon from them. See Task 14.
            flags.arbChain = arbChain
            flags.arbSnapshotURL = arbSnapshotURL
            flags.parentChainRPC = parentRPC
            flags.parentChainBeacon = parentBeacon
            return runGenCmd(cmd.OutOrStdout(), &flags, sedgeAction, services)
        },
    }

    // L1 flags (reused)
    cmd.Flags().StringVarP(&flags.consensusName, "consensus", "c", "", "L1 consensus client (lighthouse|lodestar|teku|prysm|nimbus). Use 'name:image' to override.")
    cmd.Flags().StringVarP(&flags.executionName, "execution", "e", "", "L1 execution client (geth|nethermind|besu|erigon). Use 'name:image' to override.")
    cmd.Flags().StringVar(&flags.checkpointSyncUrl, "checkpoint-sync-url", "", "L1 checkpoint sync URL.")
    cmd.Flags().StringArrayVar(&flags.elExtraFlags, "el-extra-flag", []string{}, "Extra flags for L1 EL.")
    cmd.Flags().StringArrayVar(&flags.clExtraFlags, "cl-extra-flag", []string{}, "Extra flags for L1 CL.")
    cmd.Flags().StringSliceVar(&flags.fallbackEL, "fallback-execution-urls", []string{}, "Fallback EL URLs for the L1 CL.")
    cmd.Flags().StringVar(&flags.executionApiUrl, "execution-api-url", "", "External L1 EL RPC URL (skips bundled L1 EL).")
    cmd.Flags().StringVar(&flags.consensusApiUrl, "consensus-url", "", "External L1 CL beacon URL (skips bundled L1 CL).")
    cmd.Flags().StringVar(&flags.feeRecipient, "fee-recipient", "", "L1 validator fee recipient.")
    cmd.Flags().StringVar(&flags.jwtPath, "jwt-secret-path", "", "Path to the JWT secret file.")
    cmd.Flags().BoolVar(&flags.mapAllPorts, "map-all", false, "Map all client ports to the host.")
    cmd.Flags().BoolVar(&flags.latestVersion, "latest", false, "Use :latest tags for L1 clients (L2 clients stay pinned).")

    // Arbitrum-specific flags
    cmd.Flags().StringVar(&arbChain, "chain", "", "Arbitrum chain (arbitrum-one|arbitrum-sepolia). Inferred from -n if omitted.")
    cmd.Flags().StringVar(&flags.arbExecutionName, "arb-execution", "", "Arbitrum EL client image override (default nethermind-arbitrum).")
    cmd.Flags().StringVar(&flags.arbitrumName, "arb-image", "", "Nitro image override.")
    cmd.Flags().StringArrayVar(&flags.arbExtraFlags, "arb-extra-flag", []string{}, "Extra flags for nitro.")
    cmd.Flags().StringArrayVar(&flags.arbElExtraFlags, "arb-el-extra-flag", []string{}, "Extra flags for nethermind-arbitrum.")
    cmd.Flags().StringVar(&arbSnapshotURL, "snapshot-url", "", "Nitro --init.url (when empty, --init.empty=true is used).")
    cmd.Flags().StringVar(&parentRPC, "parent-chain-rpc-url", "", "External L1 RPC URL for nitro (required in external-L1 mode).")
    cmd.Flags().StringVar(&parentBeacon, "parent-chain-beacon-url", "", "External L1 beacon URL for nitro (required in external-L1 mode).")

    cmd.Flags().SortFlags = false
    return cmd
}
```

**Step 2: Extend `GenCmdFlags`**

In `cli/generate.go` add to `GenCmdFlags`:

```go
arbExecutionName  string
arbitrumName      string
arbChain          string
arbSnapshotURL    string
parentChainRPC    string
parentChainBeacon string
arbExtraFlags     []string
arbElExtraFlags   []string
```

And the corresponding accessors on `*GenCmdFlags` if the `ClientFlags` interface in `cli/factory/flags.go` requires them (mirror `GetOptimismName`, `GetL2ExecutionName`).

**Step 3: Have `runGenCmd` populate `GenData`**

Find `runGenCmd` (likely in `cli/generate.go`). Where it constructs `GenData{}`, add the new fields driven by `flags.arbChain`, `flags.arbSnapshotURL`, `flags.parentChainRPC`, `flags.parentChainBeacon`, `flags.arbExtraFlags`, `flags.arbElExtraFlags`.

**Step 4: Build**

```bash
go build ./...
```

Expected: PASS.

**Step 5: Commit**

```bash
git add cli/sub_gen.go cli/generate.go cli/factory/flags.go
git commit -S -m "feat(arbitrum): arb-full-node subcommand and flags"
```

---

## Task 14: Register the subcommand on `GenerateCmd`

**Files:**
- Modify: `cli/generate.go` (the `GenerateCmd` function around line 169)

**Step 1: Add the registration**

In `GenerateCmd`, after `cmd.AddCommand(SurgeFullNodeSubCmd(sedgeAction))`:

```go
cmd.AddCommand(ArbFullNodeSubCmd(sedgeAction))
```

**Step 2: Build and run help**

```bash
make compile-sedge
./build/sedge generate arb-full-node --help 2>&1 | head -20
```

Expected: shows the new subcommand with all flags.

**Step 3: Commit**

```bash
git add cli/generate.go
git commit -S -m "feat(arbitrum): wire arb-full-node into GenerateCmd"
```

---

## Task 15: Offline generation smoke test

**Files:** none modified (verification only)

**Step 1: Generate a sepolia stack offline (placeholder URLs)**

```bash
rm -rf /tmp/sedge-arb-sepolia
./build/sedge generate arb-full-node -n sepolia \
  --execution-api-url=http://placeholder:8545 \
  --consensus-url=http://placeholder:5052 \
  --parent-chain-rpc-url=http://placeholder:8545 \
  --parent-chain-beacon-url=http://placeholder:5052 \
  -e nethermind \
  -p /tmp/sedge-arb-sepolia
```

Expected: `Generation of files successfully, happy staking!`

**Step 2: Inspect the generated artifacts**

```bash
cat /tmp/sedge-arb-sepolia/.env
cat /tmp/sedge-arb-sepolia/docker-compose.yml
```

Expected in `.env`:
- `EC_L2_IMAGE_VERSION=nethermind/nethermind-arbitrum:0.1.0-alpha`
- `ARB_IMAGE_VERSION=offchainlabs/nitro-node:v3.10.0-rc.2-746bda2`
- `ARB_CHAINSPEC=arbitrum-sepolia`
- `PARENT_CHAIN_RPC_URL=http://placeholder:8545`

Expected in `docker-compose.yml`:
- Two services: `arbexecution`, `arbitrum-init`, `arbitrum` — no `execution` / `consensus` (external-L1 mode).
- `arbexecution` has the healthcheck.
- `arbitrum-init` `depends_on: arbexecution: condition: service_healthy`.
- `arbitrum` `depends_on: arbitrum-init: condition: service_completed_successfully`.
- nitro command lines include `--chain.id=421614` and `--init.empty=true` (since no `--snapshot-url` was passed).

**Step 3: Try the chain inference paths**

```bash
./build/sedge generate arb-full-node --chain arbitrum-one \
  --execution-api-url=http://placeholder:8545 --consensus-url=http://placeholder:5052 \
  --parent-chain-rpc-url=http://placeholder:8545 --parent-chain-beacon-url=http://placeholder:5052 \
  -p /tmp/sedge-arb-one
grep "^NETWORK\|^ARB_CHAINSPEC\|^CHAIN_ID" /tmp/sedge-arb-one/.env
```

Expected: `ARB_CHAINSPEC=arbitrum-mainnet`, chain id 42161 visible in the docker-compose command line.

Try the mismatch case:

```bash
./build/sedge generate arb-full-node -n mainnet --chain arbitrum-sepolia 2>&1 | tail -3
```

Expected: error mentioning `arbitrum chain does not match selected L1 network`.

**Step 4: Run the test suite again**

```bash
make test-no-e2e
```

Expected: PASS. If any existing test regressed (likely candidates: `internal/pkg/generate/generate_scripts_test.go` table-driven generator tests, `cli/cli_test.go`), fix before moving on.

**Step 5: Commit if anything was patched during the offline smoke test**

```bash
# only if changes were needed
git add -A
git commit -S -m "fix(arbitrum): <whatever was tweaked>"
```

---

## Task 16: End-to-end sync verification

**Files:** none modified (verification only)

**Prerequisite:** Marcos provides:
- L1 EL RPC URL (sepolia)
- L1 beacon URL (sepolia)
Both reachable from this host's docker network.

**Step 1: Generate against real L1 endpoints**

```bash
L1_RPC=<provided by Marcos>
L1_BEACON=<provided by Marcos>
rm -rf /tmp/sedge-arb-e2e
./build/sedge generate arb-full-node -n sepolia \
  --execution-api-url="$L1_RPC" \
  --consensus-url="$L1_BEACON" \
  --parent-chain-rpc-url="$L1_RPC" \
  --parent-chain-beacon-url="$L1_BEACON" \
  -p /tmp/sedge-arb-e2e
```

**Step 2: Start the stack**

```bash
./build/sedge run -p /tmp/sedge-arb-e2e
```

**Step 3: Watch the init service complete**

```bash
docker logs -f sedge-arbitrum-init 2>&1 | head -50
```

Expected: nitro logs `--init.empty=true`, sets up empty state, exits with code 0. The `sedge-arbitrum` container should then start.

**Step 4: Watch the main nitro service**

```bash
docker logs -f sedge-arbitrum 2>&1 | sed 's/\x1b\[[0-9;]*m//g' | grep -iE "Created block|peer|error|sync" | head -40
```

Expected: nitro connects to `arbexecution`, fetches batches from the L1 RPC, and starts producing log lines like `Created block` / chain head advancing. No `Exception` / `connection refused`.

**Step 5: Watch nethermind-arbitrum**

```bash
docker logs sedge-arbexecution-client 2>&1 | sed 's/\x1b\[[0-9;]*m//g' | grep -iE "loading configuration|engine|block|exception" | head -20
```

Expected: `Loading configuration from /nethermind/configs/arbitrum-sepolia.json`, engine-API requests landing from nitro, blocks being applied.

**Step 6: Shutdown**

```bash
./build/sedge down -p /tmp/sedge-arb-e2e
```

**Step 7: Final commit**

If any template tweaks were needed during the smoke test:

```bash
git add -A
git commit -S -m "fix(arbitrum): smoke-test adjustments"
```

Otherwise nothing to commit — the branch is ready for review.

---

## Done state

After Task 16, on `add-arbitrum-chain`:

- `make test-no-e2e` is green.
- `sedge generate arb-full-node -n sepolia` (with external L1 flags) produces a working stack that syncs arbitrum-sepolia.
- All commits are signed.
- No docs/README/CHANGELOG churn (deferred — internal branch).

The branch is *not* pushed yet (per maintainer's "keep the branch local for now").
