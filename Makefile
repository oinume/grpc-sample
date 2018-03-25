APP = grpc-sample
VENDOR_DIR = vendor
PROTO_GEN_DIR = proto-gen
GRPC_GATEWAY_REPO = github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
GO_GET ?= go get
VENDOR_DIR = vendor

all: build

.PHONY: setup
setup: install-commands

.PHONY: install-commands
install-commands:
	$(GO_GET) github.com/golang/protobuf/protoc-gen-go
	$(GO_GET) github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

.PHONY: build
build:
	go build -o bin/$(APP) github.com/oinume/grpc-sample/cmd

clean:
	${RM} bin/$(APP)


.PHONY: proto/go
proto/go:
	rm -rf $(PROTO_GEN_DIR)/go && mkdir -p $(PROTO_GEN_DIR)/go
	protoc -I/usr/local/include -I. \
  		-I$(GOPATH)/src \
  		-I$(VENDOR_DIR)/$(GRPC_GATEWAY_REPO) \
  		--go_out=plugins=grpc:$(PROTO_GEN_DIR)/go \
  		proto/api/v1/*.proto
	protoc -I/usr/local/include -I. -I$(GOPATH)/src -I$(VENDOR_DIR)/$(GRPC_GATEWAY_REPO) \
		--grpc-gateway_out=logtostderr=true:$(PROTO_GEN_DIR)/go \
		proto/api/v1/*.proto
