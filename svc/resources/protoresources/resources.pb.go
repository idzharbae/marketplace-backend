// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resources.proto

package protoresources

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

type UploadPhotoReq struct {
	File                 []byte   `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	FileName             string   `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadPhotoReq) Reset()         { *m = UploadPhotoReq{} }
func (m *UploadPhotoReq) String() string { return proto.CompactTextString(m) }
func (*UploadPhotoReq) ProtoMessage()    {}
func (*UploadPhotoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{0}
}

func (m *UploadPhotoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadPhotoReq.Unmarshal(m, b)
}
func (m *UploadPhotoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadPhotoReq.Marshal(b, m, deterministic)
}
func (m *UploadPhotoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadPhotoReq.Merge(m, src)
}
func (m *UploadPhotoReq) XXX_Size() int {
	return xxx_messageInfo_UploadPhotoReq.Size(m)
}
func (m *UploadPhotoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadPhotoReq.DiscardUnknown(m)
}

var xxx_messageInfo_UploadPhotoReq proto.InternalMessageInfo

func (m *UploadPhotoReq) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *UploadPhotoReq) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

type UploadPhotoResp struct {
	FileUrl              string   `protobuf:"bytes,1,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadPhotoResp) Reset()         { *m = UploadPhotoResp{} }
func (m *UploadPhotoResp) String() string { return proto.CompactTextString(m) }
func (*UploadPhotoResp) ProtoMessage()    {}
func (*UploadPhotoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{1}
}

func (m *UploadPhotoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadPhotoResp.Unmarshal(m, b)
}
func (m *UploadPhotoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadPhotoResp.Marshal(b, m, deterministic)
}
func (m *UploadPhotoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadPhotoResp.Merge(m, src)
}
func (m *UploadPhotoResp) XXX_Size() int {
	return xxx_messageInfo_UploadPhotoResp.Size(m)
}
func (m *UploadPhotoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadPhotoResp.DiscardUnknown(m)
}

var xxx_messageInfo_UploadPhotoResp proto.InternalMessageInfo

func (m *UploadPhotoResp) GetFileUrl() string {
	if m != nil {
		return m.FileUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*UploadPhotoReq)(nil), "protoresources.UploadPhotoReq")
	proto.RegisterType((*UploadPhotoResp)(nil), "protoresources.UploadPhotoResp")
}

func init() {
	proto.RegisterFile("resources.proto", fileDescriptor_cf1b13971fe4c19d)
}

var fileDescriptor_cf1b13971fe4c19d = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4a, 0x2d, 0xce,
	0x2f, 0x2d, 0x4a, 0x4e, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x03, 0x53, 0x70,
	0x51, 0x29, 0xae, 0xa4, 0xc4, 0xe2, 0x54, 0x88, 0x9c, 0x92, 0x23, 0x17, 0x5f, 0x68, 0x41, 0x4e,
	0x7e, 0x62, 0x4a, 0x40, 0x46, 0x7e, 0x49, 0x7e, 0x50, 0x6a, 0xa1, 0x90, 0x10, 0x17, 0x4b, 0x5a,
	0x66, 0x4e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x98, 0x2d, 0x24, 0xcd, 0xc5, 0x09,
	0xa2, 0xe3, 0xf3, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x38, 0x40, 0x02,
	0x7e, 0x89, 0xb9, 0xa9, 0x4a, 0x3a, 0x5c, 0xfc, 0x28, 0x46, 0x14, 0x17, 0x08, 0x49, 0x72, 0x81,
	0xa5, 0xe3, 0x4b, 0x8b, 0x72, 0xc0, 0xe6, 0x70, 0x06, 0xb1, 0x83, 0xf8, 0xa1, 0x45, 0x39, 0x46,
	0x19, 0x5c, 0x22, 0xbe, 0x89, 0x45, 0xd9, 0xa9, 0x25, 0x05, 0x39, 0x89, 0xc9, 0xa9, 0x41, 0x30,
	0x47, 0x09, 0x05, 0x70, 0x71, 0x23, 0x99, 0x22, 0x24, 0xa7, 0x87, 0xea, 0x68, 0x3d, 0x54, 0x57,
	0x4a, 0xc9, 0xe3, 0x95, 0x2f, 0x2e, 0x50, 0x62, 0x48, 0x62, 0x03, 0xab, 0x30, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xef, 0x96, 0xcb, 0x21, 0x10, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MarketplaceResourcesClient is the client API for MarketplaceResources service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MarketplaceResourcesClient interface {
	UploadPhoto(ctx context.Context, in *UploadPhotoReq, opts ...grpc.CallOption) (*UploadPhotoResp, error)
}

type marketplaceResourcesClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketplaceResourcesClient(cc grpc.ClientConnInterface) MarketplaceResourcesClient {
	return &marketplaceResourcesClient{cc}
}

func (c *marketplaceResourcesClient) UploadPhoto(ctx context.Context, in *UploadPhotoReq, opts ...grpc.CallOption) (*UploadPhotoResp, error) {
	out := new(UploadPhotoResp)
	err := c.cc.Invoke(ctx, "/protoresources.MarketplaceResources/UploadPhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketplaceResourcesServer is the server API for MarketplaceResources service.
type MarketplaceResourcesServer interface {
	UploadPhoto(context.Context, *UploadPhotoReq) (*UploadPhotoResp, error)
}

// UnimplementedMarketplaceResourcesServer can be embedded to have forward compatible implementations.
type UnimplementedMarketplaceResourcesServer struct {
}

func (*UnimplementedMarketplaceResourcesServer) UploadPhoto(ctx context.Context, req *UploadPhotoReq) (*UploadPhotoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadPhoto not implemented")
}

func RegisterMarketplaceResourcesServer(s *grpc.Server, srv MarketplaceResourcesServer) {
	s.RegisterService(&_MarketplaceResources_serviceDesc, srv)
}

func _MarketplaceResources_UploadPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadPhotoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceResourcesServer).UploadPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoresources.MarketplaceResources/UploadPhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceResourcesServer).UploadPhoto(ctx, req.(*UploadPhotoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _MarketplaceResources_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protoresources.MarketplaceResources",
	HandlerType: (*MarketplaceResourcesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadPhoto",
			Handler:    _MarketplaceResources_UploadPhoto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resources.proto",
}