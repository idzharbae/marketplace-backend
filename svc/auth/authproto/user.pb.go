// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package authproto

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

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UserName             string   `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	PhotoUrl             string   `protobuf:"bytes,6,opt,name=photo_url,json=photoUrl,proto3" json:"photo_url,omitempty"`
	Type                 int32    `protobuf:"varint,7,opt,name=type,proto3" json:"type,omitempty"`
	Password             string   `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetPhotoUrl() string {
	if m != nil {
		return m.PhotoUrl
	}
	return ""
}

func (m *User) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "authproto.User")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8f, 0xbf, 0x0a, 0x83, 0x30,
	0x10, 0xc6, 0x89, 0xff, 0xaa, 0x37, 0x74, 0x08, 0x1d, 0x8e, 0x76, 0x91, 0x4e, 0x4e, 0x5d, 0xfa,
	0x1e, 0x1d, 0x04, 0x67, 0x49, 0x31, 0xa0, 0xa0, 0x26, 0x5c, 0x22, 0xa5, 0x6f, 0xd7, 0x47, 0x2b,
	0x77, 0xba, 0x7d, 0xdf, 0xef, 0x17, 0xbe, 0x70, 0x00, 0x5b, 0xb0, 0xf4, 0xf0, 0xe4, 0xa2, 0xd3,
	0x95, 0xd9, 0xe2, 0x28, 0xf1, 0xfe, 0x53, 0x90, 0x75, 0xc1, 0x92, 0x3e, 0x43, 0x32, 0x0d, 0xa8,
	0x6a, 0xd5, 0xa4, 0x6d, 0x32, 0x0d, 0x5a, 0x43, 0xb6, 0x9a, 0xc5, 0x62, 0x52, 0xab, 0xa6, 0x6a,
	0x25, 0xeb, 0x1b, 0x54, 0xbc, 0xd2, 0x8b, 0x48, 0x45, 0x94, 0x0c, 0x5e, 0x2c, 0x2f, 0x90, 0xdb,
	0xc5, 0x4c, 0x33, 0x66, 0x22, 0xf6, 0xc2, 0xd4, 0x8f, 0x6e, 0xb5, 0x98, 0xef, 0x54, 0x0a, 0x0f,
	0xf9, 0xd1, 0x45, 0xd7, 0x6f, 0x34, 0x63, 0xb1, 0x0f, 0x09, 0xe8, 0x68, 0xe6, 0x9f, 0xe3, 0xd7,
	0x5b, 0x3c, 0xd5, 0xaa, 0xc9, 0x5b, 0xc9, 0xfa, 0x0a, 0xa5, 0x37, 0x21, 0x7c, 0x1c, 0x0d, 0x58,
	0x1e, 0xef, 0x8f, 0xfe, 0x2e, 0xe4, 0x92, 0xe7, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x82, 0x35, 0xda,
	0x1e, 0xe2, 0x00, 0x00, 0x00,
}