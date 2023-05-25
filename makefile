# Variables
BINARY_NAME := fairyport
INSTALL_PATH := $(HOME)/go/bin
HERMES_RELAYER := hermes
HERMES_INSTALLED := $(shell command -v $(HERMES_RELAYER) 2> /dev/null)

# Default target
all: build

# Check for cargo
check_cargo:
	@command -v cargo >/dev/null 2>&1 || { echo >&2 "cargo is required but not installed. Please install Rust and Cargo and add it to your PATH."; exit 1; }

# Check if Hermes relayer is installed
check_hermes: check_cargo
ifdef HERMES_INSTALLED
	@echo "Hermes relayer is already installed. Version:"
	@$(HERMES_RELAYER) --version
else
	@echo "Hermes relayer is not installed. Installing..."
	@cargo install ibc-relayer-cli --bin hermes --locked
endif

# Build target
build: check_hermes
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME)

# Install target
install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
	@mkdir -p $(INSTALL_PATH)
	@cp $(BINARY_NAME) $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "Removing $(BINARY_NAME) from project directory..."
	@rm -f $(BINARY_NAME)

# Uninstall target
uninstall:
	@echo "Uninstalling $(BINARY_NAME) from $(INSTALL_PATH)..."
	@rm -f $(INSTALL_PATH)/$(BINARY_NAME)

# Clean target
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)

.PHONY: all check_cargo install_ibc build install uninstall clean
