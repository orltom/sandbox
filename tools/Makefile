.PHONY: all
all: install

.PHONY: install
install: bin/go-test-sum bin/go-licenses bin/goimports bin/golangci-lint

.PHONY: clean
clean:
	rm -rf ./bin

bin/gotestsum: go.mod ## Install go-junit-report locally if necessary.
	$(call go-install-tool,gotest.tools/gotestsum)

bin/go-licenses: go.mod ## Install go-license locally if necessary
	$(call go-install-tool,github.com/google/go-licenses)

bin/goimports: go.mod ## Install goimports locally if necessary.
	$(call go-install-tool,golang.org/x/tools/cmd/goimports)

bin/golangci-lint: go.mod ## Install golangci-lint if necessary
	$(call go-install-tool,github.com/golangci/golangci-lint/cmd/golangci-lint)

PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
# go-install-tool will 'go install' $2 into the tools module's bin folder
define go-install-tool
@echo "Installing $(1)"
GOBIN=$(PROJECT_DIR)/bin go install $(1)
endef