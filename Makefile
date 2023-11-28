#!/usr/bin/make -f

VERSION=$(shell git describe --tags --always)
NAME=github-topics
PACKAGE=github.com/nicksantamaria/$(NAME)
LGFLAGS="-extldflags "-static""

# Build binaries for linux/amd64 and darwin/amd64
build:
	gox -os='linux darwin' -arch='amd64 arm64' -output='bin/$(NAME)_{{.Arch}}_{{.OS}}' -ldflags=$(LGFLAGS) $(PACKAGE)


# Run all lint checking with exit codes for CI
lint:
	golint -set_exit_status .

# Run tests with coverage reporting
test:
	go test -cover ./...

