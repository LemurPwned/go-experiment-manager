// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/msg.proto

package msg

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExperimentID string `protobuf:"bytes,1,opt,name=experimentID,proto3" json:"experimentID,omitempty"`
	MetricBody   string `protobuf:"bytes,2,opt,name=metricBody,proto3" json:"metricBody,omitempty"`
	CreatedAt    int64  `protobuf:"varint,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_proto_msg_proto_rawDescGZIP(), []int{0}
}

func (x *Metric) GetExperimentID() string {
	if x != nil {
		return x.ExperimentID
	}
	return ""
}

func (x *Metric) GetMetricBody() string {
	if x != nil {
		return x.MetricBody
	}
	return ""
}

func (x *Metric) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type MetricsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MetricsReply) Reset() {
	*x = MetricsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsReply) ProtoMessage() {}

func (x *MetricsReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsReply.ProtoReflect.Descriptor instead.
func (*MetricsReply) Descriptor() ([]byte, []int) {
	return file_proto_msg_proto_rawDescGZIP(), []int{1}
}

func (x *MetricsReply) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *MetricsReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_msg_proto protoreflect.FileDescriptor

var file_proto_msg_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x42,
	0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x48, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x3f, 0x0a,
	0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x65, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a,
	0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x0b, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x1a, 0x11, 0x2e, 0x6d, 0x73, 0x67, 0x2e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_msg_proto_rawDescOnce sync.Once
	file_proto_msg_proto_rawDescData = file_proto_msg_proto_rawDesc
)

func file_proto_msg_proto_rawDescGZIP() []byte {
	file_proto_msg_proto_rawDescOnce.Do(func() {
		file_proto_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_msg_proto_rawDescData)
	})
	return file_proto_msg_proto_rawDescData
}

var file_proto_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_msg_proto_goTypes = []interface{}{
	(*Metric)(nil),       // 0: msg.Metric
	(*MetricsReply)(nil), // 1: msg.MetricsReply
}
var file_proto_msg_proto_depIdxs = []int32{
	0, // 0: msg.MetricSerice.SendMetrics:input_type -> msg.Metric
	1, // 1: msg.MetricSerice.SendMetrics:output_type -> msg.MetricsReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_msg_proto_init() }
func file_proto_msg_proto_init() {
	if File_proto_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
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
		file_proto_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsReply); i {
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
			RawDescriptor: file_proto_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_msg_proto_goTypes,
		DependencyIndexes: file_proto_msg_proto_depIdxs,
		MessageInfos:      file_proto_msg_proto_msgTypes,
	}.Build()
	File_proto_msg_proto = out.File
	file_proto_msg_proto_rawDesc = nil
	file_proto_msg_proto_goTypes = nil
	file_proto_msg_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MetricSericeClient is the client API for MetricSerice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricSericeClient interface {
	SendMetrics(ctx context.Context, in *Metric, opts ...grpc.CallOption) (*MetricsReply, error)
}

type metricSericeClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricSericeClient(cc grpc.ClientConnInterface) MetricSericeClient {
	return &metricSericeClient{cc}
}

func (c *metricSericeClient) SendMetrics(ctx context.Context, in *Metric, opts ...grpc.CallOption) (*MetricsReply, error) {
	out := new(MetricsReply)
	err := c.cc.Invoke(ctx, "/msg.MetricSerice/SendMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricSericeServer is the server API for MetricSerice service.
type MetricSericeServer interface {
	SendMetrics(context.Context, *Metric) (*MetricsReply, error)
}

// UnimplementedMetricSericeServer can be embedded to have forward compatible implementations.
type UnimplementedMetricSericeServer struct {
}

func (*UnimplementedMetricSericeServer) SendMetrics(context.Context, *Metric) (*MetricsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMetrics not implemented")
}

func RegisterMetricSericeServer(s *grpc.Server, srv MetricSericeServer) {
	s.RegisterService(&_MetricSerice_serviceDesc, srv)
}

func _MetricSerice_SendMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Metric)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricSericeServer).SendMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.MetricSerice/SendMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricSericeServer).SendMetrics(ctx, req.(*Metric))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricSerice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "msg.MetricSerice",
	HandlerType: (*MetricSericeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMetrics",
			Handler:    _MetricSerice_SendMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/msg.proto",
}
