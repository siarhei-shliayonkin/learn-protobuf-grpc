package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/siarhei-shliayonkin/learn-protobuf-grpc/pkg/proto"
)

// The client code for the gRPC client. Contains the code for streaming on both
// client-side and server-side.

const svcAddress = ":8090"

func main() {
	// Create a client connection to the server.
	conn, err := grpc.Dial(svcAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	// Get the client
	client := pb.NewPersonServiceClient(conn)

	// send to server
	_ = bulkAdd(client)
	time.Sleep(2 * time.Second)

	// receive from server
	fmt.Println("--")
	_ = bulkGet(client)
}

var testData = map[string]string{
	"John": "Smith",
	"Jane": "Smith",
	// "Mary": "Smith",
}

// bulkAdd sends bulk data to the server using a stream to the PersonServiceClient.
// It receives a client of type PersonServiceClient and returns an error if one occurs.
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

	// Close the stream and receive the response from the server.
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("error closing and receiving response: %v", err)
		return err
	}
	log.Println("stream closed")

	// Handle the response from the server.
	log.Printf("Response: %v", res)
	return nil
}

// bulkGet receives multiple Person messages from the server via a stream.
//
// client is a PersonServiceClient that contains a BulkGet method.
// It returns an error if the BulkGet method fails or if there is an error
// receiving a Person message.
// It returns nil if the BulkGet method completes successfully.
func bulkGet(client pb.PersonServiceClient) error {
	var (
		stream pb.PersonService_BulkGetClient
		err    error
	)
	log.Println("calling BulkGet")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call the BulkGet function on the client.
	stream, err = client.BulkGet(ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("error calling BulkGet: %v", err)
		return err
	}
	defer func() {
		// Close the stream.
		if err := stream.CloseSend(); err != nil {
			log.Printf("error closing stream: %v", err)
		}
		log.Println("stream closed")
	}()

	// Receive multiple Person messages from the server via the stream.
	for {
		person, err := stream.Recv()
		if err == io.EOF {
			// End of stream.
			break
		}
		if err != nil {
			log.Printf("error receiving person: %v", err)
			return err
		}

		// Handle the received Person message.
		log.Printf("Received: %v %v", person.FirstName, person.LastName)
	}

	return nil
}
