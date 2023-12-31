// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: proto/timesquared/timesquared.proto

package timesquared

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
	TimeSquaredService_HealthCheck_FullMethodName = "/timesquared.TimeSquaredService/HealthCheck"
)

// TimeSquaredServiceClient is the client API for TimeSquaredService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeSquaredServiceClient interface {
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type timeSquaredServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeSquaredServiceClient(cc grpc.ClientConnInterface) TimeSquaredServiceClient {
	return &timeSquaredServiceClient{cc}
}

func (c *timeSquaredServiceClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, TimeSquaredService_HealthCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeSquaredServiceServer is the server API for TimeSquaredService service.
// All implementations must embed UnimplementedTimeSquaredServiceServer
// for forward compatibility
type TimeSquaredServiceServer interface {
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	mustEmbedUnimplementedTimeSquaredServiceServer()
}

// UnimplementedTimeSquaredServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTimeSquaredServiceServer struct {
}

func (UnimplementedTimeSquaredServiceServer) HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedTimeSquaredServiceServer) mustEmbedUnimplementedTimeSquaredServiceServer() {}

// UnsafeTimeSquaredServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeSquaredServiceServer will
// result in compilation errors.
type UnsafeTimeSquaredServiceServer interface {
	mustEmbedUnimplementedTimeSquaredServiceServer()
}

func RegisterTimeSquaredServiceServer(s grpc.ServiceRegistrar, srv TimeSquaredServiceServer) {
	s.RegisterService(&TimeSquaredService_ServiceDesc, srv)
}

func _TimeSquaredService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSquaredServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSquaredService_HealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSquaredServiceServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TimeSquaredService_ServiceDesc is the grpc.ServiceDesc for TimeSquaredService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimeSquaredService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "timesquared.TimeSquaredService",
	HandlerType: (*TimeSquaredServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _TimeSquaredService_HealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/timesquared/timesquared.proto",
}
