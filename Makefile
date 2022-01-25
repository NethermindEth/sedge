.DEFAULT_GOAL 	:= help

compile: ## compile:
	@mkdir -p bin
	@go build -o build/1click cmd/1Click/main.go

run: ## run
	@./build/1click

run-cli: compile ## run randomized cli
	@./build/1click cli -r --config ./config.yaml

listClients: compile ## Run cmd listClients with ./config.yaml
	@./build/1click listClients --config ./config.yaml

all: compile run ## build and run

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

gomod_tidy:
	 go mod tidy

gofmt:
	go fmt -x ./...
