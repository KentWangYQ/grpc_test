// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

/*
Package storage is a generated protocol buffer package.

It is generated from these files:
	storage.proto

It has these top-level messages:
	User
	Data
	DataSummary
*/
package storage

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

type User struct {
	Id   int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Data struct {
	User *User  `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *Data) Reset()                    { *m = Data{} }
func (m *Data) String() string            { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()               {}
func (*Data) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Data) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Data) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type DataSummary struct {
	MessageCount int32 `protobuf:"varint,2,opt,name=messageCount" json:"messageCount,omitempty"`
}

func (m *DataSummary) Reset()                    { *m = DataSummary{} }
func (m *DataSummary) String() string            { return proto.CompactTextString(m) }
func (*DataSummary) ProtoMessage()               {}
func (*DataSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DataSummary) GetMessageCount() int32 {
	if m != nil {
		return m.MessageCount
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "storage.User")
	proto.RegisterType((*Data)(nil), "storage.Data")
	proto.RegisterType((*DataSummary)(nil), "storage.DataSummary")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Storage service

type StorageClient interface {
	GetFirstData(ctx context.Context, in *User, opts ...grpc.CallOption) (*Data, error)
	GetData(ctx context.Context, in *User, opts ...grpc.CallOption) (Storage_GetDataClient, error)
	InsertData(ctx context.Context, opts ...grpc.CallOption) (Storage_InsertDataClient, error)
	InsertAndShowAllData(ctx context.Context, opts ...grpc.CallOption) (Storage_InsertAndShowAllDataClient, error)
}

type storageClient struct {
	cc *grpc.ClientConn
}

func NewStorageClient(cc *grpc.ClientConn) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) GetFirstData(ctx context.Context, in *User, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := grpc.Invoke(ctx, "/storage.Storage/GetFirstData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) GetData(ctx context.Context, in *User, opts ...grpc.CallOption) (Storage_GetDataClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Storage_serviceDesc.Streams[0], c.cc, "/storage.Storage/GetData", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageGetDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Storage_GetDataClient interface {
	Recv() (*Data, error)
	grpc.ClientStream
}

type storageGetDataClient struct {
	grpc.ClientStream
}

func (x *storageGetDataClient) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageClient) InsertData(ctx context.Context, opts ...grpc.CallOption) (Storage_InsertDataClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Storage_serviceDesc.Streams[1], c.cc, "/storage.Storage/InsertData", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageInsertDataClient{stream}
	return x, nil
}

type Storage_InsertDataClient interface {
	Send(*Data) error
	CloseAndRecv() (*DataSummary, error)
	grpc.ClientStream
}

type storageInsertDataClient struct {
	grpc.ClientStream
}

func (x *storageInsertDataClient) Send(m *Data) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageInsertDataClient) CloseAndRecv() (*DataSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(DataSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageClient) InsertAndShowAllData(ctx context.Context, opts ...grpc.CallOption) (Storage_InsertAndShowAllDataClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Storage_serviceDesc.Streams[2], c.cc, "/storage.Storage/InsertAndShowAllData", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageInsertAndShowAllDataClient{stream}
	return x, nil
}

type Storage_InsertAndShowAllDataClient interface {
	Send(*Data) error
	Recv() (*Data, error)
	grpc.ClientStream
}

type storageInsertAndShowAllDataClient struct {
	grpc.ClientStream
}

func (x *storageInsertAndShowAllDataClient) Send(m *Data) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageInsertAndShowAllDataClient) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Storage service

type StorageServer interface {
	GetFirstData(context.Context, *User) (*Data, error)
	GetData(*User, Storage_GetDataServer) error
	InsertData(Storage_InsertDataServer) error
	InsertAndShowAllData(Storage_InsertAndShowAllDataServer) error
}

func RegisterStorageServer(s *grpc.Server, srv StorageServer) {
	s.RegisterService(&_Storage_serviceDesc, srv)
}

func _Storage_GetFirstData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).GetFirstData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.Storage/GetFirstData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).GetFirstData(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_GetData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageServer).GetData(m, &storageGetDataServer{stream})
}

type Storage_GetDataServer interface {
	Send(*Data) error
	grpc.ServerStream
}

type storageGetDataServer struct {
	grpc.ServerStream
}

func (x *storageGetDataServer) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func _Storage_InsertData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServer).InsertData(&storageInsertDataServer{stream})
}

type Storage_InsertDataServer interface {
	SendAndClose(*DataSummary) error
	Recv() (*Data, error)
	grpc.ServerStream
}

type storageInsertDataServer struct {
	grpc.ServerStream
}

func (x *storageInsertDataServer) SendAndClose(m *DataSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageInsertDataServer) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Storage_InsertAndShowAllData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServer).InsertAndShowAllData(&storageInsertAndShowAllDataServer{stream})
}

type Storage_InsertAndShowAllDataServer interface {
	Send(*Data) error
	Recv() (*Data, error)
	grpc.ServerStream
}

type storageInsertAndShowAllDataServer struct {
	grpc.ServerStream
}

func (x *storageInsertAndShowAllDataServer) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageInsertAndShowAllDataServer) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Storage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storage.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFirstData",
			Handler:    _Storage_GetFirstData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetData",
			Handler:       _Storage_GetData_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "InsertData",
			Handler:       _Storage_InsertData_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "InsertAndShowAllData",
			Handler:       _Storage_InsertAndShowAllData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "storage.proto",
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xbb, 0x31, 0x35, 0x38, 0x6d, 0x45, 0x86, 0x1e, 0x4a, 0x4f, 0x75, 0x4f, 0x41, 0x21,
	0xd4, 0x8a, 0x17, 0x3d, 0x15, 0xc5, 0xe2, 0x35, 0xc1, 0x1f, 0xb0, 0x92, 0x21, 0x06, 0xba, 0x59,
	0xd9, 0xd9, 0x20, 0xfe, 0x54, 0xff, 0x8d, 0xec, 0x36, 0x22, 0x09, 0x1e, 0x7a, 0x9b, 0xb7, 0xf3,
	0xde, 0x37, 0x0f, 0x16, 0x66, 0xec, 0x8c, 0x55, 0x15, 0x65, 0x1f, 0xd6, 0x38, 0x83, 0x49, 0x27,
	0xe5, 0x15, 0xc4, 0xaf, 0x4c, 0x16, 0xcf, 0x21, 0xaa, 0xcb, 0x85, 0x58, 0x89, 0x74, 0x9c, 0x47,
	0x75, 0x89, 0x08, 0x71, 0xa3, 0x34, 0x2d, 0xa2, 0x95, 0x48, 0xcf, 0xf2, 0x30, 0xcb, 0x07, 0x88,
	0x9f, 0x94, 0x53, 0x78, 0x09, 0x71, 0xcb, 0x64, 0x83, 0x7b, 0xb2, 0x99, 0x65, 0xbf, 0x68, 0x0f,
	0xca, 0xc3, 0x0a, 0x2f, 0xe0, 0x44, 0x73, 0xd5, 0xa5, 0xfd, 0x28, 0x6f, 0x60, 0xe2, 0xc3, 0x45,
	0xab, 0xb5, 0xb2, 0x5f, 0x28, 0x61, 0xaa, 0x89, 0x59, 0x55, 0xf4, 0x68, 0xda, 0xc6, 0x05, 0xe7,
	0x38, 0xef, 0xbd, 0x6d, 0xbe, 0x05, 0x24, 0xc5, 0x81, 0x8d, 0x19, 0x4c, 0x77, 0xe4, 0x9e, 0x6b,
	0xcb, 0x2e, 0x74, 0xe8, 0x5f, 0x5d, 0xfe, 0x49, 0xbf, 0x95, 0x23, 0xbc, 0x86, 0x64, 0x47, 0x47,
	0x59, 0xd7, 0x02, 0xef, 0x00, 0x5e, 0x1a, 0x26, 0x3b, 0xf4, 0x7b, 0xb9, 0x9c, 0xf7, 0x64, 0xd7,
	0x5f, 0x8e, 0x52, 0x81, 0xf7, 0x30, 0x3f, 0xc4, 0xb6, 0x4d, 0x59, 0xbc, 0x9b, 0xcf, 0xed, 0x7e,
	0xff, 0x1f, 0x60, 0x78, 0x30, 0x15, 0x6b, 0xf1, 0x76, 0x1a, 0xfe, 0xe1, 0xf6, 0x27, 0x00, 0x00,
	0xff, 0xff, 0xde, 0xf0, 0x57, 0x8a, 0x98, 0x01, 0x00, 0x00,
}
