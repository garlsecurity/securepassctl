NAME = spctl
HARDWARE = $(shell uname -m)
VERSION ?= 0.0.0
BUILD_TAG ?= dev
BUILD_DIR ?= build

CWD := $(shell pwd)
GOPATH ?= $(CWD)/GOPATH
GOLINT := $(GOPATH)/bin/golint
LDFLAGS += -X=main.Version=$(VERSION)

all: $(BUILD_DIR)/linux/amd64 $(BUILD_DIR)/linux/386 \
	$(BUILD_DIR)/darwin/amd64 $(BUILD_DIR)/darwin/386 \
	$(BUILD_DIR)/windows/amd64 $(BUILD_DIR)/windows/386

$(BUILD_DIR)/%:
	GOOS=$(word 2,$(subst /, ,$@)) GOARCH=$(word 3,$(subst /, ,$@)) go build -v -ldflags="$(LDFLAGS)" -o $@/spctl ../securepassctl/spctl

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
	gh-release create garlsecurity/securepassctl $(VERSION) $(shell git rev-parse --abbrev-ref HEAD)

test: deps
	go vet ./...
	go test -cover -v ./...
	$(GOLINT) ./...

dist-clean: clean
	rm -f build-deps-stamp check-deps-stamp deps-stamp

clean:
	rm -rf $(BUILD_DIR)/


.PHONY: release deps clean test
