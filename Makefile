# Directories
BIN_DIR := bin
BIN := $(BIN_DIR)/flyer-splicer

SRC := $(shell find cmd internal -type f -name '*.go')

.PHONY: all run dev clean

# Default target: build the binary
all: $(BIN)

# Build binary if sources changed
$(BIN): $(SRC)
	@printf "Building...\n-----\n"
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN) ./cmd

# Run binary
run: $(BIN)
	@printf "Running...\n-----\n"
	@./$(BIN)

# Development mode (run without building)
dev:
	@printf "Developing...\n-----\n"
	@go run ./cmd

# Clean output directory
clean:
	@printf "Cleaning...\n-----\n"
	@rm -rf $(BIN_DIR)
