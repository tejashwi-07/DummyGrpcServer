// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: proto/indriyas/indriyas.proto

package indriyas

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

const (
	IndriyasService_Healthcheck_FullMethodName = "/indriyas.IndriyasService/Healthcheck"
)

// IndriyasServiceClient is the client API for IndriyasService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IndriyasServiceClient interface {
	// Healthcheck operation
	Healthcheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type indriyasServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIndriyasServiceClient(cc grpc.ClientConnInterface) IndriyasServiceClient {
	return &indriyasServiceClient{cc}
}

func (c *indriyasServiceClient) Healthcheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, IndriyasService_Healthcheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndriyasServiceServer is the server API for IndriyasService service.
// All implementations must embed UnimplementedIndriyasServiceServer
// for forward compatibility
type IndriyasServiceServer interface {
	// Healthcheck operation
	Healthcheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	mustEmbedUnimplementedIndriyasServiceServer()
}

// UnimplementedIndriyasServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIndriyasServiceServer struct {
}

func (UnimplementedIndriyasServiceServer) Healthcheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthcheck not implemented")
}
func (UnimplementedIndriyasServiceServer) mustEmbedUnimplementedIndriyasServiceServer() {}

// UnsafeIndriyasServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IndriyasServiceServer will
// result in compilation errors.
type UnsafeIndriyasServiceServer interface {
	mustEmbedUnimplementedIndriyasServiceServer()
}

func RegisterIndriyasServiceServer(s grpc.ServiceRegistrar, srv IndriyasServiceServer) {
	s.RegisterService(&IndriyasService_ServiceDesc, srv)
}

func _IndriyasService_Healthcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndriyasServiceServer).Healthcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IndriyasService_Healthcheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndriyasServiceServer).Healthcheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IndriyasService_ServiceDesc is the grpc.ServiceDesc for IndriyasService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IndriyasService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "indriyas.IndriyasService",
	HandlerType: (*IndriyasServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healthcheck",
			Handler:    _IndriyasService_Healthcheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/indriyas/indriyas.proto",
}