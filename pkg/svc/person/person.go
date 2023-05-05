package person

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/siarhei-shliayonkin/learn-protobuf-grpc/pkg/proto"
)

// service is the server API for Service service.
type service struct {
	pb.PersonServiceServer
}

// NewService returns a new instance of the PersonServiceServer.
func NewService() pb.PersonServiceServer {
	return &service{}
}

// Add adds a new person to the service.
//
// ctx: The context for the function.
// person: The person to add to the service.
// returns: The added person and an error.
func (s *service) Add(ctx context.Context, person *pb.Person) (*pb.Person, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

// List returns a list of persons.
//
// Takes a context and an empty protocol buffer as input.
// Returns a PersonList protocol buffer and an error.
func (s *service) List(ctx context.Context, empty *emptypb.Empty) (*pb.PersonList, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
