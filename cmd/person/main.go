package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"

	pb "github.com/siarhei-shliayonkin/learn-protobuf-grpc/pkg/proto"
	"github.com/siarhei-shliayonkin/learn-protobuf-grpc/pkg/svc/person"
)

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterPersonServiceServer(grpcServer, person.NewService())

	// create a HTTP server mux
	// mux := http.NewServeMux()
	// mux.HandleFunc("/person", healthCheckHandler)

	// create gRPC-Web server
	wrappedGrpc := grpcweb.WrapServer(grpcServer)

	// create HTTP server
	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if wrappedGrpc.IsGrpcWebRequest(r) {
			wrappedGrpc.ServeHTTP(w, r)
			return
		}
		http.DefaultServeMux.ServeHTTP(w, r)
	})

	// start server
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", httpMux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	// handle HTTP requests
}

// func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
// 	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
// 			grpcServer.ServeHTTP(w, r)
// 		} else {
// 			otherHandler.ServeHTTP(w, r)
// 		}
// 	}), &http2.Server{})
// }
