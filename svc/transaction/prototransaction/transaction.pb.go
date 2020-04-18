// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction.proto

package prototransaction

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AddToCartReq struct {
	ProductId            int64    `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	QuantityKg           int32    `protobuf:"varint,2,opt,name=quantity_kg,json=quantityKg,proto3" json:"quantity_kg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddToCartReq) Reset()         { *m = AddToCartReq{} }
func (m *AddToCartReq) String() string { return proto.CompactTextString(m) }
func (*AddToCartReq) ProtoMessage()    {}
func (*AddToCartReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0}
}

func (m *AddToCartReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddToCartReq.Unmarshal(m, b)
}
func (m *AddToCartReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddToCartReq.Marshal(b, m, deterministic)
}
func (m *AddToCartReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddToCartReq.Merge(m, src)
}
func (m *AddToCartReq) XXX_Size() int {
	return xxx_messageInfo_AddToCartReq.Size(m)
}
func (m *AddToCartReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddToCartReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddToCartReq proto.InternalMessageInfo

func (m *AddToCartReq) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *AddToCartReq) GetQuantityKg() int32 {
	if m != nil {
		return m.QuantityKg
	}
	return 0
}

type AddToCartResp struct {
	Product              *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddToCartResp) Reset()         { *m = AddToCartResp{} }
func (m *AddToCartResp) String() string { return proto.CompactTextString(m) }
func (*AddToCartResp) ProtoMessage()    {}
func (*AddToCartResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{1}
}

func (m *AddToCartResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddToCartResp.Unmarshal(m, b)
}
func (m *AddToCartResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddToCartResp.Marshal(b, m, deterministic)
}
func (m *AddToCartResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddToCartResp.Merge(m, src)
}
func (m *AddToCartResp) XXX_Size() int {
	return xxx_messageInfo_AddToCartResp.Size(m)
}
func (m *AddToCartResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddToCartResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddToCartResp proto.InternalMessageInfo

func (m *AddToCartResp) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

type CheckoutReq struct {
	Products             []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CheckoutReq) Reset()         { *m = CheckoutReq{} }
func (m *CheckoutReq) String() string { return proto.CompactTextString(m) }
func (*CheckoutReq) ProtoMessage()    {}
func (*CheckoutReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{2}
}

func (m *CheckoutReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckoutReq.Unmarshal(m, b)
}
func (m *CheckoutReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckoutReq.Marshal(b, m, deterministic)
}
func (m *CheckoutReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckoutReq.Merge(m, src)
}
func (m *CheckoutReq) XXX_Size() int {
	return xxx_messageInfo_CheckoutReq.Size(m)
}
func (m *CheckoutReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckoutReq.DiscardUnknown(m)
}

var xxx_messageInfo_CheckoutReq proto.InternalMessageInfo

func (m *CheckoutReq) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

type CheckoutResp struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckoutResp) Reset()         { *m = CheckoutResp{} }
func (m *CheckoutResp) String() string { return proto.CompactTextString(m) }
func (*CheckoutResp) ProtoMessage()    {}
func (*CheckoutResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{3}
}

func (m *CheckoutResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckoutResp.Unmarshal(m, b)
}
func (m *CheckoutResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckoutResp.Marshal(b, m, deterministic)
}
func (m *CheckoutResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckoutResp.Merge(m, src)
}
func (m *CheckoutResp) XXX_Size() int {
	return xxx_messageInfo_CheckoutResp.Size(m)
}
func (m *CheckoutResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckoutResp.DiscardUnknown(m)
}

var xxx_messageInfo_CheckoutResp proto.InternalMessageInfo

func (m *CheckoutResp) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type FulfillReq struct {
	OrderId              int64    `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	PaymentAmount        int64    `protobuf:"varint,2,opt,name=payment_amount,json=paymentAmount,proto3" json:"payment_amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FulfillReq) Reset()         { *m = FulfillReq{} }
func (m *FulfillReq) String() string { return proto.CompactTextString(m) }
func (*FulfillReq) ProtoMessage()    {}
func (*FulfillReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{4}
}

func (m *FulfillReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FulfillReq.Unmarshal(m, b)
}
func (m *FulfillReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FulfillReq.Marshal(b, m, deterministic)
}
func (m *FulfillReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FulfillReq.Merge(m, src)
}
func (m *FulfillReq) XXX_Size() int {
	return xxx_messageInfo_FulfillReq.Size(m)
}
func (m *FulfillReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FulfillReq.DiscardUnknown(m)
}

var xxx_messageInfo_FulfillReq proto.InternalMessageInfo

func (m *FulfillReq) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *FulfillReq) GetPaymentAmount() int64 {
	if m != nil {
		return m.PaymentAmount
	}
	return 0
}

type FulfillResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FulfillResp) Reset()         { *m = FulfillResp{} }
func (m *FulfillResp) String() string { return proto.CompactTextString(m) }
func (*FulfillResp) ProtoMessage()    {}
func (*FulfillResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{5}
}

func (m *FulfillResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FulfillResp.Unmarshal(m, b)
}
func (m *FulfillResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FulfillResp.Marshal(b, m, deterministic)
}
func (m *FulfillResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FulfillResp.Merge(m, src)
}
func (m *FulfillResp) XXX_Size() int {
	return xxx_messageInfo_FulfillResp.Size(m)
}
func (m *FulfillResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FulfillResp.DiscardUnknown(m)
}

var xxx_messageInfo_FulfillResp proto.InternalMessageInfo

func (m *FulfillResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*AddToCartReq)(nil), "prototransaction.AddToCartReq")
	proto.RegisterType((*AddToCartResp)(nil), "prototransaction.AddToCartResp")
	proto.RegisterType((*CheckoutReq)(nil), "prototransaction.CheckoutReq")
	proto.RegisterType((*CheckoutResp)(nil), "prototransaction.CheckoutResp")
	proto.RegisterType((*FulfillReq)(nil), "prototransaction.FulfillReq")
	proto.RegisterType((*FulfillResp)(nil), "prototransaction.FulfillResp")
}

func init() {
	proto.RegisterFile("transaction.proto", fileDescriptor_2cc4e03d2c28c490)
}

var fileDescriptor_2cc4e03d2c28c490 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0xad, 0x04, 0x5b, 0xa6, 0x60, 0x74, 0x0f, 0x5a, 0x1a, 0xf9, 0x93, 0x26, 0x46, 0x2e,
	0x72, 0x80, 0x78, 0xf4, 0x40, 0x20, 0x46, 0x42, 0x44, 0x53, 0xb9, 0x93, 0xb5, 0xad, 0xd8, 0x50,
	0xba, 0xcb, 0xee, 0xf6, 0xc0, 0xa7, 0xd7, 0xb0, 0xdd, 0xd2, 0xc6, 0x26, 0x3d, 0x35, 0xf3, 0xde,
	0xeb, 0x2f, 0x33, 0x6f, 0xe1, 0x5a, 0x30, 0x1c, 0x73, 0xec, 0x89, 0x90, 0xc4, 0x43, 0xca, 0x88,
	0x20, 0xe8, 0x4a, 0x7e, 0x0a, 0xba, 0xdd, 0xa2, 0x8c, 0xf8, 0x89, 0x27, 0xd2, 0x80, 0x6d, 0x12,
	0xe6, 0x07, 0x2c, 0x1d, 0x9c, 0x25, 0x34, 0x27, 0xbe, 0xbf, 0x22, 0x53, 0xcc, 0x84, 0x1b, 0xec,
	0x51, 0x07, 0x40, 0xa5, 0xd7, 0xa1, 0x6f, 0x69, 0x7d, 0x6d, 0x50, 0x73, 0x1b, 0x4a, 0x99, 0xfb,
	0xa8, 0x07, 0xe6, 0x3e, 0xc1, 0xb1, 0x08, 0xc5, 0x61, 0xbd, 0xdd, 0x58, 0xe7, 0x7d, 0x6d, 0x50,
	0x77, 0x21, 0x93, 0x16, 0x1b, 0x67, 0x06, 0xad, 0x02, 0x8f, 0x53, 0x34, 0x06, 0x5d, 0xfd, 0x2e,
	0x69, 0xe6, 0xa8, 0x3d, 0xfc, 0xbf, 0xe0, 0xf0, 0x23, 0x0d, 0xb8, 0x59, 0xd2, 0x99, 0x81, 0x39,
	0xfd, 0x09, 0xbc, 0x2d, 0x49, 0xe4, 0x52, 0x4f, 0x60, 0x28, 0x87, 0x5b, 0x5a, 0xbf, 0x56, 0x0d,
	0x39, 0x45, 0x9d, 0x67, 0x68, 0xe6, 0x14, 0x4e, 0xd1, 0x23, 0xd4, 0xe5, 0xe9, 0x6a, 0x91, 0xdb,
	0x32, 0xe3, 0xfd, 0x68, 0xbb, 0x69, 0xca, 0x59, 0x02, 0xbc, 0x24, 0xd1, 0x77, 0x18, 0x45, 0xc7,
	0x1d, 0xda, 0x60, 0x48, 0x39, 0xaf, 0x45, 0x97, 0xf3, 0xdc, 0x47, 0xf7, 0x70, 0x49, 0xf1, 0x61,
	0x17, 0xc4, 0x62, 0x8d, 0x77, 0x24, 0x89, 0x85, 0xec, 0xa5, 0xe6, 0xb6, 0x94, 0x3a, 0x91, 0xa2,
	0xf3, 0x00, 0xe6, 0x89, 0xc7, 0x29, 0xb2, 0x40, 0xff, 0x4c, 0x3c, 0x2f, 0xe0, 0x5c, 0xf2, 0x0c,
	0x37, 0x1b, 0x47, 0xbf, 0x1a, 0xdc, 0xbc, 0x61, 0xb6, 0x0d, 0x04, 0x8d, 0xb0, 0x17, 0xac, 0xf2,
	0x05, 0xd1, 0x12, 0x1a, 0xa7, 0x7a, 0x51, 0xb7, 0x7c, 0x40, 0xf1, 0x2d, 0xed, 0x5e, 0xa5, 0xcf,
	0xa9, 0x73, 0x86, 0x16, 0x60, 0x64, 0x15, 0xa1, 0x4e, 0x39, 0x5e, 0x78, 0x04, 0xbb, 0x5b, 0x65,
	0x4b, 0xd8, 0x2b, 0xe8, 0xea, 0x40, 0x74, 0x57, 0x0e, 0xe7, 0x5d, 0xda, 0x9d, 0x0a, 0xf7, 0x48,
	0xfa, 0xba, 0x90, 0xfe, 0xf8, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xed, 0x2e, 0x5f, 0xdf, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MarketplaceTransactionClient is the client API for MarketplaceTransaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MarketplaceTransactionClient interface {
	AddToCart(ctx context.Context, in *AddToCartReq, opts ...grpc.CallOption) (*AddToCartResp, error)
	Checkout(ctx context.Context, in *CheckoutReq, opts ...grpc.CallOption) (*CheckoutResp, error)
	Fulfill(ctx context.Context, in *FulfillReq, opts ...grpc.CallOption) (*FulfillResp, error)
}

type marketplaceTransactionClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketplaceTransactionClient(cc grpc.ClientConnInterface) MarketplaceTransactionClient {
	return &marketplaceTransactionClient{cc}
}

func (c *marketplaceTransactionClient) AddToCart(ctx context.Context, in *AddToCartReq, opts ...grpc.CallOption) (*AddToCartResp, error) {
	out := new(AddToCartResp)
	err := c.cc.Invoke(ctx, "/prototransaction.MarketplaceTransaction/AddToCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceTransactionClient) Checkout(ctx context.Context, in *CheckoutReq, opts ...grpc.CallOption) (*CheckoutResp, error) {
	out := new(CheckoutResp)
	err := c.cc.Invoke(ctx, "/prototransaction.MarketplaceTransaction/Checkout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceTransactionClient) Fulfill(ctx context.Context, in *FulfillReq, opts ...grpc.CallOption) (*FulfillResp, error) {
	out := new(FulfillResp)
	err := c.cc.Invoke(ctx, "/prototransaction.MarketplaceTransaction/Fulfill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketplaceTransactionServer is the server API for MarketplaceTransaction service.
type MarketplaceTransactionServer interface {
	AddToCart(context.Context, *AddToCartReq) (*AddToCartResp, error)
	Checkout(context.Context, *CheckoutReq) (*CheckoutResp, error)
	Fulfill(context.Context, *FulfillReq) (*FulfillResp, error)
}

// UnimplementedMarketplaceTransactionServer can be embedded to have forward compatible implementations.
type UnimplementedMarketplaceTransactionServer struct {
}

func (*UnimplementedMarketplaceTransactionServer) AddToCart(ctx context.Context, req *AddToCartReq) (*AddToCartResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToCart not implemented")
}
func (*UnimplementedMarketplaceTransactionServer) Checkout(ctx context.Context, req *CheckoutReq) (*CheckoutResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checkout not implemented")
}
func (*UnimplementedMarketplaceTransactionServer) Fulfill(ctx context.Context, req *FulfillReq) (*FulfillResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fulfill not implemented")
}

func RegisterMarketplaceTransactionServer(s *grpc.Server, srv MarketplaceTransactionServer) {
	s.RegisterService(&_MarketplaceTransaction_serviceDesc, srv)
}

func _MarketplaceTransaction_AddToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToCartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceTransactionServer).AddToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prototransaction.MarketplaceTransaction/AddToCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceTransactionServer).AddToCart(ctx, req.(*AddToCartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceTransaction_Checkout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceTransactionServer).Checkout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prototransaction.MarketplaceTransaction/Checkout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceTransactionServer).Checkout(ctx, req.(*CheckoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceTransaction_Fulfill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FulfillReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceTransactionServer).Fulfill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prototransaction.MarketplaceTransaction/Fulfill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceTransactionServer).Fulfill(ctx, req.(*FulfillReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _MarketplaceTransaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "prototransaction.MarketplaceTransaction",
	HandlerType: (*MarketplaceTransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddToCart",
			Handler:    _MarketplaceTransaction_AddToCart_Handler,
		},
		{
			MethodName: "Checkout",
			Handler:    _MarketplaceTransaction_Checkout_Handler,
		},
		{
			MethodName: "Fulfill",
			Handler:    _MarketplaceTransaction_Fulfill_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction.proto",
}
