#@IgnoreInspection BashAddShebang

export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

export APP=openflag

export APP_VERSION=v0.1.0

export BUILD_INFO_PKG="github.com/OpenFlag/OpenFlag/pkg/version"

export LDFLAGS="-w -s -X '$(BUILD_INFO_PKG).AppVersion=$(APP_VERSION)' -X '$(BUILD_INFO_PKG).Date=$$(date)' -X '$(BUILD_INFO_PKG).BuildVersion=$$(git rev-parse HEAD | cut -c 1-8)' -X '$(BUILD_INFO_PKG).VCSRef=$$(git rev-parse --abbrev-ref HEAD)'"

all: format lint build

run-version:
	go run -ldflags $(LDFLAGS) ./cmd/openflag version

run-migrate:
	go run -ldflags $(LDFLAGS) ./cmd/openflag migrate --path ./internal/app/openflag/migrations

run-server:
	go run -ldflags $(LDFLAGS) ./cmd/openflag server

build:
	go build -ldflags $(LDFLAGS)  ./cmd/openflag

install:
	go install -ldflags $(LDFLAGS) ./cmd/openflag

release:
	./scripts/release.sh $(APP) $(APP_VERSION) $(LDFLAGS)

check-formatter:
	which goimports || GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports

format: check-formatter
	find $(ROOT) -type f -name "*.go" -not -path "$(ROOT)/vendor/*" | xargs -n 1 -I R goimports -w R
	find $(ROOT) -type f -name "*.go" -not -path "$(ROOT)/vendor/*" | xargs -n 1 -I R gofmt -s -w R

check-linter:
	which golangci-lint || GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint@v1.23.8

lint: check-linter
	golangci-lint run $(ROOT)/...

test:
	go test -v -race -p 1 ./...

ci-test:
	go test -v -race -p 1 -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -func coverage.txt

up:
	docker-compose up -d redis postgres mysql

down:
	docker-compose down
