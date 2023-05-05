package main

import (
	"net/http"

	_ "github.com/gogo/googleapis/google/api"
	"google.golang.org/grpc"

	pb "github.com/siarhei-shliayonkin/learn-protobuf-grpc/pkg/proto"
	"github.com/siarhei-shliayonkin/learn-protobuf-grpc/pkg/svc/person"
)

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterPersonServiceServer(grpcServer, person.NewService())

	// create gRPC gateway server
	// gwMux := runtime.NewServeMux()
	// err := pb.Register MyServiceHandlerFromEndpoint(
	// 	context.Background(),
	// 	gwMux,
	// 	"localhost:8080",
	// 	[]grpc.DialOption{grpc.WithInsecure()},
	// )
	// if err != nil {
	// 	log.Fatalf("failed to register gRPC gateway: %v", err)
	// }

	// // create HTTP server
	// httpMux := http.NewServeMux()
	// httpMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
	// 		gwMux.ServeHTTP(w, r)
	// 		return
	// 	}
	// 	http.DefaultServeMux.ServeHTTP(w, r)
	// })

	// // start server
	// fmt.Println("Starting server on :8080")
	// if err := http.ListenAndServe(":8080", httpMux); err != nil {
	// 	log.Fatalf("failed to start server: %v", err)
	// }
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	// handle HTTP requests
}
