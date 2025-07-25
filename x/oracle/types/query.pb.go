// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/oracle/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("elys/oracle/query.proto", fileDescriptor_2afc8ed4b42b5980) }

var fileDescriptor_2afc8ed4b42b5980 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x31, 0x4e, 0xc4, 0x30,
	0x10, 0x45, 0x93, 0x02, 0x90, 0x96, 0x0e, 0x21, 0x01, 0xab, 0x95, 0x0f, 0x80, 0x44, 0x46, 0x0b,
	0x12, 0x07, 0xa0, 0xa2, 0xa5, 0xa5, 0x89, 0xc6, 0xd9, 0x59, 0x63, 0x91, 0x78, 0x8c, 0xed, 0x5d,
	0xc8, 0x2d, 0x38, 0x16, 0x65, 0x4a, 0x4a, 0x94, 0x5c, 0x04, 0x25, 0x4e, 0xa4, 0x20, 0xb6, 0xf3,
	0xf8, 0xfd, 0xf9, 0xf3, 0xff, 0xe2, 0x82, 0xca, 0xda, 0x03, 0x3b, 0x2c, 0x4a, 0x82, 0xb7, 0x1d,
	0xb9, 0x3a, 0xb3, 0x8e, 0x03, 0x9f, 0x9d, 0xf6, 0x20, 0x8b, 0x60, 0x79, 0xae, 0x58, 0xf1, 0xf0,
	0x0f, 0xfd, 0x2b, 0x4a, 0x96, 0x2b, 0xc5, 0xac, 0x4a, 0x02, 0xb4, 0x1a, 0xd0, 0x18, 0x0e, 0x18,
	0x34, 0x1b, 0x3f, 0xd2, 0xeb, 0x82, 0x7d, 0xc5, 0x1e, 0x24, 0xfa, 0xd1, 0x19, 0xf6, 0x6b, 0x49,
	0x01, 0xd7, 0x60, 0x51, 0x69, 0x33, 0x88, 0x47, 0xed, 0xe5, 0x3c, 0x85, 0x45, 0x87, 0xd5, 0xe4,
	0xb2, 0x9a, 0x13, 0x89, 0x66, 0x93, 0x5b, 0xa7, 0x0b, 0x3a, 0x44, 0xd1, 0x7b, 0x0a, 0xb9, 0x36,
	0xdb, 0x29, 0xdf, 0x9f, 0x6e, 0xf3, 0x35, 0xf1, 0x0f, 0xe4, 0x5b, 0xa2, 0x0d, 0xb9, 0x91, 0x5f,
	0xc5, 0xe8, 0x79, 0x6c, 0x1c, 0x87, 0x88, 0x6e, 0x4f, 0x16, 0x47, 0x4f, 0x7d, 0x97, 0x87, 0xc7,
	0xaf, 0x56, 0xa4, 0x4d, 0x2b, 0xd2, 0x9f, 0x56, 0xa4, 0x9f, 0x9d, 0x48, 0x9a, 0x4e, 0x24, 0xdf,
	0x9d, 0x48, 0x9e, 0x33, 0xa5, 0xc3, 0xcb, 0x4e, 0x66, 0x05, 0x57, 0xd0, 0x1f, 0xba, 0x31, 0x14,
	0xde, 0xd9, 0xbd, 0x0e, 0x03, 0xec, 0xef, 0xe1, 0x63, 0x3a, 0x1d, 0x6a, 0x4b, 0x5e, 0x1e, 0x0f,
	0xce, 0x77, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x65, 0xd7, 0xb9, 0x43, 0x8b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var Query_serviceDesc = _Query_serviceDesc
var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "elys.oracle.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "elys/oracle/query.proto",
}
