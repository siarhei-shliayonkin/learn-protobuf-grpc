package main

import (
	"log"
	"net"
	"net/http"

	_ "github.com/gogo/googleapis/google/api"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"

	pb "github.com/siarhei-shliayonkin/learn-protobuf-grpc/pkg/proto"
	"github.com/siarhei-shliayonkin/learn-protobuf-grpc/pkg/svc/person"
)

const (
	port = ":8090"
)

func main() {
	// serve swagger via http
	file := "./pkg/proto/person.swagger.json"
	http.HandleFunc("/swagger-ui", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		http.ServeFile(w, r, file)
	})
	go func() {
		log.Println("swagger server ready on :8081")
		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			log.Fatal(err)
		}
		// run external swagger UI
		// docker run --name sw --rm -p 80:8080 -e SWAGGER_JSON_URL=http://127.0.0.1:8081/swagger-ui swaggerapi/swagger-ui
	}()

	s, err := initServer(port)
	if err != nil {
		log.Fatalf("failed to init server: %v", err)
	}
	// start gRPC server
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
