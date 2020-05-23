// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

package catalogproto

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

type Product struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ShopId               int32    `protobuf:"varint,2,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Quantity             int32    `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	PricePerKg           int32    `protobuf:"varint,5,opt,name=price_per_kg,json=pricePerKg,proto3" json:"price_per_kg,omitempty"`
	StockKg              float32  `protobuf:"fixed32,6,opt,name=stock_kg,json=stockKg,proto3" json:"stock_kg,omitempty"`
	CreatedAt            int64    `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Slug                 string   `protobuf:"bytes,9,opt,name=slug,proto3" json:"slug,omitempty"`
	PhotoUrl             string   `protobuf:"bytes,10,opt,name=photo_url,json=photoUrl,proto3" json:"photo_url,omitempty"`
	Description          string   `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	Category             string   `protobuf:"bytes,12,opt,name=category,proto3" json:"category,omitempty"`
	TotalReviews         int32    `protobuf:"varint,13,opt,name=total_reviews,json=totalReviews,proto3" json:"total_reviews,omitempty"`
	AverageRating        float32  `protobuf:"fixed32,14,opt,name=average_rating,json=averageRating,proto3" json:"average_rating,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{0}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Product) GetShopId() int32 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *Product) GetPricePerKg() int32 {
	if m != nil {
		return m.PricePerKg
	}
	return 0
}

func (m *Product) GetStockKg() float32 {
	if m != nil {
		return m.StockKg
	}
	return 0
}

func (m *Product) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Product) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Product) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Product) GetPhotoUrl() string {
	if m != nil {
		return m.PhotoUrl
	}
	return ""
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Product) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Product) GetTotalReviews() int32 {
	if m != nil {
		return m.TotalReviews
	}
	return 0
}

func (m *Product) GetAverageRating() float32 {
	if m != nil {
		return m.AverageRating
	}
	return 0
}

func init() {
	proto.RegisterType((*Product)(nil), "catalogproto.Product")
}

func init() {
	proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5)
}

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcd, 0x6a, 0xeb, 0x30,
	0x10, 0x85, 0xb1, 0xf3, 0x63, 0x7b, 0x62, 0x67, 0xa1, 0xcd, 0xd5, 0xbd, 0x97, 0x82, 0x69, 0x29,
	0x78, 0xd5, 0x4d, 0x9f, 0x20, 0xcb, 0x92, 0x4d, 0x30, 0x74, 0x2d, 0x54, 0x49, 0x28, 0x22, 0xae,
	0xa5, 0xca, 0xe3, 0x94, 0xbc, 0x49, 0x1f, 0xb7, 0x78, 0x9c, 0x94, 0xee, 0xe6, 0x7c, 0x9f, 0x04,
	0x73, 0x18, 0xa8, 0x42, 0xf4, 0x7a, 0x54, 0xf8, 0x14, 0xa2, 0x47, 0xcf, 0x4a, 0x25, 0x51, 0x76,
	0xde, 0x52, 0xba, 0xff, 0x5a, 0x40, 0x76, 0x98, 0x3d, 0xdb, 0x42, 0xea, 0x34, 0x4f, 0xea, 0xa4,
	0x59, 0xb5, 0xa9, 0xd3, 0xec, 0x0f, 0x64, 0xc3, 0xd1, 0x07, 0xe1, 0x34, 0x4f, 0x09, 0xae, 0xa7,
	0xf8, 0xa2, 0x19, 0x83, 0x65, 0x2f, 0xdf, 0x0d, 0x5f, 0xd4, 0x49, 0x53, 0xb4, 0x34, 0xb3, 0x7f,
	0x90, 0x7f, 0x8c, 0xb2, 0x47, 0x87, 0x17, 0xbe, 0xa4, 0xd7, 0x3f, 0x99, 0xd5, 0x50, 0x86, 0xe8,
	0x94, 0x11, 0xc1, 0x44, 0x71, 0xb2, 0x7c, 0x45, 0x1e, 0x88, 0x1d, 0x4c, 0xdc, 0x5b, 0xf6, 0x17,
	0xf2, 0x01, 0xbd, 0x3a, 0x4d, 0x76, 0x5d, 0x27, 0x4d, 0xda, 0x66, 0x94, 0xf7, 0x96, 0xdd, 0x01,
	0xa8, 0x68, 0x24, 0x1a, 0x2d, 0x24, 0xf2, 0xac, 0x4e, 0x9a, 0x45, 0x5b, 0x5c, 0xc9, 0x0e, 0x27,
	0x3d, 0x06, 0x7d, 0xd3, 0xf9, 0xac, 0xaf, 0x64, 0x87, 0xd3, 0xaa, 0x43, 0x37, 0x5a, 0x5e, 0xcc,
	0xab, 0x4e, 0x33, 0xfb, 0x0f, 0x45, 0x38, 0x7a, 0xf4, 0x62, 0x8c, 0x1d, 0x07, 0x12, 0x39, 0x81,
	0xd7, 0xd8, 0xb1, 0x1a, 0x36, 0xda, 0x0c, 0x2a, 0xba, 0x80, 0xce, 0xf7, 0x7c, 0x43, 0xfa, 0x37,
	0x9a, 0x9a, 0x2a, 0x89, 0xc6, 0xfa, 0x78, 0xe1, 0xe5, 0xfc, 0xfb, 0x96, 0xd9, 0x03, 0x54, 0xe8,
	0x51, 0x76, 0x22, 0x9a, 0xb3, 0x33, 0x9f, 0x03, 0xaf, 0xa8, 0x6a, 0x49, 0xb0, 0x9d, 0x19, 0x7b,
	0x84, 0xad, 0x3c, 0x9b, 0x28, 0xad, 0x11, 0x51, 0xa2, 0xeb, 0x2d, 0xdf, 0x52, 0xe5, 0xea, 0x4a,
	0x5b, 0x82, 0x6f, 0x6b, 0xba, 0xd0, 0xf3, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0x51, 0x40,
	0x62, 0xc0, 0x01, 0x00, 0x00,
}
