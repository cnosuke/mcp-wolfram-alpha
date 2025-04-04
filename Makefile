NAME     := mcp-greeting
VERSION  := $(shell git describe --tags 2>/dev/null)
REVISION := $(shell git rev-parse --short HEAD 2>/dev/null)
SRCS    := $(shell find . -type f -name '*.go' -o -name 'go.*')
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\""

bin/$(NAME): $(SRCS)
	go build $(LDFLAGS) -o bin/$(NAME) main.go

.PHONY: test deps inspect clean

deps:
	go mod download

inspect:
	golangci-lint run

clean:
	rm -rf bin/* dist/*

test:
	go test -v ./...
