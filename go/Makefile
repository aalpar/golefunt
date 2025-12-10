# Elefunt Go Test Suite Makefile
# Port of the classic elementary function test package by W.J. Cody

VERSION=$(shell cat VERSION)
GITSHA=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.GitSHA=$(GITSHA)"

.PHONY: all clean test build

# Build all test programs
all: build

TESTS = sincos exp log tan sqrt asin atan sinh tanh power

build:
	@mkdir -p bin
	@for test in $(TESTS); do \
		echo "Building $$test ($(VERSION)-$(GITSHA))..."; \
		go build $(LDFLAGS) -o bin/$$test ./$$test; \
	done

# Run all tests
test: build
	@echo "Running all elefunt tests..."
	@echo ""
	@for test in $(TESTS); do \
		echo "========================================"; \
		echo "Running $$test test"; \
		echo "========================================"; \
		./bin/$$test; \
		echo ""; \
	done

# Run individual tests
test-sincos: build
	./bin/sincos

test-exp: build
	./bin/exp

test-log: build
	./bin/log

test-tan: build
	./bin/tan

test-sqrt: build
	./bin/sqrt

test-asin: build
	./bin/asin

test-atan: build
	./bin/atan

test-sinh: build
	./bin/sinh

test-tanh: build
	./bin/tanh

test-power: build
	./bin/power

# Clean build artifacts
clean:
	rm -rf bin/
	go clean ./...

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Module management
tidy:
	go mod tidy -x

modclean:
	go clean -x -modcache
