# Makefile for zapfluent
#
# This Makefile provides a set of common commands to build, test, and lint the project.

# Ensure that go-installed binaries are available in the PATH
export PATH := $(shell go env GOPATH)/bin:$(PATH)

# Define the binary name for the linter
GOLANGCI_LINT := golangci-lint

# ==============================================================================
# Help Target
# ==============================================================================

.PHONY: help
help: ## ✨ Show this help message
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# ==============================================================================
# Development Targets
# ==============================================================================

.PHONY: build
build: ## 🔨 Build the project
	@echo ">> building project..."
	@go build ./...

.PHONY: test
test: ## 🧪 Run all tests
	@echo ">> running tests..."
	@go test ./...

.PHONY: lint
lint: tools ## 🔍 Run linter
	@echo ">> running linter..."
	@$(GOLANGCI_LINT) run ./...

# ==============================================================================
# Tooling Targets
# ==============================================================================

.PHONY: tools
tools: $(GOLANGCI_LINT) ## 🛠️ Install development tools

$(GOLANGCI_LINT):
	@echo ">> checking for $(GOLANGCI_LINT)..."
	@command -v $(GOLANGCI_LINT) >/dev/null 2>&1 || \
		(echo "   -> $(GOLANGCI_LINT) not found, installing..." && \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)

# ==============================================================================
# Housekeeping Targets
# ==============================================================================

.PHONY: clean
clean: ## 🧹 Clean build artifacts
	@echo ">> cleaning up..."
	@# This project is a library, so there are no build artifacts to clean by default.
	@# This target is here for convention.

# Set the default target to 'help'
.DEFAULT_GOAL := help
