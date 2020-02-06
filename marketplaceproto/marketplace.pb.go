// Code generated by protoc-gen-go. DO NOT EDIT.
// source: marketplace.proto

package marketplaceproto

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MarketplaceClient is the client API for Marketplace service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MarketplaceClient interface {
	ListProducts(ctx context.Context, in *ListProductsReq, opts ...grpc.CallOption) (*ListProductsResp, error)
}

type marketplaceClient struct {
	cc *grpc.ClientConn
}

func NewMarketplaceClient(cc *grpc.ClientConn) MarketplaceClient {
	return &marketplaceClient{cc}
}

func (c *marketplaceClient) ListProducts(ctx context.Context, in *ListProductsReq, opts ...grpc.CallOption) (*ListProductsResp, error) {
	out := new(ListProductsResp)
	err := c.cc.Invoke(ctx, "/marketplaceproto.Marketplace/ListProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketplaceServer is the server API for Marketplace service.
type MarketplaceServer interface {
	ListProducts(context.Context, *ListProductsReq) (*ListProductsResp, error)
}

func RegisterMarketplaceServer(s *grpc.Server, srv MarketplaceServer) {
	s.RegisterService(&_Marketplace_serviceDesc, srv)
}

func _Marketplace_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketplaceproto.Marketplace/ListProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServer).ListProducts(ctx, req.(*ListProductsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Marketplace_serviceDesc = grpc.ServiceDesc{
	ServiceName: "marketplaceproto.Marketplace",
	HandlerType: (*MarketplaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProducts",
			Handler:    _Marketplace_ListProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "marketplace.proto",
}

func init() { proto.RegisterFile("marketplace.proto", fileDescriptor_marketplace_2bf955cfcc370ff0) }

var fileDescriptor_marketplace_2bf955cfcc370ff0 = []byte{
	// 103 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x4d, 0x2c, 0xca,
	0x4e, 0x2d, 0x29, 0xc8, 0x49, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x40,
	0x12, 0x02, 0x8b, 0x48, 0xf1, 0x16, 0x14, 0xe5, 0xa7, 0x94, 0x26, 0x97, 0x40, 0x14, 0x18, 0xa5,
	0x71, 0x71, 0xfb, 0x22, 0x94, 0x08, 0x85, 0x73, 0xf1, 0xf8, 0x64, 0x16, 0x97, 0x04, 0x40, 0xd4,
	0x14, 0x0b, 0x29, 0xea, 0xa1, 0x1b, 0xa0, 0x87, 0x2c, 0x1f, 0x94, 0x5a, 0x28, 0xa5, 0x44, 0x48,
	0x49, 0x71, 0x81, 0x12, 0x43, 0x12, 0x1b, 0x58, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x87,
	0xb0, 0xcb, 0x64, 0xa4, 0x00, 0x00, 0x00,
}