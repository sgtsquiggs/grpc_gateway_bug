// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: svcone/svcone.proto

package proto

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

type Table int32

const (
	Table_TABLE_INVALID    Table = 0
	Table_LEDGERS          Table = 1
	Table_TRANSACTIONS     Table = 2
	Table_TRADES           Table = 3
	Table_CORPORATEACTIONS Table = 4
)

// Enum value maps for Table.
var (
	Table_name = map[int32]string{
		0: "TABLE_INVALID",
		1: "LEDGERS",
		2: "TRANSACTIONS",
		3: "TRADES",
		4: "CORPORATEACTIONS",
	}
	Table_value = map[string]int32{
		"TABLE_INVALID":    0,
		"LEDGERS":          1,
		"TRANSACTIONS":     2,
		"TRADES":           3,
		"CORPORATEACTIONS": 4,
	}
)

func (x Table) Enum() *Table {
	p := new(Table)
	*p = x
	return p
}

func (x Table) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Table) Descriptor() protoreflect.EnumDescriptor {
	return file_svcone_svcone_proto_enumTypes[0].Descriptor()
}

func (Table) Type() protoreflect.EnumType {
	return &file_svcone_svcone_proto_enumTypes[0]
}

func (x Table) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Table.Descriptor instead.
func (Table) EnumDescriptor() ([]byte, []int) {
	return file_svcone_svcone_proto_rawDescGZIP(), []int{0}
}

type ExportDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table Table `protobuf:"varint,1,opt,name=table,proto3,enum=svcone.Table" json:"table,omitempty"`
}

func (x *ExportDataRequest) Reset() {
	*x = ExportDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svcone_svcone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportDataRequest) ProtoMessage() {}

func (x *ExportDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svcone_svcone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportDataRequest.ProtoReflect.Descriptor instead.
func (*ExportDataRequest) Descriptor() ([]byte, []int) {
	return file_svcone_svcone_proto_rawDescGZIP(), []int{0}
}

func (x *ExportDataRequest) GetTable() Table {
	if x != nil {
		return x.Table
	}
	return Table_TABLE_INVALID
}

type ExportDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ExportDataResponse) Reset() {
	*x = ExportDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svcone_svcone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportDataResponse) ProtoMessage() {}

func (x *ExportDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svcone_svcone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportDataResponse.ProtoReflect.Descriptor instead.
func (*ExportDataResponse) Descriptor() ([]byte, []int) {
	return file_svcone_svcone_proto_rawDescGZIP(), []int{1}
}

func (x *ExportDataResponse) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_svcone_svcone_proto protoreflect.FileDescriptor

var file_svcone_svcone_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x76, 0x63, 0x6f, 0x6e, 0x65, 0x2f, 0x73, 0x76, 0x63, 0x6f, 0x6e, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x76, 0x63, 0x6f, 0x6e, 0x65, 0x22, 0x38, 0x0a,
	0x11, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0d, 0x2e, 0x73, 0x76, 0x63, 0x6f, 0x6e, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x28, 0x0a, 0x12, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x2a, 0x5b, 0x0a, 0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x41,
	0x42, 0x4c, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x4c, 0x45, 0x44, 0x47, 0x45, 0x52, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x54, 0x52, 0x41, 0x44, 0x45, 0x53, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x52, 0x50,
	0x4f, 0x52, 0x41, 0x54, 0x45, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x04, 0x32, 0x4d,
	0x0a, 0x06, 0x53, 0x56, 0x43, 0x4f, 0x6e, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x2e, 0x73, 0x76, 0x63, 0x6f, 0x6e, 0x65, 0x2e,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x76, 0x63, 0x6f, 0x6e, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x36, 0x5a,
	0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x67, 0x74, 0x73,
	0x71, 0x75, 0x69, 0x67, 0x67, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x5f, 0x62, 0x75, 0x67, 0x2f, 0x73, 0x76, 0x63, 0x2f, 0x6f, 0x6e, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svcone_svcone_proto_rawDescOnce sync.Once
	file_svcone_svcone_proto_rawDescData = file_svcone_svcone_proto_rawDesc
)

func file_svcone_svcone_proto_rawDescGZIP() []byte {
	file_svcone_svcone_proto_rawDescOnce.Do(func() {
		file_svcone_svcone_proto_rawDescData = protoimpl.X.CompressGZIP(file_svcone_svcone_proto_rawDescData)
	})
	return file_svcone_svcone_proto_rawDescData
}

var file_svcone_svcone_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_svcone_svcone_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_svcone_svcone_proto_goTypes = []interface{}{
	(Table)(0),                 // 0: svcone.Table
	(*ExportDataRequest)(nil),  // 1: svcone.ExportDataRequest
	(*ExportDataResponse)(nil), // 2: svcone.ExportDataResponse
}
var file_svcone_svcone_proto_depIdxs = []int32{
	0, // 0: svcone.ExportDataRequest.table:type_name -> svcone.Table
	1, // 1: svcone.SVCOne.ExportData:input_type -> svcone.ExportDataRequest
	2, // 2: svcone.SVCOne.ExportData:output_type -> svcone.ExportDataResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_svcone_svcone_proto_init() }
func file_svcone_svcone_proto_init() {
	if File_svcone_svcone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svcone_svcone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportDataRequest); i {
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
		file_svcone_svcone_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportDataResponse); i {
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
			RawDescriptor: file_svcone_svcone_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svcone_svcone_proto_goTypes,
		DependencyIndexes: file_svcone_svcone_proto_depIdxs,
		EnumInfos:         file_svcone_svcone_proto_enumTypes,
		MessageInfos:      file_svcone_svcone_proto_msgTypes,
	}.Build()
	File_svcone_svcone_proto = out.File
	file_svcone_svcone_proto_rawDesc = nil
	file_svcone_svcone_proto_goTypes = nil
	file_svcone_svcone_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SVCOneClient is the client API for SVCOne service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SVCOneClient interface {
	ExportData(ctx context.Context, in *ExportDataRequest, opts ...grpc.CallOption) (*ExportDataResponse, error)
}

type sVCOneClient struct {
	cc grpc.ClientConnInterface
}

func NewSVCOneClient(cc grpc.ClientConnInterface) SVCOneClient {
	return &sVCOneClient{cc}
}

func (c *sVCOneClient) ExportData(ctx context.Context, in *ExportDataRequest, opts ...grpc.CallOption) (*ExportDataResponse, error) {
	out := new(ExportDataResponse)
	err := c.cc.Invoke(ctx, "/svcone.SVCOne/ExportData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SVCOneServer is the server API for SVCOne service.
type SVCOneServer interface {
	ExportData(context.Context, *ExportDataRequest) (*ExportDataResponse, error)
}

// UnimplementedSVCOneServer can be embedded to have forward compatible implementations.
type UnimplementedSVCOneServer struct {
}

func (*UnimplementedSVCOneServer) ExportData(context.Context, *ExportDataRequest) (*ExportDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportData not implemented")
}

func RegisterSVCOneServer(s *grpc.Server, srv SVCOneServer) {
	s.RegisterService(&_SVCOne_serviceDesc, srv)
}

func _SVCOne_ExportData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SVCOneServer).ExportData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svcone.SVCOne/ExportData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SVCOneServer).ExportData(ctx, req.(*ExportDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SVCOne_serviceDesc = grpc.ServiceDesc{
	ServiceName: "svcone.SVCOne",
	HandlerType: (*SVCOneServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExportData",
			Handler:    _SVCOne_ExportData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svcone/svcone.proto",
}
