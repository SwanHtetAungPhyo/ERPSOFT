# Variables
APP_NAME := esfor
GO_FILES := $(shell find . -name '*.go' -not -path "./vendor/*")
GO_MOD := go.mod
GO_SUM := go.sum
DOCKER_IMAGE := $(APP_NAME):latest

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	@echo "Building the application..."
	@go build -o $(APP_NAME) main.go

# Run the application
.PHONY: run
run: build
	@echo "Running the application..."
	@./$(APP_NAME)

# Test the application
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./...

# Clean the build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -f $(APP_NAME)

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@go mod tidy

# Update dependencies
.PHONY: update-deps
update-deps:
	@echo "Updating dependencies..."
	@go get -u ./...
	@go mod tidy

# Lint the code
.PHONY: lint
lint:
	@echo "Linting the code..."
	@golangci-lint run

# Format the code
.PHONY: fmt
fmt:
	@echo "Formatting the code..."
	@go fmt ./...

# Build Docker image
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE) .

# Run Docker container
.PHONY: docker-run
docker-run:
	@echo "Running Docker container..."
	@docker run --rm -p 8006:8006 $(DOCKER_IMAGE)

# Stop Docker container
.PHONY: docker-stop
docker-stop:
	@echo "Stopping Docker container..."
	@docker stop $(APP_NAME)

# Help
.PHONY: help
help:
    @echo "Usage:"
	@echo "  make [target]"
	@echo ""
    @echo "Targets:"
	@echo "  all            Build the application (default)"
	@echo "  build          Build the application"
	@echo "  run            Run the application"
	@echo "  test           Run tests"
	@echo "  clean          Clean build artifacts"
	@echo "  deps           Install dependencies"
	@echo "  update-deps    Update dependencies"
	@echo "  lint           Lint the code"
	@echo "  fmt            Format the code"
	@echo "  docker-build   Build Docker image"
	@echo "  docker-run     Run Docker container"
	@echo "  docker-stop    Stop Docker container"
	@echo "  help           Show this help message"