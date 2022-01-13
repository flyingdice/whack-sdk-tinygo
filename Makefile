.DEFAULT_GOAL := help

# Go vars.
GOPRIVATE := github.com/flyingdice
GOPATH    ?= $(shell go env GOPATH)
GOBIN     ?= $(GOPATH)/bin
GOROOT    ?= $(shell tinygo info | grep GOROOT | cut -d ' ' -f 7)

# Project vars.
VERSION  ?= $(shell git describe --tags)
ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: clean
clean: $(wildcard examples/*) ## Build all examples.
	@for e in $^; do \
		rm -rf $(ROOT_DIR)/build/$$e/$$(basename $$e).wasm ; \
	done;

.PHONY: build
build: $(wildcard examples/*) ## Build all examples.
	@for e in $^ ; do \
		mkdir -p $(ROOT_DIR)/build/$$e ; \
		cd $$e ; \
		tinygo build -target wasi -o $(ROOT_DIR)/build/$$e/$$(basename $$e).wasm ; \
		cd $(ROOT_DIR) ; \
	done; \

.PHONY: lint
lint: ## Run golang linter.
	@golangci-lint run

.PHONY: modules
modules: ## Tidy up and vendor go modules.
	@GOPRIVATE=$(GOPRIVATE) go mod tidy
	@GOPRIVATE=$(GOPRIVATE) go mod vendor

.PHONY: help
help: ## Print Makefile usage.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
