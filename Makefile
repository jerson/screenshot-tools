APP_VERSION?=latest
PACKAGER?=packr2
BUILD?=packr2 build -ldflags="-w -s"
NAME?=screenshot-tools
UPX?=upx

default: build

build: generate format vet
	$(BUILD) -o $(NAME) main.go
	$(UPX) $(NAME)

build-all: clean build-linux build-windows build-osx

build-windows: generate format vet
	GOOS=windows GOARCH=amd64 $(BUILD) -o $(NAME)_win.exe main.go
	$(UPX) $(NAME)_win.exe

build-linux: generate format vet
	GOOS=linux GOARCH=amd64 $(BUILD) -o $(NAME)_linux main.go
	$(UPX) $(NAME)_linux

build-osx: generate format vet
	GOOS=darwin GOARCH=amd64 $(BUILD) -o $(NAME)_osx main.go
	$(UPX) $(NAME)_osx

clean:
	$(PACKAGER) clean
	rm -rf assets/*.zip
	rm -rf $(NAME)
	rm -rf $(NAME)*

generate:
	go generate
	packr2

test:
	go test ./...

format:
	go fmt ./...

vet:
	go vet ./...
