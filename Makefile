#参考 - 現場で使える実践テクニック みんなのGO言語
#メタ情報
NAME := example-golang-jwt-auth-server
VERSION := $(gobump show -r)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.revision=$(REVISION)"

export GO111MODULE=on

## Install dependencies
.PHONY: deps
deps:
	go get -v -d

# 開発に必用な依存をインストールする
## Setup
.PHONY: deps
devel-deps: deps
	GO111MODULE=off go get \
		golang.org/x/lint/golint \
		github.com/motemen/gobump/cmd/gobump \
		github.com/Songmu/make2help/cmd/make2help

# テストの実行
## Run tests
.PHONY: test
test: deps
	go test ./...

## Lint
.PHONY: lint
lint: devel-deps
	go vet ./...
	golint -set_exit_status ./...

## build binaries ex. make bin/myproj
bin/%: ./src/main.go deps
	go build -ldflags $(LDFLAGS) -o $@ $<

## build server binary
.PHONY: build
build: bin/app 

## build binaries ex. make client/myproj
client/%: ./testApp/main.go deps
	go build -ldflags $(LDFLAGS) -o $@ $<

## build client binary
.PHONY: cbuild
cbuild: client/app 


##Show heop
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)