# Explain protobuf/gRPC

## Preinstall

- Proto-compiler (/generator) and grpc package.

```bash
go install github.com/golang/protobuf/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get -u google.golang.org/grpc

experimental: go get google.golang.org/genproto/...
```

File: `./pkg/proto/person.proto`

```proto
syntax = "proto3";
package person;

import "google/protobuf/empty.proto";

option go_package = "github.com/siarhei-shliayonkin/learn-protobuf-grpc/pkg/proto;pb";

message Person {
    string first_name = 1;
    string last_name = 2;
}

message PersonList {
    repeated Person person = 1;
}

service PersonService {
    rpc Add (Person) returns (Person) {}
    rpc List (google.protobuf.Empty) returns (PersonList) {}
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
  person.proto
```

```bash
make proto

...
├── pkg
│   └── proto
│       ├── person_grpc.pb.go
│       ├── person.pb.go
│       └── person.proto
```

Implementation of the service

```go
// service is the server API for Service service.
type service struct {
 pb.PersonServiceServer
}

// NewService returns a new instance of the PersonServiceServer.
func NewService() pb.PersonServiceServer {
 return &service{}
}
```

Starting gRPC server

```go
func startGrpcServer() {
 grpcServer := grpc.NewServer()
 pb.RegisterPersonServiceServer(grpcServer, person.NewService())

 lis, err := net.Listen("tcp", port)
 if err != nil {
  log.Fatalf("failed to listen: %v", err)
 }

 log.Printf("Listening on %s", port)
 if err := grpcServer.Serve(lis); err != nil {
  log.Fatalf("failed to serve: %v", err)
 }
}
```

Init gRPC client

```go
func initTestClient(srvAddress string) (pb.PersonServiceClient, *grpc.ClientConn) {
 clientOptions := []grpc.DialOption{grpc.WithInsecure()}
 clientConnection, err := grpc.DialContext(
  context.Background(),
  srvAddress,
  clientOptions...,
 )
 if err != nil {
  println("failed to init client connection: %v", err)
  os.Exit(1)
 }
 return pb.NewPersonServiceClient(clientConnection), clientConnection
}
```
