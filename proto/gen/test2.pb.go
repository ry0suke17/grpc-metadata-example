// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.19.4
// source: test2.proto

package gen

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Test2Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Test2Request) Reset() {
	*x = Test2Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test2Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test2Request) ProtoMessage() {}

func (x *Test2Request) ProtoReflect() protoreflect.Message {
	mi := &file_test2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test2Request.ProtoReflect.Descriptor instead.
func (*Test2Request) Descriptor() ([]byte, []int) {
	return file_test2_proto_rawDescGZIP(), []int{0}
}

type Test2Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Test2Response) Reset() {
	*x = Test2Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test2Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test2Response) ProtoMessage() {}

func (x *Test2Response) ProtoReflect() protoreflect.Message {
	mi := &file_test2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test2Response.ProtoReflect.Descriptor instead.
func (*Test2Response) Descriptor() ([]byte, []int) {
	return file_test2_proto_rawDescGZIP(), []int{1}
}

func (x *Test2Response) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_test2_proto protoreflect.FileDescriptor

var file_test2_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a,
	0x0c, 0x54, 0x65, 0x73, 0x74, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x23, 0x0a,
	0x0d, 0x54, 0x65, 0x73, 0x74, 0x32, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x32, 0x31, 0x0a, 0x05, 0x54, 0x65, 0x73, 0x74, 0x32, 0x12, 0x28, 0x0a, 0x05, 0x54,
	0x65, 0x73, 0x74, 0x32, 0x12, 0x0d, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x32, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x32, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test2_proto_rawDescOnce sync.Once
	file_test2_proto_rawDescData = file_test2_proto_rawDesc
)

func file_test2_proto_rawDescGZIP() []byte {
	file_test2_proto_rawDescOnce.Do(func() {
		file_test2_proto_rawDescData = protoimpl.X.CompressGZIP(file_test2_proto_rawDescData)
	})
	return file_test2_proto_rawDescData
}

var file_test2_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test2_proto_goTypes = []interface{}{
	(*Test2Request)(nil),  // 0: Test2Request
	(*Test2Response)(nil), // 1: Test2Response
}
var file_test2_proto_depIdxs = []int32{
	0, // 0: Test2.Test2:input_type -> Test2Request
	1, // 1: Test2.Test2:output_type -> Test2Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test2_proto_init() }
func file_test2_proto_init() {
	if File_test2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test2Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test2Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test2_proto_goTypes,
		DependencyIndexes: file_test2_proto_depIdxs,
		MessageInfos:      file_test2_proto_msgTypes,
	}.Build()
	File_test2_proto = out.File
	file_test2_proto_rawDesc = nil
	file_test2_proto_goTypes = nil
	file_test2_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// Test2Client is the client API for Test2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Test2Client interface {
	Test2(ctx context.Context, in *Test2Request, opts ...grpc.CallOption) (*Test2Response, error)
}

type test2Client struct {
	cc grpc.ClientConnInterface
}

func NewTest2Client(cc grpc.ClientConnInterface) Test2Client {
	return &test2Client{cc}
}

func (c *test2Client) Test2(ctx context.Context, in *Test2Request, opts ...grpc.CallOption) (*Test2Response, error) {
	out := new(Test2Response)
	err := c.cc.Invoke(ctx, "/Test2/Test2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Test2Server is the server API for Test2 service.
type Test2Server interface {
	Test2(context.Context, *Test2Request) (*Test2Response, error)
}

// UnimplementedTest2Server can be embedded to have forward compatible implementations.
type UnimplementedTest2Server struct {
}

func (*UnimplementedTest2Server) Test2(context.Context, *Test2Request) (*Test2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test2 not implemented")
}

func RegisterTest2Server(s *grpc.Server, srv Test2Server) {
	s.RegisterService(&_Test2_serviceDesc, srv)
}

func _Test2_Test2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Test2Server).Test2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Test2/Test2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Test2Server).Test2(ctx, req.(*Test2Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Test2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Test2",
	HandlerType: (*Test2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test2",
			Handler:    _Test2_Test2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test2.proto",
}
