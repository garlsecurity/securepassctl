NAME = spctl
HARDWARE = $(shell uname -m)
VERSION ?= 0.0.0
BUILD_TAG ?= dev

build:
	mkdir -p build/linux  && GOOS=linux  go build -a -ldflags "-X main.Version=$(VERSION)" -o build/linux/$(NAME) ./securepass/spctl
	mkdir -p build/darwin && GOOS=darwin go build -a -ldflags "-X main.Version=$(VERSION)" -o build/darwin/$(NAME) ./securepass/spctl

deps:
	go get -u -v github.com/codegangsta/cli
	go get -u -v gopkg.in/ini.v1
	go get -u -v github.com/progrium/gh-release/...
	go get -u -v github.com/golang/lint/golint
	go get -u -v ./... || true

release: build
	rm -rf release && mkdir release
	tar -zcf release/$(NAME)_$(VERSION)_linux_$(HARDWARE).tgz -C build/linux $(NAME)
	tar -zcf release/$(NAME)_$(VERSION)_darwin_$(HARDWARE).tgz -C build/darwin $(NAME)
	gh-release create garlsecurity/go-securepass $(VERSION) $(shell git rev-parse --abbrev-ref HEAD)


test:
	go vet ./...
	go test -cover -v ./...
	golint ./...

clean:
	rm -rf build/

.PHONY: build release deps clean test
