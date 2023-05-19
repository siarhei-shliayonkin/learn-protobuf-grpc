SRC_DIR  := $(shell pwd)
PKG_LIST := ./pkg/...
PROTO_DIR := pkg/proto
APP_SERVER := ./cmd/server/...
APP_CLIENT := ./cmd/client/...
BIN_DIR := ./bin

PROTOBUF_ARGS =	 -I=. -I=$(SRC_DIR)/vendor -I=$(GOPATH)/src/github.com/googleapis/googleapis
PROTOBUF_ARGS += --go_out=. --go_opt paths=source_relative
PROTOBUF_ARGS += --go-grpc_out=. --go-grpc_opt paths=source_relative
PROTOBUF_ARGS += --grpc-gateway_out=. --grpc-gateway_opt paths=source_relative
PROTOBUF_ARGS += --swagger_out=logtostderr=true:.

.PHONY: all
all: proto vendor lint build test

# needs to install once for generating .proto with rpc option
.PHONY: googleapis
googleapis:
	cd ${GOPATH}/src/github.com && mkdir -p googleapis/googleapis && cd googleapis/googleapis && \
		git init && git remote add origin https://github.com/googleapis/googleapis && git fetch && \
		git checkout origin/master -- *.proto

.PHONY: proto
proto:
	@echo "generating proto"
	@protoc $(PROTOBUF_ARGS) $(PROTO_DIR)/person.proto

.PHONY: vendor
vendor:
	@echo "updating vendor"
	@go mod tidy
	@go mod vendor

.PHONY: lint
lint:
	@golint $(PKG_LIST)

.PHONY: fmt
fmt:
	@gofmt $(PKG_LIST)

.PHONY: bin
bin:
	@mkdir -p $(BIN_DIR)

.PHONY: build
build: bin build-server build-client

.PHONY: build-server
build-server:
	@echo "building server"
	@go build -o $(BIN_DIR)/server $(APP_SERVER)

.PHONY: build-client
build-client:
	@echo "building client"
	@go build -o $(BIN_DIR)/client $(APP_CLIENT)

.PHONY: test
test:
	@echo "run tests"
	@go test \
		./cmd/server/... \
		./pkg/...
