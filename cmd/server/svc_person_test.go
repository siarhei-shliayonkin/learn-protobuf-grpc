package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/siarhei-shliayonkin/learn-protobuf-grpc/pkg/proto"
)

func TestSvc(t *testing.T) {
	ctx := context.Background()
	protoExpected := &pb.Person{FirstName: "John", LastName: "Doe"}
	// Add
	protoActual, err := client.Add(ctx, protoExpected)
	require.NoError(t, err)
	require.True(t, proto.Equal(protoExpected, protoActual))
	// Get
	protoActual, err = client.Get(ctx, &pb.PersonRequest{FirstName: "John"})
	require.NoError(t, err)
	require.True(t, proto.Equal(protoExpected, protoActual))
	// List
	personList, err := client.List(ctx, &emptypb.Empty{})
	require.NoError(t, err)
	p := personList.Person
	require.Equal(t, 1, len(p))
	require.Equal(t, "John", p[0].FirstName)
	// Delete
	_, err = client.Delete(ctx, &pb.PersonRequest{FirstName: "John"})
	require.NoError(t, err)
}
