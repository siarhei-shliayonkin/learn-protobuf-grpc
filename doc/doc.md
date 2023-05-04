# Explain protobuf/gRPC

## Preinstall

- Proto-compiler (/generator) and grpc package.

```bash
go install github.com/golang/protobuf/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get -u google.golang.org/grpc
```

File: `./pkg/proto/example.proto`

```proto
syntax = "proto3";
import "google/protobuf/empty.proto";

package learn_protobuf_grpc;
option go_package = "github.com/siarhei-shliayonkin/learn-protobuf-grpc";

message Person {
    string first_name = 1;
    string last_name = 2;
}

message PersonList {
    repeated Person person = 1;
}

service PersonService {
    rpc Add (Person) returns (Person);
    rpc List (google.protobuf.Empty) returns (PersonList);
}
```

Makefile:

```make
.PHONY: proto
proto:
 @echo "generating proto"
 @protoc \
  --proto_path=$(PROTO_DIR) \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  --go_out=$(PROTO_DIR) \
  --go-grpc_out=$(PROTO_DIR) \
  example.proto
```

```bash
make proto

...
├── pkg
│   └── proto
│       ├── example_grpc.pb.go
│       ├── example.pb.go
│       └── example.proto
```
