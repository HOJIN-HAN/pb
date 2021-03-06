// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.12.0
// source: greeting.proto

package greetinggrpc

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type GreetingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GreetingRequest) Reset() {
	*x = GreetingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greeting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingRequest) ProtoMessage() {}

func (x *GreetingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greeting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingRequest.ProtoReflect.Descriptor instead.
func (*GreetingRequest) Descriptor() ([]byte, []int) {
	return file_greeting_proto_rawDescGZIP(), []int{0}
}

func (x *GreetingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GreetingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GreetingReply) Reset() {
	*x = GreetingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greeting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingReply) ProtoMessage() {}

func (x *GreetingReply) ProtoReflect() protoreflect.Message {
	mi := &file_greeting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingReply.ProtoReflect.Descriptor instead.
func (*GreetingReply) Descriptor() ([]byte, []int) {
	return file_greeting_proto_rawDescGZIP(), []int{1}
}

func (x *GreetingReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_greeting_proto protoreflect.FileDescriptor

var file_greeting_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x67, 0x72, 0x70, 0x63, 0x22, 0x25,
	0x0a, 0x0f, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0xa0, 0x01, 0x0a, 0x08, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x48, 0x0a,
	0x08, 0x53, 0x61, 0x79, 0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greeting_proto_rawDescOnce sync.Once
	file_greeting_proto_rawDescData = file_greeting_proto_rawDesc
)

func file_greeting_proto_rawDescGZIP() []byte {
	file_greeting_proto_rawDescOnce.Do(func() {
		file_greeting_proto_rawDescData = protoimpl.X.CompressGZIP(file_greeting_proto_rawDescData)
	})
	return file_greeting_proto_rawDescData
}

var file_greeting_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_greeting_proto_goTypes = []interface{}{
	(*GreetingRequest)(nil), // 0: greetinggrpc.GreetingRequest
	(*GreetingReply)(nil),   // 1: greetinggrpc.GreetingReply
}
var file_greeting_proto_depIdxs = []int32{
	0, // 0: greetinggrpc.Greeting.SayGreet:input_type -> greetinggrpc.GreetingRequest
	0, // 1: greetinggrpc.Greeting.CountGreet:input_type -> greetinggrpc.GreetingRequest
	1, // 2: greetinggrpc.Greeting.SayGreet:output_type -> greetinggrpc.GreetingReply
	1, // 3: greetinggrpc.Greeting.CountGreet:output_type -> greetinggrpc.GreetingReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_greeting_proto_init() }
func file_greeting_proto_init() {
	if File_greeting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greeting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingRequest); i {
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
		file_greeting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingReply); i {
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
			RawDescriptor: file_greeting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greeting_proto_goTypes,
		DependencyIndexes: file_greeting_proto_depIdxs,
		MessageInfos:      file_greeting_proto_msgTypes,
	}.Build()
	File_greeting_proto = out.File
	file_greeting_proto_rawDesc = nil
	file_greeting_proto_goTypes = nil
	file_greeting_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreetingClient is the client API for Greeting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetingClient interface {
	SayGreet(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingReply, error)
	CountGreet(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingReply, error)
}

type greetingClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetingClient(cc grpc.ClientConnInterface) GreetingClient {
	return &greetingClient{cc}
}

func (c *greetingClient) SayGreet(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingReply, error) {
	out := new(GreetingReply)
	err := c.cc.Invoke(ctx, "/greetinggrpc.Greeting/SayGreet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetingClient) CountGreet(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingReply, error) {
	out := new(GreetingReply)
	err := c.cc.Invoke(ctx, "/greetinggrpc.Greeting/CountGreet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetingServer is the server API for Greeting service.
type GreetingServer interface {
	SayGreet(context.Context, *GreetingRequest) (*GreetingReply, error)
	CountGreet(context.Context, *GreetingRequest) (*GreetingReply, error)
}

// UnimplementedGreetingServer can be embedded to have forward compatible implementations.
type UnimplementedGreetingServer struct {
}

func (*UnimplementedGreetingServer) SayGreet(context.Context, *GreetingRequest) (*GreetingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayGreet not implemented")
}
func (*UnimplementedGreetingServer) CountGreet(context.Context, *GreetingRequest) (*GreetingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountGreet not implemented")
}

func RegisterGreetingServer(s *grpc.Server, srv GreetingServer) {
	s.RegisterService(&_Greeting_serviceDesc, srv)
}

func _Greeting_SayGreet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingServer).SayGreet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greetinggrpc.Greeting/SayGreet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingServer).SayGreet(ctx, req.(*GreetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeting_CountGreet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingServer).CountGreet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greetinggrpc.Greeting/CountGreet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingServer).CountGreet(ctx, req.(*GreetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeting_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greetinggrpc.Greeting",
	HandlerType: (*GreetingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayGreet",
			Handler:    _Greeting_SayGreet_Handler,
		},
		{
			MethodName: "CountGreet",
			Handler:    _Greeting_CountGreet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greeting.proto",
}
