CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-edtf
	cp  *.go src/github.com/whosonfirst/go-edtf/
	cp -r bnf src/github.com/whosonfirst/go-edtf/
	if test -d vendor; then cp -r vendor/* src/; fi

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:
	@echo "no deps"

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt *.go
	go fmt bnf/*.go
	go fmt cmd/*.go

bin: 	self
	if test -d bin; then rm -rf bin/*; fi
	@GOPATH=$(GOPATH) go build -o bin/parse cmd/parse.go	
