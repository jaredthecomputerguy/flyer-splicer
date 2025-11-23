# Directories
BIN_DIR := bin
BIN := $(BIN_DIR)/flyer-splicer

# Source files
SRC := $(wildcard *.go)

.PHONY: all run dev clean

# Default target: build the binary
all: $(BIN)

# Build binary if sources changed
$(BIN): $(SRC)
	@echo "Building..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN) $(SRC)

# Run binary
run: $(BIN)
	@echo "Running..."
	@./$(BIN)

# Development mode (quick run without building)
dev:
	@echo "Developing..."
	@go run main.go

# Clean output directory
clean:
	@echo "Cleaning..."
	@rm -rf $(BIN_DIR)
