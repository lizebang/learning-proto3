// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/lizebang/learning-proto3/pkg_b/file_b.proto

package pkg_b

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	pkg_c "github.com/lizebang/learning-proto3/pkg_c"
	math "math"
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

// MsgA from public import github.com/lizebang/learning-proto3/pkg_c/file_c.proto
type MsgA = pkg_c.MsgA

func init() {
	proto.RegisterFile("github.com/lizebang/learning-proto3/pkg_b/file_b.proto", fileDescriptor_0b505f3417945aa2)
}

var fileDescriptor_0b505f3417945aa2 = []byte{
	// 96 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xc9, 0xac, 0x4a, 0x4d, 0x4a, 0xcc, 0x4b, 0xd7,
	0xcf, 0x49, 0x4d, 0x2c, 0xca, 0xcb, 0xcc, 0x4b, 0xd7, 0x2d, 0x28, 0xca, 0x2f, 0xc9, 0x37, 0xd6,
	0x2f, 0xc8, 0x4e, 0x8f, 0x4f, 0xd2, 0x4f, 0xcb, 0xcc, 0x49, 0x8d, 0x4f, 0xd2, 0x03, 0x8b, 0x09,
	0xb1, 0x82, 0xc5, 0xa4, 0x88, 0xd6, 0x9e, 0x0c, 0xd1, 0x9e, 0x0c, 0xd1, 0x1e, 0xc0, 0x90, 0xc4,
	0x06, 0x91, 0x04, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xf9, 0x31, 0x91, 0x81, 0x00, 0x00, 0x00,
}
