
.PHONY: all

all: help

## Build:
tidy: ## Tidy project
	go mod tidy

clean: ## Cleans temporary folder
	rm -rf mock

build: tidy clean gen-mock ## Builds project
	go build ./...

## Generations:
gen-mock: ## Generates mock objects
	go generate ./...

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  make <target>'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    %-20s%s\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  %s\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
