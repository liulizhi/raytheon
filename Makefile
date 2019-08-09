# basic go cmd

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GO111MODULE=on


# binary names
BINARY_NAME=bin/raytheon
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_MAC=$(BINARY_NAME)_mac
BINARY_WIN=$(BINARY_NAME)_win

assemble: build_linux

all: test build
build:
	GOPROXY=https://goproxy.io $(GOBUILD) -o $(BINARY_NAME) -v -gcflags "all=-trimpath=${GOPATH}/src" .

test:
	GOPROXY=https://goproxy.io $(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_UNIX)
	rm -f $(BINARY_MAC)
	rm -f $(BINARY_WIN)

run:
	GOPROXY=https://goproxy.io $(GOBUILD) -o $(BINARY_NAME) -v -gcflags "all=-trimpath=${GOPATH}/src" .
	./$(BINARY_NAME)

# Cross compilation
build_linux:
	GOPROXY=https://goproxy.io CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v -gcflags "all=-trimpath=${GOPATH}/src" .

build_win:
	GOPROXY=https://goproxy.io CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_WIN) -v -gcflags "all=-trimpath=${GOPATH}/src" .

build_mac:
	GOPROXY=https://goproxy.io CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_MAC) -v -gcflags "all=-trimpath=${GOPATH}/src" .

start_server:
	$(BINARY_UNIX)