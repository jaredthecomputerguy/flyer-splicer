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
	@printf "Building...\n-----\n"
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN) $(SRC)

# Run binary
run: $(BIN)
	@printf "Running...\n-----\n"
	@./$(BIN)

# Development mode (quick run without building)
dev:
	@printf "Developing...\n-----\n"
	@go run main.go

# Clean output directory
clean:
	@printf "Cleaning...\n-----\n"
	@rm -rf $(BIN_DIR)
