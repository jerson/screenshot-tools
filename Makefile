APP_VERSION?=latest
PACKAGER?=packr2
BUILD?=go build -ldflags="-w -s"
NAME?=screenshot-tools

default: generate format vet build

build: generate
	$(BUILD) -o $(NAME) main.go

build-in-docker:
	apt update && apt install -y zip
	make packr2
	make build

build-all: clean build-linux build-osx build-windows

build-windows:
	GOOS=windows GOARCH=amd64 TAG=main \
	ARGS="-e NAME=screenshot-tools_win.exe" \
	CMD="make build-in-docker" ./cross_build.sh

build-linux:
	GOOS=linux GOARCH=amd64 TAG=main \
	ARGS="-e NAME=screenshot-tools_linux" \
	CMD="make build-in-docker" ./cross_build.sh

build-osx:
	GOOS=darwin GOARCH=amd64 TAG=darwin \
	ARGS="-e NAME=screenshot-tools_osx" \
	CMD="make build-in-docker" ./cross_build.sh

clean:
	$(PACKAGER) clean
	rm -rf assets/*.zip
	rm -rf $(NAME)
	rm -rf $(NAME)*

generate:
	go generate
	packr2

packr2:
	GO111MODULE=off go get -u github.com/gobuffalo/packr/v2/packr2

test:
	go test ./...

format:
	go fmt ./...

vet:
	go vet ./...
