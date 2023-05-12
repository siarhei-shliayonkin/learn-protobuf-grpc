package main

import (
	"context"
	"os"
	"testing"

	"google.golang.org/grpc"

	pb "github.com/siarhei-shliayonkin/learn-protobuf-grpc/pkg/proto"
)

var client pb.PersonServiceClient

func TestMain(m *testing.M) {
	ts := startTestServer()
	var clientConnection *grpc.ClientConn
	client, clientConnection = initTestClient(ts.listener.Addr().String())

	exitCode := m.Run()
	clientConnection.Close()
	ts.server.GracefulStop()
	os.Exit(exitCode)
}

func startTestServer() *Srv {
	s, err := initServer(port)
	if err != nil {
		println("failed to init server")
		os.Exit(1)
	}

	go func() {
		if err := s.server.Serve(s.listener); err != nil {
			println("failed to serve: %v", err)
			os.Exit(1)
		}
	}()
	return s
}

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
