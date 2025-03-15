# Go variables
GO ?= go
GOBUILD ?= $(GO) build

# Files
MAIN_FILE ?= ./cmd/app/main.go

# Install dependencies
.PHONY: deps
deps:
	$(GO) mod tidy
	$(GO) mod download

# Build the application
.PHONY: build
build:
	@echo "Building the application..."
	$(GOBUILD) -o bin/main $(MAIN_FILE)
