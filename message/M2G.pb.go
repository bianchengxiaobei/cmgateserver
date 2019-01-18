// Code generated by protoc-gen-go. DO NOT EDIT.
// source: M2G.proto

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
func (*M2G_RegisterGate) ProtoMessage()    {}
func (*M2G_RegisterGate) Descriptor() ([]byte, []int) {
	return fileDescriptor_M2G_07fac8ac190ce268, []int{0}
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

type M2G_LoginSuccessNotifyGate struct {
	ServerId             int32    `protobuf:"varint,1,opt,name=serverId,proto3" json:"serverId,omitempty"`
	RoleId               int64    `protobuf:"varint,2,opt,name=roleId,proto3" json:"roleId,omitempty"`
	UserId               int64    `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2G_LoginSuccessNotifyGate) Reset()         { *m = M2G_LoginSuccessNotifyGate{} }
func (m *M2G_LoginSuccessNotifyGate) String() string { return proto.CompactTextString(m) }
func (*M2G_LoginSuccessNotifyGate) ProtoMessage()    {}
func (*M2G_LoginSuccessNotifyGate) Descriptor() ([]byte, []int) {
	return fileDescriptor_M2G_07fac8ac190ce268, []int{1}
}
func (m *M2G_LoginSuccessNotifyGate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2G_LoginSuccessNotifyGate.Unmarshal(m, b)
}
func (m *M2G_LoginSuccessNotifyGate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2G_LoginSuccessNotifyGate.Marshal(b, m, deterministic)
}
func (dst *M2G_LoginSuccessNotifyGate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2G_LoginSuccessNotifyGate.Merge(dst, src)
}
func (m *M2G_LoginSuccessNotifyGate) XXX_Size() int {
	return xxx_messageInfo_M2G_LoginSuccessNotifyGate.Size(m)
}
func (m *M2G_LoginSuccessNotifyGate) XXX_DiscardUnknown() {
	xxx_messageInfo_M2G_LoginSuccessNotifyGate.DiscardUnknown(m)
}

var xxx_messageInfo_M2G_LoginSuccessNotifyGate proto.InternalMessageInfo

func (m *M2G_LoginSuccessNotifyGate) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *M2G_LoginSuccessNotifyGate) GetRoleId() int64 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *M2G_LoginSuccessNotifyGate) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type M2G_RoleQuitGate struct {
	RoleId               int64    `protobuf:"varint,1,opt,name=roleId,proto3" json:"roleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2G_RoleQuitGate) Reset()         { *m = M2G_RoleQuitGate{} }
func (m *M2G_RoleQuitGate) String() string { return proto.CompactTextString(m) }
func (*M2G_RoleQuitGate) ProtoMessage()    {}
func (*M2G_RoleQuitGate) Descriptor() ([]byte, []int) {
	return fileDescriptor_M2G_07fac8ac190ce268, []int{2}
}
func (m *M2G_RoleQuitGate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2G_RoleQuitGate.Unmarshal(m, b)
}
func (m *M2G_RoleQuitGate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2G_RoleQuitGate.Marshal(b, m, deterministic)
}
func (dst *M2G_RoleQuitGate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2G_RoleQuitGate.Merge(dst, src)
}
func (m *M2G_RoleQuitGate) XXX_Size() int {
	return xxx_messageInfo_M2G_RoleQuitGate.Size(m)
}
func (m *M2G_RoleQuitGate) XXX_DiscardUnknown() {
	xxx_messageInfo_M2G_RoleQuitGate.DiscardUnknown(m)
}

var xxx_messageInfo_M2G_RoleQuitGate proto.InternalMessageInfo

func (m *M2G_RoleQuitGate) GetRoleId() int64 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

// 关闭session
type M2G_CloseSession struct {
	RoleId               int64    `protobuf:"varint,1,opt,name=roleId,proto3" json:"roleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2G_CloseSession) Reset()         { *m = M2G_CloseSession{} }
func (m *M2G_CloseSession) String() string { return proto.CompactTextString(m) }
func (*M2G_CloseSession) ProtoMessage()    {}
func (*M2G_CloseSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_M2G_07fac8ac190ce268, []int{3}
}
func (m *M2G_CloseSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2G_CloseSession.Unmarshal(m, b)
}
func (m *M2G_CloseSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2G_CloseSession.Marshal(b, m, deterministic)
}
func (dst *M2G_CloseSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2G_CloseSession.Merge(dst, src)
}
func (m *M2G_CloseSession) XXX_Size() int {
	return xxx_messageInfo_M2G_CloseSession.Size(m)
}
func (m *M2G_CloseSession) XXX_DiscardUnknown() {
	xxx_messageInfo_M2G_CloseSession.DiscardUnknown(m)
}

var xxx_messageInfo_M2G_CloseSession proto.InternalMessageInfo

func (m *M2G_CloseSession) GetRoleId() int64 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func init() {
	proto.RegisterType((*M2G_RegisterGate)(nil), "message.M2G_RegisterGate")
	proto.RegisterType((*M2G_LoginSuccessNotifyGate)(nil), "message.M2G_LoginSuccessNotifyGate")
	proto.RegisterType((*M2G_RoleQuitGate)(nil), "message.M2G_RoleQuitGate")
	proto.RegisterType((*M2G_CloseSession)(nil), "message.M2G_CloseSession")
}

func init() { proto.RegisterFile("M2G.proto", fileDescriptor_M2G_07fac8ac190ce268) }

var fileDescriptor_M2G_07fac8ac190ce268 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xf4, 0x35, 0x72, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x55, 0x52,
	0xe2, 0x12, 0xf0, 0x35, 0x72, 0x8f, 0x0f, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x49, 0x2d, 0x72, 0x4f,
	0x2c, 0x49, 0x15, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62,
	0xca, 0x4c, 0x51, 0xca, 0xe0, 0x92, 0x02, 0xa9, 0xf1, 0xc9, 0x4f, 0xcf, 0xcc, 0x0b, 0x2e, 0x4d,
	0x4e, 0x4e, 0x2d, 0x2e, 0xf6, 0xcb, 0x2f, 0xc9, 0x4c, 0xab, 0x04, 0xab, 0x96, 0xe2, 0xe2, 0x28,
	0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xf2, 0x84, 0xe9, 0x81, 0xf3, 0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0xf2,
	0x73, 0x52, 0x3d, 0x53, 0x24, 0x98, 0x14, 0x18, 0x35, 0x98, 0x83, 0xa0, 0x3c, 0x90, 0x78, 0x69,
	0x31, 0x58, 0x07, 0x33, 0x44, 0x1c, 0xc2, 0x53, 0xd2, 0x82, 0xba, 0x26, 0x3f, 0x27, 0x35, 0xb0,
	0x34, 0xb3, 0x04, 0x6c, 0x3e, 0xc2, 0x0c, 0x46, 0x64, 0x33, 0x60, 0x6a, 0x9d, 0x73, 0xf2, 0x8b,
	0x53, 0x83, 0x53, 0x8b, 0x8b, 0x33, 0xf3, 0xf3, 0x70, 0xa9, 0x4d, 0x62, 0x03, 0xfb, 0xda, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x92, 0xe0, 0x3d, 0x02, 0x01, 0x00, 0x00,
}
