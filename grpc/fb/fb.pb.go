// Code generated by protoc-gen-go.
// source: fb.proto
// DO NOT EDIT!

/*
Package fb is a generated protocol buffer package.

It is generated from these files:
	fb.proto

It has these top-level messages:
	Payload
*/
package fb

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

type Payload struct {
	Data []byte `protobuf:"bytes,1,opt,name=Data,json=data,proto3" json:"Data,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Payload)(nil), "fb.Payload")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Worker service

type WorkerClient interface {
	Hello(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Payload, error)
}

type workerClient struct {
	cc *grpc.ClientConn
}

func NewWorkerClient(cc *grpc.ClientConn) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) Hello(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Payload, error) {
	out := new(Payload)
	err := grpc.Invoke(ctx, "/fb.Worker/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Worker service

type WorkerServer interface {
	Hello(context.Context, *Payload) (*Payload, error)
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fb.Worker/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Hello(ctx, req.(*Payload))
	}
	return interceptor(ctx, in, info, handler)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fb.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Worker_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 104 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x48, 0x4b, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4a, 0x4b, 0x52, 0x92, 0xe5, 0x62, 0x0f, 0x48, 0xac, 0xcc,
	0xc9, 0x4f, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0x71, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x09, 0x62, 0x49, 0x01, 0xb2, 0x8d, 0x74, 0xb9, 0xd8, 0xc2, 0xf3, 0x8b, 0xb2, 0x53, 0x8b,
	0x84, 0x94, 0xb9, 0x58, 0x3d, 0x52, 0x73, 0x72, 0xf2, 0x85, 0xb8, 0xf5, 0x80, 0x06, 0x40, 0xf5,
	0x48, 0x21, 0x73, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0x06, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x81, 0xdc, 0x6e, 0x22, 0x64, 0x00, 0x00, 0x00,
}