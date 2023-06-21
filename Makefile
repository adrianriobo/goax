# Go and compilation related variables
BUILD_DIR ?= out
SOURCE_DIRS = cmd pkg test
SOURCES := $(shell git ls-files  *.go ":^vendor")

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

$(BUILD_DIR)/windows-amd64/goax.exe: $(SOURCES)
	CC=clang GOARCH=amd64 GOOS=windows go build -o $(BUILD_DIR)/windows-amd64/goax.exe ./cmd

.PHONY: build-darwin
build-darwin: clean $(BUILD_DIR)/darwin-amd64/goax

$(BUILD_DIR)/darwin-amd64/goax: $(SOURCES)
	CC=clang GOARCH=amd64 GOOS=darwin go build -o $(BUILD_DIR)/darwin-amd64/goax ./cmd

.PHONY: build-windows
build-windows: clean $(BUILD_DIR)/windows-amd64/goax.exe