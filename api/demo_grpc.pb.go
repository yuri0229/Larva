// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// DemoGrpcClient is the client API for DemoGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DemoGrpcClient interface {
	Detail(ctx context.Context, in *Req, opts ...grpc.CallOption) (*ItemResp, error)
}

type demoGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewDemoGrpcClient(cc grpc.ClientConnInterface) DemoGrpcClient {
	return &demoGrpcClient{cc}
}

func (c *demoGrpcClient) Detail(ctx context.Context, in *Req, opts ...grpc.CallOption) (*ItemResp, error) {
	out := new(ItemResp)
	err := c.cc.Invoke(ctx, "/api.DemoGrpc/Detail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoGrpcServer is the server API for DemoGrpc service.
// All implementations must embed UnimplementedDemoGrpcServer
// for forward compatibility
type DemoGrpcServer interface {
	Detail(context.Context, *Req) (*ItemResp, error)
	mustEmbedUnimplementedDemoGrpcServer()
}

// UnimplementedDemoGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedDemoGrpcServer struct {
}

func (UnimplementedDemoGrpcServer) Detail(context.Context, *Req) (*ItemResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}
func (UnimplementedDemoGrpcServer) mustEmbedUnimplementedDemoGrpcServer() {}

// UnsafeDemoGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DemoGrpcServer will
// result in compilation errors.
type UnsafeDemoGrpcServer interface {
	mustEmbedUnimplementedDemoGrpcServer()
}

func RegisterDemoGrpcServer(s grpc.ServiceRegistrar, srv DemoGrpcServer) {
	s.RegisterService(&DemoGrpc_ServiceDesc, srv)
}

func _DemoGrpc_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoGrpcServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DemoGrpc/Detail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoGrpcServer).Detail(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// DemoGrpc_ServiceDesc is the grpc.ServiceDesc for DemoGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DemoGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.DemoGrpc",
	HandlerType: (*DemoGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Detail",
			Handler:    _DemoGrpc_Detail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}
