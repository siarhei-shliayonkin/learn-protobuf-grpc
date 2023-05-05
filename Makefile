# SRC_DIR  := $(shell pwd)
PKG_LIST := ./pkg/...
PROTO_DIR := ./pkg/proto
APP_PERSON := ./cmd/person/...
BIN_DIR := ./bin

.PHONY: all
all: vendor lint proto build

#
.PHONY: proto
proto:
	@echo "generating proto"
	@protoc --go_opt=paths=source_relative \
		--go_out=plugins=grpc:. \
		$(PROTO_DIR)/example.proto

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
build: bin person

.PHONY: person
person:
	@go build -o $(BIN_DIR)/person $(APP_PERSON)
