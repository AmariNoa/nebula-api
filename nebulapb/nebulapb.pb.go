// Code generated by protoc-gen-go.
// source: nebulapb.proto
// DO NOT EDIT!

/*
Package nebulapb is a generated protocol buffer package.

It is generated from these files:
	nebulapb.proto

It has these top-level messages:
	Empty
	StreamRequest
	EntryStreamResponse
	QuitEntryStreamRequest
	ServerEntry
	ServerStatus
	GetServerEntryRequest
	GetServerEntryResponse
	AddServerEntryRequest
	AddServerEntryResponse
	RemoveServerEntryRequest
	RemoveServerEntryResponse
*/
package nebulapb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamType int32

const (
	StreamType_QUIT     StreamType = 0
	StreamType_CONNECT  StreamType = 1
	StreamType_RESTORED StreamType = 2
	StreamType_DISPATCH StreamType = 3
	StreamType_SYNC     StreamType = 4
	StreamType_REMOVE   StreamType = 5
)

var StreamType_name = map[int32]string{
	0: "QUIT",
	1: "CONNECT",
	2: "RESTORED",
	3: "DISPATCH",
	4: "SYNC",
	5: "REMOVE",
}
var StreamType_value = map[string]int32{
	"QUIT":     0,
	"CONNECT":  1,
	"RESTORED": 2,
	"DISPATCH": 3,
	"SYNC":     4,
	"REMOVE":   5,
}

func (x StreamType) String() string {
	return proto.EnumName(StreamType_name, int32(x))
}
func (StreamType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StreamRequest struct {
	Name string     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type StreamType `protobuf:"varint,2,opt,name=type,enum=nebulapb.StreamType" json:"type,omitempty"`
}

func (m *StreamRequest) Reset()                    { *m = StreamRequest{} }
func (m *StreamRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()               {}
func (*StreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StreamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StreamRequest) GetType() StreamType {
	if m != nil {
		return m.Type
	}
	return StreamType_QUIT
}

type EntryStreamResponse struct {
	Type   StreamType   `protobuf:"varint,1,opt,name=type,enum=nebulapb.StreamType" json:"type,omitempty"`
	Target string       `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	Entry  *ServerEntry `protobuf:"bytes,3,opt,name=entry" json:"entry,omitempty"`
}

func (m *EntryStreamResponse) Reset()                    { *m = EntryStreamResponse{} }
func (m *EntryStreamResponse) String() string            { return proto.CompactTextString(m) }
func (*EntryStreamResponse) ProtoMessage()               {}
func (*EntryStreamResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EntryStreamResponse) GetType() StreamType {
	if m != nil {
		return m.Type
	}
	return StreamType_QUIT
}

func (m *EntryStreamResponse) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *EntryStreamResponse) GetEntry() *ServerEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type QuitEntryStreamRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *QuitEntryStreamRequest) Reset()                    { *m = QuitEntryStreamRequest{} }
func (m *QuitEntryStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*QuitEntryStreamRequest) ProtoMessage()               {}
func (*QuitEntryStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *QuitEntryStreamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// --
// Server Entry
// --
type ServerEntry struct {
	Name        string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	DisplayName string        `protobuf:"bytes,2,opt,name=displayName" json:"displayName,omitempty"`
	Address     string        `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	Port        int32         `protobuf:"varint,4,opt,name=port" json:"port,omitempty"`
	Motd        string        `protobuf:"bytes,5,opt,name=motd" json:"motd,omitempty"`
	Status      *ServerStatus `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
}

func (m *ServerEntry) Reset()                    { *m = ServerEntry{} }
func (m *ServerEntry) String() string            { return proto.CompactTextString(m) }
func (*ServerEntry) ProtoMessage()               {}
func (*ServerEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ServerEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServerEntry) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ServerEntry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ServerEntry) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ServerEntry) GetMotd() string {
	if m != nil {
		return m.Motd
	}
	return ""
}

func (m *ServerEntry) GetStatus() *ServerStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ServerStatus struct {
	Online  bool                  `protobuf:"varint,1,opt,name=online" json:"online,omitempty"`
	Version *ServerStatus_Version `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Players *ServerStatus_Players `protobuf:"bytes,3,opt,name=players" json:"players,omitempty"`
	// i don't have any idea... (need map)
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Favicon     string `protobuf:"bytes,5,opt,name=favicon" json:"favicon,omitempty"`
}

func (m *ServerStatus) Reset()                    { *m = ServerStatus{} }
func (m *ServerStatus) String() string            { return proto.CompactTextString(m) }
func (*ServerStatus) ProtoMessage()               {}
func (*ServerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ServerStatus) GetOnline() bool {
	if m != nil {
		return m.Online
	}
	return false
}

func (m *ServerStatus) GetVersion() *ServerStatus_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *ServerStatus) GetPlayers() *ServerStatus_Players {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *ServerStatus) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ServerStatus) GetFavicon() string {
	if m != nil {
		return m.Favicon
	}
	return ""
}

type ServerStatus_Version struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Protocol int32  `protobuf:"varint,2,opt,name=protocol" json:"protocol,omitempty"`
}

func (m *ServerStatus_Version) Reset()                    { *m = ServerStatus_Version{} }
func (m *ServerStatus_Version) String() string            { return proto.CompactTextString(m) }
func (*ServerStatus_Version) ProtoMessage()               {}
func (*ServerStatus_Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *ServerStatus_Version) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServerStatus_Version) GetProtocol() int32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

type ServerStatus_Players struct {
	Max    int32 `protobuf:"varint,1,opt,name=max" json:"max,omitempty"`
	Online int32 `protobuf:"varint,2,opt,name=online" json:"online,omitempty"`
}

func (m *ServerStatus_Players) Reset()                    { *m = ServerStatus_Players{} }
func (m *ServerStatus_Players) String() string            { return proto.CompactTextString(m) }
func (*ServerStatus_Players) ProtoMessage()               {}
func (*ServerStatus_Players) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 1} }

func (m *ServerStatus_Players) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *ServerStatus_Players) GetOnline() int32 {
	if m != nil {
		return m.Online
	}
	return 0
}

//
// ServerEntry
//
type GetServerEntryRequest struct {
}

func (m *GetServerEntryRequest) Reset()                    { *m = GetServerEntryRequest{} }
func (m *GetServerEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetServerEntryRequest) ProtoMessage()               {}
func (*GetServerEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type GetServerEntryResponse struct {
	Entry []*ServerEntry `protobuf:"bytes,1,rep,name=entry" json:"entry,omitempty"`
}

func (m *GetServerEntryResponse) Reset()                    { *m = GetServerEntryResponse{} }
func (m *GetServerEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*GetServerEntryResponse) ProtoMessage()               {}
func (*GetServerEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetServerEntryResponse) GetEntry() []*ServerEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type AddServerEntryRequest struct {
	Entry *ServerEntry `protobuf:"bytes,1,opt,name=entry" json:"entry,omitempty"`
}

func (m *AddServerEntryRequest) Reset()                    { *m = AddServerEntryRequest{} }
func (m *AddServerEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*AddServerEntryRequest) ProtoMessage()               {}
func (*AddServerEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AddServerEntryRequest) GetEntry() *ServerEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type AddServerEntryResponse struct {
}

func (m *AddServerEntryResponse) Reset()                    { *m = AddServerEntryResponse{} }
func (m *AddServerEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*AddServerEntryResponse) ProtoMessage()               {}
func (*AddServerEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type RemoveServerEntryRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *RemoveServerEntryRequest) Reset()                    { *m = RemoveServerEntryRequest{} }
func (m *RemoveServerEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveServerEntryRequest) ProtoMessage()               {}
func (*RemoveServerEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RemoveServerEntryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RemoveServerEntryResponse struct {
}

func (m *RemoveServerEntryResponse) Reset()                    { *m = RemoveServerEntryResponse{} }
func (m *RemoveServerEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveServerEntryResponse) ProtoMessage()               {}
func (*RemoveServerEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*Empty)(nil), "nebulapb.Empty")
	proto.RegisterType((*StreamRequest)(nil), "nebulapb.StreamRequest")
	proto.RegisterType((*EntryStreamResponse)(nil), "nebulapb.EntryStreamResponse")
	proto.RegisterType((*QuitEntryStreamRequest)(nil), "nebulapb.QuitEntryStreamRequest")
	proto.RegisterType((*ServerEntry)(nil), "nebulapb.ServerEntry")
	proto.RegisterType((*ServerStatus)(nil), "nebulapb.ServerStatus")
	proto.RegisterType((*ServerStatus_Version)(nil), "nebulapb.ServerStatus.Version")
	proto.RegisterType((*ServerStatus_Players)(nil), "nebulapb.ServerStatus.Players")
	proto.RegisterType((*GetServerEntryRequest)(nil), "nebulapb.GetServerEntryRequest")
	proto.RegisterType((*GetServerEntryResponse)(nil), "nebulapb.GetServerEntryResponse")
	proto.RegisterType((*AddServerEntryRequest)(nil), "nebulapb.AddServerEntryRequest")
	proto.RegisterType((*AddServerEntryResponse)(nil), "nebulapb.AddServerEntryResponse")
	proto.RegisterType((*RemoveServerEntryRequest)(nil), "nebulapb.RemoveServerEntryRequest")
	proto.RegisterType((*RemoveServerEntryResponse)(nil), "nebulapb.RemoveServerEntryResponse")
	proto.RegisterEnum("nebulapb.StreamType", StreamType_name, StreamType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Nebula service

type NebulaClient interface {
	// -
	// MISC
	// -
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// -
	// STREAM
	// -
	// API -> Bungee (Notify Server Status)
	EntryStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Nebula_EntryStreamClient, error)
	QuitEntryStream(ctx context.Context, in *QuitEntryStreamRequest, opts ...grpc.CallOption) (*Empty, error)
	// -
	// ServerEntry
	// -
	// API -> Bungee (ServerEntry)
	GetServerEntry(ctx context.Context, in *GetServerEntryRequest, opts ...grpc.CallOption) (*GetServerEntryResponse, error)
	// API <- App
	AddServerEntry(ctx context.Context, in *AddServerEntryRequest, opts ...grpc.CallOption) (*AddServerEntryResponse, error)
	// API <- App
	RemoveServerEntry(ctx context.Context, in *RemoveServerEntryRequest, opts ...grpc.CallOption) (*RemoveServerEntryResponse, error)
}

type nebulaClient struct {
	cc *grpc.ClientConn
}

func NewNebulaClient(cc *grpc.ClientConn) NebulaClient {
	return &nebulaClient{cc}
}

func (c *nebulaClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/nebulapb.Nebula/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nebulaClient) EntryStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Nebula_EntryStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Nebula_serviceDesc.Streams[0], c.cc, "/nebulapb.Nebula/EntryStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &nebulaEntryStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Nebula_EntryStreamClient interface {
	Recv() (*EntryStreamResponse, error)
	grpc.ClientStream
}

type nebulaEntryStreamClient struct {
	grpc.ClientStream
}

func (x *nebulaEntryStreamClient) Recv() (*EntryStreamResponse, error) {
	m := new(EntryStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nebulaClient) QuitEntryStream(ctx context.Context, in *QuitEntryStreamRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/nebulapb.Nebula/QuitEntryStream", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nebulaClient) GetServerEntry(ctx context.Context, in *GetServerEntryRequest, opts ...grpc.CallOption) (*GetServerEntryResponse, error) {
	out := new(GetServerEntryResponse)
	err := grpc.Invoke(ctx, "/nebulapb.Nebula/GetServerEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nebulaClient) AddServerEntry(ctx context.Context, in *AddServerEntryRequest, opts ...grpc.CallOption) (*AddServerEntryResponse, error) {
	out := new(AddServerEntryResponse)
	err := grpc.Invoke(ctx, "/nebulapb.Nebula/AddServerEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nebulaClient) RemoveServerEntry(ctx context.Context, in *RemoveServerEntryRequest, opts ...grpc.CallOption) (*RemoveServerEntryResponse, error) {
	out := new(RemoveServerEntryResponse)
	err := grpc.Invoke(ctx, "/nebulapb.Nebula/RemoveServerEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Nebula service

type NebulaServer interface {
	// -
	// MISC
	// -
	Ping(context.Context, *Empty) (*Empty, error)
	// -
	// STREAM
	// -
	// API -> Bungee (Notify Server Status)
	EntryStream(*StreamRequest, Nebula_EntryStreamServer) error
	QuitEntryStream(context.Context, *QuitEntryStreamRequest) (*Empty, error)
	// -
	// ServerEntry
	// -
	// API -> Bungee (ServerEntry)
	GetServerEntry(context.Context, *GetServerEntryRequest) (*GetServerEntryResponse, error)
	// API <- App
	AddServerEntry(context.Context, *AddServerEntryRequest) (*AddServerEntryResponse, error)
	// API <- App
	RemoveServerEntry(context.Context, *RemoveServerEntryRequest) (*RemoveServerEntryResponse, error)
}

func RegisterNebulaServer(s *grpc.Server, srv NebulaServer) {
	s.RegisterService(&_Nebula_serviceDesc, srv)
}

func _Nebula_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NebulaServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebulapb.Nebula/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NebulaServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nebula_EntryStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NebulaServer).EntryStream(m, &nebulaEntryStreamServer{stream})
}

type Nebula_EntryStreamServer interface {
	Send(*EntryStreamResponse) error
	grpc.ServerStream
}

type nebulaEntryStreamServer struct {
	grpc.ServerStream
}

func (x *nebulaEntryStreamServer) Send(m *EntryStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Nebula_QuitEntryStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuitEntryStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NebulaServer).QuitEntryStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebulapb.Nebula/QuitEntryStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NebulaServer).QuitEntryStream(ctx, req.(*QuitEntryStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
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

var _Nebula_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nebulapb.Nebula",
	HandlerType: (*NebulaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Nebula_Ping_Handler,
		},
		{
			MethodName: "QuitEntryStream",
			Handler:    _Nebula_QuitEntryStream_Handler,
		},
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
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EntryStream",
			Handler:       _Nebula_EntryStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "nebulapb.proto",
}

func init() { proto.RegisterFile("nebulapb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 676 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x8d, 0x1b, 0x3b, 0x49, 0x6f, 0xfa, 0xb5, 0xfe, 0x06, 0x9a, 0x1a, 0x23, 0xc0, 0x32, 0x9b,
	0xa8, 0x20, 0x0b, 0xa5, 0x1b, 0x58, 0xb6, 0xa9, 0x81, 0x2e, 0x9a, 0xa6, 0x93, 0xb4, 0x12, 0x1b,
	0x24, 0x37, 0x19, 0x5a, 0x4b, 0x89, 0x6d, 0x3c, 0x93, 0x88, 0xec, 0x79, 0x25, 0x5e, 0x81, 0x37,
	0xe0, 0x7d, 0xd0, 0xfc, 0xb8, 0x76, 0x5d, 0x87, 0xec, 0xe6, 0xce, 0x9c, 0x7b, 0xe6, 0xdc, 0x3b,
	0x67, 0x2e, 0xec, 0x46, 0xe4, 0x66, 0x31, 0x0b, 0x92, 0x1b, 0x2f, 0x49, 0x63, 0x16, 0xa3, 0x56,
	0x16, 0xbb, 0x4d, 0x30, 0xfc, 0x79, 0xc2, 0x56, 0xee, 0x39, 0xfc, 0x37, 0x62, 0x29, 0x09, 0xe6,
	0x98, 0x7c, 0x5f, 0x10, 0xca, 0x10, 0x02, 0x3d, 0x0a, 0xe6, 0xc4, 0xd2, 0x1c, 0xad, 0xbb, 0x8d,
	0xc5, 0x1a, 0x75, 0x41, 0x67, 0xab, 0x84, 0x58, 0x5b, 0x8e, 0xd6, 0xdd, 0xed, 0x3d, 0xf5, 0xee,
	0x69, 0x65, 0xea, 0x78, 0x95, 0x10, 0x2c, 0x10, 0xee, 0x4f, 0x0d, 0x9e, 0xf8, 0x11, 0x4b, 0x57,
	0x19, 0x29, 0x4d, 0xe2, 0x88, 0xe6, 0x0c, 0xda, 0x26, 0x06, 0xd4, 0x81, 0x06, 0x0b, 0xd2, 0x5b,
	0xc2, 0xc4, 0x6d, 0xdb, 0x58, 0x45, 0xe8, 0x0d, 0x18, 0x84, 0x13, 0x5b, 0x75, 0x47, 0xeb, 0xb6,
	0x7b, 0xfb, 0x05, 0x0a, 0x92, 0x2e, 0x49, 0x2a, 0x6e, 0xc5, 0x12, 0xe3, 0xbe, 0x85, 0xce, 0xe5,
	0x22, 0x64, 0x0f, 0x94, 0xac, 0x2d, 0xcf, 0xfd, 0xa5, 0x41, 0xbb, 0x40, 0x52, 0xd9, 0x02, 0x07,
	0xda, 0xd3, 0x90, 0x26, 0xb3, 0x60, 0x35, 0xe0, 0x47, 0x52, 0x5b, 0x71, 0x0b, 0x59, 0xd0, 0x0c,
	0xa6, 0xd3, 0x94, 0x50, 0x2a, 0x24, 0x6e, 0xe3, 0x2c, 0xe4, 0x7c, 0x49, 0x9c, 0x32, 0x4b, 0x77,
	0xb4, 0xae, 0x81, 0xc5, 0x9a, 0xef, 0xcd, 0x63, 0x36, 0xb5, 0x0c, 0x79, 0x07, 0x5f, 0x23, 0x0f,
	0x1a, 0x94, 0x05, 0x6c, 0x41, 0xad, 0x86, 0xa8, 0xb1, 0x53, 0xae, 0x71, 0x24, 0x4e, 0xb1, 0x42,
	0xb9, 0xbf, 0xb7, 0x60, 0xa7, 0x78, 0xc0, 0x7b, 0x17, 0x47, 0xb3, 0x30, 0x92, 0xd2, 0x5b, 0x58,
	0x45, 0xe8, 0x3d, 0x34, 0x97, 0x24, 0xa5, 0x61, 0x1c, 0x09, 0xe1, 0xed, 0xde, 0xcb, 0x6a, 0x66,
	0xef, 0x5a, 0xa2, 0x70, 0x06, 0xe7, 0x99, 0xbc, 0x40, 0x92, 0x52, 0xd5, 0xf7, 0x75, 0x99, 0x43,
	0x89, 0xc2, 0x19, 0x5c, 0x34, 0x8c, 0xd0, 0x49, 0x1a, 0x26, 0x8c, 0xdf, 0xab, 0xab, 0x86, 0xe5,
	0x5b, 0xbc, 0x61, 0xdf, 0x82, 0x65, 0x38, 0x89, 0x23, 0xd5, 0x85, 0x2c, 0xb4, 0x3f, 0x40, 0x53,
	0x29, 0xa9, 0x7c, 0x0b, 0x1b, 0x5a, 0xc2, 0xcf, 0x93, 0x78, 0x26, 0xea, 0x31, 0xf0, 0x7d, 0x6c,
	0x1f, 0x41, 0x53, 0x49, 0x41, 0x26, 0xd4, 0xe7, 0xc1, 0x0f, 0x91, 0x69, 0x60, 0xbe, 0x2c, 0xf4,
	0x47, 0xa6, 0xa9, 0xc8, 0x3d, 0x80, 0xfd, 0x4f, 0x84, 0x15, 0x7d, 0x24, 0xdd, 0xe2, 0xfa, 0xd0,
	0x29, 0x1f, 0x28, 0x43, 0xdf, 0xdb, 0x51, 0x73, 0xea, 0x1b, 0xed, 0x78, 0x0a, 0xfb, 0xc7, 0xd3,
	0xe9, 0x63, 0xfe, 0x22, 0xcb, 0x66, 0x53, 0x5b, 0xd0, 0x29, 0xb3, 0x48, 0x31, 0xae, 0x07, 0x16,
	0x26, 0xf3, 0x78, 0x49, 0x2a, 0xae, 0xa8, 0x32, 0xfc, 0x73, 0x78, 0x56, 0x81, 0x97, 0x64, 0x87,
	0x57, 0x00, 0xf9, 0xa7, 0x44, 0x2d, 0xd0, 0x2f, 0xaf, 0xce, 0xc6, 0x66, 0x0d, 0xb5, 0xa1, 0xd9,
	0xbf, 0x18, 0x0c, 0xfc, 0xfe, 0xd8, 0xd4, 0xd0, 0x0e, 0xb4, 0xb0, 0x3f, 0x1a, 0x5f, 0x60, 0xff,
	0xd4, 0xdc, 0xe2, 0xd1, 0xe9, 0xd9, 0x68, 0x78, 0x3c, 0xee, 0x7f, 0x36, 0xeb, 0x3c, 0x65, 0xf4,
	0x65, 0xd0, 0x37, 0x75, 0x04, 0xd0, 0xc0, 0xfe, 0xf9, 0xc5, 0xb5, 0x6f, 0x1a, 0xbd, 0x3f, 0x75,
	0x68, 0x0c, 0x44, 0x75, 0xe8, 0x10, 0xf4, 0x61, 0x18, 0xdd, 0xa2, 0xbd, 0xbc, 0x5c, 0x31, 0x8c,
	0xec, 0xf2, 0x86, 0x5b, 0x43, 0x67, 0xd0, 0x2e, 0xfc, 0x62, 0x74, 0x50, 0x9e, 0x1c, 0xaa, 0x4c,
	0xfb, 0x45, 0x21, 0xf5, 0xf1, 0xfc, 0x71, 0x6b, 0xef, 0x34, 0xf4, 0x11, 0xf6, 0x4a, 0x43, 0x01,
	0x39, 0x79, 0x56, 0xf5, 0xbc, 0xa8, 0x92, 0x74, 0x05, 0xbb, 0x0f, 0x4d, 0x81, 0x5e, 0xe5, 0xa0,
	0x4a, 0x1f, 0xd9, 0xce, 0x7a, 0x40, 0x26, 0x90, 0xd3, 0x3e, 0x7c, 0xde, 0x22, 0x6d, 0xa5, 0x7d,
	0x8a, 0xb4, 0x6b, 0x9c, 0x51, 0x43, 0x5f, 0xe1, 0xff, 0x47, 0x6f, 0x8d, 0xdc, 0x3c, 0x71, 0x9d,
	0x71, 0xec, 0xd7, 0xff, 0xc4, 0x64, 0xfc, 0x27, 0x5d, 0xb0, 0x22, 0xc2, 0x3c, 0xba, 0x8a, 0x26,
	0x77, 0xec, 0x2e, 0x0c, 0xbc, 0x20, 0x09, 0x55, 0xe6, 0xc9, 0x8e, 0x7c, 0xf0, 0x21, 0xff, 0x9c,
	0xf4, 0xa6, 0x21, 0x3e, 0xe9, 0xd1, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xdd, 0x39, 0x87,
	0x96, 0x06, 0x00, 0x00,
}
