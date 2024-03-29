// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// CpfValidatorClient is the client API for CpfValidator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CpfValidatorClient interface {
	Validate(ctx context.Context, in *CpfRequest, opts ...grpc.CallOption) (*CpfResponse, error)
}

type cpfValidatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCpfValidatorClient(cc grpc.ClientConnInterface) CpfValidatorClient {
	return &cpfValidatorClient{cc}
}

func (c *cpfValidatorClient) Validate(ctx context.Context, in *CpfRequest, opts ...grpc.CallOption) (*CpfResponse, error) {
	out := new(CpfResponse)
	err := c.cc.Invoke(ctx, "/grpc.CpfValidator/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CpfValidatorServer is the server API for CpfValidator service.
// All implementations must embed UnimplementedCpfValidatorServer
// for forward compatibility
type CpfValidatorServer interface {
	Validate(context.Context, *CpfRequest) (*CpfResponse, error)
	mustEmbedUnimplementedCpfValidatorServer()
}

// UnimplementedCpfValidatorServer must be embedded to have forward compatible implementations.
type UnimplementedCpfValidatorServer struct {
}

func (UnimplementedCpfValidatorServer) Validate(context.Context, *CpfRequest) (*CpfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedCpfValidatorServer) mustEmbedUnimplementedCpfValidatorServer() {}

// UnsafeCpfValidatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CpfValidatorServer will
// result in compilation errors.
type UnsafeCpfValidatorServer interface {
	mustEmbedUnimplementedCpfValidatorServer()
}

func RegisterCpfValidatorServer(s grpc.ServiceRegistrar, srv CpfValidatorServer) {
	s.RegisterService(&CpfValidator_ServiceDesc, srv)
}

func _CpfValidator_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CpfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CpfValidatorServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.CpfValidator/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CpfValidatorServer).Validate(ctx, req.(*CpfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CpfValidator_ServiceDesc is the grpc.ServiceDesc for CpfValidator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CpfValidator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.CpfValidator",
	HandlerType: (*CpfValidatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validate",
			Handler:    _CpfValidator_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cpf_service.proto",
}
