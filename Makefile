.DEFAULT_GOAL 	:= help
.PHONY: compile run run-cli test coverage clients logs all gomod_tidy go_fmt help

compile: ## compile:
	@mkdir -p build
	@go build -o build/sedge cmd/main.go

compile-linux: ## compile:
	@mkdir -p build
	@env GOOS=linux go build -o build/sedge cmd/main.go

run: ## run
	@./build/sedge

run-cli: compile ## run cli
	@./build/sedge cli --config ./config.yaml

test: ## run tests
	@mkdir -p coverage
	@go test -coverprofile=coverage/coverage.out -covermode=count ./...

codecov-test: ## unit tests with coverage using the courtney tool
	@mkdir -p coverage
	@courtney/courtney -v -o coverage/coverage.out ./...

install-deps: ## Install some project dependencies
	@git clone https://github.com/stdevMac/courtney
	@(cd courtney && go get  ./... && go build courtney.go)
	@go get ./...

coverage: coverage/coverage.out ## show tests coverage
	@go tool cover -html=coverage/coverage.out -o coverage/coverage.html

clients: compile ## Run cmd clients with ./config.yaml
	@./build/sedge clients --config ./config.yaml

logs: compile ## run cmd logs with ./config.yaml
	@./build/sedge logs --config ./config.yaml

all: compile run ## build and run

gomod_tidy: ## go mod tidy
	 go mod tidy

gofmt: ## go fmt
	go fmt -x ./...

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'