NAME = spctl
HARDWARE = $(shell uname -m)
VERSION ?= 0.0.0
BUILD_TAG ?= dev
BUILD_DIR ?= build

CWD := $(shell pwd)
GOPATH ?= $(CWD)/GOPATH
GOX := $(GOPATH)/bin/gox
GOLINT := $(GOPATH)/bin/golint


build: build/linux build/darwin build/windows

$(BUILD_DIR)/%: build-deps deps
	$(GOX) -os=$(subst build/,,$@) -output="$(BUILD_DIR)/{{.OS}}/{{.Arch}}/{{.Dir}}" ./securepass/spctl

build-deps: build-deps-stamp
build-deps-stamp:
	go get github.com/mitchellh/gox
	touch $@

deps: deps-stamp
deps-stamp:
	go get -u -v github.com/codegangsta/cli
	go get -u -v gopkg.in/ini.v1
	go get -u -v github.com/progrium/gh-release/...
	go get -u -v github.com/golang/lint/golint
	go get -d -v ./... || true
	touch $@

release: build
	rm -rf release && mkdir release
	tar -zcf release/$(NAME)_$(VERSION)_linux_$(HARDWARE).tgz -C build/linux $(NAME)
	tar -zcf release/$(NAME)_$(VERSION)_darwin_$(HARDWARE).tgz -C build/darwin $(NAME)
	gh-release create garlsecurity/go-securepass $(VERSION) $(shell git rev-parse --abbrev-ref HEAD)

test: deps
	go vet ./...
	go test -cover -v ./...
	$(GOLINT) ./...

dist-clean: clean
	rm -rf $(BUILD_DIR)/

clean:
	rm -f build-deps-stamp check-deps-stamp deps-stamp


.PHONY: build-deps release deps clean test
