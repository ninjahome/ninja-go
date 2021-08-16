// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.12.3
// source: cmdService.proto

package cmd

import (
	context "context"
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

type WSInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Online   bool   `protobuf:"varint,1,opt,name=online,proto3" json:"online,omitempty"`
	Local    bool   `protobuf:"varint,2,opt,name=local,proto3" json:"local,omitempty"`
	UserAddr string `protobuf:"bytes,3,opt,name=userAddr,proto3" json:"userAddr,omitempty"`
}

func (x *WSInfoReq) Reset() {
	*x = WSInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WSInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WSInfoReq) ProtoMessage() {}

func (x *WSInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WSInfoReq.ProtoReflect.Descriptor instead.
func (*WSInfoReq) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{0}
}

func (x *WSInfoReq) GetOnline() bool {
	if x != nil {
		return x.Online
	}
	return false
}

func (x *WSInfoReq) GetLocal() bool {
	if x != nil {
		return x.Local
	}
	return false
}

func (x *WSInfoReq) GetUserAddr() string {
	if x != nil {
		return x.UserAddr
	}
	return ""
}

type ThreadGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List       bool   `protobuf:"varint,2,opt,name=list,proto3" json:"list,omitempty"`
	ThreadName string `protobuf:"bytes,1,opt,name=threadName,proto3" json:"threadName,omitempty"`
}

func (x *ThreadGroup) Reset() {
	*x = ThreadGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThreadGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadGroup) ProtoMessage() {}

func (x *ThreadGroup) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadGroup.ProtoReflect.Descriptor instead.
func (*ThreadGroup) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{1}
}

func (x *ThreadGroup) GetList() bool {
	if x != nil {
		return x.List
	}
	return false
}

func (x *ThreadGroup) GetThreadName() string {
	if x != nil {
		return x.ThreadName
	}
	return ""
}

type TopicMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *TopicMsg) Reset() {
	*x = TopicMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicMsg) ProtoMessage() {}

func (x *TopicMsg) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicMsg.ProtoReflect.Descriptor instead.
func (*TopicMsg) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{2}
}

func (x *TopicMsg) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *TopicMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ShowPeer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *ShowPeer) Reset() {
	*x = ShowPeer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowPeer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowPeer) ProtoMessage() {}

func (x *ShowPeer) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowPeer.ProtoReflect.Descriptor instead.
func (*ShowPeer) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{3}
}

func (x *ShowPeer) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{4}
}

func (x *CommonResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_cmdService_proto protoreflect.FileDescriptor

var file_cmdService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6d, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x70, 0x62, 0x73, 0x2e, 0x63, 0x6d, 0x64, 0x22, 0x55, 0x0a, 0x09, 0x57,
	0x53, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x22, 0x41, 0x0a, 0x0b, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x08, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x73,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x20, 0x0a, 0x08, 0x53, 0x68, 0x6f,
	0x77, 0x50, 0x65, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x22, 0x0a, 0x0e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32,
	0x8e, 0x02, 0x0a, 0x0a, 0x43, 0x6d, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f,
	0x0a, 0x0f, 0x50, 0x32, 0x70, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x73,
	0x67, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x4d, 0x73, 0x67, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x0c, 0x50, 0x32, 0x70, 0x53, 0x68, 0x6f, 0x77, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12,
	0x11, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x50, 0x65,
	0x65, 0x72, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x0e, 0x53, 0x68, 0x6f, 0x77, 0x41, 0x6c, 0x6c, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x12,
	0x14, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x63, 0x6d, 0x64, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3e, 0x0a, 0x0d, 0x57, 0x65, 0x62, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x57, 0x53, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x63, 0x6d, 0x64, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e,
	0x69, 0x6e, 0x6a, 0x61, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x2d, 0x67,
	0x6f, 0x2f, 0x70, 0x62, 0x73, 0x2f, 0x63, 0x6d, 0x64, 0x3b, 0x63, 0x6d, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmdService_proto_rawDescOnce sync.Once
	file_cmdService_proto_rawDescData = file_cmdService_proto_rawDesc
)

func file_cmdService_proto_rawDescGZIP() []byte {
	file_cmdService_proto_rawDescOnce.Do(func() {
		file_cmdService_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmdService_proto_rawDescData)
	})
	return file_cmdService_proto_rawDescData
}

var file_cmdService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cmdService_proto_goTypes = []interface{}{
	(*WSInfoReq)(nil),      // 0: pbs.cmd.WSInfoReq
	(*ThreadGroup)(nil),    // 1: pbs.cmd.ThreadGroup
	(*TopicMsg)(nil),       // 2: pbs.cmd.TopicMsg
	(*ShowPeer)(nil),       // 3: pbs.cmd.ShowPeer
	(*CommonResponse)(nil), // 4: pbs.cmd.CommonResponse
}
var file_cmdService_proto_depIdxs = []int32{
	2, // 0: pbs.cmd.CmdService.P2pSendTopicMsg:input_type -> pbs.cmd.TopicMsg
	3, // 1: pbs.cmd.CmdService.P2pShowPeers:input_type -> pbs.cmd.ShowPeer
	1, // 2: pbs.cmd.CmdService.ShowAllThreads:input_type -> pbs.cmd.ThreadGroup
	0, // 3: pbs.cmd.CmdService.WebSocketInfo:input_type -> pbs.cmd.WSInfoReq
	4, // 4: pbs.cmd.CmdService.P2pSendTopicMsg:output_type -> pbs.cmd.CommonResponse
	4, // 5: pbs.cmd.CmdService.P2pShowPeers:output_type -> pbs.cmd.CommonResponse
	4, // 6: pbs.cmd.CmdService.ShowAllThreads:output_type -> pbs.cmd.CommonResponse
	4, // 7: pbs.cmd.CmdService.WebSocketInfo:output_type -> pbs.cmd.CommonResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmdService_proto_init() }
func file_cmdService_proto_init() {
	if File_cmdService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmdService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WSInfoReq); i {
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
		file_cmdService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThreadGroup); i {
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
		file_cmdService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicMsg); i {
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
		file_cmdService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowPeer); i {
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
		file_cmdService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
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
			RawDescriptor: file_cmdService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmdService_proto_goTypes,
		DependencyIndexes: file_cmdService_proto_depIdxs,
		MessageInfos:      file_cmdService_proto_msgTypes,
	}.Build()
	File_cmdService_proto = out.File
	file_cmdService_proto_rawDesc = nil
	file_cmdService_proto_goTypes = nil
	file_cmdService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CmdServiceClient is the client API for CmdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CmdServiceClient interface {
	P2PSendTopicMsg(ctx context.Context, in *TopicMsg, opts ...grpc.CallOption) (*CommonResponse, error)
	P2PShowPeers(ctx context.Context, in *ShowPeer, opts ...grpc.CallOption) (*CommonResponse, error)
	ShowAllThreads(ctx context.Context, in *ThreadGroup, opts ...grpc.CallOption) (*CommonResponse, error)
	WebSocketInfo(ctx context.Context, in *WSInfoReq, opts ...grpc.CallOption) (*CommonResponse, error)
}

type cmdServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCmdServiceClient(cc grpc.ClientConnInterface) CmdServiceClient {
	return &cmdServiceClient{cc}
}

func (c *cmdServiceClient) P2PSendTopicMsg(ctx context.Context, in *TopicMsg, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pbs.cmd.CmdService/P2pSendTopicMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdServiceClient) P2PShowPeers(ctx context.Context, in *ShowPeer, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pbs.cmd.CmdService/P2pShowPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdServiceClient) ShowAllThreads(ctx context.Context, in *ThreadGroup, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pbs.cmd.CmdService/ShowAllThreads", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdServiceClient) WebSocketInfo(ctx context.Context, in *WSInfoReq, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pbs.cmd.CmdService/WebSocketInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CmdServiceServer is the server API for CmdService service.
type CmdServiceServer interface {
	P2PSendTopicMsg(context.Context, *TopicMsg) (*CommonResponse, error)
	P2PShowPeers(context.Context, *ShowPeer) (*CommonResponse, error)
	ShowAllThreads(context.Context, *ThreadGroup) (*CommonResponse, error)
	WebSocketInfo(context.Context, *WSInfoReq) (*CommonResponse, error)
}

// UnimplementedCmdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCmdServiceServer struct {
}

func (*UnimplementedCmdServiceServer) P2PSendTopicMsg(context.Context, *TopicMsg) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method P2PSendTopicMsg not implemented")
}
func (*UnimplementedCmdServiceServer) P2PShowPeers(context.Context, *ShowPeer) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method P2PShowPeers not implemented")
}
func (*UnimplementedCmdServiceServer) ShowAllThreads(context.Context, *ThreadGroup) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAllThreads not implemented")
}
func (*UnimplementedCmdServiceServer) WebSocketInfo(context.Context, *WSInfoReq) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WebSocketInfo not implemented")
}

func RegisterCmdServiceServer(s *grpc.Server, srv CmdServiceServer) {
	s.RegisterService(&_CmdService_serviceDesc, srv)
}

func _CmdService_P2PSendTopicMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServiceServer).P2PSendTopicMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbs.cmd.CmdService/P2PSendTopicMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServiceServer).P2PSendTopicMsg(ctx, req.(*TopicMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmdService_P2PShowPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowPeer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServiceServer).P2PShowPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbs.cmd.CmdService/P2PShowPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServiceServer).P2PShowPeers(ctx, req.(*ShowPeer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmdService_ShowAllThreads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServiceServer).ShowAllThreads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbs.cmd.CmdService/ShowAllThreads",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServiceServer).ShowAllThreads(ctx, req.(*ThreadGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmdService_WebSocketInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WSInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServiceServer).WebSocketInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbs.cmd.CmdService/WebSocketInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServiceServer).WebSocketInfo(ctx, req.(*WSInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CmdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbs.cmd.CmdService",
	HandlerType: (*CmdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "P2pSendTopicMsg",
			Handler:    _CmdService_P2PSendTopicMsg_Handler,
		},
		{
			MethodName: "P2pShowPeers",
			Handler:    _CmdService_P2PShowPeers_Handler,
		},
		{
			MethodName: "ShowAllThreads",
			Handler:    _CmdService_ShowAllThreads_Handler,
		},
		{
			MethodName: "WebSocketInfo",
			Handler:    _CmdService_WebSocketInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmdService.proto",
}
