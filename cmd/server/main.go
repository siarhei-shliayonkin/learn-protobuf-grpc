package main

import (
	"log"
	"net"

	_ "github.com/gogo/googleapis/google/api"
	"google.golang.org/grpc"

	pb "github.com/siarhei-shliayonkin/learn-protobuf-grpc/pkg/proto"
	"github.com/siarhei-shliayonkin/learn-protobuf-grpc/pkg/svc/person"
)

const (
	port = ":8090"
)

func main() {
	s, err := initServer(port)
	if err != nil {
		log.Fatalf("failed to init server: %v", err)
	}

	// start server
	log.Printf("Listening on %s", port)
	if err := s.server.Serve(s.listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type Srv struct {
	listener net.Listener
	server   *grpc.Server
}

// initServer initializes a new gRPC server on the given port and returns a pointer
// to the created server. An error is returned if the server could not be created.
//
// port: port number the server will listen on.
//
// *Srv: pointer to the created server.
// error: error, if any, encountered during server creation.
func initServer(port string) (*Srv, error) {
	grpcServer := grpc.NewServer()
	pb.RegisterPersonServiceServer(grpcServer, person.NewService())

	listener, err := net.Listen("tcp4", port)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return nil, err
	}

	return &Srv{
		listener: listener,
		server:   grpcServer,
	}, nil
}
