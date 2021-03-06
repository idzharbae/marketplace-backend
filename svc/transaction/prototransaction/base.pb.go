// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

package prototransaction

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
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Pagination struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pagination) Reset()         { *m = Pagination{} }
func (m *Pagination) String() string { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()    {}
func (*Pagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{0}
}

func (m *Pagination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pagination.Unmarshal(m, b)
}
func (m *Pagination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pagination.Marshal(b, m, deterministic)
}
func (m *Pagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pagination.Merge(m, src)
}
func (m *Pagination) XXX_Size() int {
	return xxx_messageInfo_Pagination.Size(m)
}
func (m *Pagination) XXX_DiscardUnknown() {
	xxx_messageInfo_Pagination.DiscardUnknown(m)
}

var xxx_messageInfo_Pagination proto.InternalMessageInfo

func (m *Pagination) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Pagination) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Pagination)(nil), "prototransaction.Pagination")
	proto.RegisterType((*Empty)(nil), "prototransaction.Empty")
}

func init() {
	proto.RegisterFile("base.proto", fileDescriptor_db1b6b0986796150)
}

var fileDescriptor_db1b6b0986796150 = []byte{
	// 108 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x00, 0x53, 0x25, 0x45, 0x89, 0x79, 0xc5, 0x89,
	0xc9, 0x25, 0x99, 0xf9, 0x79, 0x4a, 0x66, 0x5c, 0x5c, 0x01, 0x89, 0xe9, 0x99, 0x79, 0x89, 0x20,
	0x9e, 0x90, 0x10, 0x17, 0x4b, 0x41, 0x62, 0x7a, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10,
	0x98, 0x2d, 0x24, 0xc2, 0xc5, 0x9a, 0x93, 0x99, 0x9b, 0x59, 0x22, 0xc1, 0x04, 0x16, 0x84, 0x70,
	0x94, 0xd8, 0xb9, 0x58, 0x5d, 0x73, 0x0b, 0x4a, 0x2a, 0x93, 0xd8, 0xc0, 0x46, 0x1a, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xbe, 0x1c, 0x43, 0xcf, 0x67, 0x00, 0x00, 0x00,
}
