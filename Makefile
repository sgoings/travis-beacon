ifndef GOPATH
$(error No GOPATH set)
endif

export GO15VENDOREXPERIMENT=1

VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null)
MAIN_BIN := bin/beacon
MAIN_GO := beacon.go

.PHONY: build
build: $(MAIN_GO)
	go build -o $(MAIN_BIN) -ldflags "-X main.version=${VERSION}" $<

.PHONY: test
test:
	go test -v ./

.PHONY: test-dependencies
test-dependencies:
	go get -u github.com/verdverm/frisby

.PHONY: run
run: build
	$(MAIN_BIN)
