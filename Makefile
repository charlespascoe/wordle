BINARY_NAME=wordle
BINARY_LINUX=$(BINARY_NAME)-linux
SHELL=/bin/bash
GOLANG_MINOR_VERSION=17
ALL_GO_FILES=find . -name '*.go' -not -path './vendor*' -not -path './.*'

.PHONY: all
## : Same as 'make goversion init download build', recommended after checking out
all: goversion build

.PHONY: help
## help: Prints this help
help:
	@sed -ne 's/^##/make/p' $(MAKEFILE_LIST) | column -c2 -t -s ':' | sort

.PHONY: build
## build: Builds binary for current OS
build:
	go build -o $(BINARY_NAME) -v ./*.go

.PHONY: linux
## linux: Builds Linux-compatible binary
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_LINUX) -v ./*.go

.PHONY: clean
## clean: Clean up build artifacts
clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_LINUX)

.PHONY: run
## run: Builds and executes the program
run: build
	./$(BINARY_NAME)

.PHONY: fmt
## fmt: Run go fmt on all source files
fmt:
	go fmt ./...

.PHONY: vet
## vet: Run go vet on all source files
vet:
	go vet ./...

.PHONY: fmtchk
## fmtchk: Checks code formatting, exits with non-zero exit code if there are formatting issues.
fmtchk:
	@issues="$$($(ALL_GO_FILES) | xargs gofmt -l)"; \
	if [ -n "$$issues" ]; then \
		echo -e "ERROR: Formatting issues in the following files:\n$$issues"; \
		echo "Run 'make fmt' and stage the changes before attempting to commit."; \
		exit 1; \
	fi

.PHONY: goversion
## goversion: Checks minimum version of Go is installed
goversion:
	@if [ -z "$$(which go)" ]; then \
		echo "go either not installed or not in PATH"; \
		exit 1; \
	fi
	@minor="$$(go version | egrep -o 'go1.([0-9]+)' | egrep -o '[0-9]+$$')"; \
	if [ -z "$$minor" ]; then \
		echo "Could not determine Go version (1.$(GOLANG_MINOR_VERSION) or above required)"; \
		exit 1; \
	fi; \
	if [ $$minor -lt $(GOLANG_MINOR_VERSION) ]; then \
		echo "Go version 1.$(GOLANG_MINOR_VERSION) or above required (version 1.$$minor installed)"; \
		exit 1; \
	fi
