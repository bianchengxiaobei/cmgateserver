// Code generated by protoc-gen-go. DO NOT EDIT.
// source: G2C.proto

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

type G2C_CharacterInfo struct {
	Role                 *Role    `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *G2C_CharacterInfo) Reset()         { *m = G2C_CharacterInfo{} }
func (m *G2C_CharacterInfo) String() string { return proto.CompactTextString(m) }
func (*G2C_CharacterInfo) ProtoMessage()    {}
func (*G2C_CharacterInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_G2C_8ecf644d67efc838, []int{0}
}
func (m *G2C_CharacterInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_G2C_CharacterInfo.Unmarshal(m, b)
}
func (m *G2C_CharacterInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_G2C_CharacterInfo.Marshal(b, m, deterministic)
}
func (dst *G2C_CharacterInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_G2C_CharacterInfo.Merge(dst, src)
}
func (m *G2C_CharacterInfo) XXX_Size() int {
	return xxx_messageInfo_G2C_CharacterInfo.Size(m)
}
func (m *G2C_CharacterInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_G2C_CharacterInfo.DiscardUnknown(m)
}

var xxx_messageInfo_G2C_CharacterInfo proto.InternalMessageInfo

func (m *G2C_CharacterInfo) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type Role struct {
	RoleId               int64    `protobuf:"varint,1,opt,name=roleId,proto3" json:"roleId,omitempty"`
	NickName             string   `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"`
	AvatarId             int32    `protobuf:"varint,3,opt,name=avatarId,proto3" json:"avatarId,omitempty"`
	Level                int32    `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	Exp                  int32    `protobuf:"varint,5,opt,name=exp,proto3" json:"exp,omitempty"`
	Gold                 int32    `protobuf:"varint,6,opt,name=gold,proto3" json:"gold,omitempty"`
	Diam                 int32    `protobuf:"varint,7,opt,name=diam,proto3" json:"diam,omitempty"`
	MaxBagNum            int32    `protobuf:"varint,8,opt,name=maxBagNum,proto3" json:"maxBagNum,omitempty"`
	RankScore            int32    `protobuf:"varint,9,opt,name=rankScore,proto3" json:"rankScore,omitempty"`
	Sex                  int32    `protobuf:"varint,10,opt,name=sex,proto3" json:"sex,omitempty"`
	Sign                 string   `protobuf:"bytes,11,opt,name=sign,proto3" json:"sign,omitempty"`
	Qq                   int32    `protobuf:"varint,12,opt,name=qq,proto3" json:"qq,omitempty"`
	Phone                int32    `protobuf:"varint,13,opt,name=phone,proto3" json:"phone,omitempty"`
	MaillAddress         string   `protobuf:"bytes,14,opt,name=maillAddress,proto3" json:"maillAddress,omitempty"`
	Weixin               string   `protobuf:"bytes,15,opt,name=weixin,proto3" json:"weixin,omitempty"`
	GuideFinished        bool     `protobuf:"varint,16,opt,name=guideFinished,proto3" json:"guideFinished,omitempty"`
	SaijiId              int32    `protobuf:"varint,17,opt,name=saijiId,proto3" json:"saijiId,omitempty"`
	SelfMotionOpen       bool     `protobuf:"varint,18,opt,name=selfMotionOpen,proto3" json:"selfMotionOpen,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_G2C_8ecf644d67efc838, []int{1}
}
func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (dst *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(dst, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetRoleId() int64 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *Role) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *Role) GetAvatarId() int32 {
	if m != nil {
		return m.AvatarId
	}
	return 0
}

func (m *Role) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Role) GetExp() int32 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func (m *Role) GetGold() int32 {
	if m != nil {
		return m.Gold
	}
	return 0
}

func (m *Role) GetDiam() int32 {
	if m != nil {
		return m.Diam
	}
	return 0
}

func (m *Role) GetMaxBagNum() int32 {
	if m != nil {
		return m.MaxBagNum
	}
	return 0
}

func (m *Role) GetRankScore() int32 {
	if m != nil {
		return m.RankScore
	}
	return 0
}

func (m *Role) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *Role) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func (m *Role) GetQq() int32 {
	if m != nil {
		return m.Qq
	}
	return 0
}

func (m *Role) GetPhone() int32 {
	if m != nil {
		return m.Phone
	}
	return 0
}

func (m *Role) GetMaillAddress() string {
	if m != nil {
		return m.MaillAddress
	}
	return ""
}

func (m *Role) GetWeixin() string {
	if m != nil {
		return m.Weixin
	}
	return ""
}

func (m *Role) GetGuideFinished() bool {
	if m != nil {
		return m.GuideFinished
	}
	return false
}

func (m *Role) GetSaijiId() int32 {
	if m != nil {
		return m.SaijiId
	}
	return 0
}

func (m *Role) GetSelfMotionOpen() bool {
	if m != nil {
		return m.SelfMotionOpen
	}
	return false
}

// 强制客户端退出游戏
type G2C_QuitGame struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *G2C_QuitGame) Reset()         { *m = G2C_QuitGame{} }
func (m *G2C_QuitGame) String() string { return proto.CompactTextString(m) }
func (*G2C_QuitGame) ProtoMessage()    {}
func (*G2C_QuitGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_G2C_8ecf644d67efc838, []int{2}
}
func (m *G2C_QuitGame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_G2C_QuitGame.Unmarshal(m, b)
}
func (m *G2C_QuitGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_G2C_QuitGame.Marshal(b, m, deterministic)
}
func (dst *G2C_QuitGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_G2C_QuitGame.Merge(dst, src)
}
func (m *G2C_QuitGame) XXX_Size() int {
	return xxx_messageInfo_G2C_QuitGame.Size(m)
}
func (m *G2C_QuitGame) XXX_DiscardUnknown() {
	xxx_messageInfo_G2C_QuitGame.DiscardUnknown(m)
}

var xxx_messageInfo_G2C_QuitGame proto.InternalMessageInfo

// 即将关闭服务器
type G2C_NeedCloseServer struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *G2C_NeedCloseServer) Reset()         { *m = G2C_NeedCloseServer{} }
func (m *G2C_NeedCloseServer) String() string { return proto.CompactTextString(m) }
func (*G2C_NeedCloseServer) ProtoMessage()    {}
func (*G2C_NeedCloseServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_G2C_8ecf644d67efc838, []int{3}
}
func (m *G2C_NeedCloseServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_G2C_NeedCloseServer.Unmarshal(m, b)
}
func (m *G2C_NeedCloseServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_G2C_NeedCloseServer.Marshal(b, m, deterministic)
}
func (dst *G2C_NeedCloseServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_G2C_NeedCloseServer.Merge(dst, src)
}
func (m *G2C_NeedCloseServer) XXX_Size() int {
	return xxx_messageInfo_G2C_NeedCloseServer.Size(m)
}
func (m *G2C_NeedCloseServer) XXX_DiscardUnknown() {
	xxx_messageInfo_G2C_NeedCloseServer.DiscardUnknown(m)
}

var xxx_messageInfo_G2C_NeedCloseServer proto.InternalMessageInfo

func init() {
	proto.RegisterType((*G2C_CharacterInfo)(nil), "message.G2C_CharacterInfo")
	proto.RegisterType((*Role)(nil), "message.Role")
	proto.RegisterType((*G2C_QuitGame)(nil), "message.G2C_QuitGame")
	proto.RegisterType((*G2C_NeedCloseServer)(nil), "message.G2C_NeedCloseServer")
}

func init() { proto.RegisterFile("G2C.proto", fileDescriptor_G2C_8ecf644d67efc838) }

var fileDescriptor_G2C_8ecf644d67efc838 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc7, 0x95, 0x7e, 0x77, 0xb6, 0x2d, 0xbb, 0xe6, 0x43, 0x23, 0xc4, 0xa1, 0x44, 0x08, 0xf5,
	0xd4, 0x43, 0x91, 0xb8, 0x43, 0x25, 0xaa, 0x1e, 0x28, 0x22, 0xfb, 0x00, 0xc8, 0xd4, 0xb3, 0xa9,
	0x59, 0xc7, 0x4e, 0xec, 0xb4, 0xe4, 0x39, 0x79, 0x22, 0xe4, 0x49, 0x29, 0xea, 0xde, 0xe6, 0xf7,
	0xfb, 0x8f, 0x3c, 0x1a, 0x6b, 0x60, 0xbc, 0x59, 0xad, 0x97, 0xa5, 0x77, 0xb5, 0x13, 0xc3, 0x82,
	0x42, 0x90, 0x39, 0xa5, 0x1f, 0xe1, 0x6e, 0xb3, 0x5a, 0xff, 0x58, 0x1f, 0xa4, 0x97, 0xfb, 0x9a,
	0xfc, 0xd6, 0x3e, 0x38, 0xf1, 0x16, 0x7a, 0xde, 0x19, 0xc2, 0x64, 0x9e, 0x2c, 0x6e, 0x56, 0xd3,
	0xe5, 0xb9, 0x79, 0x99, 0x39, 0x43, 0x19, 0x47, 0xe9, 0x9f, 0x2e, 0xf4, 0x22, 0x8a, 0x57, 0x30,
	0x88, 0x62, 0xab, 0xb8, 0xbb, 0x9b, 0x9d, 0x49, 0xbc, 0x86, 0x91, 0xd5, 0xfb, 0xc7, 0x9d, 0x2c,
	0x08, 0x3b, 0xf3, 0x64, 0x31, 0xce, 0x2e, 0x1c, 0x33, 0x79, 0x92, 0xb5, 0xf4, 0x5b, 0x85, 0xdd,
	0x79, 0xb2, 0xe8, 0x67, 0x17, 0x16, 0x2f, 0xa0, 0x6f, 0xe8, 0x44, 0x06, 0x7b, 0x1c, 0xb4, 0x20,
	0x6e, 0xa1, 0x4b, 0x4d, 0x89, 0x7d, 0x76, 0xb1, 0x14, 0x02, 0x7a, 0xb9, 0x33, 0x0a, 0x07, 0xac,
	0xb8, 0x8e, 0x4e, 0x69, 0x59, 0xe0, 0xb0, 0x75, 0xb1, 0x16, 0x6f, 0x60, 0x5c, 0xc8, 0xe6, 0xb3,
	0xcc, 0x77, 0xc7, 0x02, 0x47, 0x1c, 0xfc, 0x17, 0x31, 0xf5, 0xd2, 0x3e, 0xde, 0xef, 0x9d, 0x27,
	0x1c, 0xb7, 0xe9, 0x45, 0xc4, 0xa9, 0x81, 0x1a, 0x84, 0x76, 0x6a, 0xa0, 0x26, 0x4e, 0x08, 0x3a,
	0xb7, 0x78, 0xc3, 0x1b, 0x71, 0x2d, 0x66, 0xd0, 0xa9, 0x2a, 0x9c, 0x70, 0x53, 0xa7, 0xaa, 0xe2,
	0x06, 0xe5, 0xc1, 0x59, 0xc2, 0x69, 0xbb, 0x01, 0x83, 0x48, 0x61, 0x52, 0x48, 0x6d, 0xcc, 0x27,
	0xa5, 0x3c, 0x85, 0x80, 0x33, 0x7e, 0xe1, 0xca, 0xc5, 0xbf, 0xfc, 0x4d, 0xba, 0xd1, 0x16, 0x9f,
	0x71, 0x7a, 0x26, 0xf1, 0x0e, 0xa6, 0xf9, 0x51, 0x2b, 0xfa, 0xa2, 0xad, 0x0e, 0x07, 0x52, 0x78,
	0x3b, 0x4f, 0x16, 0xa3, 0xec, 0x5a, 0x0a, 0x84, 0x61, 0x90, 0xfa, 0x97, 0xde, 0x2a, 0xbc, 0xe3,
	0xc9, 0xff, 0x50, 0xbc, 0x87, 0x59, 0x20, 0xf3, 0xf0, 0xd5, 0xd5, 0xda, 0xd9, 0x6f, 0x25, 0x59,
	0x14, 0xfc, 0xc0, 0x13, 0x9b, 0xce, 0x60, 0x12, 0x8f, 0xe1, 0xfb, 0x51, 0xd7, 0x1b, 0x59, 0x50,
	0xfa, 0x12, 0x9e, 0x47, 0xde, 0x11, 0xa9, 0xb5, 0x71, 0x81, 0xee, 0xc9, 0x9f, 0xc8, 0xff, 0x1c,
	0xf0, 0x0d, 0x7d, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x85, 0x48, 0xb4, 0x9c, 0x50, 0x02, 0x00,
	0x00,
}
