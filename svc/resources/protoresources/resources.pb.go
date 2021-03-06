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
	FileExt              string   `protobuf:"bytes,2,opt,name=file_ext,json=fileExt,proto3" json:"file_ext,omitempty"`
	OwnerId              int64    `protobuf:"varint,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
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

func (m *UploadPhotoReq) GetFileExt() string {
	if m != nil {
		return m.FileExt
	}
	return ""
}

func (m *UploadPhotoReq) GetOwnerId() int64 {
	if m != nil {
		return m.OwnerId
	}
	return 0
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

type DeletePhotoReq struct {
	FileUrl              string   `protobuf:"bytes,1,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePhotoReq) Reset()         { *m = DeletePhotoReq{} }
func (m *DeletePhotoReq) String() string { return proto.CompactTextString(m) }
func (*DeletePhotoReq) ProtoMessage()    {}
func (*DeletePhotoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{2}
}

func (m *DeletePhotoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePhotoReq.Unmarshal(m, b)
}
func (m *DeletePhotoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePhotoReq.Marshal(b, m, deterministic)
}
func (m *DeletePhotoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePhotoReq.Merge(m, src)
}
func (m *DeletePhotoReq) XXX_Size() int {
	return xxx_messageInfo_DeletePhotoReq.Size(m)
}
func (m *DeletePhotoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePhotoReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePhotoReq proto.InternalMessageInfo

func (m *DeletePhotoReq) GetFileUrl() string {
	if m != nil {
		return m.FileUrl
	}
	return ""
}

func (m *DeletePhotoReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type DeletePhotoResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePhotoResp) Reset()         { *m = DeletePhotoResp{} }
func (m *DeletePhotoResp) String() string { return proto.CompactTextString(m) }
func (*DeletePhotoResp) ProtoMessage()    {}
func (*DeletePhotoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{3}
}

func (m *DeletePhotoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePhotoResp.Unmarshal(m, b)
}
func (m *DeletePhotoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePhotoResp.Marshal(b, m, deterministic)
}
func (m *DeletePhotoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePhotoResp.Merge(m, src)
}
func (m *DeletePhotoResp) XXX_Size() int {
	return xxx_messageInfo_DeletePhotoResp.Size(m)
}
func (m *DeletePhotoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePhotoResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePhotoResp proto.InternalMessageInfo

func (m *DeletePhotoResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*UploadPhotoReq)(nil), "protoresources.UploadPhotoReq")
	proto.RegisterType((*UploadPhotoResp)(nil), "protoresources.UploadPhotoResp")
	proto.RegisterType((*DeletePhotoReq)(nil), "protoresources.DeletePhotoReq")
	proto.RegisterType((*DeletePhotoResp)(nil), "protoresources.DeletePhotoResp")
}

func init() {
	proto.RegisterFile("resources.proto", fileDescriptor_cf1b13971fe4c19d)
}

var fileDescriptor_cf1b13971fe4c19d = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x9b, 0x56, 0xba, 0x75, 0x94, 0x5d, 0x08, 0x82, 0xab, 0x07, 0x5d, 0x72, 0x5a, 0x50,
	0xf6, 0xa0, 0xaf, 0x50, 0x0f, 0x3d, 0x08, 0x25, 0xd0, 0x8b, 0x97, 0xb2, 0xee, 0x8e, 0x58, 0x0c,
	0x26, 0x66, 0xb2, 0xd8, 0x67, 0xf3, 0xe9, 0x24, 0xd1, 0x2e, 0x8d, 0x8b, 0x3d, 0x4d, 0xfe, 0xcc,
	0xf0, 0xcd, 0x3f, 0x3f, 0x64, 0x16, 0x49, 0x77, 0xb6, 0x41, 0xaa, 0x8c, 0xd5, 0x4e, 0xf3, 0x34,
	0x94, 0xfe, 0x57, 0x3c, 0x41, 0xba, 0x32, 0x4a, 0xd7, 0xed, 0xf2, 0x55, 0x3b, 0x2d, 0xf1, 0x83,
	0x73, 0x38, 0x7a, 0xd9, 0x28, 0xcc, 0x59, 0xc1, 0xca, 0x53, 0x19, 0xde, 0xfc, 0x02, 0x66, 0xbe,
	0xae, 0x71, 0xeb, 0xf2, 0x71, 0xc1, 0xca, 0x63, 0x99, 0x78, 0xfd, 0xb0, 0x75, 0xbe, 0xa5, 0x3f,
	0xdf, 0xd1, 0xae, 0x37, 0x6d, 0x3e, 0x29, 0x58, 0x39, 0x91, 0x49, 0xd0, 0x8b, 0x56, 0xdc, 0x42,
	0x16, 0xb1, 0xc9, 0xf4, 0xa0, 0xce, 0xaa, 0xb0, 0xe0, 0x17, 0xb4, 0xb2, 0x4a, 0xcc, 0x21, 0x9d,
	0xa3, 0x42, 0x87, 0xbd, 0x93, 0xff, 0x87, 0xf9, 0x39, 0x24, 0x1d, 0xfd, 0x2c, 0x1d, 0x87, 0xa5,
	0x53, 0x2f, 0x17, 0xad, 0xb8, 0x81, 0x2c, 0xa2, 0x90, 0xe1, 0x39, 0x24, 0xd4, 0x35, 0x0d, 0x12,
	0x05, 0xca, 0x4c, 0xee, 0xe4, 0xdd, 0x17, 0x83, 0xb3, 0xc7, 0xda, 0xbe, 0xa1, 0x33, 0xaa, 0x6e,
	0x50, 0xee, 0x52, 0xe1, 0x4b, 0x38, 0xd9, 0x73, 0xce, 0xaf, 0xaa, 0x38, 0xb5, 0x2a, 0x8e, 0xec,
	0xf2, 0xfa, 0x60, 0x9f, 0x8c, 0x18, 0x79, 0xe2, 0x9e, 0xaf, 0x21, 0x31, 0x3e, 0x7d, 0x48, 0xfc,
	0x73, 0x94, 0x18, 0x3d, 0x4f, 0xc3, 0xc4, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xb1,
	0xd0, 0xe9, 0xe3, 0x01, 0x00, 0x00,
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
	DeletePhoto(ctx context.Context, in *DeletePhotoReq, opts ...grpc.CallOption) (*DeletePhotoResp, error)
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

func (c *marketplaceResourcesClient) DeletePhoto(ctx context.Context, in *DeletePhotoReq, opts ...grpc.CallOption) (*DeletePhotoResp, error) {
	out := new(DeletePhotoResp)
	err := c.cc.Invoke(ctx, "/protoresources.MarketplaceResources/DeletePhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketplaceResourcesServer is the server API for MarketplaceResources service.
type MarketplaceResourcesServer interface {
	UploadPhoto(context.Context, *UploadPhotoReq) (*UploadPhotoResp, error)
	DeletePhoto(context.Context, *DeletePhotoReq) (*DeletePhotoResp, error)
}

// UnimplementedMarketplaceResourcesServer can be embedded to have forward compatible implementations.
type UnimplementedMarketplaceResourcesServer struct {
}

func (*UnimplementedMarketplaceResourcesServer) UploadPhoto(ctx context.Context, req *UploadPhotoReq) (*UploadPhotoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadPhoto not implemented")
}
func (*UnimplementedMarketplaceResourcesServer) DeletePhoto(ctx context.Context, req *DeletePhotoReq) (*DeletePhotoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePhoto not implemented")
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

func _MarketplaceResources_DeletePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePhotoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceResourcesServer).DeletePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoresources.MarketplaceResources/DeletePhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceResourcesServer).DeletePhoto(ctx, req.(*DeletePhotoReq))
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
		{
			MethodName: "DeletePhoto",
			Handler:    _MarketplaceResources_DeletePhoto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resources.proto",
}
