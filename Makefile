.PHONY: all build run clean

# Build output
BUILD_DIR := output

# App name
APP_NAME := blog-platform

all: build run

build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@docker-compose build

run:
	@echo "Running $(APP_NAME)..."
	@docker-compose up -d

# Clean build output
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
