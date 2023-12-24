.PHONY: all check build run test clean

# Build output
BUILD_DIR := output

# App name
APP_NAME := blog-platform

all: check build run

build: check
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@docker-compose build

run:
	@echo "Running $(APP_NAME)..."
	@docker-compose up -d

test: check
	@echo "Testing $(APP_NAME)..."
	@go test -v ./...

# Clean build output
clean:
	@echo "Cleaning up..."
	@docker-compose down
	@rm -rf $(BUILD_DIR)

check:
	@echo "Checking go and docker enviroment!"
	@which go > /dev/null || (echo "Go is not installed or not in PATH"; exit 1)
	@which docker-compose > /dev/null || (echo "Docker Compose is not installed or not in PATH"; exit 1)
	@echo "Go and docker is ready!"

