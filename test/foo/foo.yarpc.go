// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test/foo/foo.proto

/*
Package foo is a generated protocol buffer package.

It is generated from these files:
	test/foo/foo.proto

It has these top-level messages:
	Request
	Response
*/
package foo

import (
	context "context"
	yarpc "github.com/influxdata/yarpc"
)

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/influxdata/yarpc/yarpcproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ yarpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the yarpc package it is being compiled against.
const _ = yarpc.SupportPackageIsVersion1

// Client API for Foo service

type FooClient interface {
	UnaryMethod(ctx context.Context, in *Request) (*Response, error)
	ServerStreamMethod(ctx context.Context, in *Request) (Foo_ServerStreamMethodClient, error)
}

type fooClient struct {
	cc *yarpc.ClientConn
}

func NewFooClient(cc *yarpc.ClientConn) FooClient {
	return &fooClient{cc}
}

func (c *fooClient) UnaryMethod(ctx context.Context, in *Request) (*Response, error) {
	out := new(Response)
	err := yarpc.Invoke(ctx, 0x0000, in, out, c.cc)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fooClient) ServerStreamMethod(ctx context.Context, in *Request) (Foo_ServerStreamMethodClient, error) {
	stream, err := yarpc.NewClientStream(ctx, &_Foo_serviceDesc.Streams[0], c.cc, 0x0001)
	if err != nil {
		return nil, err
	}
	x := &fooServerStreamMethodClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	return x, nil
}

type Foo_ServerStreamMethodClient interface {
	Recv() (*Response, error)
	yarpc.ClientStream
}

type fooServerStreamMethodClient struct {
	yarpc.ClientStream
}

func (x *fooServerStreamMethodClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Foo service

type FooServer interface {
	UnaryMethod(context.Context, *Request) (*Response, error)
	ServerStreamMethod(*Request, Foo_ServerStreamMethodServer) error
}

func RegisterFooServer(s *yarpc.Server, srv FooServer) {
	s.RegisterService(&_Foo_serviceDesc, srv)
}

func _Foo_UnaryMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(FooServer).UnaryMethod(ctx, in)
}

func _Foo_ServerStreamMethod_Handler(srv interface{}, stream yarpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FooServer).ServerStreamMethod(m, &fooServerStreamMethodServer{stream})
}

type Foo_ServerStreamMethodServer interface {
	Send(*Response) error
	yarpc.ServerStream
}

type fooServerStreamMethodServer struct {
	yarpc.ServerStream
}

func (x *fooServerStreamMethodServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _Foo_serviceDesc = yarpc.ServiceDesc{
	ServiceName: "foo.Foo",
	Index:       0,
	HandlerType: (*FooServer)(nil),
	Methods: []yarpc.MethodDesc{
		{
			MethodName: "UnaryMethod",
			Index:      0,
			Handler:    _Foo_UnaryMethod_Handler,
		},
	},
	Streams: []yarpc.StreamDesc{
		{
			StreamName:    "ServerStreamMethod",
			Index:         1,
			Handler:       _Foo_ServerStreamMethod_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "test/foo/foo.proto",
}
