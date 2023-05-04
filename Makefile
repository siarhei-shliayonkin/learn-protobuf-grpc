# SRC_DIR  := $(shell pwd)
PKG_LIST := ./pkg/...
PROTO_DIR := ./pkg/proto

.PHONY: all
all: vendor lint proto

#
.PHONY: proto
proto:
	@echo "generating proto"
	@protoc --go_opt=paths=source_relative \
		--go_out=$(PROTO_DIR) \
		--go-grpc_out=$(PROTO_DIR) \
		--go-grpc_opt=paths=source_relative \
		--proto_path=$(PROTO_DIR) \
		example.proto

.PHONY: vendor
vendor:
	@echo "updating vendor"
	@go mod vendor
	@go mod tidy

.PHONY: lint
lint:
	@golint $(PKG_LIST)

.PHONY: fmt
fmt:
	@gofmt $(PKG_LIST)
