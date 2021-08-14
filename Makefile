NAME?=screenshot-tools
BUILD?=go build -ldflags="-w -s"
PKGER_BIN?=$(shell which packr2)

ifeq ($(PREFIX),)
    PREFIX := /usr/local
endif
ifeq ($(PKGER_BIN),)
    PKGER_BIN := $(shell go env GOPATH)/bin/packr2
endif

default: generate format vet build

.PHONY: build
build:
	$(BUILD) -o $(NAME) main.go main-packr.go

build-all: clean generate build-linux build-osx build-windows

build-windows:
	GOOS=windows GOARCH=amd64 TAG=main \
	ARGS="-e NAME=screenshot-tools_win.exe" \
	CMD="make build" ./cross_build.sh

build-linux:
	GOOS=linux GOARCH=amd64 TAG=main \
	ARGS="-e NAME=screenshot-tools_linux" \
	CMD="make build" ./cross_build.sh

build-osx:
	GOOS=darwin GOARCH=amd64 TAG=darwin \
	ARGS="-e NAME=screenshot-tools_osx" \
	CMD="make build" ./cross_build.sh

install: deps generate build
	sudo mv screenshot-tools $(PREFIX)/bin/screenshot-tools

clean:
	$(PKGER_BIN) clean
	rm -rf assets/*.zip
	rm -rf $(NAME)
	rm -rf $(NAME)*

generate:
	go generate
	$(PKGER_BIN)

deps:
	GO111MODULE=off go get -u github.com/gobuffalo/packr/v2/packr2
	go mod download

test:
	go test ./...

format:
	go fmt ./...

vet:
	go vet ./...
