package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/siarhei-shliayonkin/learn-protobuf-grpc/pkg/proto"
)

func TestSvcAdd(t *testing.T) {
	ctx := context.Background()
	protoExpected := &pb.Person{FirstName: "John", LastName: "Doe"}

	protoActual, err := client.Add(ctx, protoExpected)
	require.NoError(t, err)
	require.True(t, proto.Equal(protoExpected, protoActual))
}

func TestSvcList(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client.Add(ctx, &pb.Person{FirstName: "John", LastName: "Doe"})

	protoActual, err := client.List(ctx, &emptypb.Empty{})
	p := protoActual.Person
	require.NoError(t, err)
	require.Equal(t, 1, len(p))
	require.Equal(t, "John", p[0].FirstName)
}
