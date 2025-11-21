.PHONY: help generate start test down clean

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

generate: ## Generate API code from OpenAPI spec
	@echo "Generating API code from openapi.yaml..."
	@oapi-codegen -config oapi-codegen.yaml openapi.yaml
	@echo "✓ Code generated successfully"

start: ## Start everything (databases + server)
	@echo "Starting databases..."
	@docker-compose up -d
	@echo "Waiting for databases to be ready..."
	@sleep 3
	@echo "✓ Databases ready"
	@echo ""
	@echo "Starting server..."
	@export DATABASE_URL="postgres://developer:devpassword@localhost:5432/mini_arke_db?sslmode=disable" && \
	go run cmd/server/main.go

test: ## Run tests
	@echo "Ensuring test database is running..."
	@docker-compose up -d postgres-test
	@sleep 2
	@echo "Running tests..."
	@go test -v ./internal/api/...
	@echo "✓ Tests completed"

down: ## Stop all databases
	@echo "Stopping databases..."
	@docker-compose down
	@echo "✓ Databases stopped"

clean: down ## Stop databases and clean up
	@echo "Cleaning up..."
	@docker-compose down -v
	@echo "✓ Cleanup complete"