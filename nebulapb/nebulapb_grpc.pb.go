// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: nebulapb.proto

package nebulapb

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

// NebulaClient is the client API for Nebula service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NebulaClient interface {
	// API -> Bungee (ServerEntry)
	GetServerEntry(ctx context.Context, in *GetServerEntryRequest, opts ...grpc.CallOption) (*GetServerEntryResponse, error)
	// API <- App
	AddServerEntry(ctx context.Context, in *AddServerEntryRequest, opts ...grpc.CallOption) (*AddServerEntryResponse, error)
	// API <- App
	RemoveServerEntry(ctx context.Context, in *RemoveServerEntryRequest, opts ...grpc.CallOption) (*RemoveServerEntryResponse, error)
	// API -> Bungee (BungeeEntry)
	GetBungeeEntry(ctx context.Context, in *GetBungeeEntryRequest, opts ...grpc.CallOption) (*GetBungeeEntryResponse, error)
	// API <- App
	SetMotd(ctx context.Context, in *SetMotdRequest, opts ...grpc.CallOption) (*SetMotdResponse, error)
	// API <- App
	SetFavicon(ctx context.Context, in *SetFaviconRequest, opts ...grpc.CallOption) (*SetFaviconResponse, error)
	// API <- App
	SetLockdown(ctx context.Context, in *SetLockdownRequest, opts ...grpc.CallOption) (*SetLockdownResponse, error)
	// API <- Bungee / Server
	IPLookup(ctx context.Context, in *IPLookupRequest, opts ...grpc.CallOption) (*IPLookupResponse, error)
}

type nebulaClient struct {
	cc grpc.ClientConnInterface
}

func NewNebulaClient(cc grpc.ClientConnInterface) NebulaClient {
	return &nebulaClient{cc}
}

func (c *nebulaClient) GetServerEntry(ctx context.Context, in *GetServerEntryRequest, opts ...grpc.CallOption) (*GetServerEntryResponse, error) {
	out := new(GetServerEntryResponse)
	err := c.cc.Invoke(ctx, "/nebulapb.Nebula/GetServerEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nebulaClient) AddServerEntry(ctx context.Context, in *AddServerEntryRequest, opts ...grpc.CallOption) (*AddServerEntryResponse, error) {
	out := new(AddServerEntryResponse)
	err := c.cc.Invoke(ctx, "/nebulapb.Nebula/AddServerEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nebulaClient) RemoveServerEntry(ctx context.Context, in *RemoveServerEntryRequest, opts ...grpc.CallOption) (*RemoveServerEntryResponse, error) {
	out := new(RemoveServerEntryResponse)
	err := c.cc.Invoke(ctx, "/nebulapb.Nebula/RemoveServerEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nebulaClient) GetBungeeEntry(ctx context.Context, in *GetBungeeEntryRequest, opts ...grpc.CallOption) (*GetBungeeEntryResponse, error) {
	out := new(GetBungeeEntryResponse)
	err := c.cc.Invoke(ctx, "/nebulapb.Nebula/GetBungeeEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nebulaClient) SetMotd(ctx context.Context, in *SetMotdRequest, opts ...grpc.CallOption) (*SetMotdResponse, error) {
	out := new(SetMotdResponse)
	err := c.cc.Invoke(ctx, "/nebulapb.Nebula/SetMotd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nebulaClient) SetFavicon(ctx context.Context, in *SetFaviconRequest, opts ...grpc.CallOption) (*SetFaviconResponse, error) {
	out := new(SetFaviconResponse)
	err := c.cc.Invoke(ctx, "/nebulapb.Nebula/SetFavicon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nebulaClient) SetLockdown(ctx context.Context, in *SetLockdownRequest, opts ...grpc.CallOption) (*SetLockdownResponse, error) {
	out := new(SetLockdownResponse)
	err := c.cc.Invoke(ctx, "/nebulapb.Nebula/SetLockdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nebulaClient) IPLookup(ctx context.Context, in *IPLookupRequest, opts ...grpc.CallOption) (*IPLookupResponse, error) {
	out := new(IPLookupResponse)
	err := c.cc.Invoke(ctx, "/nebulapb.Nebula/IPLookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NebulaServer is the server API for Nebula service.
// All implementations should embed UnimplementedNebulaServer
// for forward compatibility
type NebulaServer interface {
	// API -> Bungee (ServerEntry)
	GetServerEntry(context.Context, *GetServerEntryRequest) (*GetServerEntryResponse, error)
	// API <- App
	AddServerEntry(context.Context, *AddServerEntryRequest) (*AddServerEntryResponse, error)
	// API <- App
	RemoveServerEntry(context.Context, *RemoveServerEntryRequest) (*RemoveServerEntryResponse, error)
	// API -> Bungee (BungeeEntry)
	GetBungeeEntry(context.Context, *GetBungeeEntryRequest) (*GetBungeeEntryResponse, error)
	// API <- App
	SetMotd(context.Context, *SetMotdRequest) (*SetMotdResponse, error)
	// API <- App
	SetFavicon(context.Context, *SetFaviconRequest) (*SetFaviconResponse, error)
	// API <- App
	SetLockdown(context.Context, *SetLockdownRequest) (*SetLockdownResponse, error)
	// API <- Bungee / Server
	IPLookup(context.Context, *IPLookupRequest) (*IPLookupResponse, error)
}

// UnimplementedNebulaServer should be embedded to have forward compatible implementations.
type UnimplementedNebulaServer struct {
}

func (UnimplementedNebulaServer) GetServerEntry(context.Context, *GetServerEntryRequest) (*GetServerEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerEntry not implemented")
}
func (UnimplementedNebulaServer) AddServerEntry(context.Context, *AddServerEntryRequest) (*AddServerEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddServerEntry not implemented")
}
func (UnimplementedNebulaServer) RemoveServerEntry(context.Context, *RemoveServerEntryRequest) (*RemoveServerEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveServerEntry not implemented")
}
func (UnimplementedNebulaServer) GetBungeeEntry(context.Context, *GetBungeeEntryRequest) (*GetBungeeEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBungeeEntry not implemented")
}
func (UnimplementedNebulaServer) SetMotd(context.Context, *SetMotdRequest) (*SetMotdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMotd not implemented")
}
func (UnimplementedNebulaServer) SetFavicon(context.Context, *SetFaviconRequest) (*SetFaviconResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFavicon not implemented")
}
func (UnimplementedNebulaServer) SetLockdown(context.Context, *SetLockdownRequest) (*SetLockdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLockdown not implemented")
}
func (UnimplementedNebulaServer) IPLookup(context.Context, *IPLookupRequest) (*IPLookupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IPLookup not implemented")
}

// UnsafeNebulaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NebulaServer will
// result in compilation errors.
type UnsafeNebulaServer interface {
	mustEmbedUnimplementedNebulaServer()
}

func RegisterNebulaServer(s grpc.ServiceRegistrar, srv NebulaServer) {
	s.RegisterService(&Nebula_ServiceDesc, srv)
}

func _Nebula_GetServerEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NebulaServer).GetServerEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebulapb.Nebula/GetServerEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NebulaServer).GetServerEntry(ctx, req.(*GetServerEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nebula_AddServerEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddServerEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NebulaServer).AddServerEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebulapb.Nebula/AddServerEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NebulaServer).AddServerEntry(ctx, req.(*AddServerEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nebula_RemoveServerEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveServerEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NebulaServer).RemoveServerEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebulapb.Nebula/RemoveServerEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NebulaServer).RemoveServerEntry(ctx, req.(*RemoveServerEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nebula_GetBungeeEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBungeeEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NebulaServer).GetBungeeEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebulapb.Nebula/GetBungeeEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NebulaServer).GetBungeeEntry(ctx, req.(*GetBungeeEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nebula_SetMotd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMotdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NebulaServer).SetMotd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebulapb.Nebula/SetMotd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NebulaServer).SetMotd(ctx, req.(*SetMotdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nebula_SetFavicon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFaviconRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NebulaServer).SetFavicon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebulapb.Nebula/SetFavicon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NebulaServer).SetFavicon(ctx, req.(*SetFaviconRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nebula_SetLockdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLockdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NebulaServer).SetLockdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebulapb.Nebula/SetLockdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NebulaServer).SetLockdown(ctx, req.(*SetLockdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nebula_IPLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPLookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NebulaServer).IPLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebulapb.Nebula/IPLookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NebulaServer).IPLookup(ctx, req.(*IPLookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Nebula_ServiceDesc is the grpc.ServiceDesc for Nebula service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Nebula_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nebulapb.Nebula",
	HandlerType: (*NebulaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServerEntry",
			Handler:    _Nebula_GetServerEntry_Handler,
		},
		{
			MethodName: "AddServerEntry",
			Handler:    _Nebula_AddServerEntry_Handler,
		},
		{
			MethodName: "RemoveServerEntry",
			Handler:    _Nebula_RemoveServerEntry_Handler,
		},
		{
			MethodName: "GetBungeeEntry",
			Handler:    _Nebula_GetBungeeEntry_Handler,
		},
		{
			MethodName: "SetMotd",
			Handler:    _Nebula_SetMotd_Handler,
		},
		{
			MethodName: "SetFavicon",
			Handler:    _Nebula_SetFavicon_Handler,
		},
		{
			MethodName: "SetLockdown",
			Handler:    _Nebula_SetLockdown_Handler,
		},
		{
			MethodName: "IPLookup",
			Handler:    _Nebula_IPLookup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nebulapb.proto",
}
