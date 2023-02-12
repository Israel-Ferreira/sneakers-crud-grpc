// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/sneaker.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SneakerServiceClient is the client API for SneakerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SneakerServiceClient interface {
	GetAll(ctx context.Context, in *SneakerGetAllReq, opts ...grpc.CallOption) (SneakerService_GetAllClient, error)
	GetById(ctx context.Context, in *SneakerGetByIdReq, opts ...grpc.CallOption) (*Sneaker, error)
	DeleteById(ctx context.Context, in *SneakerGetByIdReq, opts ...grpc.CallOption) (*DeleteSneakerMsg, error)
	AddSneaker(ctx context.Context, in *SneakerReq, opts ...grpc.CallOption) (*Sneaker, error)
	UpdateSneaker(ctx context.Context, in *SneakerReq, opts ...grpc.CallOption) (*UpdateSneakerMsg, error)
	GetAvailableSizes(ctx context.Context, in *SneakerGetByIdReq, opts ...grpc.CallOption) (SneakerService_GetAvailableSizesClient, error)
}

type sneakerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSneakerServiceClient(cc grpc.ClientConnInterface) SneakerServiceClient {
	return &sneakerServiceClient{cc}
}

func (c *sneakerServiceClient) GetAll(ctx context.Context, in *SneakerGetAllReq, opts ...grpc.CallOption) (SneakerService_GetAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &SneakerService_ServiceDesc.Streams[0], "/sneakers.SneakerService/GetAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &sneakerServiceGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SneakerService_GetAllClient interface {
	Recv() (*Sneaker, error)
	grpc.ClientStream
}

type sneakerServiceGetAllClient struct {
	grpc.ClientStream
}

func (x *sneakerServiceGetAllClient) Recv() (*Sneaker, error) {
	m := new(Sneaker)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sneakerServiceClient) GetById(ctx context.Context, in *SneakerGetByIdReq, opts ...grpc.CallOption) (*Sneaker, error) {
	out := new(Sneaker)
	err := c.cc.Invoke(ctx, "/sneakers.SneakerService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sneakerServiceClient) DeleteById(ctx context.Context, in *SneakerGetByIdReq, opts ...grpc.CallOption) (*DeleteSneakerMsg, error) {
	out := new(DeleteSneakerMsg)
	err := c.cc.Invoke(ctx, "/sneakers.SneakerService/DeleteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sneakerServiceClient) AddSneaker(ctx context.Context, in *SneakerReq, opts ...grpc.CallOption) (*Sneaker, error) {
	out := new(Sneaker)
	err := c.cc.Invoke(ctx, "/sneakers.SneakerService/AddSneaker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sneakerServiceClient) UpdateSneaker(ctx context.Context, in *SneakerReq, opts ...grpc.CallOption) (*UpdateSneakerMsg, error) {
	out := new(UpdateSneakerMsg)
	err := c.cc.Invoke(ctx, "/sneakers.SneakerService/UpdateSneaker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sneakerServiceClient) GetAvailableSizes(ctx context.Context, in *SneakerGetByIdReq, opts ...grpc.CallOption) (SneakerService_GetAvailableSizesClient, error) {
	stream, err := c.cc.NewStream(ctx, &SneakerService_ServiceDesc.Streams[1], "/sneakers.SneakerService/GetAvailableSizes", opts...)
	if err != nil {
		return nil, err
	}
	x := &sneakerServiceGetAvailableSizesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SneakerService_GetAvailableSizesClient interface {
	Recv() (*AvailableSize, error)
	grpc.ClientStream
}

type sneakerServiceGetAvailableSizesClient struct {
	grpc.ClientStream
}

func (x *sneakerServiceGetAvailableSizesClient) Recv() (*AvailableSize, error) {
	m := new(AvailableSize)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SneakerServiceServer is the server API for SneakerService service.
// All implementations must embed UnimplementedSneakerServiceServer
// for forward compatibility
type SneakerServiceServer interface {
	GetAll(*SneakerGetAllReq, SneakerService_GetAllServer) error
	GetById(context.Context, *SneakerGetByIdReq) (*Sneaker, error)
	DeleteById(context.Context, *SneakerGetByIdReq) (*DeleteSneakerMsg, error)
	AddSneaker(context.Context, *SneakerReq) (*Sneaker, error)
	UpdateSneaker(context.Context, *SneakerReq) (*UpdateSneakerMsg, error)
	GetAvailableSizes(*SneakerGetByIdReq, SneakerService_GetAvailableSizesServer) error
	mustEmbedUnimplementedSneakerServiceServer()
}

// UnimplementedSneakerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSneakerServiceServer struct {
}

func (UnimplementedSneakerServiceServer) GetAll(*SneakerGetAllReq, SneakerService_GetAllServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedSneakerServiceServer) GetById(context.Context, *SneakerGetByIdReq) (*Sneaker, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedSneakerServiceServer) DeleteById(context.Context, *SneakerGetByIdReq) (*DeleteSneakerMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
}
func (UnimplementedSneakerServiceServer) AddSneaker(context.Context, *SneakerReq) (*Sneaker, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSneaker not implemented")
}
func (UnimplementedSneakerServiceServer) UpdateSneaker(context.Context, *SneakerReq) (*UpdateSneakerMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSneaker not implemented")
}
func (UnimplementedSneakerServiceServer) GetAvailableSizes(*SneakerGetByIdReq, SneakerService_GetAvailableSizesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAvailableSizes not implemented")
}
func (UnimplementedSneakerServiceServer) mustEmbedUnimplementedSneakerServiceServer() {}

// UnsafeSneakerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SneakerServiceServer will
// result in compilation errors.
type UnsafeSneakerServiceServer interface {
	mustEmbedUnimplementedSneakerServiceServer()
}

func RegisterSneakerServiceServer(s grpc.ServiceRegistrar, srv SneakerServiceServer) {
	s.RegisterService(&SneakerService_ServiceDesc, srv)
}

func _SneakerService_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SneakerGetAllReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SneakerServiceServer).GetAll(m, &sneakerServiceGetAllServer{stream})
}

type SneakerService_GetAllServer interface {
	Send(*Sneaker) error
	grpc.ServerStream
}

type sneakerServiceGetAllServer struct {
	grpc.ServerStream
}

func (x *sneakerServiceGetAllServer) Send(m *Sneaker) error {
	return x.ServerStream.SendMsg(m)
}

func _SneakerService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SneakerGetByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SneakerServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sneakers.SneakerService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SneakerServiceServer).GetById(ctx, req.(*SneakerGetByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SneakerService_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SneakerGetByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SneakerServiceServer).DeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sneakers.SneakerService/DeleteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SneakerServiceServer).DeleteById(ctx, req.(*SneakerGetByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SneakerService_AddSneaker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SneakerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SneakerServiceServer).AddSneaker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sneakers.SneakerService/AddSneaker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SneakerServiceServer).AddSneaker(ctx, req.(*SneakerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SneakerService_UpdateSneaker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SneakerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SneakerServiceServer).UpdateSneaker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sneakers.SneakerService/UpdateSneaker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SneakerServiceServer).UpdateSneaker(ctx, req.(*SneakerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SneakerService_GetAvailableSizes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SneakerGetByIdReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SneakerServiceServer).GetAvailableSizes(m, &sneakerServiceGetAvailableSizesServer{stream})
}

type SneakerService_GetAvailableSizesServer interface {
	Send(*AvailableSize) error
	grpc.ServerStream
}

type sneakerServiceGetAvailableSizesServer struct {
	grpc.ServerStream
}

func (x *sneakerServiceGetAvailableSizesServer) Send(m *AvailableSize) error {
	return x.ServerStream.SendMsg(m)
}

// SneakerService_ServiceDesc is the grpc.ServiceDesc for SneakerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SneakerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sneakers.SneakerService",
	HandlerType: (*SneakerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _SneakerService_GetById_Handler,
		},
		{
			MethodName: "DeleteById",
			Handler:    _SneakerService_DeleteById_Handler,
		},
		{
			MethodName: "AddSneaker",
			Handler:    _SneakerService_AddSneaker_Handler,
		},
		{
			MethodName: "UpdateSneaker",
			Handler:    _SneakerService_UpdateSneaker_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAll",
			Handler:       _SneakerService_GetAll_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAvailableSizes",
			Handler:       _SneakerService_GetAvailableSizes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/sneaker.proto",
}
