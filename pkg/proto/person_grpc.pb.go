// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: pkg/proto/person.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PersonService_Add_FullMethodName     = "/person.PersonService/Add"
	PersonService_Get_FullMethodName     = "/person.PersonService/Get"
	PersonService_List_FullMethodName    = "/person.PersonService/List"
	PersonService_Delete_FullMethodName  = "/person.PersonService/Delete"
	PersonService_BulkAdd_FullMethodName = "/person.PersonService/BulkAdd"
	PersonService_BulkGet_FullMethodName = "/person.PersonService/BulkGet"
)

// PersonServiceClient is the client API for PersonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonServiceClient interface {
	Add(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error)
	Get(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*Person, error)
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PersonList, error)
	Delete(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Streaming
	BulkAdd(ctx context.Context, opts ...grpc.CallOption) (PersonService_BulkAddClient, error)
	BulkGet(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (PersonService_BulkGetClient, error)
}

type personServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonServiceClient(cc grpc.ClientConnInterface) PersonServiceClient {
	return &personServiceClient{cc}
}

func (c *personServiceClient) Add(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, PersonService_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) Get(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, PersonService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PersonList, error) {
	out := new(PersonList)
	err := c.cc.Invoke(ctx, PersonService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) Delete(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PersonService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) BulkAdd(ctx context.Context, opts ...grpc.CallOption) (PersonService_BulkAddClient, error) {
	stream, err := c.cc.NewStream(ctx, &PersonService_ServiceDesc.Streams[0], PersonService_BulkAdd_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &personServiceBulkAddClient{stream}
	return x, nil
}

type PersonService_BulkAddClient interface {
	Send(*Person) error
	CloseAndRecv() (*StreamResponse, error)
	grpc.ClientStream
}

type personServiceBulkAddClient struct {
	grpc.ClientStream
}

func (x *personServiceBulkAddClient) Send(m *Person) error {
	return x.ClientStream.SendMsg(m)
}

func (x *personServiceBulkAddClient) CloseAndRecv() (*StreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *personServiceClient) BulkGet(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (PersonService_BulkGetClient, error) {
	stream, err := c.cc.NewStream(ctx, &PersonService_ServiceDesc.Streams[1], PersonService_BulkGet_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &personServiceBulkGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PersonService_BulkGetClient interface {
	Recv() (*Person, error)
	grpc.ClientStream
}

type personServiceBulkGetClient struct {
	grpc.ClientStream
}

func (x *personServiceBulkGetClient) Recv() (*Person, error) {
	m := new(Person)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PersonServiceServer is the server API for PersonService service.
// All implementations must embed UnimplementedPersonServiceServer
// for forward compatibility
type PersonServiceServer interface {
	Add(context.Context, *Person) (*Person, error)
	Get(context.Context, *PersonRequest) (*Person, error)
	List(context.Context, *emptypb.Empty) (*PersonList, error)
	Delete(context.Context, *PersonRequest) (*emptypb.Empty, error)
	// Streaming
	BulkAdd(PersonService_BulkAddServer) error
	BulkGet(*emptypb.Empty, PersonService_BulkGetServer) error
	mustEmbedUnimplementedPersonServiceServer()
}

// UnimplementedPersonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPersonServiceServer struct {
}

func (UnimplementedPersonServiceServer) Add(context.Context, *Person) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedPersonServiceServer) Get(context.Context, *PersonRequest) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPersonServiceServer) List(context.Context, *emptypb.Empty) (*PersonList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPersonServiceServer) Delete(context.Context, *PersonRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPersonServiceServer) BulkAdd(PersonService_BulkAddServer) error {
	return status.Errorf(codes.Unimplemented, "method BulkAdd not implemented")
}
func (UnimplementedPersonServiceServer) BulkGet(*emptypb.Empty, PersonService_BulkGetServer) error {
	return status.Errorf(codes.Unimplemented, "method BulkGet not implemented")
}
func (UnimplementedPersonServiceServer) mustEmbedUnimplementedPersonServiceServer() {}

// UnsafePersonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonServiceServer will
// result in compilation errors.
type UnsafePersonServiceServer interface {
	mustEmbedUnimplementedPersonServiceServer()
}

func RegisterPersonServiceServer(s grpc.ServiceRegistrar, srv PersonServiceServer) {
	s.RegisterService(&PersonService_ServiceDesc, srv)
}

func _PersonService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PersonService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).Add(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PersonService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).Get(ctx, req.(*PersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PersonService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PersonService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).Delete(ctx, req.(*PersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_BulkAdd_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PersonServiceServer).BulkAdd(&personServiceBulkAddServer{stream})
}

type PersonService_BulkAddServer interface {
	SendAndClose(*StreamResponse) error
	Recv() (*Person, error)
	grpc.ServerStream
}

type personServiceBulkAddServer struct {
	grpc.ServerStream
}

func (x *personServiceBulkAddServer) SendAndClose(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *personServiceBulkAddServer) Recv() (*Person, error) {
	m := new(Person)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PersonService_BulkGet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PersonServiceServer).BulkGet(m, &personServiceBulkGetServer{stream})
}

type PersonService_BulkGetServer interface {
	Send(*Person) error
	grpc.ServerStream
}

type personServiceBulkGetServer struct {
	grpc.ServerStream
}

func (x *personServiceBulkGetServer) Send(m *Person) error {
	return x.ServerStream.SendMsg(m)
}

// PersonService_ServiceDesc is the grpc.ServiceDesc for PersonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PersonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "person.PersonService",
	HandlerType: (*PersonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _PersonService_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PersonService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _PersonService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PersonService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BulkAdd",
			Handler:       _PersonService_BulkAdd_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BulkGet",
			Handler:       _PersonService_BulkGet_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/proto/person.proto",
}
