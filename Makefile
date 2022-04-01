.DEFAULT_GOAL := help

# Golang vars.
GOPRIVATE := github.com/flyingdice
GOPATH ?= $(shell go env GOPATH)
GOBIN ?= $(GOPATH)/bin

# Project vars.
VERSION ?= $(shell git describe --tags)

.PHONY: build
build: ## todo
	@env $$(cat .env) go build ./...

.PHONY: lint
lint: ## Run go linters.
	@golangci-lint run

.PHONY: test
test: ## Run test suite.
	@go test -count=1 ./...

.phony: modules
modules: ## Tidy up and vendor go modules.
	@GOPRIVATE=$(GOPRIVATE) go mod tidy
	@GOPRIVATE=$(GOPRIVATE) go mod vendor

.phony: help
help: ## Print Makefile usage.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
