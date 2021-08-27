// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.3
// source: stream.proto

package stream

import (
	contact "github.com/ninjahome/ninja-go/pbs/contact"
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

type StreamMType int32

const (
	StreamMType_MTOnlineAck   StreamMType = 0
	StreamMType_MTOnlineSync  StreamMType = 1
	StreamMType_MTContactSync StreamMType = 2
	StreamMType_MTContactAck  StreamMType = 3
	StreamMType_MTDevInfoSync StreamMType = 4
	StreamMType_MTDevInfoAck  StreamMType = 5
)

// Enum value maps for StreamMType.
var (
	StreamMType_name = map[int32]string{
		0: "MTOnlineAck",
		1: "MTOnlineSync",
		2: "MTContactSync",
		3: "MTContactAck",
		4: "MTDevInfoSync",
		5: "MTDevInfoAck",
	}
	StreamMType_value = map[string]int32{
		"MTOnlineAck":   0,
		"MTOnlineSync":  1,
		"MTContactSync": 2,
		"MTContactAck":  3,
		"MTDevInfoSync": 4,
		"MTDevInfoAck":  5,
	}
)

func (x StreamMType) Enum() *StreamMType {
	p := new(StreamMType)
	*p = x
	return p
}

func (x StreamMType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamMType) Descriptor() protoreflect.EnumDescriptor {
	return file_stream_proto_enumTypes[0].Descriptor()
}

func (StreamMType) Type() protoreflect.EnumType {
	return &file_stream_proto_enumTypes[0]
}

func (x StreamMType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamMType.Descriptor instead.
func (StreamMType) EnumDescriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{0}
}

type StreamMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MTyp StreamMType `protobuf:"varint,1,opt,name=MTyp,proto3,enum=pbs.stream.StreamMType" json:"MTyp,omitempty"`
	// Types that are assignable to Payload:
	//	*StreamMsg_OnlineSync
	//	*StreamMsg_OnlineAck
	//	*StreamMsg_ContactSync
	//	*StreamMsg_ContactAck
	//	*StreamMsg_DiSync
	//	*StreamMsg_DiAck
	Payload isStreamMsg_Payload `protobuf_oneof:"payload"`
}

func (x *StreamMsg) Reset() {
	*x = StreamMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMsg) ProtoMessage() {}

func (x *StreamMsg) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMsg.ProtoReflect.Descriptor instead.
func (*StreamMsg) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{0}
}

func (x *StreamMsg) GetMTyp() StreamMType {
	if x != nil {
		return x.MTyp
	}
	return StreamMType_MTOnlineAck
}

func (m *StreamMsg) GetPayload() isStreamMsg_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *StreamMsg) GetOnlineSync() *OnlineSync {
	if x, ok := x.GetPayload().(*StreamMsg_OnlineSync); ok {
		return x.OnlineSync
	}
	return nil
}

func (x *StreamMsg) GetOnlineAck() *OnlineMap {
	if x, ok := x.GetPayload().(*StreamMsg_OnlineAck); ok {
		return x.OnlineAck
	}
	return nil
}

func (x *StreamMsg) GetContactSync() *ContactSync {
	if x, ok := x.GetPayload().(*StreamMsg_ContactSync); ok {
		return x.ContactSync
	}
	return nil
}

func (x *StreamMsg) GetContactAck() *ContactAck {
	if x, ok := x.GetPayload().(*StreamMsg_ContactAck); ok {
		return x.ContactAck
	}
	return nil
}

func (x *StreamMsg) GetDiSync() *DevInfoSync {
	if x, ok := x.GetPayload().(*StreamMsg_DiSync); ok {
		return x.DiSync
	}
	return nil
}

func (x *StreamMsg) GetDiAck() *DevInfoAck {
	if x, ok := x.GetPayload().(*StreamMsg_DiAck); ok {
		return x.DiAck
	}
	return nil
}

type isStreamMsg_Payload interface {
	isStreamMsg_Payload()
}

type StreamMsg_OnlineSync struct {
	OnlineSync *OnlineSync `protobuf:"bytes,2,opt,name=onlineSync,proto3,oneof"`
}

type StreamMsg_OnlineAck struct {
	OnlineAck *OnlineMap `protobuf:"bytes,5,opt,name=onlineAck,proto3,oneof"`
}

type StreamMsg_ContactSync struct {
	ContactSync *ContactSync `protobuf:"bytes,3,opt,name=contactSync,proto3,oneof"`
}

type StreamMsg_ContactAck struct {
	ContactAck *ContactAck `protobuf:"bytes,4,opt,name=contactAck,proto3,oneof"`
}

type StreamMsg_DiSync struct {
	DiSync *DevInfoSync `protobuf:"bytes,6,opt,name=diSync,proto3,oneof"`
}

type StreamMsg_DiAck struct {
	DiAck *DevInfoAck `protobuf:"bytes,7,opt,name=diAck,proto3,oneof"`
}

func (*StreamMsg_OnlineSync) isStreamMsg_Payload() {}

func (*StreamMsg_OnlineAck) isStreamMsg_Payload() {}

func (*StreamMsg_ContactSync) isStreamMsg_Payload() {}

func (*StreamMsg_ContactAck) isStreamMsg_Payload() {}

func (*StreamMsg_DiSync) isStreamMsg_Payload() {}

func (*StreamMsg_DiAck) isStreamMsg_Payload() {}

type OnlineSync struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeID string `protobuf:"bytes,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
}

func (x *OnlineSync) Reset() {
	*x = OnlineSync{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineSync) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineSync) ProtoMessage() {}

func (x *OnlineSync) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineSync.ProtoReflect.Descriptor instead.
func (*OnlineSync) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{1}
}

func (x *OnlineSync) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

type DevInfoSync struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId string `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
}

func (x *DevInfoSync) Reset() {
	*x = DevInfoSync{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevInfoSync) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevInfoSync) ProtoMessage() {}

func (x *DevInfoSync) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevInfoSync.ProtoReflect.Descriptor instead.
func (*DevInfoSync) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{2}
}

func (x *DevInfoSync) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

type DevInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	DevToken string `protobuf:"bytes,2,opt,name=devToken,proto3" json:"devToken,omitempty"`
	DevTyp   int32  `protobuf:"varint,3,opt,name=devTyp,proto3" json:"devTyp,omitempty"`
}

func (x *DevInfo) Reset() {
	*x = DevInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevInfo) ProtoMessage() {}

func (x *DevInfo) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevInfo.ProtoReflect.Descriptor instead.
func (*DevInfo) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{3}
}

func (x *DevInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *DevInfo) GetDevToken() string {
	if x != nil {
		return x.DevToken
	}
	return ""
}

func (x *DevInfo) GetDevTyp() int32 {
	if x != nil {
		return x.DevTyp
	}
	return 0
}

type DevInfoAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dis []*DevInfo `protobuf:"bytes,1,rep,name=dis,proto3" json:"dis,omitempty"`
}

func (x *DevInfoAck) Reset() {
	*x = DevInfoAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevInfoAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevInfoAck) ProtoMessage() {}

func (x *DevInfoAck) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevInfoAck.ProtoReflect.Descriptor instead.
func (*DevInfoAck) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{4}
}

func (x *DevInfoAck) GetDis() []*DevInfo {
	if x != nil {
		return x.Dis
	}
	return nil
}

type OnlineMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UID []string `protobuf:"bytes,1,rep,name=UID,proto3" json:"UID,omitempty"`
}

func (x *OnlineMap) Reset() {
	*x = OnlineMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineMap) ProtoMessage() {}

func (x *OnlineMap) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineMap.ProtoReflect.Descriptor instead.
func (*OnlineMap) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{5}
}

func (x *OnlineMap) GetUID() []string {
	if x != nil {
		return x.UID
	}
	return nil
}

type ContactSync struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UID    string `protobuf:"bytes,1,opt,name=UID,proto3" json:"UID,omitempty"`
	SeqVer int64  `protobuf:"varint,2,opt,name=SeqVer,proto3" json:"SeqVer,omitempty"`
}

func (x *ContactSync) Reset() {
	*x = ContactSync{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactSync) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactSync) ProtoMessage() {}

func (x *ContactSync) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactSync.ProtoReflect.Descriptor instead.
func (*ContactSync) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{6}
}

func (x *ContactSync) GetUID() string {
	if x != nil {
		return x.UID
	}
	return ""
}

func (x *ContactSync) GetSeqVer() int64 {
	if x != nil {
		return x.SeqVer
	}
	return 0
}

type ContactAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contacts []*contact.ContactItem `protobuf:"bytes,1,rep,name=contacts,proto3" json:"contacts,omitempty"`
}

func (x *ContactAck) Reset() {
	*x = ContactAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactAck) ProtoMessage() {}

func (x *ContactAck) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactAck.ProtoReflect.Descriptor instead.
func (*ContactAck) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{7}
}

func (x *ContactAck) GetContacts() []*contact.ContactItem {
	if x != nil {
		return x.Contacts
	}
	return nil
}

var File_stream_proto protoreflect.FileDescriptor

var file_stream_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x70, 0x62, 0x73, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x03, 0x0a, 0x09, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x4d, 0x73, 0x67, 0x12, 0x2b, 0x0a, 0x04, 0x4d, 0x54, 0x79, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x4d, 0x54, 0x79, 0x70, 0x12, 0x38, 0x0a, 0x0a, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x79,
	0x6e, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x79, 0x6e, 0x63,
	0x48, 0x00, 0x52, 0x0a, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x35,
	0x0a, 0x09, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x41, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x48, 0x00, 0x52, 0x09, 0x6f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x41, 0x63, 0x6b, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x53, 0x79, 0x6e, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x73,
	0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53,
	0x79, 0x6e, 0x63, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53, 0x79,
	0x6e, 0x63, 0x12, 0x38, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x41, 0x63, 0x6b,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x41, 0x63, 0x6b, 0x48, 0x00,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x41, 0x63, 0x6b, 0x12, 0x31, 0x0a, 0x06,
	0x64, 0x69, 0x53, 0x79, 0x6e, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70,
	0x62, 0x73, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x76, 0x49, 0x6e, 0x66,
	0x6f, 0x53, 0x79, 0x6e, 0x63, 0x48, 0x00, 0x52, 0x06, 0x64, 0x69, 0x53, 0x79, 0x6e, 0x63, 0x12,
	0x2e, 0x0a, 0x05, 0x64, 0x69, 0x41, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x70, 0x62, 0x73, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x76, 0x49,
	0x6e, 0x66, 0x6f, 0x41, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x05, 0x64, 0x69, 0x41, 0x63, 0x6b, 0x42,
	0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x24, 0x0a, 0x0a, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44,
	0x22, 0x25, 0x0a, 0x0b, 0x44, 0x65, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x79, 0x6e, 0x63, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x07, 0x44, 0x65, 0x76, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x54, 0x79, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x64, 0x65, 0x76, 0x54, 0x79, 0x70, 0x22, 0x33, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x49,
	0x6e, 0x66, 0x6f, 0x41, 0x63, 0x6b, 0x12, 0x25, 0x0a, 0x03, 0x64, 0x69, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x44, 0x65, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x64, 0x69, 0x73, 0x22, 0x1d, 0x0a,
	0x09, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x55, 0x49, 0x44, 0x22, 0x37, 0x0a, 0x0b,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x55,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x65, 0x71, 0x56, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53,
	0x65, 0x71, 0x56, 0x65, 0x72, 0x22, 0x42, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x41, 0x63, 0x6b, 0x12, 0x34, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x2a, 0x7a, 0x0a, 0x0b, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x4d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x54, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x41, 0x63, 0x6b, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x54, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4d,
	0x54, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x10, 0x02, 0x12, 0x10,
	0x0a, 0x0c, 0x4d, 0x54, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x41, 0x63, 0x6b, 0x10, 0x03,
	0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x54, 0x44, 0x65, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x79, 0x6e,
	0x63, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x54, 0x44, 0x65, 0x76, 0x49, 0x6e, 0x66, 0x6f,
	0x41, 0x63, 0x6b, 0x10, 0x05, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x6e, 0x69,
	0x6e, 0x6a, 0x61, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x3b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stream_proto_rawDescOnce sync.Once
	file_stream_proto_rawDescData = file_stream_proto_rawDesc
)

func file_stream_proto_rawDescGZIP() []byte {
	file_stream_proto_rawDescOnce.Do(func() {
		file_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_stream_proto_rawDescData)
	})
	return file_stream_proto_rawDescData
}

var file_stream_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_stream_proto_goTypes = []interface{}{
	(StreamMType)(0),            // 0: pbs.stream.StreamMType
	(*StreamMsg)(nil),           // 1: pbs.stream.StreamMsg
	(*OnlineSync)(nil),          // 2: pbs.stream.OnlineSync
	(*DevInfoSync)(nil),         // 3: pbs.stream.DevInfoSync
	(*DevInfo)(nil),             // 4: pbs.stream.DevInfo
	(*DevInfoAck)(nil),          // 5: pbs.stream.DevInfoAck
	(*OnlineMap)(nil),           // 6: pbs.stream.OnlineMap
	(*ContactSync)(nil),         // 7: pbs.stream.ContactSync
	(*ContactAck)(nil),          // 8: pbs.stream.ContactAck
	(*contact.ContactItem)(nil), // 9: pbs.contact.ContactItem
}
var file_stream_proto_depIdxs = []int32{
	0, // 0: pbs.stream.StreamMsg.MTyp:type_name -> pbs.stream.StreamMType
	2, // 1: pbs.stream.StreamMsg.onlineSync:type_name -> pbs.stream.OnlineSync
	6, // 2: pbs.stream.StreamMsg.onlineAck:type_name -> pbs.stream.OnlineMap
	7, // 3: pbs.stream.StreamMsg.contactSync:type_name -> pbs.stream.ContactSync
	8, // 4: pbs.stream.StreamMsg.contactAck:type_name -> pbs.stream.ContactAck
	3, // 5: pbs.stream.StreamMsg.diSync:type_name -> pbs.stream.DevInfoSync
	5, // 6: pbs.stream.StreamMsg.diAck:type_name -> pbs.stream.DevInfoAck
	4, // 7: pbs.stream.DevInfoAck.dis:type_name -> pbs.stream.DevInfo
	9, // 8: pbs.stream.ContactAck.contacts:type_name -> pbs.contact.ContactItem
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_stream_proto_init() }
func file_stream_proto_init() {
	if File_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamMsg); i {
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
		file_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineSync); i {
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
		file_stream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevInfoSync); i {
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
		file_stream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevInfo); i {
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
		file_stream_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevInfoAck); i {
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
		file_stream_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineMap); i {
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
		file_stream_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactSync); i {
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
		file_stream_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactAck); i {
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
	file_stream_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StreamMsg_OnlineSync)(nil),
		(*StreamMsg_OnlineAck)(nil),
		(*StreamMsg_ContactSync)(nil),
		(*StreamMsg_ContactAck)(nil),
		(*StreamMsg_DiSync)(nil),
		(*StreamMsg_DiAck)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stream_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stream_proto_goTypes,
		DependencyIndexes: file_stream_proto_depIdxs,
		EnumInfos:         file_stream_proto_enumTypes,
		MessageInfos:      file_stream_proto_msgTypes,
	}.Build()
	File_stream_proto = out.File
	file_stream_proto_rawDesc = nil
	file_stream_proto_goTypes = nil
	file_stream_proto_depIdxs = nil
}
