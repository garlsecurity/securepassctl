NAME = spctl
HARDWARE = $(shell uname -m)
VERSION ?= 0.1
BUILD_TAG ?= dev
BUILD_DIR ?= build

CWD := $(shell pwd)
GOPATH ?= $(CWD)/GOPATH
GOLINT := $(GOPATH)/bin/golint
LDFLAGS_DEFAULT = -X=main.Version=$(VERSION)

ci: $(BUILD_DIR)/linux/amd64 \
	$(BUILD_DIR)/windows/amd64 \
	$(BUILD_DIR)/darwin/amd64

all: $(BUILD_DIR)/linux/amd64 $(BUILD_DIR)/linux/386 \
	$(BUILD_DIR)/darwin/amd64 $(BUILD_DIR)/darwin/386 \
	$(BUILD_DIR)/windows/amd64 $(BUILD_DIR)/windows/386

darwin: $(BUILD_DIR)/darwin/amd64 $(BUILD_DIR)/darwin/386
	lipo -create build/darwin/386/spctl build/darwin/amd64/spctl -output build/darwin/spctl
	lipo -create build/darwin/386/spssh build/darwin/amd64/spssh -output build/darwin/spssh

$(BUILD_DIR)/%: deps
	GOOS=$(word 2,$(subst /, ,$@)) GOARCH=$(word 3,$(subst /, ,$@)) go build -v -ldflags="$(LDFLAGS_DEFAULT) $(LDFLAGS)" -o $@/spctl ../securepassctl/spctl
	GOOS=$(word 2,$(subst /, ,$@)) GOARCH=$(word 3,$(subst /, ,$@)) go build -v -ldflags="$(LDFLAGS_DEFAULT) $(LDFLAGS)" -o $@/spssh ../securepassctl/spssh

deps: $(BUILD_DIR)/deps-stamp
$(BUILD_DIR)/deps-stamp:
	go get -u -v github.com/progrium/gh-release/...
	go get -u -v github.com/golang/lint/golint
	go get -u -v github.com/urfave/cli
	go get -d -v ./... || true
	mkdir -p $(BUILD_DIR)
	touch $@

release: all
	rm -rf release && mkdir release
	tar -zcf release/$(NAME)-$(VERSION)-linux.tgz -C build/linux/ amd64/$(NAME) 386/$(NAME)
	tar -zcf release/$(NAME)-$(VERSION)-darwin.tgz -C build/darwin/ $(NAME)
	mv build/windows/amd64/spctl  build/windows/spctl64.exe
	mv build/windows/386/spctl  build/windows/spctl.exe
	zip -r -j release/$(NAME)-$(VERSION)-win.zip build/windows/*exe
	gh-release create garlsecurity/securepassctl $(VERSION) $(shell git rev-parse --abbrev-ref HEAD)

test: deps
	go vet ./...
	go test -cover -v ./...
	$(GOLINT) ./...

dist-clean: clean
	rm -f $(BUILD_DIR)/deps-stamp

clean:
	rm -rf $(BUILD_DIR)/*/ -rf
	rm -rf release/*


.PHONY: release deps clean test
