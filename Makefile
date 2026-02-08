# Makefile
.PHONY: all build build-all run run-race clean clear test test-coverage deps lint fmt vet generate audit docker-build docker-up docker-down migrate-up migrate-down help

# Variables
BINARY_NAME=api
BUILD_DIR=bin
GO=go
GOFLAGS=-v

# Default target
all: audit build ## Run audit and build

build: ## Build the application
	@echo "Building..."
	$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/server

build-all: ## Build for multiple platforms
	@echo "Building for multiple platforms..."
	GOOS=linux GOARCH=amd64 $(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 ./cmd/server
	GOOS=darwin GOARCH=amd64 $(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 ./cmd/server
	GOOS=windows GOARCH=amd64 $(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe ./cmd/server

run: ## Run the application
	@go run ./cmd/server

run-race: ## Run the application with race detector
	@go run -race ./cmd/server

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out

clear: clean ## Clear build artifacts and Go caches
	@echo "Clearing Go caches..."
	$(GO) clean -cache -testcache -modcache

test: ## Run tests
	@echo "Running tests..."
	$(GO) test -v -race -coverprofile=coverage.out ./...

test-coverage: test ## Run tests with coverage report
	$(GO) tool cover -html=coverage.out

deps: ## Install dependencies
	@echo "Downloading dependencies..."
	$(GO) mod download
	$(GO) mod tidy

lint: ## Run linters (golangci-lint)
	@echo "Linting..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed"; \
		exit 1; \
	fi

fmt: ## Format code
	@echo "Formatting..."
	$(GO) fmt ./...
	@if command -v goimports >/dev/null 2>&1; then \
		goimports -w .; \
	fi

vet: ## Vet code
	@echo "Vetting..."
	$(GO) vet ./...

generate: ## Generate code
	@echo "Generating code..."
	$(GO) generate ./...

audit: fmt vet lint ## Run formatting, vetting, and linting

docker-build: ## Build docker image
	docker build -t $(BINARY_NAME):latest .

docker-up: ## Start the database
	@docker-compose up -d

docker-down: ## Stop the database
	@docker-compose down

migrate-up: ## Run database migrations
	@echo "Running migrations..."
	@for file in migrations/*.up.sql; do \
		echo "Applying $$file..."; \
		psql -h localhost -U postgres -d go_backend_db -f $$file; \
	done

migrate-down: ## Rollback database migrations
	@echo "Rolling back migrations..."
	@for file in migrations/*.down.sql; do \
		echo "Rolling back $$file..."; \
		psql -h localhost -U postgres -d go_backend_db -f $$file; \
	done

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
