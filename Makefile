# Setting SHELL to bash allows bash commands to be executed by recipes.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

PROJECT_DIR := $(shell dirname $(abspath $(firstword $(MAKEFILE_LIST))))
OUTPUT_DIRECTORY=${PROJECT_DIR}/bin
BUILD_DIR=${PROJECT_DIR}/build
TOOLS_BIN_DIR=${PROJECT_DIR}/tools/bin
REPORT_DIR=${BUILD_DIR}/reports
SCRIPT_DIR=${PROJECT_DIR}/scripts


GOIMPORTS = $(TOOLS_BIN_DIR)/goimports
GOLANGCI_LINT = $(TOOLS_BIN_DIR)/golangci-lint
GOLICENSES = $(TOOLS_BIN_DIR)/go-licenses
GOTESTSUM = $(TOOLS_BIN_DIR)/gotestsum

API_IMG = golang-http-example:latest
DB_IMG = golang-db-example:latest



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
build: build-api build-docker-images ## Build web application and docker images

.PHONY: build-api
build-api: ## Build manager & network-manager binary.
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o "$(OUTPUT_DIRECTORY)/golang-http-example-amd64" ./...

.PHONY:
download-jokes: ## Download random chuck norris jokes
	mkdir -p $(BUILD_DIR)
	rm -rf $(BUILD_DIR)/jokes.sql
	echo "CREATE DATABASE jokes;" >> $(BUILD_DIR)/jokes.sql
	echo "CREATE TABLE chuck_norris (id SERIAL NOT NULL, joke VARCHAR NOT NULL, PRIMARY KEY (id));" >> $(BUILD_DIR)/jokes.sql
	@for i in {1..10}; do \
	  JOKE=`http https://api.chucknorris.io/jokes/random | jq -r '.value' | sed "s/'/''/g"` ; \
	  echo "INSERT INTO chuck_norris(joke) VALUES ('$$JOKE');" >> $(BUILD_DIR)/jokes.sql ; \
	done

.PHONY: build-docker-images
build-docker-images: build-api-docker-image build-database-docker-image ## Build docker images

.PHONY: build-api-docker-image
build-api-docker-image: build ## Build API docker images
	$(call build-docker-image,$(API_IMG),config/docker/httpd/Dockerfile)

.PHONY: build-database-docker-image
build-database-docker-image: download-jokes ## Build DB docker images
	$(call build-docker-image,$(DB_IMG),config/docker/database/Dockerfile)

define build-docker-image
DOCKER_BUILDKIT=1 docker build --platform linux/amd64 -t "$(1)" \
		--build-arg "GIT_HASH=$$(git rev-parse HEAD)" \
		--build-arg "BUILD_DATE=$$(date --iso-8601=s)" \
		-f "$(2)" .
endef

## Helper Rules
$(TOOLS_BIN_DIR)/%: FORCE
	@$(MAKE) -sC tools bin/$*

$(REPORT_DIR):
	mkdir -p "$(REPORT_DIR)"

.PHONY: FORCE
FORCE: