// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Empty struct {
	Index                int64    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
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

func (m *Empty) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type GetIDReply struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIDReply) Reset()         { *m = GetIDReply{} }
func (m *GetIDReply) String() string { return proto.CompactTextString(m) }
func (*GetIDReply) ProtoMessage()    {}
func (*GetIDReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *GetIDReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIDReply.Unmarshal(m, b)
}
func (m *GetIDReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIDReply.Marshal(b, m, deterministic)
}
func (m *GetIDReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIDReply.Merge(m, src)
}
func (m *GetIDReply) XXX_Size() int {
	return xxx_messageInfo_GetIDReply.Size(m)
}
func (m *GetIDReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIDReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetIDReply proto.InternalMessageInfo

func (m *GetIDReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "dtalk.generator.Empty")
	proto.RegisterType((*GetIDReply)(nil), "dtalk.generator.GetIDReply")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4f, 0x29, 0x49, 0xcc, 0xc9, 0xd6, 0x4b, 0x4f, 0xcd,
	0x4b, 0x2d, 0x4a, 0x2c, 0xc9, 0x2f, 0x52, 0x92, 0xe5, 0x62, 0x75, 0xcd, 0x2d, 0x28, 0xa9, 0x14,
	0x12, 0xe1, 0x62, 0xcd, 0xcc, 0x4b, 0x49, 0xad, 0x90, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x82,
	0x70, 0x94, 0x64, 0xb8, 0xb8, 0xdc, 0x53, 0x4b, 0x3c, 0x5d, 0x82, 0x52, 0x0b, 0x72, 0x2a, 0x85,
	0xf8, 0xb8, 0x98, 0x32, 0x53, 0xa0, 0x0a, 0x98, 0x32, 0x53, 0x8c, 0x3c, 0xb9, 0x38, 0xdd, 0x61,
	0x26, 0x09, 0xd9, 0x70, 0xb1, 0x82, 0x95, 0x0a, 0x89, 0xe9, 0xa1, 0x59, 0xa2, 0x07, 0xb6, 0x41,
	0x4a, 0x1a, 0x43, 0x1c, 0x61, 0xb4, 0x13, 0x6b, 0x14, 0x73, 0x62, 0x41, 0x66, 0x12, 0x1b, 0xd8,
	0x99, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0x54, 0x9b, 0x70, 0xb3, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GeneratorClient is the client API for Generator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeneratorClient interface {
	GetID(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetIDReply, error)
}

type generatorClient struct {
	cc *grpc.ClientConn
}

func NewGeneratorClient(cc *grpc.ClientConn) GeneratorClient {
	return &generatorClient{cc}
}

func (c *generatorClient) GetID(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetIDReply, error) {
	out := new(GetIDReply)
	err := c.cc.Invoke(ctx, "/dtalk.generator.Generator/GetID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeneratorServer is the server API for Generator service.
type GeneratorServer interface {
	GetID(context.Context, *Empty) (*GetIDReply, error)
}

func RegisterGeneratorServer(s *grpc.Server, srv GeneratorServer) {
	s.RegisterService(&_Generator_serviceDesc, srv)
}

func _Generator_GetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServer).GetID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtalk.generator.Generator/GetID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServer).GetID(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Generator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dtalk.generator.Generator",
	HandlerType: (*GeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetID",
			Handler:    _Generator_GetID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
