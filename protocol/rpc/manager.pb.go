// Code generated by protoc-gen-go.
// source: manager.proto
// DO NOT EDIT!

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	manager.proto
	msg_server.proto
	router.proto

It has these top-level messages:
	LoginReq
	LoginRes
	SendP2PMsgReq
	SendP2PMsgRes
*/
package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

func init() { proto.RegisterFile("manager.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 49 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0xcc, 0x4b,
	0x4c, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x4e, 0x62,
	0x03, 0xb3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x14, 0x5f, 0x65, 0x7c, 0x1c, 0x00, 0x00,
	0x00,
}
