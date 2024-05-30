// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: pkg/pb/identity_verification.service.proto

package pb

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

// VerificationServiceClient is the client API for VerificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VerificationServiceClient interface {
	VerifyNIN(ctx context.Context, in *VerifyNINRequest, opts ...grpc.CallOption) (*VerifyNINResponse, error)
	VerifyVNIN(ctx context.Context, in *VerifyNINRequest, opts ...grpc.CallOption) (*VerifyNINResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type verificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVerificationServiceClient(cc grpc.ClientConnInterface) VerificationServiceClient {
	return &verificationServiceClient{cc}
}

func (c *verificationServiceClient) VerifyNIN(ctx context.Context, in *VerifyNINRequest, opts ...grpc.CallOption) (*VerifyNINResponse, error) {
	out := new(VerifyNINResponse)
	err := c.cc.Invoke(ctx, "/identity_verification.VerificationService/VerifyNIN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verificationServiceClient) VerifyVNIN(ctx context.Context, in *VerifyNINRequest, opts ...grpc.CallOption) (*VerifyNINResponse, error) {
	out := new(VerifyNINResponse)
	err := c.cc.Invoke(ctx, "/identity_verification.VerificationService/VerifyVNIN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verificationServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/identity_verification.VerificationService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationServiceServer is the server API for VerificationService service.
// All implementations should embed UnimplementedVerificationServiceServer
// for forward compatibility
type VerificationServiceServer interface {
	VerifyNIN(context.Context, *VerifyNINRequest) (*VerifyNINResponse, error)
	VerifyVNIN(context.Context, *VerifyNINRequest) (*VerifyNINResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

// UnimplementedVerificationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedVerificationServiceServer struct {
}

func (UnimplementedVerificationServiceServer) VerifyNIN(context.Context, *VerifyNINRequest) (*VerifyNINResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyNIN not implemented")
}
func (UnimplementedVerificationServiceServer) VerifyVNIN(context.Context, *VerifyNINRequest) (*VerifyNINResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyVNIN not implemented")
}
func (UnimplementedVerificationServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

// UnsafeVerificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VerificationServiceServer will
// result in compilation errors.
type UnsafeVerificationServiceServer interface {
	mustEmbedUnimplementedVerificationServiceServer()
}

func RegisterVerificationServiceServer(s grpc.ServiceRegistrar, srv VerificationServiceServer) {
	s.RegisterService(&VerificationService_ServiceDesc, srv)
}

func _VerificationService_VerifyNIN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyNINRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationServiceServer).VerifyNIN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity_verification.VerificationService/VerifyNIN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationServiceServer).VerifyNIN(ctx, req.(*VerifyNINRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerificationService_VerifyVNIN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyNINRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationServiceServer).VerifyVNIN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity_verification.VerificationService/VerifyVNIN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationServiceServer).VerifyVNIN(ctx, req.(*VerifyNINRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerificationService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity_verification.VerificationService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VerificationService_ServiceDesc is the grpc.ServiceDesc for VerificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VerificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "identity_verification.VerificationService",
	HandlerType: (*VerificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyNIN",
			Handler:    _VerificationService_VerifyNIN_Handler,
		},
		{
			MethodName: "VerifyVNIN",
			Handler:    _VerificationService_VerifyVNIN_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _VerificationService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/identity_verification.service.proto",
}
