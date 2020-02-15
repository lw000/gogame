// Code generated by protoc-gen-go. DO NOT EDIT.
// source: room.protos

package room

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protos package it is being compiled against.
// A compilation error at this line likely means your copy of the
// protos package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the protos package

type RequestJoinRoom struct {
	RoomId               int32    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	RoomName             string   `protobuf:"bytes,2,opt,name=roomName,proto3" json:"roomName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestJoinRoom) Reset()         { *m = RequestJoinRoom{} }
func (m *RequestJoinRoom) String() string { return proto.CompactTextString(m) }
func (*RequestJoinRoom) ProtoMessage()    {}
func (*RequestJoinRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{0}
}

func (m *RequestJoinRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestJoinRoom.Unmarshal(m, b)
}
func (m *RequestJoinRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestJoinRoom.Marshal(b, m, deterministic)
}
func (m *RequestJoinRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestJoinRoom.Merge(m, src)
}
func (m *RequestJoinRoom) XXX_Size() int {
	return xxx_messageInfo_RequestJoinRoom.Size(m)
}
func (m *RequestJoinRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestJoinRoom.DiscardUnknown(m)
}

var xxx_messageInfo_RequestJoinRoom proto.InternalMessageInfo

func (m *RequestJoinRoom) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *RequestJoinRoom) GetRoomName() string {
	if m != nil {
		return m.RoomName
	}
	return ""
}

type ReponseJoinRoom struct {
	RoomId               int32    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	RoomName             string   `protobuf:"bytes,2,opt,name=roomName,proto3" json:"roomName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReponseJoinRoom) Reset()         { *m = ReponseJoinRoom{} }
func (m *ReponseJoinRoom) String() string { return proto.CompactTextString(m) }
func (*ReponseJoinRoom) ProtoMessage()    {}
func (*ReponseJoinRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{1}
}

func (m *ReponseJoinRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReponseJoinRoom.Unmarshal(m, b)
}
func (m *ReponseJoinRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReponseJoinRoom.Marshal(b, m, deterministic)
}
func (m *ReponseJoinRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReponseJoinRoom.Merge(m, src)
}
func (m *ReponseJoinRoom) XXX_Size() int {
	return xxx_messageInfo_ReponseJoinRoom.Size(m)
}
func (m *ReponseJoinRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_ReponseJoinRoom.DiscardUnknown(m)
}

var xxx_messageInfo_ReponseJoinRoom proto.InternalMessageInfo

func (m *ReponseJoinRoom) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *ReponseJoinRoom) GetRoomName() string {
	if m != nil {
		return m.RoomName
	}
	return ""
}

type RequestLeaveRoom struct {
	RoomId               int32    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestLeaveRoom) Reset()         { *m = RequestLeaveRoom{} }
func (m *RequestLeaveRoom) String() string { return proto.CompactTextString(m) }
func (*RequestLeaveRoom) ProtoMessage()    {}
func (*RequestLeaveRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{2}
}

func (m *RequestLeaveRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestLeaveRoom.Unmarshal(m, b)
}
func (m *RequestLeaveRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestLeaveRoom.Marshal(b, m, deterministic)
}
func (m *RequestLeaveRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestLeaveRoom.Merge(m, src)
}
func (m *RequestLeaveRoom) XXX_Size() int {
	return xxx_messageInfo_RequestLeaveRoom.Size(m)
}
func (m *RequestLeaveRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestLeaveRoom.DiscardUnknown(m)
}

var xxx_messageInfo_RequestLeaveRoom proto.InternalMessageInfo

func (m *RequestLeaveRoom) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

type ReponseLeaveRoom struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReponseLeaveRoom) Reset()         { *m = ReponseLeaveRoom{} }
func (m *ReponseLeaveRoom) String() string { return proto.CompactTextString(m) }
func (*ReponseLeaveRoom) ProtoMessage()    {}
func (*ReponseLeaveRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{3}
}

func (m *ReponseLeaveRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReponseLeaveRoom.Unmarshal(m, b)
}
func (m *ReponseLeaveRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReponseLeaveRoom.Marshal(b, m, deterministic)
}
func (m *ReponseLeaveRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReponseLeaveRoom.Merge(m, src)
}
func (m *ReponseLeaveRoom) XXX_Size() int {
	return xxx_messageInfo_ReponseLeaveRoom.Size(m)
}
func (m *ReponseLeaveRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_ReponseLeaveRoom.DiscardUnknown(m)
}

var xxx_messageInfo_ReponseLeaveRoom proto.InternalMessageInfo

func (m *ReponseLeaveRoom) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestJoinRoom)(nil), "room.RequestJoinRoom")
	proto.RegisterType((*ReponseJoinRoom)(nil), "room.ReponseJoinRoom")
	proto.RegisterType((*RequestLeaveRoom)(nil), "room.RequestLeaveRoom")
	proto.RegisterType((*ReponseLeaveRoom)(nil), "room.ReponseLeaveRoom")
}

func init() { proto.RegisterFile("room.protos", fileDescriptor_c5fd27dd97284ef4) }

var fileDescriptor_c5fd27dd97284ef4 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xca, 0xcf, 0xcf,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x5c, 0xb9, 0xf8, 0x83, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xbc, 0xf2, 0x33, 0xf3, 0x82, 0xf2, 0xf3, 0x73, 0x85, 0xc4, 0xb8,
	0xd8, 0x40, 0x52, 0x9e, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x50, 0x9e, 0x90, 0x14,
	0x17, 0x07, 0x88, 0xe5, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe7,
	0x43, 0x8c, 0x29, 0xc8, 0xcf, 0x2b, 0x4e, 0xa5, 0xc8, 0x18, 0x2d, 0x2e, 0x01, 0xa8, 0x6b, 0x7c,
	0x52, 0x13, 0xcb, 0x52, 0xf1, 0x99, 0xa3, 0xa4, 0x06, 0x52, 0x0b, 0xb6, 0x12, 0xa1, 0x56, 0x88,
	0x8b, 0x25, 0x39, 0x3f, 0x25, 0x15, 0xaa, 0x12, 0xcc, 0x4e, 0x62, 0x03, 0x7b, 0xd7, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x73, 0x63, 0xa3, 0x88, 0xfc, 0x00, 0x00, 0x00,
}
