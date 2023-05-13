package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/siarhei-shliayonkin/learn-protobuf-grpc/pkg/proto"
)

// The client code for the gRPC server.
// Contains the code for gRPC streaming.

const svcAddress = ":8090"

func main() {
	// Create a client connection to the server.
	conn, err := grpc.Dial(svcAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPersonServiceClient(conn)

	_ = bulkAdd(client)
	// _ = bulkGet(client)
}

var testData = map[string]string{
	"John": "Smith",
	"Jane": "Smith",
}

func bulkAdd(client pb.PersonServiceClient) error {
	var (
		stream pb.PersonService_BulkAddClient
		err    error
	)
	log.Println("calling BulkAdd")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a stream by invoking the client.
	stream, err = client.BulkAdd(ctx)
	if err != nil {
		log.Printf("error calling BulkAdd: %v", err)
		return err
	}

	// graceful close stream
	defer func() {
		if err != nil {
			stream.CloseSend()
		}
	}()

	// send data to the server
	for k, v := range testData {
		log.Printf("sending Person item: %s %s", k, v)
		if err := stream.Send(&pb.Person{
			FirstName: k,
			LastName:  v,
		}); err != nil {
			log.Printf("error sending Person item: %v", err)
			return err
		}
	}

	// Close stream and receive the response from the server.
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("error closing and receiving response: %v", err)
		return err
	}

	// Handle the response from the server.
	log.Printf("Response: %v", res)
	return nil
}

// func bulkGet(client pb.PersonServiceClient) error {
// 	var (
// 		stream pb.PersonService_BulkGetClient
// 		err    error
// 	)
// 	log.Println("calling BulkGet")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// }
