SHELL := /bin/bash

.DEFAULT_GOAL := help

GO ?= go
BIN_DIR ?= bin

WEB_DIR ?= web/app
WEB_PM ?= npm

TAG ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo 1)
IMAGE_REGISTRY ?= rayyildiz.azurecr.io
IMAGE_NAME ?= capturetweet-server
IMAGE ?= $(IMAGE_REGISTRY)/$(IMAGE_NAME)

DOCKER_PLATFORM ?= linux/amd64
COMPRESS ?= 0

.PHONY: help generate fmt vet test test-cover build build-server run clean \
	web-install web-start web-build web-test web-codegen \
	docker-prep-backend docker-build-backend docker-push-backend docker-run-backend docker-backend

help: ## Show available targets
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z0-9_\-]+:.*##/ {printf "\033[36m%-24s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

generate: ## Run Go code generation (gqlgen, mocks, wire, ...)
	$(GO) generate ./...

fmt: ## Format Go code
	$(GO) fmt ./...

vet: ## Run `go vet`
	$(GO) vet ./...

test: ## Run Go tests
	$(GO) test ./...

test-cover: ## Run Go tests with coverage output (CI-style)
	$(GO) test -cover -coverprofile=coverage.txt -covermode=atomic ./...

build: build-server ## Build binaries into `$(BIN_DIR)`

build-server: ## Build the API server binary
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/server ./cmd/server

run: ## Run the API server locally
	$(GO) run ./cmd/server

clean: ## Remove build artifacts
	rm -rf $(BIN_DIR) cmd/server/tmp

web-install: ## Install frontend dependencies
	cd $(WEB_DIR) && $(WEB_PM) install

web-start: ## Start the frontend dev server
	cd $(WEB_DIR) && $(WEB_PM) run start

web-build: ## Build the frontend
	cd $(WEB_DIR) && $(WEB_PM) run build

web-test: ## Run frontend tests
	cd $(WEB_DIR) && $(WEB_PM) run test

web-codegen: ## Generate frontend GraphQL types
	cd $(WEB_DIR) && $(WEB_PM) run codegen

docker-prep-backend: ## Build the linux binary into cmd/server/tmp/app (Docker build context)
	@mkdir -p cmd/server/tmp
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -trimpath -ldflags="-s -w" -o cmd/server/tmp/app cmd/server/main.go
	@if [[ "$(COMPRESS)" == "1" ]]; then \
		if command -v upx >/dev/null 2>&1; then \
			upx cmd/server/tmp/app; \
		else \
			echo "upx not found; skipping compression"; \
		fi; \
	fi

docker-build-backend: docker-prep-backend ## Build backend Docker image
	docker build --platform $(DOCKER_PLATFORM) -t $(IMAGE):$(TAG) cmd/server

docker-push-backend: ## Push backend Docker image
	docker push $(IMAGE):$(TAG)

docker-run-backend: ## Run backend Docker image locally (uses .env)
	docker run --rm -p 4000:4000 --env-file .env $(IMAGE):$(TAG)

docker-backend: docker-build-backend docker-push-backend ## Build + push backend Docker image (legacy target)
