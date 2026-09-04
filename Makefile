DEBUG ?= false

COV="./coverage.txt"

PKG_VERSION ?= $(shell git describe --tags 2>/dev/null || git rev-parse --short HEAD)
DATE_FMT = +%Y-%m-%d
BUILD_DATE ?= $(shell date "$(DATE_FMT)")
GO_LDFLAGS := -X main.buildDate=$(BUILD_DATE) $(GO_LDFLAGS)
GO_LDFLAGS := $(GO_LDFLAGS) -X main.version=$(PKG_VERSION)

.PHONY: build
.DEFAULT_GOAL := build
build: ## Build the cmd applications
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a31001 ./cmd/a31001
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a31002 ./cmd/a31002
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a31003 ./cmd/a31003
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a31004 ./cmd/a31004
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a31005 ./cmd/a31005
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a31006 ./cmd/a31006
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a31007 ./cmd/a31007
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a31008 ./cmd/a31008
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a31009 ./cmd/a31009
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a31010 ./cmd/a31010
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a31011 ./cmd/a31011
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a31012 ./cmd/a31012
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a31013 ./cmd/a31013
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a31014 ./cmd/a31014
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a31015 ./cmd/a31015

.PHONY: install
install: ## Install cmd applications in the equivalent to $GOPATH/bin
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./cmd/a31001
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./cmd/a31002
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./cmd/a31003
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./cmd/a31004
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./cmd/a31005
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./cmd/a31006
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./cmd/a31007
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./cmd/a31008
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./cmd/a31009
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./cmd/a31010
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./cmd/a31011
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./cmd/a31012
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./cmd/a31013
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./cmd/a31014
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./cmd/a31015

.PHONY: test
test: ## Run tests
	go test -coverprofile $(COV) ./...

.PHONY: coverage
coverage: ## Run coverage report
	go tool cover -func $(COV)

.PHONY: vet
vet: ## Run code issues detection
	go vet ./...

.PHONY: clean
clean: ## Remove the bin folder and test files leaving the download data files intact
	rm -rf ./bin $(COV)

.PHONY: check
check: ## Validate URL links
	bash ./scripts/bash/b31001 ./configs/input/url.txt

.PHONY: download
download: check ## Downloads, renames the non-persistent ones, and moves all to $ENV_ALT_DATA_PATH
	bash ./scripts/bash/b31002 ./configs/input/url.txt
	bash ./scripts/bash/b31003

.PHONY: help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
