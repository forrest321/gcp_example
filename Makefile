# Makefile for the RSS Reader project

# Project variables
BINARY_NAME=goread
BACKEND_DIR=./backend
FRONTEND_DIR=./frontend

# Build variables
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean

# Build the backend binary
.PHONY: backend
backend:
	$(GO_BUILD) -o $(BINARY_NAME) $(BACKEND_DIR)/cmd

# Clean build files
.PHONY: clean
clean:
	$(GO_CLEAN)
	rm -f $(BINARY_NAME)

# TODO: Add frontend build steps when ready
# .PHONY: frontend
# frontend:
#     (frontend build steps go here)

# Build both backend and frontend
.PHONY: all
all: backend # frontend
