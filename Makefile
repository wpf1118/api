NOW = $(shell date "+%s")
PWD = $(shell pwd)
CURRENT_BRANCH = $(shell git symbolic-ref --short HEAD)

UNAME_S := $(shell uname -s)

GOOS?=linux
GOARCH?=amd64

GO?=go
RUNTIME?=docker
TAG?=v1

IMAGE?=czht1118/api:$(TAG)

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

.PHONY: build
build: binary
	$(RUNTIME) build -f build/Dockerfile -t $(IMAGE) .

.PHONY: push
push: build
	$(RUNTIME) push $(IMAGE)

.PHONY: deploy
deploy: binary
	@echo "+ $@"
	ssh -t wmf "rm -f /data/console/api && exit";
	scp ./dist/api wmf:/data/console;
	ssh -t wmf "cd /root/deploy-api && docker-compose restart api && exit";

.PHONY: pull
pull:
	@echo "+ $@"
	ssh -t wmf "cd /data/www/shop-admin && git pull"

.PHONY: proto
proto:
	@echo "+ $@"
	 protoc ./pkg/proto/*.proto --go_out=./

.PHONY: proto-grpc
proto-grpc:
	@echo "+ $@"
	 protoc ./pkg/proto/*.proto --go-grpc_out=./