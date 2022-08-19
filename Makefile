NOW = $(shell date "+%s")
PWD = $(shell pwd)
CURRENT_BRANCH = $(shell git symbolic-ref --short HEAD)

UNAME_S := $(shell uname -s)

GOOS?=linux
GOARCH?=amd64

GO?=go

IMAGE?=czht1118/api:latest

.PHONY: binary
binary:
	@echo "+ $@"
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -mod=mod \
		--ldflags "-s -w \
        	-X github.com/wpf1118/api/cmd.BuildTime=$(NOW) \
        	-X github.com/wpf1118/api/cmd.PWD=$(PWD) \
        	-X github.com/wpf1118/api/cmd.BRANCH=$(CURRENT_BRANCH) \
        " \
		-o dist/api github.com/wpf1118/api

build: binary
	$(RUNTIME) build -f build/Dockerfile -t $(IMAGE) .


push: build
	$(RUNTIME) push $(IMAGE)

.PHONY: proto
proto:
	@echo "+ $@"
	 protoc ./pkg/proto/*.proto --go_out=./

.PHONY: proto-grpc
proto-grpc:
	@echo "+ $@"
	 protoc ./pkg/proto/*.proto --go-grpc_out=./