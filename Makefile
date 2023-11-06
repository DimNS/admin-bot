.EXPORT_ALL_VARIABLES:
GOBIN = $(shell pwd)/bin
GOFLAGS = -mod=vendor
GO111MODULE = on

.PHONY: deps
deps:
	@go mod tidy
	@go mod vendor

.PHONY: mocks
mocks: tools
	@export PATH="$(shell pwd)/bin:$(PATH)"; mockery --config=.mockery.yaml

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: test
test:
	@go test

.PHONY: build
build:
	@GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-s -w" -o ./bin/geeksonator ./cmd/geeksonator

.PHONY: tools
tools: deps
	@go install github.com/vektra/mockery/v2@v2.35.4
	@go install github.com/goreleaser/goreleaser@v1.21.2
