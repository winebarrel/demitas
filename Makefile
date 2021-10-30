SHELL   := /bin/bash
VERSION := v0.4.0
GOOS    := $(shell go env GOOS)
GOARCH  := $(shell go env GOARCH)

.PHONY: all
all: vet build

.PHONY: build
build:
	go build -ldflags "-X main.version=$(VERSION)" ./cmd/demitas

.PHONY: vet
vet:
	go vet

.PHONY: package
package: clean vet build
	gzip demitas -c > demitas_$(VERSION)_$(GOOS)_$(GOARCH).gz
	sha1sum demitas_$(VERSION)_$(GOOS)_$(GOARCH).gz > demitas_$(VERSION)_$(GOOS)_$(GOARCH).gz.sha1sum

.PHONY: clean
clean:
	rm -f demitas
