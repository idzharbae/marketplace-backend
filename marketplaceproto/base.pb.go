// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

package marketplaceproto

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

type Pagination struct {
	Page                 int64    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pagination) Reset()         { *m = Pagination{} }
func (m *Pagination) String() string { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()    {}
func (*Pagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_140ff5433226dbde, []int{0}
}
func (m *Pagination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pagination.Unmarshal(m, b)
}
func (m *Pagination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pagination.Marshal(b, m, deterministic)
}
func (dst *Pagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pagination.Merge(dst, src)
}
func (m *Pagination) XXX_Size() int {
	return xxx_messageInfo_Pagination.Size(m)
}
func (m *Pagination) XXX_DiscardUnknown() {
	xxx_messageInfo_Pagination.DiscardUnknown(m)
}

var xxx_messageInfo_Pagination proto.InternalMessageInfo

func (m *Pagination) GetPage() int64 {
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
	return fileDescriptor_base_140ff5433226dbde, []int{1}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Pagination)(nil), "marketplaceproto.Pagination")
	proto.RegisterType((*Empty)(nil), "marketplaceproto.Empty")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor_base_140ff5433226dbde) }

var fileDescriptor_base_140ff5433226dbde = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc8, 0x4d, 0x2c, 0xca, 0x4e, 0x2d, 0x29, 0xc8,
	0x49, 0x4c, 0x4e, 0x05, 0x8b, 0x28, 0x99, 0x71, 0x71, 0x05, 0x24, 0xa6, 0x67, 0xe6, 0x25, 0x96,
	0x64, 0xe6, 0xe7, 0x09, 0x09, 0x71, 0xb1, 0x14, 0x24, 0xa6, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0x30, 0x07, 0x81, 0xd9, 0x42, 0x22, 0x5c, 0xac, 0x39, 0x99, 0xb9, 0x99, 0x25, 0x12, 0x4c, 0x0a,
	0x8c, 0x1a, 0xac, 0x41, 0x10, 0x8e, 0x12, 0x3b, 0x17, 0xab, 0x6b, 0x6e, 0x41, 0x49, 0x65, 0x12,
	0x1b, 0xd8, 0x1c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xff, 0x09, 0x7a, 0x67, 0x00,
	0x00, 0x00,
}
