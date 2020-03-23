// Code generated by protoc-gen-go. DO NOT EDIT.
// source: catalog.proto

package catalogproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type ListProductsReq struct {
	Pagination           *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListProductsReq) Reset()         { *m = ListProductsReq{} }
func (m *ListProductsReq) String() string { return proto.CompactTextString(m) }
func (*ListProductsReq) ProtoMessage()    {}
func (*ListProductsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_catalog_6ce5c586dc5fb196, []int{0}
}
func (m *ListProductsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProductsReq.Unmarshal(m, b)
}
func (m *ListProductsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProductsReq.Marshal(b, m, deterministic)
}
func (dst *ListProductsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProductsReq.Merge(dst, src)
}
func (m *ListProductsReq) XXX_Size() int {
	return xxx_messageInfo_ListProductsReq.Size(m)
}
func (m *ListProductsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProductsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListProductsReq proto.InternalMessageInfo

func (m *ListProductsReq) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type ListProductsResp struct {
	Products             []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListProductsResp) Reset()         { *m = ListProductsResp{} }
func (m *ListProductsResp) String() string { return proto.CompactTextString(m) }
func (*ListProductsResp) ProtoMessage()    {}
func (*ListProductsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_catalog_6ce5c586dc5fb196, []int{1}
}
func (m *ListProductsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProductsResp.Unmarshal(m, b)
}
func (m *ListProductsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProductsResp.Marshal(b, m, deterministic)
}
func (dst *ListProductsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProductsResp.Merge(dst, src)
}
func (m *ListProductsResp) XXX_Size() int {
	return xxx_messageInfo_ListProductsResp.Size(m)
}
func (m *ListProductsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProductsResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListProductsResp proto.InternalMessageInfo

func (m *ListProductsResp) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

type GetProductByIDReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductByIDReq) Reset()         { *m = GetProductByIDReq{} }
func (m *GetProductByIDReq) String() string { return proto.CompactTextString(m) }
func (*GetProductByIDReq) ProtoMessage()    {}
func (*GetProductByIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_catalog_6ce5c586dc5fb196, []int{2}
}
func (m *GetProductByIDReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductByIDReq.Unmarshal(m, b)
}
func (m *GetProductByIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductByIDReq.Marshal(b, m, deterministic)
}
func (dst *GetProductByIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductByIDReq.Merge(dst, src)
}
func (m *GetProductByIDReq) XXX_Size() int {
	return xxx_messageInfo_GetProductByIDReq.Size(m)
}
func (m *GetProductByIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductByIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductByIDReq proto.InternalMessageInfo

func (m *GetProductByIDReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetProductBySlugReq struct {
	Slug                 string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductBySlugReq) Reset()         { *m = GetProductBySlugReq{} }
func (m *GetProductBySlugReq) String() string { return proto.CompactTextString(m) }
func (*GetProductBySlugReq) ProtoMessage()    {}
func (*GetProductBySlugReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_catalog_6ce5c586dc5fb196, []int{3}
}
func (m *GetProductBySlugReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductBySlugReq.Unmarshal(m, b)
}
func (m *GetProductBySlugReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductBySlugReq.Marshal(b, m, deterministic)
}
func (dst *GetProductBySlugReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductBySlugReq.Merge(dst, src)
}
func (m *GetProductBySlugReq) XXX_Size() int {
	return xxx_messageInfo_GetProductBySlugReq.Size(m)
}
func (m *GetProductBySlugReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductBySlugReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductBySlugReq proto.InternalMessageInfo

func (m *GetProductBySlugReq) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

type PKReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PKReq) Reset()         { *m = PKReq{} }
func (m *PKReq) String() string { return proto.CompactTextString(m) }
func (*PKReq) ProtoMessage()    {}
func (*PKReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_catalog_6ce5c586dc5fb196, []int{4}
}
func (m *PKReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PKReq.Unmarshal(m, b)
}
func (m *PKReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PKReq.Marshal(b, m, deterministic)
}
func (dst *PKReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PKReq.Merge(dst, src)
}
func (m *PKReq) XXX_Size() int {
	return xxx_messageInfo_PKReq.Size(m)
}
func (m *PKReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PKReq.DiscardUnknown(m)
}

var xxx_messageInfo_PKReq proto.InternalMessageInfo

func (m *PKReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListShopsReq struct {
	Pagination           *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListShopsReq) Reset()         { *m = ListShopsReq{} }
func (m *ListShopsReq) String() string { return proto.CompactTextString(m) }
func (*ListShopsReq) ProtoMessage()    {}
func (*ListShopsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_catalog_6ce5c586dc5fb196, []int{5}
}
func (m *ListShopsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListShopsReq.Unmarshal(m, b)
}
func (m *ListShopsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListShopsReq.Marshal(b, m, deterministic)
}
func (dst *ListShopsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListShopsReq.Merge(dst, src)
}
func (m *ListShopsReq) XXX_Size() int {
	return xxx_messageInfo_ListShopsReq.Size(m)
}
func (m *ListShopsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListShopsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListShopsReq proto.InternalMessageInfo

func (m *ListShopsReq) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type ListShopsResp struct {
	Shops                []*Shop  `protobuf:"bytes,1,rep,name=shops,proto3" json:"shops,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListShopsResp) Reset()         { *m = ListShopsResp{} }
func (m *ListShopsResp) String() string { return proto.CompactTextString(m) }
func (*ListShopsResp) ProtoMessage()    {}
func (*ListShopsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_catalog_6ce5c586dc5fb196, []int{6}
}
func (m *ListShopsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListShopsResp.Unmarshal(m, b)
}
func (m *ListShopsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListShopsResp.Marshal(b, m, deterministic)
}
func (dst *ListShopsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListShopsResp.Merge(dst, src)
}
func (m *ListShopsResp) XXX_Size() int {
	return xxx_messageInfo_ListShopsResp.Size(m)
}
func (m *ListShopsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListShopsResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListShopsResp proto.InternalMessageInfo

func (m *ListShopsResp) GetShops() []*Shop {
	if m != nil {
		return m.Shops
	}
	return nil
}

type GetShopByPKReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShopByPKReq) Reset()         { *m = GetShopByPKReq{} }
func (m *GetShopByPKReq) String() string { return proto.CompactTextString(m) }
func (*GetShopByPKReq) ProtoMessage()    {}
func (*GetShopByPKReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_catalog_6ce5c586dc5fb196, []int{7}
}
func (m *GetShopByPKReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShopByPKReq.Unmarshal(m, b)
}
func (m *GetShopByPKReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShopByPKReq.Marshal(b, m, deterministic)
}
func (dst *GetShopByPKReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShopByPKReq.Merge(dst, src)
}
func (m *GetShopByPKReq) XXX_Size() int {
	return xxx_messageInfo_GetShopByPKReq.Size(m)
}
func (m *GetShopByPKReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShopByPKReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetShopByPKReq proto.InternalMessageInfo

func (m *GetShopByPKReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetShopBySlugReq struct {
	Slug                 string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShopBySlugReq) Reset()         { *m = GetShopBySlugReq{} }
func (m *GetShopBySlugReq) String() string { return proto.CompactTextString(m) }
func (*GetShopBySlugReq) ProtoMessage()    {}
func (*GetShopBySlugReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_catalog_6ce5c586dc5fb196, []int{8}
}
func (m *GetShopBySlugReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShopBySlugReq.Unmarshal(m, b)
}
func (m *GetShopBySlugReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShopBySlugReq.Marshal(b, m, deterministic)
}
func (dst *GetShopBySlugReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShopBySlugReq.Merge(dst, src)
}
func (m *GetShopBySlugReq) XXX_Size() int {
	return xxx_messageInfo_GetShopBySlugReq.Size(m)
}
func (m *GetShopBySlugReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShopBySlugReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetShopBySlugReq proto.InternalMessageInfo

func (m *GetShopBySlugReq) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func init() {
	proto.RegisterType((*ListProductsReq)(nil), "catalogproto.ListProductsReq")
	proto.RegisterType((*ListProductsResp)(nil), "catalogproto.ListProductsResp")
	proto.RegisterType((*GetProductByIDReq)(nil), "catalogproto.GetProductByIDReq")
	proto.RegisterType((*GetProductBySlugReq)(nil), "catalogproto.GetProductBySlugReq")
	proto.RegisterType((*PKReq)(nil), "catalogproto.PKReq")
	proto.RegisterType((*ListShopsReq)(nil), "catalogproto.ListShopsReq")
	proto.RegisterType((*ListShopsResp)(nil), "catalogproto.ListShopsResp")
	proto.RegisterType((*GetShopByPKReq)(nil), "catalogproto.GetShopByPKReq")
	proto.RegisterType((*GetShopBySlugReq)(nil), "catalogproto.GetShopBySlugReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MarketplaceCatalogClient is the client API for MarketplaceCatalog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MarketplaceCatalogClient interface {
	// Product
	ListProducts(ctx context.Context, in *ListProductsReq, opts ...grpc.CallOption) (*ListProductsResp, error)
	GetProductByID(ctx context.Context, in *GetProductByIDReq, opts ...grpc.CallOption) (*Product, error)
	GetProductBySlug(ctx context.Context, in *GetProductBySlugReq, opts ...grpc.CallOption) (*Product, error)
	CreateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error)
	UpdateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error)
	DeleteProduct(ctx context.Context, in *PKReq, opts ...grpc.CallOption) (*Empty, error)
	// Shop
	ListShops(ctx context.Context, in *ListShopsReq, opts ...grpc.CallOption) (*ListShopsResp, error)
	GetShopByID(ctx context.Context, in *GetShopByPKReq, opts ...grpc.CallOption) (*Shop, error)
	GetShopBySlug(ctx context.Context, in *GetShopBySlugReq, opts ...grpc.CallOption) (*Shop, error)
	CreateShop(ctx context.Context, in *Shop, opts ...grpc.CallOption) (*Shop, error)
	UpdateShop(ctx context.Context, in *Shop, opts ...grpc.CallOption) (*Shop, error)
	DeleteShop(ctx context.Context, in *PKReq, opts ...grpc.CallOption) (*Empty, error)
}

type marketplaceCatalogClient struct {
	cc *grpc.ClientConn
}

func NewMarketplaceCatalogClient(cc *grpc.ClientConn) MarketplaceCatalogClient {
	return &marketplaceCatalogClient{cc}
}

func (c *marketplaceCatalogClient) ListProducts(ctx context.Context, in *ListProductsReq, opts ...grpc.CallOption) (*ListProductsResp, error) {
	out := new(ListProductsResp)
	err := c.cc.Invoke(ctx, "/catalogproto.MarketplaceCatalog/ListProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceCatalogClient) GetProductByID(ctx context.Context, in *GetProductByIDReq, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/catalogproto.MarketplaceCatalog/GetProductByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceCatalogClient) GetProductBySlug(ctx context.Context, in *GetProductBySlugReq, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/catalogproto.MarketplaceCatalog/GetProductBySlug", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceCatalogClient) CreateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/catalogproto.MarketplaceCatalog/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceCatalogClient) UpdateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/catalogproto.MarketplaceCatalog/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceCatalogClient) DeleteProduct(ctx context.Context, in *PKReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/catalogproto.MarketplaceCatalog/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceCatalogClient) ListShops(ctx context.Context, in *ListShopsReq, opts ...grpc.CallOption) (*ListShopsResp, error) {
	out := new(ListShopsResp)
	err := c.cc.Invoke(ctx, "/catalogproto.MarketplaceCatalog/ListShops", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceCatalogClient) GetShopByID(ctx context.Context, in *GetShopByPKReq, opts ...grpc.CallOption) (*Shop, error) {
	out := new(Shop)
	err := c.cc.Invoke(ctx, "/catalogproto.MarketplaceCatalog/GetShopByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceCatalogClient) GetShopBySlug(ctx context.Context, in *GetShopBySlugReq, opts ...grpc.CallOption) (*Shop, error) {
	out := new(Shop)
	err := c.cc.Invoke(ctx, "/catalogproto.MarketplaceCatalog/GetShopBySlug", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceCatalogClient) CreateShop(ctx context.Context, in *Shop, opts ...grpc.CallOption) (*Shop, error) {
	out := new(Shop)
	err := c.cc.Invoke(ctx, "/catalogproto.MarketplaceCatalog/CreateShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceCatalogClient) UpdateShop(ctx context.Context, in *Shop, opts ...grpc.CallOption) (*Shop, error) {
	out := new(Shop)
	err := c.cc.Invoke(ctx, "/catalogproto.MarketplaceCatalog/UpdateShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceCatalogClient) DeleteShop(ctx context.Context, in *PKReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/catalogproto.MarketplaceCatalog/DeleteShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketplaceCatalogServer is the server API for MarketplaceCatalog service.
type MarketplaceCatalogServer interface {
	// Product
	ListProducts(context.Context, *ListProductsReq) (*ListProductsResp, error)
	GetProductByID(context.Context, *GetProductByIDReq) (*Product, error)
	GetProductBySlug(context.Context, *GetProductBySlugReq) (*Product, error)
	CreateProduct(context.Context, *Product) (*Product, error)
	UpdateProduct(context.Context, *Product) (*Product, error)
	DeleteProduct(context.Context, *PKReq) (*Empty, error)
	// Shop
	ListShops(context.Context, *ListShopsReq) (*ListShopsResp, error)
	GetShopByID(context.Context, *GetShopByPKReq) (*Shop, error)
	GetShopBySlug(context.Context, *GetShopBySlugReq) (*Shop, error)
	CreateShop(context.Context, *Shop) (*Shop, error)
	UpdateShop(context.Context, *Shop) (*Shop, error)
	DeleteShop(context.Context, *PKReq) (*Empty, error)
}

func RegisterMarketplaceCatalogServer(s *grpc.Server, srv MarketplaceCatalogServer) {
	s.RegisterService(&_MarketplaceCatalog_serviceDesc, srv)
}

func _MarketplaceCatalog_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceCatalogServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogproto.MarketplaceCatalog/ListProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceCatalogServer).ListProducts(ctx, req.(*ListProductsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceCatalog_GetProductByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceCatalogServer).GetProductByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogproto.MarketplaceCatalog/GetProductByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceCatalogServer).GetProductByID(ctx, req.(*GetProductByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceCatalog_GetProductBySlug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductBySlugReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceCatalogServer).GetProductBySlug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogproto.MarketplaceCatalog/GetProductBySlug",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceCatalogServer).GetProductBySlug(ctx, req.(*GetProductBySlugReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceCatalog_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceCatalogServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogproto.MarketplaceCatalog/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceCatalogServer).CreateProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceCatalog_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceCatalogServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogproto.MarketplaceCatalog/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceCatalogServer).UpdateProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceCatalog_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PKReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceCatalogServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogproto.MarketplaceCatalog/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceCatalogServer).DeleteProduct(ctx, req.(*PKReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceCatalog_ListShops_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListShopsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceCatalogServer).ListShops(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogproto.MarketplaceCatalog/ListShops",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceCatalogServer).ListShops(ctx, req.(*ListShopsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceCatalog_GetShopByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShopByPKReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceCatalogServer).GetShopByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogproto.MarketplaceCatalog/GetShopByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceCatalogServer).GetShopByID(ctx, req.(*GetShopByPKReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceCatalog_GetShopBySlug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShopBySlugReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceCatalogServer).GetShopBySlug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogproto.MarketplaceCatalog/GetShopBySlug",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceCatalogServer).GetShopBySlug(ctx, req.(*GetShopBySlugReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceCatalog_CreateShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Shop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceCatalogServer).CreateShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogproto.MarketplaceCatalog/CreateShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceCatalogServer).CreateShop(ctx, req.(*Shop))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceCatalog_UpdateShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Shop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceCatalogServer).UpdateShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogproto.MarketplaceCatalog/UpdateShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceCatalogServer).UpdateShop(ctx, req.(*Shop))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceCatalog_DeleteShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PKReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceCatalogServer).DeleteShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogproto.MarketplaceCatalog/DeleteShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceCatalogServer).DeleteShop(ctx, req.(*PKReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _MarketplaceCatalog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "catalogproto.MarketplaceCatalog",
	HandlerType: (*MarketplaceCatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProducts",
			Handler:    _MarketplaceCatalog_ListProducts_Handler,
		},
		{
			MethodName: "GetProductByID",
			Handler:    _MarketplaceCatalog_GetProductByID_Handler,
		},
		{
			MethodName: "GetProductBySlug",
			Handler:    _MarketplaceCatalog_GetProductBySlug_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _MarketplaceCatalog_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _MarketplaceCatalog_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _MarketplaceCatalog_DeleteProduct_Handler,
		},
		{
			MethodName: "ListShops",
			Handler:    _MarketplaceCatalog_ListShops_Handler,
		},
		{
			MethodName: "GetShopByID",
			Handler:    _MarketplaceCatalog_GetShopByID_Handler,
		},
		{
			MethodName: "GetShopBySlug",
			Handler:    _MarketplaceCatalog_GetShopBySlug_Handler,
		},
		{
			MethodName: "CreateShop",
			Handler:    _MarketplaceCatalog_CreateShop_Handler,
		},
		{
			MethodName: "UpdateShop",
			Handler:    _MarketplaceCatalog_UpdateShop_Handler,
		},
		{
			MethodName: "DeleteShop",
			Handler:    _MarketplaceCatalog_DeleteShop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalog.proto",
}

func init() { proto.RegisterFile("catalog.proto", fileDescriptor_catalog_6ce5c586dc5fb196) }

var fileDescriptor_catalog_6ce5c586dc5fb196 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x6f, 0x8f, 0xd2, 0x40,
	0x10, 0xc6, 0xcb, 0x29, 0x17, 0x6f, 0xb8, 0x9e, 0xe7, 0x5c, 0x8c, 0x64, 0xd5, 0x13, 0xd7, 0xc4,
	0xe0, 0x1b, 0x12, 0x31, 0x31, 0x67, 0x7c, 0x61, 0x3c, 0x40, 0x44, 0xfc, 0x43, 0x4a, 0xfc, 0x00,
	0x0b, 0xdd, 0x94, 0xc6, 0x4a, 0xd7, 0xee, 0xf2, 0x82, 0xef, 0xe6, 0x87, 0x33, 0xdd, 0x6d, 0x6b,
	0x5b, 0x5a, 0x12, 0xc2, 0xbb, 0x9d, 0xce, 0x33, 0xbf, 0xce, 0xcc, 0x33, 0x60, 0x2f, 0x99, 0x62,
	0x41, 0xe8, 0xf5, 0x44, 0x14, 0xaa, 0x10, 0xcf, 0x93, 0x50, 0x47, 0xc4, 0x16, 0x51, 0xe8, 0x6e,
	0x96, 0xca, 0x24, 0x09, 0xc8, 0x55, 0x28, 0xd2, 0xf7, 0x82, 0x49, 0x6e, 0xde, 0x74, 0x0a, 0xf7,
	0xbf, 0xfa, 0x52, 0xcd, 0x8c, 0x58, 0x3a, 0xfc, 0x0f, 0xde, 0x00, 0x08, 0xe6, 0xf9, 0x6b, 0xa6,
	0xfc, 0x70, 0xdd, 0x6e, 0x74, 0x1a, 0xdd, 0x56, 0xbf, 0xdd, 0xcb, 0xc3, 0x7b, 0xb3, 0x2c, 0xef,
	0xe4, 0xb4, 0x74, 0x04, 0x97, 0x45, 0x98, 0x14, 0xf8, 0x1a, 0xee, 0x25, 0x9d, 0xc8, 0x76, 0xa3,
	0x73, 0xa7, 0xdb, 0xea, 0x3f, 0x2c, 0xb1, 0x4c, 0xd6, 0xc9, 0x64, 0xf4, 0x05, 0x3c, 0x18, 0xf3,
	0x94, 0x72, 0xbb, 0x9d, 0x0c, 0xe3, 0xae, 0x2e, 0xe0, 0xc4, 0x77, 0x75, 0x37, 0x4d, 0xe7, 0xc4,
	0x77, 0xe9, 0x2b, 0xb8, 0xca, 0x8b, 0xe6, 0xc1, 0xc6, 0x8b, 0x65, 0x08, 0x77, 0x65, 0xb0, 0xf1,
	0xb4, 0xf0, 0xcc, 0xd1, 0x6f, 0xfa, 0x08, 0x9a, 0xb3, 0x69, 0x15, 0xe3, 0x33, 0x9c, 0xc7, 0xfd,
	0xce, 0x57, 0xa1, 0x38, 0x72, 0xf2, 0x77, 0x60, 0xe7, 0x48, 0x52, 0x60, 0x17, 0x9a, 0xf1, 0xc6,
	0xd3, 0x99, 0xb1, 0x48, 0x89, 0x75, 0x8e, 0x11, 0xd0, 0x0e, 0x5c, 0x8c, 0xb9, 0xae, 0xbc, 0xdd,
	0x56, 0xb7, 0xf9, 0x12, 0x2e, 0x33, 0xc5, 0x9e, 0x39, 0xfb, 0x7f, 0x4f, 0x01, 0xbf, 0xb1, 0xe8,
	0x17, 0x57, 0x22, 0x60, 0x4b, 0x3e, 0x30, 0x7f, 0xc4, 0x1f, 0x66, 0xca, 0xd4, 0x15, 0x7c, 0x5a,
	0xec, 0xa5, 0x64, 0x3f, 0xb9, 0xde, 0x97, 0x96, 0x82, 0x5a, 0xf8, 0x45, 0x77, 0x9c, 0xf3, 0x07,
	0x9f, 0x15, 0x6b, 0x76, 0xdc, 0x23, 0xd5, 0x9e, 0x53, 0x0b, 0xbf, 0xeb, 0xd9, 0x0a, 0x36, 0xe2,
	0xf3, 0x7a, 0x5a, 0x32, 0x7e, 0x3d, 0xef, 0x03, 0xd8, 0x83, 0x88, 0x33, 0xc5, 0x93, 0x4f, 0x58,
	0xad, 0xdc, 0x0b, 0xf8, 0x29, 0xdc, 0x23, 0x00, 0xef, 0xc1, 0x1e, 0xf2, 0x80, 0xff, 0x07, 0x5c,
	0x95, 0x94, 0xb1, 0xc7, 0xa4, 0xf4, 0x71, 0xf4, 0x5b, 0xa8, 0x2d, 0xb5, 0xf0, 0x13, 0x9c, 0x65,
	0x77, 0x84, 0x64, 0xd7, 0x89, 0xf4, 0x54, 0xc9, 0xe3, 0xda, 0x9c, 0xb6, 0xe8, 0x23, 0xb4, 0xb2,
	0x93, 0x99, 0x0c, 0xf1, 0xc9, 0xce, 0x46, 0x73, 0xf7, 0x46, 0x2a, 0x8e, 0x93, 0x5a, 0x38, 0x02,
	0xbb, 0x70, 0x75, 0x78, 0x5d, 0x03, 0x49, 0x3d, 0xa9, 0xc6, 0xbc, 0x05, 0x30, 0x86, 0xc4, 0x31,
	0x56, 0x68, 0xea, 0xeb, 0x8c, 0x0f, 0x07, 0xd6, 0xdd, 0x00, 0x98, 0xf5, 0xeb, 0xba, 0x03, 0x76,
	0xbf, 0x38, 0xd5, 0xe1, 0x9b, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x5d, 0x52, 0xef, 0x57,
	0x05, 0x00, 0x00,
}
