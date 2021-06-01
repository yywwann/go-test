# golang1.9 or latest
# 1. make help
# 2. make dep
# 3. make build
# ...

VERSION := $(shell echo $(shell cat version.go | grep "Version =" | cut -d '=' -f2))
APP_NAME := generator
BUILD_DIR := build
APP := ${BUILD_DIR}/${APP_NAME}
PKG_NAME := ${APP_NAME}_v${VERSION}
PKG := ${PKG_NAME}.tar.gz

LDFLAGS := -ldflags "-w -s -X gitlab.33.cn/chat/dtalk/version.GitCommit=`git rev-parse --short=8 HEAD`"

.PHONY: clean build pkg

clean: ## Remove previous build
	@rm -rf ${BUILD_DIR}
	@go clean

build: #checkgofmt ## Build the binary file
	GOOS=linux GOARCH=amd64 GO111MODULE=on GOPROXY=https://goproxy.cn,direct GOSUMDB="sum.golang.google.cn" go build -v -i $(LDFLAGS) -o $(APP) cmd/main.go

pkg: build ## Package
	mkdir -p ${PKG_NAME}/bin
	mkdir -p ${PKG_NAME}/etc
	cp ${APP} ${PKG_NAME}/bin/
	cp etc/*  ${PKG_NAME}/etc/
	tar zvcf ${PKG} ${PKG_NAME}
	rm -rf ${PKG_NAME}