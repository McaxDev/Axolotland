// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: gameapi.proto

package rpc

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
	GameAPI_WorldBackup_FullMethodName = "/GameAPI.GameAPI/WorldBackup"
	GameAPI_SendCmd_FullMethodName     = "/GameAPI.GameAPI/SendCmd"
	GameAPI_GameBind_FullMethodName    = "/GameAPI.GameAPI/GameBind"
	GameAPI_LoadJSON_FullMethodName    = "/GameAPI.GameAPI/LoadJSON"
)

// GameAPIClient is the client API for GameAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameAPIClient interface {
	WorldBackup(ctx context.Context, in *Server, opts ...grpc.CallOption) (*Boolean, error)
	SendCmd(ctx context.Context, in *CmdReq, opts ...grpc.CallOption) (*String, error)
	GameBind(ctx context.Context, in *BindReq, opts ...grpc.CallOption) (*Boolean, error)
	LoadJSON(ctx context.Context, in *SrvAndPath, opts ...grpc.CallOption) (*ByteSlice, error)
}

type gameAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewGameAPIClient(cc grpc.ClientConnInterface) GameAPIClient {
	return &gameAPIClient{cc}
}

func (c *gameAPIClient) WorldBackup(ctx context.Context, in *Server, opts ...grpc.CallOption) (*Boolean, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Boolean)
	err := c.cc.Invoke(ctx, GameAPI_WorldBackup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameAPIClient) SendCmd(ctx context.Context, in *CmdReq, opts ...grpc.CallOption) (*String, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(String)
	err := c.cc.Invoke(ctx, GameAPI_SendCmd_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameAPIClient) GameBind(ctx context.Context, in *BindReq, opts ...grpc.CallOption) (*Boolean, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Boolean)
	err := c.cc.Invoke(ctx, GameAPI_GameBind_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameAPIClient) LoadJSON(ctx context.Context, in *SrvAndPath, opts ...grpc.CallOption) (*ByteSlice, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ByteSlice)
	err := c.cc.Invoke(ctx, GameAPI_LoadJSON_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameAPIServer is the server API for GameAPI service.
// All implementations must embed UnimplementedGameAPIServer
// for forward compatibility.
type GameAPIServer interface {
	WorldBackup(context.Context, *Server) (*Boolean, error)
	SendCmd(context.Context, *CmdReq) (*String, error)
	GameBind(context.Context, *BindReq) (*Boolean, error)
	LoadJSON(context.Context, *SrvAndPath) (*ByteSlice, error)
	mustEmbedUnimplementedGameAPIServer()
}

// UnimplementedGameAPIServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGameAPIServer struct{}

func (UnimplementedGameAPIServer) WorldBackup(context.Context, *Server) (*Boolean, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WorldBackup not implemented")
}
func (UnimplementedGameAPIServer) SendCmd(context.Context, *CmdReq) (*String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCmd not implemented")
}
func (UnimplementedGameAPIServer) GameBind(context.Context, *BindReq) (*Boolean, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GameBind not implemented")
}
func (UnimplementedGameAPIServer) LoadJSON(context.Context, *SrvAndPath) (*ByteSlice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadJSON not implemented")
}
func (UnimplementedGameAPIServer) mustEmbedUnimplementedGameAPIServer() {}
func (UnimplementedGameAPIServer) testEmbeddedByValue()                 {}

// UnsafeGameAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameAPIServer will
// result in compilation errors.
type UnsafeGameAPIServer interface {
	mustEmbedUnimplementedGameAPIServer()
}

func RegisterGameAPIServer(s grpc.ServiceRegistrar, srv GameAPIServer) {
	// If the following call pancis, it indicates UnimplementedGameAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GameAPI_ServiceDesc, srv)
}

func _GameAPI_WorldBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Server)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameAPIServer).WorldBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameAPI_WorldBackup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameAPIServer).WorldBackup(ctx, req.(*Server))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameAPI_SendCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CmdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameAPIServer).SendCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameAPI_SendCmd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameAPIServer).SendCmd(ctx, req.(*CmdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameAPI_GameBind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameAPIServer).GameBind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameAPI_GameBind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameAPIServer).GameBind(ctx, req.(*BindReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameAPI_LoadJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SrvAndPath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameAPIServer).LoadJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameAPI_LoadJSON_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameAPIServer).LoadJSON(ctx, req.(*SrvAndPath))
	}
	return interceptor(ctx, in, info, handler)
}

// GameAPI_ServiceDesc is the grpc.ServiceDesc for GameAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GameAPI.GameAPI",
	HandlerType: (*GameAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WorldBackup",
			Handler:    _GameAPI_WorldBackup_Handler,
		},
		{
			MethodName: "SendCmd",
			Handler:    _GameAPI_SendCmd_Handler,
		},
		{
			MethodName: "GameBind",
			Handler:    _GameAPI_GameBind_Handler,
		},
		{
			MethodName: "LoadJSON",
			Handler:    _GameAPI_LoadJSON_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gameapi.proto",
}
