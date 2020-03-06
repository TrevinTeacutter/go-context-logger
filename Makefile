GOVERSION = $(shell go version | awk '{print $$3;}')
SOURCE_FILES?=./...

export PATH := ./bin:$(PATH)
export CGO_ENABLED := 0

clean:
	rm -rf ./dist && rm -rf ./vendor
.PHONY: clean

vendor:
	go mod vendor
.PHONY: vendor

fmt:
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done
.PHONY: fmt

lint:
	# Cannot do both outputs in single command -> https://github.com/golangci/golangci-lint/issues/481
	# Human readable output
	golangci-lint run --timeout=5m
.PHONY: lint

test:
	gotestsum -- -failfast -v -covermode count -timeout 5m ./...
.PHONY: test

tidy:
	go mod tidy
.PHONY: tidy

build:
	GOVERSION=$(GOVERSION) goreleaser release --snapshot --skip-publish --skip-sign --rm-dist --debug
.PHONY: build

snapshot:
	GOVERSION=$(GOVERSION) goreleaser release --snapshot --skip-sign --rm-dist --debug
.PHONY: snapshot

release:
	GOVERSION=$(GOVERSION) goreleaser release --rm-dist --skip-sign --debug
.PHONY: release
