SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

PROJECT_DIR := $(shell dirname $(abspath $(firstword $(MAKEFILE_LIST))))
OUTPUT_DIRECTORY=${PROJECT_DIR}/bin
BUILD_DIR=${PROJECT_DIR}/build
TOOLS_BIN_DIR=${PROJECT_DIR}/tools/bin
REPORT_DIR=${BUILD_DIR}/reports

# Tools
GOIMPORTS = $(TOOLS_BIN_DIR)/goimports
GOLANGCI_LINT = $(TOOLS_BIN_DIR)/golangci-lint
GOLICENSES = $(TOOLS_BIN_DIR)/go-licenses
GOTESTSUM = $(TOOLS_BIN_DIR)/gotestsum


.PHONY: all
all: fmt


##@ General
.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development
.PHONY: clean
clean: ## Clean output
	rm -rf "$(BUILD_DIR)" "$(TOOLS_BIN_DIR)" "${OUTPUT_DIRECTORY}"

.PHONY: gomod
gomod:
	go mod tidy
	go mod verify

.PHONY: fmt
fmt: $(GOIMPORTS) $(GCI) ## Run goimports
	@for d in $$(go list -f '{{.Dir}}' ./...); do \
		find "$$d" -maxdepth 1 -name "*.go" -not -path "*generated*" -not -path "*mocks*" -exec "$(GOIMPORTS)" -local 'crypto-cli/' -w {} \; ;\
	done


##@ Checks
.PHONY: check
check: lint go-licenses-check

.PHONY: lint
lint: $(GOLANGCI_LINT) | $(REPORT_DIR) ## Run golangci-lint against code
	$(GOLANGCI_LINT) version
	$(GOLANGCI_LINT) run -v

.PHONY: go-licenses-check
go-licenses-check: $(GOLICENSES) ## Checks for forbidden Go licenses.
	$(GOLICENSES) check ./...


##@ Build
.PHONY: build
build: ## Build manager & network-manager binary.
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o "$(OUTPUT_DIRECTORY)/golang-http-example" ./...


## Helper Rules
$(TOOLS_BIN_DIR)/%: FORCE
	@$(MAKE) -sC tools bin/$*

$(REPORT_DIR):
	mkdir -p "$(REPORT_DIR)"

.PHONY: FORCE
FORCE: