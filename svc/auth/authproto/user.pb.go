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
	CreatedAt            int64    `protobuf:"varint,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Province             string   `protobuf:"bytes,11,opt,name=province,proto3" json:"province,omitempty"`
	City                 string   `protobuf:"bytes,12,opt,name=city,proto3" json:"city,omitempty"`
	AddressDetail        string   `protobuf:"bytes,13,opt,name=address_detail,json=addressDetail,proto3" json:"address_detail,omitempty"`
	ZipCode              int32    `protobuf:"varint,14,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	Description          string   `protobuf:"bytes,15,opt,name=description,proto3" json:"description,omitempty"`
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

func (m *User) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *User) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *User) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *User) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *User) GetAddressDetail() string {
	if m != nil {
		return m.AddressDetail
	}
	return ""
}

func (m *User) GetZipCode() int32 {
	if m != nil {
		return m.ZipCode
	}
	return 0
}

func (m *User) GetDescription() string {
	if m != nil {
		return m.Description
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
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xc1, 0x6a, 0xeb, 0x30,
	0x10, 0x45, 0x71, 0x12, 0x27, 0xf6, 0xe4, 0xc5, 0x0f, 0x44, 0x17, 0xd3, 0x96, 0x82, 0x29, 0x14,
	0xbc, 0xea, 0xa6, 0x5f, 0x10, 0xda, 0x75, 0x17, 0x86, 0xac, 0x8d, 0x6a, 0x0d, 0x58, 0xe0, 0x58,
	0x42, 0x92, 0x5b, 0x92, 0xcf, 0xe8, 0x17, 0x17, 0x8d, 0x55, 0xe8, 0x6e, 0xee, 0x39, 0xf6, 0xb5,
	0x47, 0x02, 0x98, 0x3d, 0xb9, 0x67, 0xeb, 0x4c, 0x30, 0xa2, 0x94, 0x73, 0x18, 0x78, 0x7c, 0xfc,
	0x5e, 0xc3, 0xe6, 0xe4, 0xc9, 0x89, 0x0a, 0x56, 0x5a, 0x61, 0x56, 0x67, 0xcd, 0xba, 0x5d, 0x69,
	0x25, 0x04, 0x6c, 0x26, 0x79, 0x26, 0x5c, 0xd5, 0x59, 0x53, 0xb6, 0x3c, 0x8b, 0x7b, 0x28, 0x63,
	0x4b, 0xc7, 0x62, 0xcd, 0xa2, 0x88, 0xe0, 0x3d, 0xca, 0x1b, 0xc8, 0xe9, 0x2c, 0xf5, 0x88, 0x1b,
	0x16, 0x4b, 0x88, 0xd4, 0x0e, 0x66, 0x22, 0xcc, 0x17, 0xca, 0x21, 0x16, 0xd9, 0xc1, 0x04, 0xd3,
	0xcd, 0x6e, 0xc4, 0xed, 0x52, 0xc4, 0xe0, 0xe4, 0xc6, 0xf8, 0xe5, 0x70, 0xb1, 0x84, 0xbb, 0x3a,
	0x6b, 0xf2, 0x96, 0x67, 0x71, 0x07, 0x85, 0x95, 0xde, 0x7f, 0x19, 0xa7, 0xb0, 0x48, 0xcf, 0xa7,
	0x2c, 0x1e, 0x00, 0x7a, 0x47, 0x32, 0x90, 0xea, 0x64, 0xc0, 0x92, 0x37, 0x28, 0x13, 0x39, 0x86,
	0xa8, 0x67, 0xab, 0x7e, 0x35, 0x2c, 0x3a, 0x91, 0x63, 0xe0, 0x66, 0x67, 0x3e, 0xf5, 0xd4, 0x13,
	0xee, 0x53, 0x73, 0xca, 0xf1, 0x4f, 0x7a, 0x1d, 0x2e, 0xf8, 0x6f, 0x39, 0x83, 0x38, 0x8b, 0x27,
	0xa8, 0xa4, 0x52, 0x8e, 0xbc, 0xef, 0x14, 0x85, 0xb8, 0xef, 0x81, 0xed, 0x21, 0xd1, 0x37, 0x86,
	0xe2, 0x16, 0x8a, 0xab, 0xb6, 0x5d, 0x6f, 0x14, 0x61, 0xc5, 0x8b, 0xec, 0xae, 0xda, 0xbe, 0x1a,
	0x45, 0xa2, 0x86, 0xbd, 0x22, 0xdf, 0x3b, 0x6d, 0x83, 0x36, 0x13, 0xfe, 0xe7, 0xd7, 0xff, 0xa2,
	0x8f, 0x2d, 0xdf, 0xcd, 0xcb, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0xf2, 0x3b, 0x23, 0xb4,
	0x01, 0x00, 0x00,
}
