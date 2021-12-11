.DEFAULT_GOAL := help

# Go vars.
GOPRIVATE := github.com/flyingdice
GOPATH ?= $(shell go env GOPATH)
GOBIN ?= $(GOPATH)/bin

# Project vars.
VERSION ?= $(shell git describe --tags)

.PHONY: modules
modules: ## Tidy up and vendor go modules.
	@go mod tidy
	@go mod vendor

.PHONY: help
help: ## Print Makefile usage.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
