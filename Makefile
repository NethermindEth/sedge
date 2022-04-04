.DEFAULT_GOAL 	:= help
.PHONY: compile run run-cli test coverage clients logs all gomod_tidy go_fmt help

compile: ## compile:
	@mkdir -p build
	@go build -o build/1click cmd/1click/main.go

run: ## run
	@./build/1click

run-cli: compile ## run randomized cli
	@./build/1click cli -r --config ./config.yaml

test: ## run tests
	@go test ./... -coverprofile cover.out

coverage: cover.out ## show tests coverage
	@go tool cover -func cover.out

clients: compile ## Run cmd clients with ./config.yaml
	@./build/1click clients --config ./config.yaml

logs: compile ## run cmd logs with ./config.yaml
	@./build/1click logs --config ./config.yaml

all: compile run ## build and run

gomod_tidy:
	 go mod tidy

gofmt:
	go fmt -x ./...

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'