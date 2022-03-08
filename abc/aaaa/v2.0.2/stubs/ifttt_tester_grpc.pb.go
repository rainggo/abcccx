// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package funnychen_server_a

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

// SignServerClient is the client API for SignServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignServerClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type signServerClient struct {
	cc grpc.ClientConnInterface
}

func NewSignServerClient(cc grpc.ClientConnInterface) SignServerClient {
	return &signServerClient{cc}
}

func (c *signServerClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/devx_sample.funnychen_server_a.SignServer/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignServerServer is the server API for SignServer service.
// All implementations must embed UnimplementedSignServerServer
// for forward compatibility
type SignServerServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedSignServerServer()
}

// UnimplementedSignServerServer must be embedded to have forward compatible implementations.
type UnimplementedSignServerServer struct {
}

func (UnimplementedSignServerServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedSignServerServer) mustEmbedUnimplementedSignServerServer() {}

// UnsafeSignServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignServerServer will
// result in compilation errors.
type UnsafeSignServerServer interface {
	mustEmbedUnimplementedSignServerServer()
}

func RegisterSignServerServer(s grpc.ServiceRegistrar, srv SignServerServer) {
	s.RegisterService(&SignServer_ServiceDesc, srv)
}

func _SignServer_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignServerServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devx_sample.funnychen_server_a.SignServer/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignServerServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SignServer_ServiceDesc is the grpc.ServiceDesc for SignServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "devx_sample.funnychen_server_a.SignServer",
	HandlerType: (*SignServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _SignServer_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ifttt_tester.proto",
}
