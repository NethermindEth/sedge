.DEFAULT_GOAL 	:= help

compile:: ## compile:
	@mkdir -p bin
	@go build -o build/1click cmd/1Click/main.go

run: ## run
	@./build/1click

all: compile run ## build and run

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'