.DEFAULT_GOAL 	:= help
.PHONY: compile run run-cli test coverage clients logs all gomod_tidy go_fmt help

# Variables
SEDGE_VERSION = $(shell git tag | sort | tail -n 1)
LDFLAGS=-X github.com/NethermindEth/sedge/internal/utils.Version="${SEDGE_VERSION}"

# Commands
compile: install-mockgen install-abigen generate ## compile:
	@mkdir -p build
	@go build -ldflags "${LDFLAGS}" -o build/sedge cmd/sedge/main.go
	@go build -ldflags "${LDFLAGS}" -o build/lido-exporter cmd/lido-exporter/main.go

compile-linux: generate ## compile:
	@mkdir -p build
	@env GOOS=linux go build -ldflags="${LDFLAGS[*]}" -o build/sedge cmd/sedge/main.go
	@env GOOS=linux go build -ldflags="${LDFLAGS[*]}" -o build/lido-exporter cmd/lido-exporter/main.go
	
compile-sedge:
	@mkdir -p build
	@go build -ldflags "${LDFLAGS}" -o build/sedge cmd/sedge/main.go

compile-lido-exporter:
	@mkdir -p build
	@go build -ldflags "${LDFLAGS}" -o build/lido-exporter cmd/lido-exporter/main.go

compile-sedge-linux:
	@mkdir -p build
	@env GOOS=linux go build -ldflags="${LDFLAGS[*]}" -o build/sedge cmd/sedge/main.go

compile-lido-exporter-linux:
	@mkdir -p build
	@env GOOS=linux go build -ldflags="${LDFLAGS[*]}" -o build/lido-exporter cmd/lido-exporter/main.go

install: compile ## compile the binary and copy it to PATH
	@sudo cp build/sedge /usr/local/bin

run: ## run
	@./build/sedge

run-cli: compile ## run cli
	@./build/sedge cli --config ./config.yaml

generate: ## generate go files
	@abigen --abi ./internal/lido/contracts/csmodule/CSModule.abi --bin ./internal/lido/contracts/csmodule/CSModule.bin --pkg csmodule --out ./internal/lido/contracts/csmodule/CSModule.go
	@abigen --abi ./internal/lido/contracts/csfeedistributor/CSFeeDistributor.abi --bin ./internal/lido/contracts/csfeedistributor/CSFeeDistributor.bin --pkg csfeedistributor --out ./internal/lido/contracts/csfeedistributor/CSFeeDistributor.go
	@abigen --abi ./internal/lido/contracts/csaccounting/CSAccounting.abi --bin ./internal/lido/contracts/csaccounting/CSAccounting.bin --pkg csaccounting --out ./internal/lido/contracts/csaccounting/CSAccounting.go
	@abigen --abi ./internal/lido/contracts/mevboostrelaylist/MEVBoostRelayAllowedList.abi --bin ./internal/lido/contracts/mevboostrelaylist/MEVBoostRelayAllowedList.bin --pkg mevboostrelaylist --out ./internal/lido/contracts/mevboostrelaylist/MEVBoostRelayAllowedList.go
	@abigen --abi ./internal/lido/contracts/vebo/VEBO.abi --bin ./internal/lido/contracts/vebo/VEBO.bin --pkg vebo --out ./internal/lido/contracts/vebo/VEBO.go
	@go generate ./...

test: generate ## run tests
	@mkdir -p coverage
	@go test -coverprofile=coverage/coverage.out -covermode=count -timeout 25m ./...

e2e-test: generate ## Run e2e tests
	@go test -timeout 25m -count=1 ./e2e/sedge/...

e2e-test-windows: generate ## Run e2e tests on Windows
	@go test -timeout 25m -count=1 -skip TestE2E_MonitoringStack ./e2e/sedge/...

test-no-e2e: generate ## run tests excluding e2e
	@mkdir -p coverage
	@go test -coverprofile=coverage/coverage.out -covermode=count ./... -skip TestE2E

codecov-test: generate ## unit tests with coverage using the courtney tool
	@mkdir -p coverage
	@go test -coverprofile=coverage/coverage.out -covermode=count -timeout 25m ./... -skip TestE2E
	@go tool cover -html=coverage/coverage.out -o coverage/coverage.html

install-gofumpt: ## install gofumpt
	go install mvdan.cc/gofumpt@latest

install-mockgen: ## install mockgen
	go install github.com/golang/mock/mockgen@v1.6.0 

install-abigen: ## install abigen
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest

install-deps: | install-gofumpt install-mockgen install-abigen ## Install some project dependencies

coverage: ## show tests coverage
	@go tool cover -html=coverage/coverage.out -o coverage/coverage.html

clients: compile ## Run cmd clients with ./config.yaml
	@./build/sedge clients --config ./config.yaml

logs: compile ## run cmd logs with ./config.yaml
	@./build/sedge logs --config ./config.yaml

all: compile run ## build and run

gomod_tidy: ## go mod tidy
	 go mod tidy

format: ## run code formatting
	gofumpt -l -w .

# assert `gofumpt -l` produces no output
format-check: SHELL:=/bin/bash
format-check: ## check formatting
	test -z "$$(gofumpt -l . | tee >(cat 1>&2))"


help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
