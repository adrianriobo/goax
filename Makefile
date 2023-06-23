# Go and compilation related variables
BUILD_DIR ?= out
SOURCE_DIRS = cmd pkg test
SOURCES := $(shell git ls-files  *.go ":^vendor")
ARCH ?= amd64

.PHONY: clean ## Remove all build artifacts
clean: 
	rm -rf $(BUILD_DIR)

# Create and update the vendor directory
.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: cross
cross: clean $(BUILD_DIR)/windows-amd64/goax.exe

.PHONY: build-windows
build-windows: clean $(BUILD_DIR)/windows-amd64/goax.exe

$(BUILD_DIR)/windows-amd64/goax.exe: $(SOURCES)
	CC=clang GOARCH=amd64 GOOS=windows go build -o $(BUILD_DIR)/windows-amd64/goax.exe ./cmd

.PHONY: build-darwin
build-darwin: clean $(BUILD_DIR)/darwin-${ARCH}/goax

$(BUILD_DIR)/darwin-${ARCH}/goax:
	CGO_ENABLED=1 CC=clang GOARCH=${ARCH} GOOS=darwin go build -o $(BUILD_DIR)/darwin-${ARCH}/goax ./cmd
