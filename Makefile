SHELL   := /bin/bash
VERSION := v0.4.4
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
	tar zcf demitas_$(VERSION)_$(GOOS)_$(GOARCH).tar.gz demitas demitas-pf
	sha1sum demitas_$(VERSION)_$(GOOS)_$(GOARCH).tar.gz > demitas_$(VERSION)_$(GOOS)_$(GOARCH).tar.gz.sha1sum

.PHONY: clean
clean:
	rm -f demitas
