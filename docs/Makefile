.DEFAULT_GOAL 	:= help
.PHONY: docs-start docs-static docs-serve docs-publish

docs-start: ## Starts the docs's development server
	npm start

docs-static: ## Bundles your docs website into static files for production
	npm run build

docs-serve: ## Serves the built website locally
	npm run serve

docs-publish: ## Publishes the website to GitHub pages
	npm deploy

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'