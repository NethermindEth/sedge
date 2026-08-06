# AGENTS Instructions

These guidelines apply to the entire repository.

## Environment setup
- Requires **Go 1.22+** (`go version`).
- Download Go modules (once per write session):
  ```bash
  go mod download
  ```
- Install project tools for formatting and mock generation. The simplest approach is:
  ```bash
  make install-deps
  ```
  This installs `gofumpt`, `mockgen` and `abigen`. Building `abigen` compiles
  `go-ethereum` and can take several minutes. In constrained environments, install
  only what unit tests need:
  ```bash
  go install mvdan.cc/gofumpt@latest
  go install github.com/golang/mock/mockgen@v1.6.0
  ```

## Generating mocks and other code
- Generated mocks live in git-ignored folders (e.g. `mocks/`,
  `internal/monitoring/mocks`). Refresh them before running tests or when
  interfaces change. To regenerate mocks without compiling heavy contract bindings,
  run:
  ```bash
  go generate ./internal/compose \
              ./internal/pkg/services \
              ./internal/pkg/dependencies \
              ./internal/monitoring \
              ./internal/monitoring/locker \
              ./internal/ui \
              ./cli \
              ./cli/actions
  ```
- `make test-no-e2e` runs these generation steps automatically but also triggers
  contract binding generation with `abigen`. Use it when verifying the full project
  or when bindings are intentionally changed.

### Quick write-environment workflow
1. Run the `go generate` command above after any interface changes.
2. Execute tests for the packages you touched:
   ```bash
   go test ./path/to/pkg -count=1
   ```
3. When feasible, run the whole suite while skipping slow E2E tests:
   ```bash
   go test ./... -count=1 -timeout 25m -skip TestE2E
   ```
   Some packages (notably `cli` and `internal/lido/contracts`) depend on Docker or
   network services and may fail in constrained environments.

## Testing
- **Full unit test run:**
  ```bash
  make test-no-e2e
  ```
  Invokes the generation steps above (including `abigen`) and runs `go test` with
  coverage while skipping E2E tests.
- **Targeted package tests:** when full generation or Docker services are
  unavailable, test only the affected packages:
  ```bash
  go test ./configs -count=1
  ```
- **End-to-end tests:**
  ```bash
  make e2e-test
  ```
  These are slow and require Docker.

## Formatting and style
- Format code with `make format` and verify with `make format-check`.
- Follow [Conventional Commits](https://www.conventionalcommits.org/) for commit
  messages.
- Update the `CHANGELOG.md` under the **Unreleased** section for any
  user-facing changes.
