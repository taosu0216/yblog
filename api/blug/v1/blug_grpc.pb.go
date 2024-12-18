// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: blug/v1/blug.proto

package v1

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
	Blug_CreateNewFriendLink_FullMethodName = "/blug.v1.Blug/CreateNewFriendLink"
	Blug_GetFriendLinkList_FullMethodName   = "/blug.v1.Blug/GetFriendLinkList"
	Blug_RegisterUser_FullMethodName        = "/blug.v1.Blug/RegisterUser"
	Blug_UserLogin_FullMethodName           = "/blug.v1.Blug/UserLogin"
	Blug_UserList_FullMethodName            = "/blug.v1.Blug/UserList"
	Blug_GetArticleList_FullMethodName      = "/blug.v1.Blug/GetArticleList"
	Blug_GetArticleByTitle_FullMethodName   = "/blug.v1.Blug/GetArticleByTitle"
)

// BlugClient is the client API for Blug service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlugClient interface {
	// Friend service
	CreateNewFriendLink(ctx context.Context, in *CreateNewFriendLinkReq, opts ...grpc.CallOption) (*CreateNewFriendLinkResp, error)
	GetFriendLinkList(ctx context.Context, in *GetFriendLinkListReq, opts ...grpc.CallOption) (*GetFriendLinkListResp, error)
	// User service
	RegisterUser(ctx context.Context, in *RegisterUserReq, opts ...grpc.CallOption) (*RegisterUserResp, error)
	UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error)
	UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error)
	// Article service
	GetArticleList(ctx context.Context, in *GetArticleListReq, opts ...grpc.CallOption) (*GetArticleListResp, error)
	GetArticleByTitle(ctx context.Context, in *GetArticleByTitleReq, opts ...grpc.CallOption) (*GetArticleByTitleResp, error)
}

type blugClient struct {
	cc grpc.ClientConnInterface
}

func NewBlugClient(cc grpc.ClientConnInterface) BlugClient {
	return &blugClient{cc}
}

func (c *blugClient) CreateNewFriendLink(ctx context.Context, in *CreateNewFriendLinkReq, opts ...grpc.CallOption) (*CreateNewFriendLinkResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateNewFriendLinkResp)
	err := c.cc.Invoke(ctx, Blug_CreateNewFriendLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blugClient) GetFriendLinkList(ctx context.Context, in *GetFriendLinkListReq, opts ...grpc.CallOption) (*GetFriendLinkListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFriendLinkListResp)
	err := c.cc.Invoke(ctx, Blug_GetFriendLinkList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blugClient) RegisterUser(ctx context.Context, in *RegisterUserReq, opts ...grpc.CallOption) (*RegisterUserResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterUserResp)
	err := c.cc.Invoke(ctx, Blug_RegisterUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blugClient) UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserLoginResp)
	err := c.cc.Invoke(ctx, Blug_UserLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blugClient) UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserListResp)
	err := c.cc.Invoke(ctx, Blug_UserList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blugClient) GetArticleList(ctx context.Context, in *GetArticleListReq, opts ...grpc.CallOption) (*GetArticleListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetArticleListResp)
	err := c.cc.Invoke(ctx, Blug_GetArticleList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blugClient) GetArticleByTitle(ctx context.Context, in *GetArticleByTitleReq, opts ...grpc.CallOption) (*GetArticleByTitleResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetArticleByTitleResp)
	err := c.cc.Invoke(ctx, Blug_GetArticleByTitle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlugServer is the server API for Blug service.
// All implementations must embed UnimplementedBlugServer
// for forward compatibility.
type BlugServer interface {
	// Friend service
	CreateNewFriendLink(context.Context, *CreateNewFriendLinkReq) (*CreateNewFriendLinkResp, error)
	GetFriendLinkList(context.Context, *GetFriendLinkListReq) (*GetFriendLinkListResp, error)
	// User service
	RegisterUser(context.Context, *RegisterUserReq) (*RegisterUserResp, error)
	UserLogin(context.Context, *UserLoginReq) (*UserLoginResp, error)
	UserList(context.Context, *UserListReq) (*UserListResp, error)
	// Article service
	GetArticleList(context.Context, *GetArticleListReq) (*GetArticleListResp, error)
	GetArticleByTitle(context.Context, *GetArticleByTitleReq) (*GetArticleByTitleResp, error)
	mustEmbedUnimplementedBlugServer()
}

// UnimplementedBlugServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBlugServer struct{}

func (UnimplementedBlugServer) CreateNewFriendLink(context.Context, *CreateNewFriendLinkReq) (*CreateNewFriendLinkResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewFriendLink not implemented")
}
func (UnimplementedBlugServer) GetFriendLinkList(context.Context, *GetFriendLinkListReq) (*GetFriendLinkListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendLinkList not implemented")
}
func (UnimplementedBlugServer) RegisterUser(context.Context, *RegisterUserReq) (*RegisterUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedBlugServer) UserLogin(context.Context, *UserLoginReq) (*UserLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedBlugServer) UserList(context.Context, *UserListReq) (*UserListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (UnimplementedBlugServer) GetArticleList(context.Context, *GetArticleListReq) (*GetArticleListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleList not implemented")
}
func (UnimplementedBlugServer) GetArticleByTitle(context.Context, *GetArticleByTitleReq) (*GetArticleByTitleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleByTitle not implemented")
}
func (UnimplementedBlugServer) mustEmbedUnimplementedBlugServer() {}
func (UnimplementedBlugServer) testEmbeddedByValue()              {}

// UnsafeBlugServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlugServer will
// result in compilation errors.
type UnsafeBlugServer interface {
	mustEmbedUnimplementedBlugServer()
}

func RegisterBlugServer(s grpc.ServiceRegistrar, srv BlugServer) {
	// If the following call pancis, it indicates UnimplementedBlugServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Blug_ServiceDesc, srv)
}

func _Blug_CreateNewFriendLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewFriendLinkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlugServer).CreateNewFriendLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blug_CreateNewFriendLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlugServer).CreateNewFriendLink(ctx, req.(*CreateNewFriendLinkReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blug_GetFriendLinkList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendLinkListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlugServer).GetFriendLinkList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blug_GetFriendLinkList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlugServer).GetFriendLinkList(ctx, req.(*GetFriendLinkListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blug_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlugServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blug_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlugServer).RegisterUser(ctx, req.(*RegisterUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blug_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlugServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blug_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlugServer).UserLogin(ctx, req.(*UserLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blug_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlugServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blug_UserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlugServer).UserList(ctx, req.(*UserListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blug_GetArticleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlugServer).GetArticleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blug_GetArticleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlugServer).GetArticleList(ctx, req.(*GetArticleListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blug_GetArticleByTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleByTitleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlugServer).GetArticleByTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blug_GetArticleByTitle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlugServer).GetArticleByTitle(ctx, req.(*GetArticleByTitleReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Blug_ServiceDesc is the grpc.ServiceDesc for Blug service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Blug_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blug.v1.Blug",
	HandlerType: (*BlugServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewFriendLink",
			Handler:    _Blug_CreateNewFriendLink_Handler,
		},
		{
			MethodName: "GetFriendLinkList",
			Handler:    _Blug_GetFriendLinkList_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _Blug_RegisterUser_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _Blug_UserLogin_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _Blug_UserList_Handler,
		},
		{
			MethodName: "GetArticleList",
			Handler:    _Blug_GetArticleList_Handler,
		},
		{
			MethodName: "GetArticleByTitle",
			Handler:    _Blug_GetArticleByTitle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blug/v1/blug.proto",
}
