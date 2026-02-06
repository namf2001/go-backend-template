.PHONY: help run build test clean docker-up docker-down migrate-up migrate-down

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: ## Run the application
	@go run ./cmd/server

build: ## Build the application
	@go build -o bin/api ./cmd/server

test: ## Run tests
	@go test ./...

clean: ## Clean the build
	@rm -rf bin

docker-up: ## Start the database
	@docker-compose up -d

docker-down: ## Stop the database
	@docker-compose down

migrate-up: ## Run database migrations
	@echo "Running migrations..."
	psql -h localhost -U postgres -d go_backend_db -f migrations/001_create_users_table.sql

migrate-down: ## Rollback database migrations
	@echo "Rolling back migrations..."
	psql -h localhost -U postgres -d go_backend_db -c "DROP TABLE IF EXISTS users;"

deps: ## Download dependencies
	go mod download
	go mod tidy

.DEFAULT_GOAL := help

