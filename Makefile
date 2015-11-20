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

.PHONY: watch-test
watch-test: bootstrap build
	Watch -t make test

.PHONY: run
run: build
	$(MAIN_BIN)

.PHONY: watch-run
watch-run: build
	Watch -t make run

.PHONY: bootstrap
bootstrap:
	go get github.com/eaburns/Watch

.PHONY: e2e
e2e:
	_test/main.sh
