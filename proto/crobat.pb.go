// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: crobat.proto

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

type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crobat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crobat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_crobat_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type Domain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Ipv4   string `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
}

func (x *Domain) Reset() {
	*x = Domain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crobat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Domain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Domain) ProtoMessage() {}

func (x *Domain) ProtoReflect() protoreflect.Message {
	mi := &file_crobat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Domain.ProtoReflect.Descriptor instead.
func (*Domain) Descriptor() ([]byte, []int) {
	return file_crobat_proto_rawDescGZIP(), []int{1}
}

func (x *Domain) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Domain) GetIpv4() string {
	if x != nil {
		return x.Ipv4
	}
	return ""
}

type ReverseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*ReverseResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ReverseReply) Reset() {
	*x = ReverseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crobat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReverseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReverseReply) ProtoMessage() {}

func (x *ReverseReply) ProtoReflect() protoreflect.Message {
	mi := &file_crobat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReverseReply.ProtoReflect.Descriptor instead.
func (*ReverseReply) Descriptor() ([]byte, []int) {
	return file_crobat_proto_rawDescGZIP(), []int{2}
}

func (x *ReverseReply) GetResults() []*ReverseResult {
	if x != nil {
		return x.Results
	}
	return nil
}

type ReverseResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip      string    `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Domains []*Domain `protobuf:"bytes,2,rep,name=domains,proto3" json:"domains,omitempty"`
}

func (x *ReverseResult) Reset() {
	*x = ReverseResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crobat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReverseResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReverseResult) ProtoMessage() {}

func (x *ReverseResult) ProtoReflect() protoreflect.Message {
	mi := &file_crobat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReverseResult.ProtoReflect.Descriptor instead.
func (*ReverseResult) Descriptor() ([]byte, []int) {
	return file_crobat_proto_rawDescGZIP(), []int{3}
}

func (x *ReverseResult) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ReverseResult) GetDomains() []*Domain {
	if x != nil {
		return x.Domains
	}
	return nil
}

var File_crobat_proto protoreflect.FileDescriptor

var file_crobat_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x72, 0x6f, 0x62, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x34, 0x0a, 0x06, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x76,
	0x34, 0x22, 0x3e, 0x0a, 0x0c, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x22, 0x48, 0x0a, 0x0d, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x12, 0x27, 0x0a, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x52, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x32, 0xec, 0x01, 0x0a, 0x06,
	0x43, 0x72, 0x6f, 0x62, 0x61, 0x74, 0x12, 0x37, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x31, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x4c, 0x44, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x34, 0x0a, 0x0a, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x44, 0x4e, 0x53,
	0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x12, 0x40, 0x0a, 0x0f, 0x52, 0x65, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x44, 0x4e, 0x53, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_crobat_proto_rawDescOnce sync.Once
	file_crobat_proto_rawDescData = file_crobat_proto_rawDesc
)

func file_crobat_proto_rawDescGZIP() []byte {
	file_crobat_proto_rawDescOnce.Do(func() {
		file_crobat_proto_rawDescData = protoimpl.X.CompressGZIP(file_crobat_proto_rawDescData)
	})
	return file_crobat_proto_rawDescData
}

var file_crobat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_crobat_proto_goTypes = []interface{}{
	(*QueryRequest)(nil),  // 0: proto.QueryRequest
	(*Domain)(nil),        // 1: proto.Domain
	(*ReverseReply)(nil),  // 2: proto.ReverseReply
	(*ReverseResult)(nil), // 3: proto.ReverseResult
}
var file_crobat_proto_depIdxs = []int32{
	3, // 0: proto.ReverseReply.results:type_name -> proto.ReverseResult
	1, // 1: proto.ReverseResult.domains:type_name -> proto.Domain
	0, // 2: proto.Crobat.GetSubdomains:input_type -> proto.QueryRequest
	0, // 3: proto.Crobat.GetTLDs:input_type -> proto.QueryRequest
	0, // 4: proto.Crobat.ReverseDNS:input_type -> proto.QueryRequest
	0, // 5: proto.Crobat.ReverseDNSRange:input_type -> proto.QueryRequest
	1, // 6: proto.Crobat.GetSubdomains:output_type -> proto.Domain
	1, // 7: proto.Crobat.GetTLDs:output_type -> proto.Domain
	1, // 8: proto.Crobat.ReverseDNS:output_type -> proto.Domain
	3, // 9: proto.Crobat.ReverseDNSRange:output_type -> proto.ReverseResult
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_crobat_proto_init() }
func file_crobat_proto_init() {
	if File_crobat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crobat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_crobat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Domain); i {
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
		file_crobat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReverseReply); i {
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
		file_crobat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReverseResult); i {
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
			RawDescriptor: file_crobat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crobat_proto_goTypes,
		DependencyIndexes: file_crobat_proto_depIdxs,
		MessageInfos:      file_crobat_proto_msgTypes,
	}.Build()
	File_crobat_proto = out.File
	file_crobat_proto_rawDesc = nil
	file_crobat_proto_goTypes = nil
	file_crobat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CrobatClient is the client API for Crobat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CrobatClient interface {
	GetSubdomains(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (Crobat_GetSubdomainsClient, error)
	GetTLDs(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (Crobat_GetTLDsClient, error)
	ReverseDNS(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (Crobat_ReverseDNSClient, error)
	ReverseDNSRange(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (Crobat_ReverseDNSRangeClient, error)
}

type crobatClient struct {
	cc grpc.ClientConnInterface
}

func NewCrobatClient(cc grpc.ClientConnInterface) CrobatClient {
	return &crobatClient{cc}
}

func (c *crobatClient) GetSubdomains(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (Crobat_GetSubdomainsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Crobat_serviceDesc.Streams[0], "/proto.Crobat/GetSubdomains", opts...)
	if err != nil {
		return nil, err
	}
	x := &crobatGetSubdomainsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Crobat_GetSubdomainsClient interface {
	Recv() (*Domain, error)
	grpc.ClientStream
}

type crobatGetSubdomainsClient struct {
	grpc.ClientStream
}

func (x *crobatGetSubdomainsClient) Recv() (*Domain, error) {
	m := new(Domain)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *crobatClient) GetTLDs(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (Crobat_GetTLDsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Crobat_serviceDesc.Streams[1], "/proto.Crobat/GetTLDs", opts...)
	if err != nil {
		return nil, err
	}
	x := &crobatGetTLDsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Crobat_GetTLDsClient interface {
	Recv() (*Domain, error)
	grpc.ClientStream
}

type crobatGetTLDsClient struct {
	grpc.ClientStream
}

func (x *crobatGetTLDsClient) Recv() (*Domain, error) {
	m := new(Domain)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *crobatClient) ReverseDNS(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (Crobat_ReverseDNSClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Crobat_serviceDesc.Streams[2], "/proto.Crobat/ReverseDNS", opts...)
	if err != nil {
		return nil, err
	}
	x := &crobatReverseDNSClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Crobat_ReverseDNSClient interface {
	Recv() (*Domain, error)
	grpc.ClientStream
}

type crobatReverseDNSClient struct {
	grpc.ClientStream
}

func (x *crobatReverseDNSClient) Recv() (*Domain, error) {
	m := new(Domain)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *crobatClient) ReverseDNSRange(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (Crobat_ReverseDNSRangeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Crobat_serviceDesc.Streams[3], "/proto.Crobat/ReverseDNSRange", opts...)
	if err != nil {
		return nil, err
	}
	x := &crobatReverseDNSRangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Crobat_ReverseDNSRangeClient interface {
	Recv() (*ReverseResult, error)
	grpc.ClientStream
}

type crobatReverseDNSRangeClient struct {
	grpc.ClientStream
}

func (x *crobatReverseDNSRangeClient) Recv() (*ReverseResult, error) {
	m := new(ReverseResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CrobatServer is the server API for Crobat service.
type CrobatServer interface {
	GetSubdomains(*QueryRequest, Crobat_GetSubdomainsServer) error
	GetTLDs(*QueryRequest, Crobat_GetTLDsServer) error
	ReverseDNS(*QueryRequest, Crobat_ReverseDNSServer) error
	ReverseDNSRange(*QueryRequest, Crobat_ReverseDNSRangeServer) error
}

// UnimplementedCrobatServer can be embedded to have forward compatible implementations.
type UnimplementedCrobatServer struct {
}

func (*UnimplementedCrobatServer) GetSubdomains(*QueryRequest, Crobat_GetSubdomainsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSubdomains not implemented")
}
func (*UnimplementedCrobatServer) GetTLDs(*QueryRequest, Crobat_GetTLDsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTLDs not implemented")
}
func (*UnimplementedCrobatServer) ReverseDNS(*QueryRequest, Crobat_ReverseDNSServer) error {
	return status.Errorf(codes.Unimplemented, "method ReverseDNS not implemented")
}
func (*UnimplementedCrobatServer) ReverseDNSRange(*QueryRequest, Crobat_ReverseDNSRangeServer) error {
	return status.Errorf(codes.Unimplemented, "method ReverseDNSRange not implemented")
}

func RegisterCrobatServer(s *grpc.Server, srv CrobatServer) {
	s.RegisterService(&_Crobat_serviceDesc, srv)
}

func _Crobat_GetSubdomains_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CrobatServer).GetSubdomains(m, &crobatGetSubdomainsServer{stream})
}

type Crobat_GetSubdomainsServer interface {
	Send(*Domain) error
	grpc.ServerStream
}

type crobatGetSubdomainsServer struct {
	grpc.ServerStream
}

func (x *crobatGetSubdomainsServer) Send(m *Domain) error {
	return x.ServerStream.SendMsg(m)
}

func _Crobat_GetTLDs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CrobatServer).GetTLDs(m, &crobatGetTLDsServer{stream})
}

type Crobat_GetTLDsServer interface {
	Send(*Domain) error
	grpc.ServerStream
}

type crobatGetTLDsServer struct {
	grpc.ServerStream
}

func (x *crobatGetTLDsServer) Send(m *Domain) error {
	return x.ServerStream.SendMsg(m)
}

func _Crobat_ReverseDNS_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CrobatServer).ReverseDNS(m, &crobatReverseDNSServer{stream})
}

type Crobat_ReverseDNSServer interface {
	Send(*Domain) error
	grpc.ServerStream
}

type crobatReverseDNSServer struct {
	grpc.ServerStream
}

func (x *crobatReverseDNSServer) Send(m *Domain) error {
	return x.ServerStream.SendMsg(m)
}

func _Crobat_ReverseDNSRange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CrobatServer).ReverseDNSRange(m, &crobatReverseDNSRangeServer{stream})
}

type Crobat_ReverseDNSRangeServer interface {
	Send(*ReverseResult) error
	grpc.ServerStream
}

type crobatReverseDNSRangeServer struct {
	grpc.ServerStream
}

func (x *crobatReverseDNSRangeServer) Send(m *ReverseResult) error {
	return x.ServerStream.SendMsg(m)
}

var _Crobat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Crobat",
	HandlerType: (*CrobatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetSubdomains",
			Handler:       _Crobat_GetSubdomains_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetTLDs",
			Handler:       _Crobat_GetTLDs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReverseDNS",
			Handler:       _Crobat_ReverseDNS_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReverseDNSRange",
			Handler:       _Crobat_ReverseDNSRange_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "crobat.proto",
}
