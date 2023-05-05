SRC_DIR  := $(shell pwd)
# PROJECT_ROOT ?= github.com/siarhei-shliayonkin/learn-protobuf-grpc
PKG_LIST := ./pkg/...
PROTO_DIR := pkg/proto
APP_PERSON := ./cmd/person/...
BIN_DIR := ./bin

# PROTOC_PARAM = --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true,allow_patch_feature=false,allow_delete_body=true:. -I $(SRC_DIR)/vendor --validate_out="lang=go:."
# PROTOC_PARAM = \
# 	--go_out=plugins=grpc:. \
# 	--grpc-gateway_out=logtostderr=true,allow_patch_feature=false,allow_delete_body=true:. \
# 	-I $(SRC_DIR)/vendor \
# 	-I $(SRC_DIR)/$(PROTO_DIR) \
# 	--validate_out="lang=go:."
# PROTOC_PARAM := \
#    --go_out . --go_opt paths=source_relative \
#    --go-grpc_out . --go-grpc_opt paths=source_relative \
#    --grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
#    --grpc-gateway_out=logtostderr=true

# PROTOC_PARAM := \
#    --go_out $(SRC_DIR)/$(PROTO_DIR)  \
#    --go-grpc_out $(SRC_DIR)/$(PROTO_DIR) \
#    --grpc-gateway_out $(SRC_DIR)/$(PROTO_DIR) \
#    	-I $(SRC_DIR)/vendor \
# 	-I $(SRC_DIR)/$(PROTO_DIR) \
#    --grpc-gateway_out=logtostderr=true

PROTOBUF_ARGS =	 -I=. -I=$(SRC_DIR)/vendor -I=$(GOPATH)/src/github.com/googleapis/googleapis
PROTOBUF_ARGS += --go_out=. --go_opt paths=source_relative
PROTOBUF_ARGS += --go-grpc_out=. --go-grpc_opt paths=source_relative
PROTOBUF_ARGS += --grpc-gateway_out=. --grpc-gateway_opt paths=source_relative
# PROTOBUF_ARGS += --go-grpc_opt require_unimplemented_servers=false
# PROTOBUF_ARGS += --validate_out="lang=go:."
# PROTOBUF_ARGS += --grpc-gateway_opt logtostderr=true,allow_delete_body=true
# PROTOBUF_ARGS += --openapiv2_out=.
# PROTOBUF_ARGS += --openapiv2_opt allow_delete_body=true,atlas_patch=true,json_names_for_fields=false


# go install google.golang.org/protobuf/cmd/protoc-gen-go
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc


.PHONY: all
all: vendor lint proto build

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
build: bin person

.PHONY: person
person:
	@go build -o $(BIN_DIR)/person $(APP_PERSON)
