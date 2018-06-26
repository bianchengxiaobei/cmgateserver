// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RegisterGate.proto

package message

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

type M2G_RegisterGate struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2G_RegisterGate) Reset()         { *m = M2G_RegisterGate{} }
func (m *M2G_RegisterGate) String() string { return proto.CompactTextString(m) }
func (m *M2G_RegisterGate) ProtoMessage()	{}
func (m *M2G_RegisterGate) Descriptor() ([]byte, []int) {
	return fileDescriptor_RegisterGate_c6490c50607574f8, []int{0}
}
func (m *M2G_RegisterGate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2G_RegisterGate.Unmarshal(m, b)
}
func (m *M2G_RegisterGate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2G_RegisterGate.Marshal(b, m, deterministic)
}
func (dst *M2G_RegisterGate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2G_RegisterGate.Merge(dst, src)
}
func (m *M2G_RegisterGate) XXX_Size() int {
	return xxx_messageInfo_M2G_RegisterGate.Size(m)
}
func (m *M2G_RegisterGate) XXX_DiscardUnknown() {
	xxx_messageInfo_M2G_RegisterGate.DiscardUnknown(m)
}

var xxx_messageInfo_M2G_RegisterGate proto.InternalMessageInfo

func (m *M2G_RegisterGate) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*M2G_RegisterGate)(nil), "message.M2G_RegisterGate")
}

func init() { proto.RegisterFile("RegisterGate.proto", fileDescriptor_RegisterGate_c6490c50607574f8) }

var fileDescriptor_RegisterGate_c6490c50607574f8 = []byte{
	// 84 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0a, 0x4a, 0x4d, 0xcf,
	0x2c, 0x2e, 0x49, 0x2d, 0x72, 0x4f, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x55, 0x52, 0xe2, 0x12, 0xf0, 0x35, 0x72, 0x8f, 0x47,
	0x56, 0x22, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1a, 0xc4, 0x94,
	0x99, 0x92, 0xc4, 0x06, 0xd6, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x14, 0x57, 0xd5,
	0x49, 0x00, 0x00, 0x00,
}
