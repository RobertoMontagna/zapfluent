# Makefile for zapfluent
#
# This Makefile provides a set of common commands to build, test, and lint the project.
export PATH := $(shell go env GOPATH)/bin:$(PATH)

# Define binary names
GOLANGCI_LINT := golangci-lint
GO_JUNIT_REPORT := go-junit-report

# Externalize the OS name check into a variable
OS_NAME := $(shell uname -s)

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

.PHONY: lint-fix
lint-fix: tools ## Fix all auto-fixable problems
	@echo ">> running linter and applying fixes..."
	@$(GOLANGCI_LINT) run --fix ./...

.PHONY: fmt
fmt: ## 🎨 Format all Go files
	@echo ">> formatting go files..."
	@go fmt ./...

.PHONY: check-fmt
check-fmt: ## 🧐 Check if all Go files are formatted
	@echo ">> checking go files formatting..."
	@if [ -n "$$(gofmt -l .)" ]; then \
		echo "The following files are not formatted:"; \
		gofmt -l .; \
		exit 1; \
	fi

.PHONY: coverage
coverage: ## 📊 Generate test coverage report
	@echo ">> generating coverage report..."
	@go test -coverprofile=coverage.out ./...

.PHONY: test-ci
test-ci: tools ## 📜 Generate reports for CI
	@echo ">> generating reports for CI..."
	@go test -v -coverprofile=coverage.out ./... 2>&1 > test_output.log
	@ls -la
	@cat test_output.log | $(GO_JUNIT_REPORT) > report.xml
	@ls -la

.PHONY: coverage-html
coverage-html: coverage ## 🌐 View coverage report in browser
	@echo ">> opening coverage report in browser..."
	@go tool cover -html=coverage.out

# ==============================================================================
# Tooling Targets
# ==============================================================================

.PHONY: tools
tools: $(GOLANGCI_LINT) $(GO_JUNIT_REPORT) ## 🛠️ Install development tools

$(GOLANGCI_LINT):
	@echo ">> checking for $(GOLANGCI_LINT)..."
	@command -v $(GOLANGCI_LINT) >/dev/null 2>&1 || ( \
		if [ "$(OS_NAME)" = "Darwin" ]; then \
			echo "   -> $(GOLANGCI_LINT) not found, installing with Homebrew..." && \
			brew install golangci/tap/golangci-lint; \
		elif [ "$(OS_NAME)" = "Linux" ]; then \
			echo "   -> $(GOLANGCI_LINT) not found, installing with official script..." && \
			curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin; \
		else \
			echo "   -> Unsupported operating system for automatic installation." && \
			exit 1; \
		fi \
	)

$(GO_JUNIT_REPORT):
	@echo ">> checking for $(GO_JUNIT_REPORT)..."
	@command -v $(GO_JUNIT_REPORT) >/dev/null 2>&1 || \
		(echo "   -> $(GO_JUNIT_REPORT) not found, installing..." && \
		go install github.com/jstemmer/go-junit-report@latest)

# ==============================================================================
# Housekeeping Targets
# ==============================================================================

.PHONY: clean
clean: ## 🧹 Clean build artifacts
	@echo ">> cleaning up..."
	@rm -f coverage.out report.xml test_output.log
	@# This project is a library, so there are no other build artifacts to clean by default.
	@# This target is here for convention.

# Set the default target to 'help'
.DEFAULT_GOAL := help