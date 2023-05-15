package person

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/siarhei-shliayonkin/learn-protobuf-grpc/pkg/proto"
	"github.com/siarhei-shliayonkin/learn-protobuf-grpc/pkg/svc/storage"
)

// service represents the gRPC server API for Person service.
type service struct {
	pb.PersonServiceServer
	storageMap storage.Storage
}

// NewService returns a new instance of the PersonServiceServer.
func NewService() pb.PersonServiceServer {
	return &service{
		storageMap: storage.New(),
	}
}

// Add adds a new person to the service.
//
// ctx: The context for the function.
// person: The person to add to the service.
// returns: The added person and an error.
func (s *service) Add(ctx context.Context, person *pb.Person) (*pb.Person, error) {
	s.storageMap.Add(person.GetFirstName(), person.GetLastName())
	return person, status.Error(codes.OK, "OK")
}

// List returns a list of persons.
//
// Takes a context and an empty protocol buffer as input.
// Returns a PersonList protocol buffer and an error.
func (s *service) List(ctx context.Context, empty *emptypb.Empty) (*pb.PersonList, error) {
	l := s.storageMap.List()
	pl := new(pb.PersonList)
	pl.Person = make([]*pb.Person, 0, len(l))
	for k, v := range l {
		pl.Person = append(pl.Person, &pb.Person{FirstName: k, LastName: v})
	}

	return pl, status.Error(codes.OK, "OK")
}

func (s *service) Get(ctx context.Context, in *pb.PersonRequest) (*pb.Person, error) {
	v, err := s.storageMap.Get(in.FirstName)
	if err != nil {
		log.Println("can not get item from storage: ", err)
		return nil, err
	}
	return &pb.Person{FirstName: in.FirstName, LastName: v}, nil
}

func (s *service) Delete(ctx context.Context, in *pb.PersonRequest) (*emptypb.Empty, error) {
	_ = s.storageMap.Delete(in.FirstName)
	return &emptypb.Empty{}, nil
}

// BulkAdd adds multiple persons to the server via the PersonService_BulkAddServer stream.
// (Client-side RPC streaming)
//
// s: the service struct that has the implementation of the PersonServiceServer interface.
// PersonService_BulkAddServer: the server stream that receives the persons to be added.
// error: returns an error if the server fails to add the persons.
func (s *service) BulkAdd(stream pb.PersonService_BulkAddServer) error {
	log.Println("receive stream")
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			log.Println("EOF")
			return stream.SendAndClose(&pb.StreamResponse{Message: "Data received!"})
		}
		if err != nil {
			return err
		}

		s.storageMap.Add(data.GetFirstName(), data.GetLastName())
		log.Printf("Added %s %s\n", data.GetFirstName(), data.GetLastName())
	}
}

// BulkGet sends a stream of Person objects to the client.
// (Server-side RPC streaming)
//
// It takes an empty request and a stream, and returns an error.
// The function loops through a list of items, and for each key-value
// pair it sends a new Person object with the first name set to the key
// and the last name to the value.
func (s *service) BulkGet(in *emptypb.Empty, stream pb.PersonService_BulkGetServer) error {
	log.Println("send stream")
	itemsMap := s.storageMap.List()

	for k, v := range itemsMap {
		if err := stream.Send(&pb.Person{
			FirstName: k,
			LastName:  v,
		}); err != nil {
			return err
		}
	}

	return nil
}
