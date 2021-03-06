// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello/hello.proto

/*
Package hello is a generated protocol buffer package.

It is generated from these files:
	hello/hello.proto

It has these top-level messages:
	StringMessage
	SimpleMessage
*/
package hello

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type StringMessage struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *StringMessage) Reset()                    { *m = StringMessage{} }
func (m *StringMessage) String() string            { return proto.CompactTextString(m) }
func (*StringMessage) ProtoMessage()               {}
func (*StringMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StringMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SimpleMessage struct {
	Id  string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Num int64  `protobuf:"varint,2,opt,name=num" json:"num,omitempty"`
}

func (m *SimpleMessage) Reset()                    { *m = SimpleMessage{} }
func (m *SimpleMessage) String() string            { return proto.CompactTextString(m) }
func (*SimpleMessage) ProtoMessage()               {}
func (*SimpleMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SimpleMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SimpleMessage) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func init() {
	proto.RegisterType((*StringMessage)(nil), "hello.StringMessage")
	proto.RegisterType((*SimpleMessage)(nil), "hello.SimpleMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HelloService service

type HelloServiceClient interface {
	EchoPost(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
	Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) EchoPost(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := grpc.Invoke(ctx, "/hello.HelloService/EchoPost", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := grpc.Invoke(ctx, "/hello.HelloService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HelloService service

type HelloServiceServer interface {
	EchoPost(context.Context, *StringMessage) (*StringMessage, error)
	Echo(context.Context, *SimpleMessage) (*SimpleMessage, error)
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_EchoPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).EchoPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.HelloService/EchoPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).EchoPost(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.HelloService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Echo(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EchoPost",
			Handler:    _HelloService_EchoPost_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _HelloService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello/hello.proto",
}

func init() { proto.RegisterFile("hello/hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e, 0x94, 0x4c,
	0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49,
	0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31, 0x44, 0x91, 0x92, 0x2a, 0x17, 0x6f, 0x70, 0x49, 0x51, 0x66,
	0x5e, 0xba, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e,
	0x69, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0xa3, 0x64, 0xc8, 0xc5, 0x1b, 0x9c,
	0x99, 0x5b, 0x90, 0x93, 0x0a, 0x53, 0xc6, 0xc7, 0xc5, 0x94, 0x99, 0x02, 0x55, 0xc3, 0x94, 0x99,
	0x22, 0x24, 0xc0, 0xc5, 0x9c, 0x57, 0x9a, 0x2b, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1c, 0x04, 0x62,
	0x1a, 0xdd, 0x62, 0xe4, 0xe2, 0xf1, 0x00, 0xb9, 0x20, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55,
	0x28, 0x82, 0x8b, 0xc3, 0x35, 0x39, 0x23, 0x3f, 0x20, 0xbf, 0xb8, 0x44, 0x48, 0x44, 0x0f, 0xe2,
	0x52, 0x14, 0xbb, 0xa5, 0xb0, 0x8a, 0x2a, 0x29, 0x34, 0x5d, 0x7e, 0x32, 0x99, 0x49, 0x4a, 0x49,
	0x54, 0xbf, 0xcc, 0x50, 0x3f, 0xb5, 0x22, 0x11, 0xe4, 0x0c, 0xfd, 0xd4, 0xe4, 0x8c, 0xfc, 0xf8,
	0x82, 0xfc, 0xe2, 0x12, 0x2b, 0x46, 0x2d, 0xa1, 0x02, 0x2e, 0x16, 0x90, 0xc9, 0x08, 0x53, 0x91,
	0x9d, 0x2a, 0x85, 0x55, 0x54, 0xc9, 0x06, 0x6c, 0xaa, 0x19, 0xa6, 0xa9, 0xfa, 0xd5, 0x99, 0x29,
	0xb5, 0x51, 0xb2, 0x42, 0xd2, 0x58, 0x25, 0xf4, 0xab, 0xf3, 0x4a, 0x73, 0x6b, 0x9d, 0xd8, 0xa3,
	0x20, 0xa1, 0x9b, 0xc4, 0x06, 0x0e, 0x46, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x76,
	0x19, 0x7c, 0x80, 0x01, 0x00, 0x00,
}
