// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: service.proto

package service

import (
	proto "github.com/golang/protobuf/proto"
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

type SrvMsgType int32

const (
	SrvMsgType_Unknown   SrvMsgType = 0
	SrvMsgType_Online    SrvMsgType = 80
	SrvMsgType_ACK       SrvMsgType = 81
	SrvMsgType_CryptoMsg SrvMsgType = 82
	SrvMsgType_OnlineACK SrvMsgType = 83
)

// Enum value maps for SrvMsgType.
var (
	SrvMsgType_name = map[int32]string{
		0:  "Unknown",
		80: "Online",
		81: "ACK",
		82: "CryptoMsg",
		83: "OnlineACK",
	}
	SrvMsgType_value = map[string]int32{
		"Unknown":   0,
		"Online":    80,
		"ACK":       81,
		"CryptoMsg": 82,
		"OnlineACK": 83,
	}
)

func (x SrvMsgType) Enum() *SrvMsgType {
	p := new(SrvMsgType)
	*p = x
	return p
}

func (x SrvMsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SrvMsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_service_proto_enumTypes[0].Descriptor()
}

func (SrvMsgType) Type() protoreflect.EnumType {
	return &file_service_proto_enumTypes[0]
}

func (x SrvMsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SrvMsgType.Descriptor instead.
func (SrvMsgType) EnumDescriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

type P2PMsgType int32

const (
	P2PMsgType_P2pUnknown   P2PMsgType = 0
	P2PMsgType_P2pOnline    P2PMsgType = 1
	P2PMsgType_P2pCryptoMsg P2PMsgType = 2
)

// Enum value maps for P2PMsgType.
var (
	P2PMsgType_name = map[int32]string{
		0: "P2pUnknown",
		1: "P2pOnline",
		2: "P2pCryptoMsg",
	}
	P2PMsgType_value = map[string]int32{
		"P2pUnknown":   0,
		"P2pOnline":    1,
		"P2pCryptoMsg": 2,
	}
)

func (x P2PMsgType) Enum() *P2PMsgType {
	p := new(P2PMsgType)
	*p = x
	return p
}

func (x P2PMsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (P2PMsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_service_proto_enumTypes[1].Descriptor()
}

func (P2PMsgType) Type() protoreflect.EnumType {
	return &file_service_proto_enumTypes[1]
}

func (x P2PMsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use P2PMsgType.Descriptor instead.
func (P2PMsgType) EnumDescriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

type WSOnlineAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool  `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Seq     int64 `protobuf:"varint,2,opt,name=Seq,proto3" json:"Seq,omitempty"`
}

func (x *WSOnlineAck) Reset() {
	*x = WSOnlineAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WSOnlineAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WSOnlineAck) ProtoMessage() {}

func (x *WSOnlineAck) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WSOnlineAck.ProtoReflect.Descriptor instead.
func (*WSOnlineAck) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *WSOnlineAck) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *WSOnlineAck) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

type WSAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Seq  int64  `protobuf:"varint,3,opt,name=Seq,proto3" json:"Seq,omitempty"`
	MSG  string `protobuf:"bytes,2,opt,name=MSG,proto3" json:"MSG,omitempty"`
}

func (x *WSAck) Reset() {
	*x = WSAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WSAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WSAck) ProtoMessage() {}

func (x *WSAck) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WSAck.ProtoReflect.Descriptor instead.
func (*WSAck) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *WSAck) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *WSAck) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *WSAck) GetMSG() string {
	if x != nil {
		return x.MSG
	}
	return ""
}

type WSOnline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UID      string `protobuf:"bytes,1,opt,name=UID,proto3" json:"UID,omitempty"`
	UnixTime int64  `protobuf:"varint,3,opt,name=UnixTime,proto3" json:"UnixTime,omitempty"`
}

func (x *WSOnline) Reset() {
	*x = WSOnline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WSOnline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WSOnline) ProtoMessage() {}

func (x *WSOnline) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WSOnline.ProtoReflect.Descriptor instead.
func (*WSOnline) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *WSOnline) GetUID() string {
	if x != nil {
		return x.UID
	}
	return ""
}

func (x *WSOnline) GetUnixTime() int64 {
	if x != nil {
		return x.UnixTime
	}
	return 0
}

type WSCryptoMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From     string `protobuf:"bytes,1,opt,name=From,proto3" json:"From,omitempty"`
	To       string `protobuf:"bytes,3,opt,name=To,proto3" json:"To,omitempty"`
	PayLoad  []byte `protobuf:"bytes,5,opt,name=PayLoad,proto3" json:"PayLoad,omitempty"`
	UnixTime int64  `protobuf:"varint,6,opt,name=UnixTime,proto3" json:"UnixTime,omitempty"`
}

func (x *WSCryptoMsg) Reset() {
	*x = WSCryptoMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WSCryptoMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WSCryptoMsg) ProtoMessage() {}

func (x *WSCryptoMsg) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WSCryptoMsg.ProtoReflect.Descriptor instead.
func (*WSCryptoMsg) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *WSCryptoMsg) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *WSCryptoMsg) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *WSCryptoMsg) GetPayLoad() []byte {
	if x != nil {
		return x.PayLoad
	}
	return nil
}

func (x *WSCryptoMsg) GetUnixTime() int64 {
	if x != nil {
		return x.UnixTime
	}
	return 0
}

type CliOnlineMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash []byte `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Sig  []byte `protobuf:"bytes,2,opt,name=Sig,proto3" json:"Sig,omitempty"`
	// Types that are assignable to Payload:
	//	*CliOnlineMsg_Online
	//	*CliOnlineMsg_OlAck
	Payload isCliOnlineMsg_Payload `protobuf_oneof:"Payload"`
}

func (x *CliOnlineMsg) Reset() {
	*x = CliOnlineMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CliOnlineMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CliOnlineMsg) ProtoMessage() {}

func (x *CliOnlineMsg) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CliOnlineMsg.ProtoReflect.Descriptor instead.
func (*CliOnlineMsg) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *CliOnlineMsg) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *CliOnlineMsg) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

func (m *CliOnlineMsg) GetPayload() isCliOnlineMsg_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *CliOnlineMsg) GetOnline() *WSOnline {
	if x, ok := x.GetPayload().(*CliOnlineMsg_Online); ok {
		return x.Online
	}
	return nil
}

func (x *CliOnlineMsg) GetOlAck() *WSOnlineAck {
	if x, ok := x.GetPayload().(*CliOnlineMsg_OlAck); ok {
		return x.OlAck
	}
	return nil
}

type isCliOnlineMsg_Payload interface {
	isCliOnlineMsg_Payload()
}

type CliOnlineMsg_Online struct {
	Online *WSOnline `protobuf:"bytes,3,opt,name=online,proto3,oneof"`
}

type CliOnlineMsg_OlAck struct {
	OlAck *WSOnlineAck `protobuf:"bytes,4,opt,name=olAck,proto3,oneof"`
}

func (*CliOnlineMsg_Online) isCliOnlineMsg_Payload() {}

func (*CliOnlineMsg_OlAck) isCliOnlineMsg_Payload() {}

type P2PMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgTyp P2PMsgType `protobuf:"varint,1,opt,name=msgTyp,proto3,enum=pbs.P2PMsgType" json:"msgTyp,omitempty"`
	// Types that are assignable to Payload:
	//	*P2PMsg_Online
	//	*P2PMsg_Msg
	Payload isP2PMsg_Payload `protobuf_oneof:"payload"`
}

func (x *P2PMsg) Reset() {
	*x = P2PMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *P2PMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*P2PMsg) ProtoMessage() {}

func (x *P2PMsg) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use P2PMsg.ProtoReflect.Descriptor instead.
func (*P2PMsg) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *P2PMsg) GetMsgTyp() P2PMsgType {
	if x != nil {
		return x.MsgTyp
	}
	return P2PMsgType_P2pUnknown
}

func (m *P2PMsg) GetPayload() isP2PMsg_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *P2PMsg) GetOnline() *WSOnline {
	if x, ok := x.GetPayload().(*P2PMsg_Online); ok {
		return x.Online
	}
	return nil
}

func (x *P2PMsg) GetMsg() *WSCryptoMsg {
	if x, ok := x.GetPayload().(*P2PMsg_Msg); ok {
		return x.Msg
	}
	return nil
}

type isP2PMsg_Payload interface {
	isP2PMsg_Payload()
}

type P2PMsg_Online struct {
	Online *WSOnline `protobuf:"bytes,2,opt,name=online,proto3,oneof"`
}

type P2PMsg_Msg struct {
	Msg *WSCryptoMsg `protobuf:"bytes,3,opt,name=msg,proto3,oneof"`
}

func (*P2PMsg_Online) isP2PMsg_Payload() {}

func (*P2PMsg_Msg) isP2PMsg_Payload() {}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x70, 0x62, 0x73, 0x22, 0x39, 0x0a, 0x0b, 0x57, 0x53, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x41, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x53, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x53, 0x65, 0x71, 0x22,
	0x3f, 0x0a, 0x05, 0x57, 0x53, 0x41, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x53, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x53, 0x65, 0x71, 0x12, 0x10,
	0x0a, 0x03, 0x4d, 0x53, 0x47, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x53, 0x47,
	0x22, 0x38, 0x0a, 0x08, 0x57, 0x53, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x55, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x55, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x67, 0x0a, 0x0b, 0x57, 0x53,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x54, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x18, 0x0a,
	0x07, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x50, 0x61, 0x79, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x6e, 0x69, 0x78, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x55, 0x6e, 0x69, 0x78, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x4f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x69, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x53, 0x69, 0x67, 0x12, 0x27, 0x0a, 0x06, 0x6f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x73,
	0x2e, 0x57, 0x53, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x6f, 0x6c, 0x41, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x57, 0x53, 0x4f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x41, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x05, 0x6f, 0x6c, 0x41, 0x63, 0x6b, 0x42, 0x09, 0x0a,
	0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x06, 0x50, 0x32, 0x70,
	0x4d, 0x73, 0x67, 0x12, 0x27, 0x0a, 0x06, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x50, 0x32, 0x70, 0x4d, 0x73, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x12, 0x27, 0x0a, 0x06,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70,
	0x62, 0x73, 0x2e, 0x57, 0x53, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x48, 0x00, 0x52, 0x06, 0x6f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x24, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x57, 0x53, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x4d, 0x73, 0x67, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x09, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2a, 0x4c, 0x0a, 0x0a, 0x53, 0x72, 0x76, 0x4d, 0x73, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x50, 0x12, 0x07, 0x0a,
	0x03, 0x41, 0x43, 0x4b, 0x10, 0x51, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x4d, 0x73, 0x67, 0x10, 0x52, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x41,
	0x43, 0x4b, 0x10, 0x53, 0x2a, 0x3d, 0x0a, 0x0a, 0x50, 0x32, 0x70, 0x4d, 0x73, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x32, 0x70, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x32, 0x70, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x10,
	0x01, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x32, 0x70, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4d, 0x73,
	0x67, 0x10, 0x02, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_proto_goTypes = []interface{}{
	(SrvMsgType)(0),      // 0: pbs.SrvMsgType
	(P2PMsgType)(0),      // 1: pbs.P2pMsgType
	(*WSOnlineAck)(nil),  // 2: pbs.WSOnlineAck
	(*WSAck)(nil),        // 3: pbs.WSAck
	(*WSOnline)(nil),     // 4: pbs.WSOnline
	(*WSCryptoMsg)(nil),  // 5: pbs.WSCryptoMsg
	(*CliOnlineMsg)(nil), // 6: pbs.CliOnlineMsg
	(*P2PMsg)(nil),       // 7: pbs.P2pMsg
}
var file_service_proto_depIdxs = []int32{
	4, // 0: pbs.CliOnlineMsg.online:type_name -> pbs.WSOnline
	2, // 1: pbs.CliOnlineMsg.olAck:type_name -> pbs.WSOnlineAck
	1, // 2: pbs.P2pMsg.msgTyp:type_name -> pbs.P2pMsgType
	4, // 3: pbs.P2pMsg.online:type_name -> pbs.WSOnline
	5, // 4: pbs.P2pMsg.msg:type_name -> pbs.WSCryptoMsg
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WSOnlineAck); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WSAck); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WSOnline); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WSCryptoMsg); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CliOnlineMsg); i {
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
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*P2PMsg); i {
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
	file_service_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*CliOnlineMsg_Online)(nil),
		(*CliOnlineMsg_OlAck)(nil),
	}
	file_service_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*P2PMsg_Online)(nil),
		(*P2PMsg_Msg)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		EnumInfos:         file_service_proto_enumTypes,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
