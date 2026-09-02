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
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60100 ./examples/cmd/a60100
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60101 ./examples/cmd/a60101
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60102 ./examples/cmd/a60102
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60103 ./examples/cmd/a60103
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60104 ./examples/cmd/a60104
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60105 ./examples/cmd/a60105
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60106 ./examples/cmd/a60106
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60107 ./examples/cmd/a60107
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60108 ./examples/cmd/a60108
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60109 ./examples/cmd/a60109
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60110 ./examples/cmd/a60110
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60111 ./examples/cmd/a60111
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60112 ./examples/cmd/a60112
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60113 ./examples/cmd/a60113
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60114 ./examples/cmd/a60114
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60115 ./examples/cmd/a60115
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60116 ./examples/cmd/a60116
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60117 ./examples/cmd/a60117
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60118 ./examples/cmd/a60118
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60119 ./examples/cmd/a60119
	go build -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" -o ./bin/a60120 ./examples/cmd/a60120

.PHONY: install
install: ## Install cmd applications in the equivalent to $GOPATH/bin
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60100
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60101
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60102
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60103
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60104
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60105
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60106
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60107
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60108
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60109
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60110
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60111
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60112
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60113
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60114
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60115
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60116
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60117
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60118
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60119
	GO111MODULE=on go install -trimpath -ldflags "$(GO_LDFLAGS) -X main.debugMode=$(DEBUG)" ./examples/cmd/a60120


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
	rm -rf ./bin ./test/data/btcusd_one_min_data.csv.gz $(COV) $(FUNC)/j* $(ACTION)/j* $(ARTIFACT)/b* $(EXEC)/*.tsv

.PHONY: check
check: ## Validate URL links
	bash ./scripts/automating/b60108 ./docs/go-interlace/d60100/urls.txt

.PHONY: download
download: check ## Download files in the URL list, renames the non-persistent ones, and moves them to $ENV_ALT_DATA_PATH
	bash ./scripts/automating/b60106 ./docs/go-interlace/d60100/urls.txt
	bash ./scripts/automating/b60107

.PHONY: help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
