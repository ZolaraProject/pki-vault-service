// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: pkivault.proto

package pkivaultrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PkiVaultService_GetUsers_FullMethodName       = "/pkivaultrpc.pkiVaultService/GetUsers"
	PkiVaultService_GetUserProfile_FullMethodName = "/pkivaultrpc.pkiVaultService/GetUserProfile"
	PkiVaultService_CreateUser_FullMethodName     = "/pkivaultrpc.pkiVaultService/CreateUser"
	PkiVaultService_UpdateUser_FullMethodName     = "/pkivaultrpc.pkiVaultService/UpdateUser"
	PkiVaultService_DeleteUser_FullMethodName     = "/pkivaultrpc.pkiVaultService/DeleteUser"
)

// PkiVaultServiceClient is the client API for PkiVaultService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PkiVaultServiceClient interface {
	GetUsers(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserList, error)
	GetUserProfile(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserInList, error)
	CreateUser(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*Response, error)
	UpdateUser(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*Response, error)
	DeleteUser(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*Response, error)
}

type pkiVaultServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPkiVaultServiceClient(cc grpc.ClientConnInterface) PkiVaultServiceClient {
	return &pkiVaultServiceClient{cc}
}

func (c *pkiVaultServiceClient) GetUsers(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserList)
	err := c.cc.Invoke(ctx, PkiVaultService_GetUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pkiVaultServiceClient) GetUserProfile(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserInList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserInList)
	err := c.cc.Invoke(ctx, PkiVaultService_GetUserProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pkiVaultServiceClient) CreateUser(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, PkiVaultService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pkiVaultServiceClient) UpdateUser(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, PkiVaultService_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pkiVaultServiceClient) DeleteUser(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, PkiVaultService_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PkiVaultServiceServer is the server API for PkiVaultService service.
// All implementations must embed UnimplementedPkiVaultServiceServer
// for forward compatibility.
type PkiVaultServiceServer interface {
	GetUsers(context.Context, *UserRequest) (*UserList, error)
	GetUserProfile(context.Context, *GetUserRequest) (*UserInList, error)
	CreateUser(context.Context, *UserCreateRequest) (*Response, error)
	UpdateUser(context.Context, *UserUpdateRequest) (*Response, error)
	DeleteUser(context.Context, *UserDeleteRequest) (*Response, error)
	mustEmbedUnimplementedPkiVaultServiceServer()
}

// UnimplementedPkiVaultServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPkiVaultServiceServer struct{}

func (UnimplementedPkiVaultServiceServer) GetUsers(context.Context, *UserRequest) (*UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedPkiVaultServiceServer) GetUserProfile(context.Context, *GetUserRequest) (*UserInList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedPkiVaultServiceServer) CreateUser(context.Context, *UserCreateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedPkiVaultServiceServer) UpdateUser(context.Context, *UserUpdateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedPkiVaultServiceServer) DeleteUser(context.Context, *UserDeleteRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedPkiVaultServiceServer) mustEmbedUnimplementedPkiVaultServiceServer() {}
func (UnimplementedPkiVaultServiceServer) testEmbeddedByValue()                         {}

// UnsafePkiVaultServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PkiVaultServiceServer will
// result in compilation errors.
type UnsafePkiVaultServiceServer interface {
	mustEmbedUnimplementedPkiVaultServiceServer()
}

func RegisterPkiVaultServiceServer(s grpc.ServiceRegistrar, srv PkiVaultServiceServer) {
	// If the following call pancis, it indicates UnimplementedPkiVaultServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PkiVaultService_ServiceDesc, srv)
}

func _PkiVaultService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PkiVaultServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PkiVaultService_GetUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PkiVaultServiceServer).GetUsers(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PkiVaultService_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PkiVaultServiceServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PkiVaultService_GetUserProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PkiVaultServiceServer).GetUserProfile(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PkiVaultService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PkiVaultServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PkiVaultService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PkiVaultServiceServer).CreateUser(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PkiVaultService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PkiVaultServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PkiVaultService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PkiVaultServiceServer).UpdateUser(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PkiVaultService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PkiVaultServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PkiVaultService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PkiVaultServiceServer).DeleteUser(ctx, req.(*UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PkiVaultService_ServiceDesc is the grpc.ServiceDesc for PkiVaultService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PkiVaultService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pkivaultrpc.pkiVaultService",
	HandlerType: (*PkiVaultServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUsers",
			Handler:    _PkiVaultService_GetUsers_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _PkiVaultService_GetUserProfile_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _PkiVaultService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _PkiVaultService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _PkiVaultService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkivault.proto",
}
