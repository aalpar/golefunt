# ELEFUNT - Elementary Function Testing
# Root Makefile for building and running Fortran and Go versions

SUBDIRS = fortran go

.PHONY: all build test clean fortran go

# Build both Fortran and Go versions
all: build

build: build-fortran build-go

build-fortran:
	@echo "========================================"
	@echo "Building Fortran version"
	@echo "========================================"
	$(MAKE) -C fortran all

build-go:
	@echo "========================================"
	@echo "Building Go version"
	@echo "========================================"
	$(MAKE) -C go build

# Run all tests (Fortran then Go)
test: test-fortran test-go

test-fortran: build-fortran
	@echo "========================================"
	@echo "Running Fortran tests"
	@echo "========================================"
	$(MAKE) -C fortran test

test-go: build-go
	@echo "========================================"
	@echo "Running Go tests"
	@echo "========================================"
	$(MAKE) -C go test

# Clean both
clean:
	$(MAKE) -C fortran clean
	$(MAKE) -C go clean

# Individual subdirectory targets
fortran:
	$(MAKE) -C fortran

go:
	$(MAKE) -C go
