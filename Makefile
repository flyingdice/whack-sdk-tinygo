.DEFAULT_GOAL := help

# Go vars.
GOPRIVATE := github.com/flyingdice
GOPATH    ?= $(shell go env GOPATH)
GOBIN     ?= $(GOPATH)/bin
GOROOT    ?= $(shell tinygo info | grep GOROOT | cut -d ' ' -f 7)

# Project vars.
VERSION  ?= $(shell git describe --tags)
ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: build
build: $(wildcard examples/*)
	@mkdir -p $(ROOT_DIR)/build/$^ && \
	cd $^ && \
	tinygo build \
		-target wasm \
		-gc conservative \
		-no-debug \
		-wasm-abi generic \
		-o $(ROOT_DIR)/build/$^/$(shell basename $^).wasm \
		.

.PHONY: modules
modules: ## Tidy up and vendor go modules.
	@GOPRIVATE=$(GOPRIVATE) go mod tidy
	@GOPRIVATE=$(GOPRIVATE) go mod vendor

.PHONY: help
help: ## Print Makefile usage.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
