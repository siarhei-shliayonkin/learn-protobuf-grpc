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

## gRPC advantages

gRPC offers several advantages over HTTP-based communication protocols like REST. Some of the main advantages are:

1. Faster performance: gRPC is built on top of HTTP/2, which supports multiplexed streams, server-side streaming, client-side streaming, and flow control. This makes gRPC more efficient in terms of network usage and allows it to deliver faster performance than REST.

2. Better resource utilization: gRPC uses Protocol Buffers for data serialization, which is a binary format. This means that it can transmit data faster and with lower overhead than REST, which often uses JSON or XML. Additionally, gRPC typically requires fewer CPU cycles, which translates to better resource utilization.

3. Strongly typed contracts: gRPC uses Protocol Buffers to define its APIs, which means that it has strong typing and well-defined contracts. This makes it easier to maintain and evolve the API over time, which is particularly important for large and complex systems.

4. Support for bidirectional streaming: gRPC supports bidirectional streaming, which allows clients and servers to send and receive multiple messages over a single connection. This enables real-time communication and can be useful in building interactive applications and systems.

5. Supports multiple programming languages: gRPC has ready-to-use support for many programming languages, such as Java, Python, C++, and Go, among others.

Overall, gRPC offers many advantages over HTTP-based communication protocols like REST, making it a solid choice for building distributed systems that require high performance, low latency, and bidirectional communication.

## potential disadvantages of using gRPC over HTTP

Sure, here are some potential disadvantages of using gRPC over HTTP:

1. Limited backward compatibility: gRPC relies heavily on Protocol Buffers for API definition, which can make it harder to evolve your API over time or to maintain backward compatibility with previous versions. This can be a challenge if you have clients or services relying on an older version of your API.

2. Higher learning curve: Compared to REST, gRPC has a steeper learning curve, especially if you're not familiar with the Protocol Buffers format or the gRPC ecosystem. This can make it harder to onboard new developers or integrate with existing systems.

3. Limited availability of third-party tools: While gRPC is gaining popularity, it might still be harder to find third-party tools and libraries that support it compared to REST. This could be an issue if you need to integrate with other systems or services that don't support gRPC.

4. Limited browser support: While some modern web browsers support gRPC, it's not as widely supported as HTTP and REST, which could be an issue if you need to build client-side applications that consume your API.

5. Increased complexity: gRPC has many powerful features, such as bidirectional streaming, error handling, and flow control, which can make it more powerful and efficient than REST. However, these features also increase the complexity of your system, which can make it harder to develop, test, and debug.

## Here are some popular use cases and applications that use gRPC bidirectional streaming

1. Real-time communication: Bidirectional streaming is well-suited for real-time communication scenarios, such as chat applications, multiplayer games, and collaboration tools.

2. Distributed computing: Bidirectional streaming can be used to facilitate distributed computing workflows, such as task processing and job orchestration, where multiple nodes need to communicate with each other in real-time.

3. IoT and edge computing: Bidirectional streaming enables devices and sensors to communicate with cloud-based services in real-time, which can be useful in IoT and edge computing scenarios where low latency and fast response times are important.

4. Financial services: Bidirectional streaming can be used to support real-time trading platforms and risk management systems in the financial services industry.

5. Video and audio streaming: Bidirectional streaming can be used to support video and audio streaming applications, where data flows in both directions and the client-server relationship is more dynamic.

Overall, bidirectional streaming is a powerful feature of gRPC that is well-suited for a wide range of applications and use cases. It enables real-time communication and can be used to build efficient and scalable distributed systems.

## Examples of well-known apps that use bidirectional streaming

There are several well-known apps that use bidirectional streaming in gRPC. For example:

1. Google Cloud Spanner: A globally distributed, strongly consistent relational database that uses gRPC and bidirectional streaming to provide high performance and low latency.

2. Netflix: Netflix uses gRPC for its data transfer layer, which includes bidirectional streaming to support real-time communication between the client and server.

3. Uber: Uber uses gRPC and bidirectional streaming for its backend services to enable real-time communication between components.

4. Square: Square uses bidirectional streaming in gRPC to build real-time financial systems that require low latency and high throughput.

5. The Envoy Proxy: Envoy is a popular open-source service proxy that uses bidirectional streaming in gRPC to provide efficient and scalable communication between microservices.

## Connections over http

The HTTP protocol does not necessarily establish a new connection each time a client accesses the server. In HTTP/1.0, connections were typically short-lived and a new connection would be created for each request/response pair, but this approach was inefficient and led to slower performance.

In HTTP/1.1, keep-alive connections were introduced, which allowed connections to be reused for multiple request/response pairs. This approach improved performance by reducing the overhead of establishing new connections for each request/response pair.

HTTP/2 further improved on connection management by using a multiplexed model that allows multiple requests and responses to be sent over a single connection simultaneously. This approach further reduces the overhead of connection establishment and improves performance.

So, whether a new connection is established each time a client accesses the server depends on the version of HTTP being used and the connection settings.

## How popular is the HTTP/2 protocol? How widely is it used? Which browsers and applications support it?

HTTP/2 is a widely used protocol that has gained popularity since its introduction in 2015. Most major web browsers, including Chrome, Firefox, Safari, and Edge, support HTTP/2, with about 97% of web browsers having the capability. Additionally, many web servers, such as Apache and NGINX, support HTTP/2.

HTTP/2 brings many improvements in terms of performance, including better use of available network resources through multiplexing and header compression. These improvements make HTTP/2 well-suited for modern web applications that require high performance and low latency.

According to recent statistics, about 39% of the top 10 million websites support HTTP/2, indicating that it is becoming increasingly popular for large-scale web applications. However, it is worth noting that some legacy web applications may not yet support HTTP/2, especially those that rely on older versions of web servers or browsers.

Overall, HTTP/2 is a widely adopted protocol that provides significant performance benefits for modern web applications. It is supported by most major web browsers and web servers, and it is likely to become even more popular in the coming years as more web applications adopt it for improved performance.
