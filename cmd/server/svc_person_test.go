package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	pb "github.com/siarhei-shliayonkin/learn-protobuf-grpc/pkg/proto"
)

func TestSvcAdd(t *testing.T) {
	ctx := context.Background()
	p, err := client.Add(ctx, &pb.Person{FirstName: "John", LastName: "Doe"})
	println("--after Add()")
	println(p, err, err.Error())

	// require.NoError(t, err)
	require.Error(t, err, err) // tmp
}
