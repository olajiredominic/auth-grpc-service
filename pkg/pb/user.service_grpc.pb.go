// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: pkg/pb/user.service.proto

package pb

import (
	context "context"
	model "github.com/lerryjay/auth-grpc-service/pkg/pb/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*model.User, error)
	CreateUser(ctx context.Context, in *model.User, opts ...grpc.CallOption) (*model.User, error)
	UpdateUser(ctx context.Context, in *model.User, opts ...grpc.CallOption) (*model.User, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateUserIDImage(ctx context.Context, in *UpdateIDImageRequest, opts ...grpc.CallOption) (*UpdateIDImageResponse, error)
	UpdateUserIDNumber(ctx context.Context, in *UpdateIDNumberRequest, opts ...grpc.CallOption) (*UpdateIDNumberResponse, error)
	UpdateUserSelfie(ctx context.Context, in *UpdateSelfieRequest, opts ...grpc.CallOption) (*UpdateSelfieResponse, error)
	VerifyUser(ctx context.Context, in *VerifyUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateUserIDType(ctx context.Context, in *UpdateIDTypeRequest, opts ...grpc.CallOption) (*UpdateIDTypeResponse, error)
	UpdateUserProfilePicture(ctx context.Context, in *UpdateProfilePictureRequest, opts ...grpc.CallOption) (*model.User, error)
	UpdateUserAddress(ctx context.Context, in *UpdateUserAddressRequest, opts ...grpc.CallOption) (*model.Address, error)
	UpdateUserVerificationNames(ctx context.Context, in *UpdateUserNamesRequest, opts ...grpc.CallOption) (*UpdateUserNamesResponse, error)
	GetUserAddress(ctx context.Context, in *GetUserAddressRequest, opts ...grpc.CallOption) (*GetUserAddressResponse, error)
	GetUserStats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error)
	UpdateUserHostingStatus(ctx context.Context, in *model.User, opts ...grpc.CallOption) (*model.User, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*model.User, error) {
	out := new(model.User)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *model.User, opts ...grpc.CallOption) (*model.User, error) {
	out := new(model.User)
	err := c.cc.Invoke(ctx, "/user.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *model.User, opts ...grpc.CallOption) (*model.User, error) {
	out := new(model.User)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserIDImage(ctx context.Context, in *UpdateIDImageRequest, opts ...grpc.CallOption) (*UpdateIDImageResponse, error) {
	out := new(UpdateIDImageResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUserIDImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserIDNumber(ctx context.Context, in *UpdateIDNumberRequest, opts ...grpc.CallOption) (*UpdateIDNumberResponse, error) {
	out := new(UpdateIDNumberResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUserIDNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserSelfie(ctx context.Context, in *UpdateSelfieRequest, opts ...grpc.CallOption) (*UpdateSelfieResponse, error) {
	out := new(UpdateSelfieResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUserSelfie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) VerifyUser(ctx context.Context, in *VerifyUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/VerifyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserIDType(ctx context.Context, in *UpdateIDTypeRequest, opts ...grpc.CallOption) (*UpdateIDTypeResponse, error) {
	out := new(UpdateIDTypeResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUserIDType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserProfilePicture(ctx context.Context, in *UpdateProfilePictureRequest, opts ...grpc.CallOption) (*model.User, error) {
	out := new(model.User)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUserProfilePicture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserAddress(ctx context.Context, in *UpdateUserAddressRequest, opts ...grpc.CallOption) (*model.Address, error) {
	out := new(model.Address)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUserAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserVerificationNames(ctx context.Context, in *UpdateUserNamesRequest, opts ...grpc.CallOption) (*UpdateUserNamesResponse, error) {
	out := new(UpdateUserNamesResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUserVerificationNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserAddress(ctx context.Context, in *GetUserAddressRequest, opts ...grpc.CallOption) (*GetUserAddressResponse, error) {
	out := new(GetUserAddressResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUserAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserStats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error) {
	out := new(StatsResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUserStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserHostingStatus(ctx context.Context, in *model.User, opts ...grpc.CallOption) (*model.User, error) {
	out := new(model.User)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUserHostingStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	GetUser(context.Context, *GetUserRequest) (*model.User, error)
	CreateUser(context.Context, *model.User) (*model.User, error)
	UpdateUser(context.Context, *model.User) (*model.User, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error)
	UpdateUserIDImage(context.Context, *UpdateIDImageRequest) (*UpdateIDImageResponse, error)
	UpdateUserIDNumber(context.Context, *UpdateIDNumberRequest) (*UpdateIDNumberResponse, error)
	UpdateUserSelfie(context.Context, *UpdateSelfieRequest) (*UpdateSelfieResponse, error)
	VerifyUser(context.Context, *VerifyUserRequest) (*emptypb.Empty, error)
	UpdateUserIDType(context.Context, *UpdateIDTypeRequest) (*UpdateIDTypeResponse, error)
	UpdateUserProfilePicture(context.Context, *UpdateProfilePictureRequest) (*model.User, error)
	UpdateUserAddress(context.Context, *UpdateUserAddressRequest) (*model.Address, error)
	UpdateUserVerificationNames(context.Context, *UpdateUserNamesRequest) (*UpdateUserNamesResponse, error)
	GetUserAddress(context.Context, *GetUserAddressRequest) (*GetUserAddressResponse, error)
	GetUserStats(context.Context, *StatsRequest) (*StatsResponse, error)
	UpdateUserHostingStatus(context.Context, *model.User) (*model.User, error)
}

// UnimplementedUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*model.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) CreateUser(context.Context, *model.User) (*model.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *model.User) (*model.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserIDImage(context.Context, *UpdateIDImageRequest) (*UpdateIDImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserIDImage not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserIDNumber(context.Context, *UpdateIDNumberRequest) (*UpdateIDNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserIDNumber not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserSelfie(context.Context, *UpdateSelfieRequest) (*UpdateSelfieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserSelfie not implemented")
}
func (UnimplementedUserServiceServer) VerifyUser(context.Context, *VerifyUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserIDType(context.Context, *UpdateIDTypeRequest) (*UpdateIDTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserIDType not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserProfilePicture(context.Context, *UpdateProfilePictureRequest) (*model.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfilePicture not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserAddress(context.Context, *UpdateUserAddressRequest) (*model.Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserAddress not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserVerificationNames(context.Context, *UpdateUserNamesRequest) (*UpdateUserNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserVerificationNames not implemented")
}
func (UnimplementedUserServiceServer) GetUserAddress(context.Context, *GetUserAddressRequest) (*GetUserAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAddress not implemented")
}
func (UnimplementedUserServiceServer) GetUserStats(context.Context, *StatsRequest) (*StatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserStats not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserHostingStatus(context.Context, *model.User) (*model.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserHostingStatus not implemented")
}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*model.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*model.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserIDImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIDImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserIDImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUserIDImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserIDImage(ctx, req.(*UpdateIDImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserIDNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIDNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserIDNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUserIDNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserIDNumber(ctx, req.(*UpdateIDNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserSelfie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSelfieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserSelfie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUserSelfie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserSelfie(ctx, req.(*UpdateSelfieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_VerifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).VerifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/VerifyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).VerifyUser(ctx, req.(*VerifyUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserIDType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIDTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserIDType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUserIDType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserIDType(ctx, req.(*UpdateIDTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserProfilePicture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfilePictureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserProfilePicture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUserProfilePicture",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserProfilePicture(ctx, req.(*UpdateProfilePictureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUserAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserAddress(ctx, req.(*UpdateUserAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserVerificationNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserVerificationNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUserVerificationNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserVerificationNames(ctx, req.(*UpdateUserNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUserAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserAddress(ctx, req.(*GetUserAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUserStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserStats(ctx, req.(*StatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserHostingStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserHostingStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUserHostingStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserHostingStatus(ctx, req.(*model.User))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUsers",
			Handler:    _UserService_ListUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "UpdateUserIDImage",
			Handler:    _UserService_UpdateUserIDImage_Handler,
		},
		{
			MethodName: "UpdateUserIDNumber",
			Handler:    _UserService_UpdateUserIDNumber_Handler,
		},
		{
			MethodName: "UpdateUserSelfie",
			Handler:    _UserService_UpdateUserSelfie_Handler,
		},
		{
			MethodName: "VerifyUser",
			Handler:    _UserService_VerifyUser_Handler,
		},
		{
			MethodName: "UpdateUserIDType",
			Handler:    _UserService_UpdateUserIDType_Handler,
		},
		{
			MethodName: "UpdateUserProfilePicture",
			Handler:    _UserService_UpdateUserProfilePicture_Handler,
		},
		{
			MethodName: "UpdateUserAddress",
			Handler:    _UserService_UpdateUserAddress_Handler,
		},
		{
			MethodName: "UpdateUserVerificationNames",
			Handler:    _UserService_UpdateUserVerificationNames_Handler,
		},
		{
			MethodName: "GetUserAddress",
			Handler:    _UserService_GetUserAddress_Handler,
		},
		{
			MethodName: "GetUserStats",
			Handler:    _UserService_GetUserStats_Handler,
		},
		{
			MethodName: "UpdateUserHostingStatus",
			Handler:    _UserService_UpdateUserHostingStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/user.service.proto",
}
