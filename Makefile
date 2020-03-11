NAME?=screenshot-tools

default: generate format vet build

build:
	go build -ldflags="-w -s" -o $(NAME) main.go

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

clean:
	packr2 clean
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
