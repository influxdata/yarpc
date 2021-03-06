// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: yarpcproto/yarpc.proto

/*
Package yarpcproto is a generated protocol buffer package.

It is generated from these files:
	yarpcproto/yarpc.proto

It has these top-level messages:
*/
package yarpcproto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

var E_YarpcServiceIndex = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         50000,
	Name:          "yarpcproto.yarpc_service_index",
	Tag:           "varint,50000,opt,name=yarpc_service_index,json=yarpcServiceIndex",
	Filename:      "yarpcproto/yarpc.proto",
}

var E_YarpcMethodIndex = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         50000,
	Name:          "yarpcproto.yarpc_method_index",
	Tag:           "varint,50000,opt,name=yarpc_method_index,json=yarpcMethodIndex",
	Filename:      "yarpcproto/yarpc.proto",
}

func init() {
	proto.RegisterExtension(E_YarpcServiceIndex)
	proto.RegisterExtension(E_YarpcMethodIndex)
}

func init() { proto.RegisterFile("yarpcproto/yarpc.proto", fileDescriptorYarpc) }

var fileDescriptorYarpc = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xab, 0x4c, 0x2c, 0x2a,
	0x48, 0x2e, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x07, 0x33, 0xf5, 0xc0, 0x6c, 0x21, 0x2e, 0x84, 0xb8,
	0x94, 0x42, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x98, 0x97, 0x54, 0x9a, 0xa6, 0x9f, 0x92,
	0x5a, 0x9c, 0x5c, 0x94, 0x59, 0x50, 0x92, 0x5f, 0x04, 0x51, 0x6d, 0x15, 0xc8, 0x25, 0x0c, 0x56,
	0x1f, 0x5f, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x1a, 0x9f, 0x99, 0x97, 0x92, 0x5a, 0x21, 0x24,
	0xaf, 0x07, 0xd1, 0xa9, 0x07, 0xd3, 0xa9, 0x17, 0x0c, 0x91, 0xf7, 0x2f, 0x28, 0xc9, 0xcc, 0xcf,
	0x2b, 0x96, 0xb8, 0xd0, 0xc6, 0xac, 0xc0, 0xa8, 0xc1, 0x1b, 0x24, 0x08, 0xd6, 0x0d, 0x95, 0xf4,
	0x04, 0xe9, 0xb5, 0xf2, 0xe3, 0x12, 0x82, 0x18, 0x99, 0x9b, 0x5a, 0x92, 0x91, 0x9f, 0x02, 0x35,
	0x51, 0x0e, 0xc3, 0x44, 0x5f, 0xb0, 0x34, 0xba, 0x81, 0x02, 0x60, 0xbd, 0x10, 0x39, 0xb0, 0x79,
	0x4e, 0x02, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84,
	0xc7, 0x72, 0x0c, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x27, 0x26, 0x95, 0xfb, 0x00, 0x00,
	0x00,
}
