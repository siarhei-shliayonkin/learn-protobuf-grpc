PKG_LIST := ./pkg/...

.PHONY: all
all: vendor lint proto

#
.PHONY: proto
proto:
	@echo "building proto"
	@protoc --go_out=. \
		--go_opt=paths=source_relative pkg/proto/example.proto

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
