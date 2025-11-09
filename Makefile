# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Project variables
DOCKER_IMAGE = gomatic/modern-go-application
BUILD_DIR ?= bin
VERSION ?= $(shell git describe --tags --exact-match 2>/dev/null || git symbolic-ref -q --short HEAD)
COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE ?= $(shell date +%FT%T%z)
LDFLAGS = -X main.version=${VERSION} -X main.commitHash=${COMMIT_HASH} -X main.buildDate=${BUILD_DATE}

export CGO_ENABLED ?= 0

## Development Environment

.PHONY: up
up: docker-compose.override.yml config.up ## Set up the development environment
	docker-compose up -d

.PHONY: down
down: ## Destroy the development environment
	docker-compose down --volumes --remove-orphans --rmi local

.PHONY: start
start: docker-compose.override.yml ## Start docker development environment
	docker-compose up -d

.PHONY: stop
stop: ## Stop docker development environment
	docker-compose stop

docker-compose.override.yml:
	cp docker-compose.override.yml.dist $@

config.up: config/template.up config/deployment.up ## Generate development configuration
	go tool up template process -i config/deployment.up -o $@

.PHONY: config-production
config-production: config/template.up config/production.up ## Generate production configuration
	go tool up template process -i config/production.up -o config.production.out.up

## Build

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
DIST_BINARY = dist/mga-${GOOS}-${GOARCH}$(if $(filter windows,${GOOS}),.exe,)
BIN_TARGET = ${BUILD_DIR}/mga-${GOOS}-${GOARCH}$(if $(filter windows,${GOOS}),.exe,)

${BUILD_DIR}:
	@mkdir -p $@

${DIST_BINARY}:
	go tool goreleaser build --single-target --snapshot --clean

${BIN_TARGET}: ${BUILD_DIR} ${DIST_BINARY}
	cp ${DIST_BINARY} $@

.PHONY: build
build: ${BIN_TARGET} ## Build binary for current platform only

.PHONY: build-all
build-all: ## Build binaries for all platforms
	go tool goreleaser build --snapshot --clean

.PHONY: release
release: ## Create a release with goreleaser
	go tool goreleaser release --clean

.PHONY: release-snapshot
release-snapshot: ## Create a snapshot release (no git tag required)
	go tool goreleaser release --snapshot --clean

.PHONY: clean
clean: ## Clean builds
	rm -rf ${BUILD_DIR}/mga* dist/

## Test

.PHONY: test
test: ## Run tests
	go tool gotestsum --format short -- ./...

.PHONY: test-verbose
test-verbose: ## Run tests with verbose output
	go tool gotestsum --format short-verbose -- -v ./...

.PHONY: test-integration
test-integration: ## Run integration tests
	go test -run ^TestIntegration ./...

.PHONY: test-functional
test-functional: ## Run functional tests
	go test -run ^TestFunctional ./...

## Code Quality

.PHONY: lint
lint: ## Run linter
	GOFLAGS="-buildvcs=false" go tool golangci-lint run

.PHONY: check
check: lint test ## Run tests and linters

## Code Generation

.PHONY: generate
generate: ## Generate code
	go generate ./...

.PHONY: graphql
graphql: ## Generate GraphQL code
	go tool gqlgen

.PHONY: proto
proto: ## Generate protobuf code (requires buf: https://buf.build/docs/installation)
	buf generate

## Docker

.PHONY: docker
docker: ## Build a Docker image
	docker build -t ${DOCKER_IMAGE}:${VERSION} .

.PHONY: docker-debug
docker-debug: ## Build a Docker image with remote debugging capabilities
	docker build -t ${DOCKER_IMAGE}:${VERSION}-debug --build-arg BUILD_TARGET=debug .

## Utilities

.PHONY: tidy
tidy: ## Tidy and verify dependencies
	go mod tidy
	go mod verify

.PHONY: fmt
fmt: ## Format code
	go tool gofumpt -l -w .

.PHONY: vet
vet: ## Run go vet
	go vet ./...

# Variable outputting/exporting rules
var-%: ; @echo $($*)
varexport-%: ; @echo $*=$($*)
