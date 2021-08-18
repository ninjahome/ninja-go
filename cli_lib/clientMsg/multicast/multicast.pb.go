// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.3
// source: multicast.proto

package multicast

import (
	unicast "github.com/ninjahome/ninja-go/cli_lib/clientMsg/unicast"
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

type GroupMessageType int32

const (
	GroupMessageType_CreateGroupT  GroupMessageType = 0
	GroupMessageType_QuitGroupT    GroupMessageType = 1
	GroupMessageType_KickOutUserT  GroupMessageType = 2
	GroupMessageType_JoinGroupT    GroupMessageType = 3
	GroupMessageType_DismisGroupT  GroupMessageType = 4
	GroupMessageType_SyncGroupAckT GroupMessageType = 5
	GroupMessageType_BanTalkingT   GroupMessageType = 6
	GroupMessageType_ChatMessageT  GroupMessageType = 7
)

// Enum value maps for GroupMessageType.
var (
	GroupMessageType_name = map[int32]string{
		0: "CreateGroupT",
		1: "QuitGroupT",
		2: "KickOutUserT",
		3: "JoinGroupT",
		4: "DismisGroupT",
		5: "SyncGroupAckT",
		6: "BanTalkingT",
		7: "ChatMessageT",
	}
	GroupMessageType_value = map[string]int32{
		"CreateGroupT":  0,
		"QuitGroupT":    1,
		"KickOutUserT":  2,
		"JoinGroupT":    3,
		"DismisGroupT":  4,
		"SyncGroupAckT": 5,
		"BanTalkingT":   6,
		"ChatMessageT":  7,
	}
)

func (x GroupMessageType) Enum() *GroupMessageType {
	p := new(GroupMessageType)
	*p = x
	return p
}

func (x GroupMessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GroupMessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_multicast_proto_enumTypes[0].Descriptor()
}

func (GroupMessageType) Type() protoreflect.EnumType {
	return &file_multicast_proto_enumTypes[0]
}

func (x GroupMessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GroupMessageType.Descriptor instead.
func (GroupMessageType) EnumDescriptor() ([]byte, []int) {
	return file_multicast_proto_rawDescGZIP(), []int{0}
}

type GroupDesc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupName  string   `protobuf:"bytes,1,opt,name=GroupName,proto3" json:"GroupName,omitempty"`
	GroupId    string   `protobuf:"bytes,2,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	GroupOwner string   `protobuf:"bytes,3,opt,name=GroupOwner,proto3" json:"GroupOwner,omitempty"`
	NickName   []string `protobuf:"bytes,4,rep,name=NickName,proto3" json:"NickName,omitempty"`
}

func (x *GroupDesc) Reset() {
	*x = GroupDesc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multicast_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupDesc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupDesc) ProtoMessage() {}

func (x *GroupDesc) ProtoReflect() protoreflect.Message {
	mi := &file_multicast_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupDesc.ProtoReflect.Descriptor instead.
func (*GroupDesc) Descriptor() ([]byte, []int) {
	return file_multicast_proto_rawDescGZIP(), []int{0}
}

func (x *GroupDesc) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *GroupDesc) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *GroupDesc) GetGroupOwner() string {
	if x != nil {
		return x.GroupOwner
	}
	return ""
}

func (x *GroupDesc) GetNickName() []string {
	if x != nil {
		return x.NickName
	}
	return nil
}

type JoinGroupDesc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupInfo  *GroupDesc `protobuf:"bytes,1,opt,name=groupInfo,proto3" json:"groupInfo,omitempty"`
	BanTalking bool       `protobuf:"varint,2,opt,name=banTalking,proto3" json:"banTalking,omitempty"`
	NewID      []string   `protobuf:"bytes,3,rep,name=NewID,proto3" json:"NewID,omitempty"`
}

func (x *JoinGroupDesc) Reset() {
	*x = JoinGroupDesc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multicast_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinGroupDesc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGroupDesc) ProtoMessage() {}

func (x *JoinGroupDesc) ProtoReflect() protoreflect.Message {
	mi := &file_multicast_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGroupDesc.ProtoReflect.Descriptor instead.
func (*JoinGroupDesc) Descriptor() ([]byte, []int) {
	return file_multicast_proto_rawDescGZIP(), []int{1}
}

func (x *JoinGroupDesc) GetGroupInfo() *GroupDesc {
	if x != nil {
		return x.GroupInfo
	}
	return nil
}

func (x *JoinGroupDesc) GetBanTalking() bool {
	if x != nil {
		return x.BanTalking
	}
	return false
}

func (x *JoinGroupDesc) GetNewID() []string {
	if x != nil {
		return x.NewID
	}
	return nil
}

type QuitGroupDesc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuitId  string `protobuf:"bytes,1,opt,name=QuitId,proto3" json:"QuitId,omitempty"`
	GroupId string `protobuf:"bytes,2,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
}

func (x *QuitGroupDesc) Reset() {
	*x = QuitGroupDesc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multicast_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuitGroupDesc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuitGroupDesc) ProtoMessage() {}

func (x *QuitGroupDesc) ProtoReflect() protoreflect.Message {
	mi := &file_multicast_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuitGroupDesc.ProtoReflect.Descriptor instead.
func (*QuitGroupDesc) Descriptor() ([]byte, []int) {
	return file_multicast_proto_rawDescGZIP(), []int{2}
}

func (x *QuitGroupDesc) GetQuitId() string {
	if x != nil {
		return x.QuitId
	}
	return ""
}

func (x *QuitGroupDesc) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type KickUserDesc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId    string   `protobuf:"bytes,1,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	KickUserId []string `protobuf:"bytes,2,rep,name=KickUserId,proto3" json:"KickUserId,omitempty"`
}

func (x *KickUserDesc) Reset() {
	*x = KickUserDesc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multicast_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickUserDesc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickUserDesc) ProtoMessage() {}

func (x *KickUserDesc) ProtoReflect() protoreflect.Message {
	mi := &file_multicast_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickUserDesc.ProtoReflect.Descriptor instead.
func (*KickUserDesc) Descriptor() ([]byte, []int) {
	return file_multicast_proto_rawDescGZIP(), []int{3}
}

func (x *KickUserDesc) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *KickUserDesc) GetKickUserId() []string {
	if x != nil {
		return x.KickUserId
	}
	return nil
}

type SyncGroupAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupInfo  *GroupDesc `protobuf:"bytes,1,opt,name=groupInfo,proto3" json:"groupInfo,omitempty"`
	BanTalking bool       `protobuf:"varint,2,opt,name=BanTalking,proto3" json:"BanTalking,omitempty"`
	MemberId   []string   `protobuf:"bytes,3,rep,name=MemberId,proto3" json:"MemberId,omitempty"`
}

func (x *SyncGroupAck) Reset() {
	*x = SyncGroupAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multicast_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncGroupAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncGroupAck) ProtoMessage() {}

func (x *SyncGroupAck) ProtoReflect() protoreflect.Message {
	mi := &file_multicast_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncGroupAck.ProtoReflect.Descriptor instead.
func (*SyncGroupAck) Descriptor() ([]byte, []int) {
	return file_multicast_proto_rawDescGZIP(), []int{4}
}

func (x *SyncGroupAck) GetGroupInfo() *GroupDesc {
	if x != nil {
		return x.GroupInfo
	}
	return nil
}

func (x *SyncGroupAck) GetBanTalking() bool {
	if x != nil {
		return x.BanTalking
	}
	return false
}

func (x *SyncGroupAck) GetMemberId() []string {
	if x != nil {
		return x.MemberId
	}
	return nil
}

type ChatMesageDesc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string               `protobuf:"bytes,1,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	ChatMsg *unicast.ChatMessage `protobuf:"bytes,2,opt,name=chatMsg,proto3" json:"chatMsg,omitempty"`
}

func (x *ChatMesageDesc) Reset() {
	*x = ChatMesageDesc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multicast_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMesageDesc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMesageDesc) ProtoMessage() {}

func (x *ChatMesageDesc) ProtoReflect() protoreflect.Message {
	mi := &file_multicast_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMesageDesc.ProtoReflect.Descriptor instead.
func (*ChatMesageDesc) Descriptor() ([]byte, []int) {
	return file_multicast_proto_rawDescGZIP(), []int{5}
}

func (x *ChatMesageDesc) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *ChatMesageDesc) GetChatMsg() *unicast.ChatMessage {
	if x != nil {
		return x.ChatMsg
	}
	return nil
}

type BanTalkingMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	Banned  bool   `protobuf:"varint,2,opt,name=Banned,proto3" json:"Banned,omitempty"`
}

func (x *BanTalkingMsg) Reset() {
	*x = BanTalkingMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multicast_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BanTalkingMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BanTalkingMsg) ProtoMessage() {}

func (x *BanTalkingMsg) ProtoReflect() protoreflect.Message {
	mi := &file_multicast_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BanTalkingMsg.ProtoReflect.Descriptor instead.
func (*BanTalkingMsg) Descriptor() ([]byte, []int) {
	return file_multicast_proto_rawDescGZIP(), []int{6}
}

func (x *BanTalkingMsg) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *BanTalkingMsg) GetBanned() bool {
	if x != nil {
		return x.Banned
	}
	return false
}

type GroupMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupMsgTyp GroupMessageType `protobuf:"varint,1,opt,name=GroupMsgTyp,proto3,enum=GroupMessageType" json:"GroupMsgTyp,omitempty"`
	// Types that are assignable to Payload:
	//	*GroupMessage_GroupInfo
	//	*GroupMessage_JoinGroupInfo
	//	*GroupMessage_GroupId
	//	*GroupMessage_QuitGroupInfo
	//	*GroupMessage_SyncGroupAck
	//	*GroupMessage_ChatMsg
	//	*GroupMessage_KickId
	//	*GroupMessage_BanTalking
	Payload isGroupMessage_Payload `protobuf_oneof:"payload"`
}

func (x *GroupMessage) Reset() {
	*x = GroupMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multicast_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMessage) ProtoMessage() {}

func (x *GroupMessage) ProtoReflect() protoreflect.Message {
	mi := &file_multicast_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMessage.ProtoReflect.Descriptor instead.
func (*GroupMessage) Descriptor() ([]byte, []int) {
	return file_multicast_proto_rawDescGZIP(), []int{7}
}

func (x *GroupMessage) GetGroupMsgTyp() GroupMessageType {
	if x != nil {
		return x.GroupMsgTyp
	}
	return GroupMessageType_CreateGroupT
}

func (m *GroupMessage) GetPayload() isGroupMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *GroupMessage) GetGroupInfo() *GroupDesc {
	if x, ok := x.GetPayload().(*GroupMessage_GroupInfo); ok {
		return x.GroupInfo
	}
	return nil
}

func (x *GroupMessage) GetJoinGroupInfo() *JoinGroupDesc {
	if x, ok := x.GetPayload().(*GroupMessage_JoinGroupInfo); ok {
		return x.JoinGroupInfo
	}
	return nil
}

func (x *GroupMessage) GetGroupId() string {
	if x, ok := x.GetPayload().(*GroupMessage_GroupId); ok {
		return x.GroupId
	}
	return ""
}

func (x *GroupMessage) GetQuitGroupInfo() *QuitGroupDesc {
	if x, ok := x.GetPayload().(*GroupMessage_QuitGroupInfo); ok {
		return x.QuitGroupInfo
	}
	return nil
}

func (x *GroupMessage) GetSyncGroupAck() *SyncGroupAck {
	if x, ok := x.GetPayload().(*GroupMessage_SyncGroupAck); ok {
		return x.SyncGroupAck
	}
	return nil
}

func (x *GroupMessage) GetChatMsg() *ChatMesageDesc {
	if x, ok := x.GetPayload().(*GroupMessage_ChatMsg); ok {
		return x.ChatMsg
	}
	return nil
}

func (x *GroupMessage) GetKickId() *KickUserDesc {
	if x, ok := x.GetPayload().(*GroupMessage_KickId); ok {
		return x.KickId
	}
	return nil
}

func (x *GroupMessage) GetBanTalking() *BanTalkingMsg {
	if x, ok := x.GetPayload().(*GroupMessage_BanTalking); ok {
		return x.BanTalking
	}
	return nil
}

type isGroupMessage_Payload interface {
	isGroupMessage_Payload()
}

type GroupMessage_GroupInfo struct {
	GroupInfo *GroupDesc `protobuf:"bytes,2,opt,name=groupInfo,proto3,oneof"`
}

type GroupMessage_JoinGroupInfo struct {
	JoinGroupInfo *JoinGroupDesc `protobuf:"bytes,3,opt,name=joinGroupInfo,proto3,oneof"`
}

type GroupMessage_GroupId struct {
	GroupId string `protobuf:"bytes,4,opt,name=GroupId,proto3,oneof"`
}

type GroupMessage_QuitGroupInfo struct {
	QuitGroupInfo *QuitGroupDesc `protobuf:"bytes,5,opt,name=quitGroupInfo,proto3,oneof"`
}

type GroupMessage_SyncGroupAck struct {
	SyncGroupAck *SyncGroupAck `protobuf:"bytes,6,opt,name=syncGroupAck,proto3,oneof"`
}

type GroupMessage_ChatMsg struct {
	ChatMsg *ChatMesageDesc `protobuf:"bytes,7,opt,name=chatMsg,proto3,oneof"`
}

type GroupMessage_KickId struct {
	KickId *KickUserDesc `protobuf:"bytes,8,opt,name=kickId,proto3,oneof"`
}

type GroupMessage_BanTalking struct {
	BanTalking *BanTalkingMsg `protobuf:"bytes,9,opt,name=banTalking,proto3,oneof"`
}

func (*GroupMessage_GroupInfo) isGroupMessage_Payload() {}

func (*GroupMessage_JoinGroupInfo) isGroupMessage_Payload() {}

func (*GroupMessage_GroupId) isGroupMessage_Payload() {}

func (*GroupMessage_QuitGroupInfo) isGroupMessage_Payload() {}

func (*GroupMessage_SyncGroupAck) isGroupMessage_Payload() {}

func (*GroupMessage_ChatMsg) isGroupMessage_Payload() {}

func (*GroupMessage_KickId) isGroupMessage_Payload() {}

func (*GroupMessage_BanTalking) isGroupMessage_Payload() {}

var File_multicast_proto protoreflect.FileDescriptor

var file_multicast_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69,
	0x6e, 0x6a, 0x61, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x2d, 0x67, 0x6f,
	0x2f, 0x63, 0x6c, 0x69, 0x5f, 0x6c, 0x69, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d,
	0x73, 0x67, 0x2f, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x6f, 0x0a, 0x0d, 0x4a, 0x6f, 0x69,
	0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x73, 0x63, 0x12, 0x28, 0x0a, 0x09, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x73, 0x63, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x61, 0x6e, 0x54, 0x61, 0x6c, 0x6b, 0x69,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x62, 0x61, 0x6e, 0x54, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x4e, 0x65, 0x77, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x4e, 0x65, 0x77, 0x49, 0x44, 0x22, 0x41, 0x0a, 0x0d, 0x51, 0x75,
	0x69, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x73, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x51,
	0x75, 0x69, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x51, 0x75, 0x69,
	0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x48, 0x0a,
	0x0c, 0x4b, 0x69, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x12, 0x18, 0x0a,
	0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x4b, 0x69, 0x63, 0x6b, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x4b, 0x69, 0x63,
	0x6b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x74, 0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x6b, 0x12, 0x28, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x44, 0x65, 0x73, 0x63, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x61, 0x6e, 0x54, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x42, 0x61, 0x6e, 0x54, 0x61, 0x6c, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22, 0x52, 0x0a,
	0x0e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x73, 0x63, 0x12,
	0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x73,
	0x67, 0x22, 0x41, 0x0a, 0x0d, 0x42, 0x61, 0x6e, 0x54, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x4d,
	0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x42, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x42, 0x61,
	0x6e, 0x6e, 0x65, 0x64, 0x22, 0xc3, 0x03, 0x0a, 0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x73,
	0x67, 0x54, 0x79, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x12, 0x2a, 0x0a, 0x09, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x73, 0x63, 0x48, 0x00, 0x52, 0x09, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x0d, 0x6a, 0x6f, 0x69, 0x6e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x73, 0x63, 0x48, 0x00, 0x52,
	0x0d, 0x6a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a,
	0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x0d, 0x71, 0x75,
	0x69, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x51, 0x75, 0x69, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x73,
	0x63, 0x48, 0x00, 0x52, 0x0d, 0x71, 0x75, 0x69, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x33, 0x0a, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41,
	0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x6b, 0x12, 0x2b, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x4d,
	0x73, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d,
	0x65, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x73, 0x63, 0x48, 0x00, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x74, 0x4d, 0x73, 0x67, 0x12, 0x27, 0x0a, 0x06, 0x6b, 0x69, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x73, 0x63, 0x48, 0x00, 0x52, 0x06, 0x6b, 0x69, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x30, 0x0a,
	0x0a, 0x62, 0x61, 0x6e, 0x54, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x42, 0x61, 0x6e, 0x54, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x4d, 0x73,
	0x67, 0x48, 0x00, 0x52, 0x0a, 0x62, 0x61, 0x6e, 0x54, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x42,
	0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2a, 0x9e, 0x01, 0x0a, 0x10, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x10, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x51, 0x75, 0x69, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x10,
	0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x54, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x6d, 0x69, 0x73, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x54, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x79, 0x6e, 0x63, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x63, 0x6b, 0x54, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x61, 0x6e, 0x54,
	0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x68, 0x61,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x10, 0x07, 0x42, 0x45, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x68,
	0x6f, 0x6d, 0x65, 0x2f, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x69,
	0x5f, 0x6c, 0x69, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x2f, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x3b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61,
	0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_multicast_proto_rawDescOnce sync.Once
	file_multicast_proto_rawDescData = file_multicast_proto_rawDesc
)

func file_multicast_proto_rawDescGZIP() []byte {
	file_multicast_proto_rawDescOnce.Do(func() {
		file_multicast_proto_rawDescData = protoimpl.X.CompressGZIP(file_multicast_proto_rawDescData)
	})
	return file_multicast_proto_rawDescData
}

var file_multicast_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_multicast_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_multicast_proto_goTypes = []interface{}{
	(GroupMessageType)(0),       // 0: GroupMessageType
	(*GroupDesc)(nil),           // 1: GroupDesc
	(*JoinGroupDesc)(nil),       // 2: JoinGroupDesc
	(*QuitGroupDesc)(nil),       // 3: QuitGroupDesc
	(*KickUserDesc)(nil),        // 4: KickUserDesc
	(*SyncGroupAck)(nil),        // 5: SyncGroupAck
	(*ChatMesageDesc)(nil),      // 6: ChatMesageDesc
	(*BanTalkingMsg)(nil),       // 7: BanTalkingMsg
	(*GroupMessage)(nil),        // 8: GroupMessage
	(*unicast.ChatMessage)(nil), // 9: ChatMessage
}
var file_multicast_proto_depIdxs = []int32{
	1,  // 0: JoinGroupDesc.groupInfo:type_name -> GroupDesc
	1,  // 1: SyncGroupAck.groupInfo:type_name -> GroupDesc
	9,  // 2: ChatMesageDesc.chatMsg:type_name -> ChatMessage
	0,  // 3: GroupMessage.GroupMsgTyp:type_name -> GroupMessageType
	1,  // 4: GroupMessage.groupInfo:type_name -> GroupDesc
	2,  // 5: GroupMessage.joinGroupInfo:type_name -> JoinGroupDesc
	3,  // 6: GroupMessage.quitGroupInfo:type_name -> QuitGroupDesc
	5,  // 7: GroupMessage.syncGroupAck:type_name -> SyncGroupAck
	6,  // 8: GroupMessage.chatMsg:type_name -> ChatMesageDesc
	4,  // 9: GroupMessage.kickId:type_name -> KickUserDesc
	7,  // 10: GroupMessage.banTalking:type_name -> BanTalkingMsg
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_multicast_proto_init() }
func file_multicast_proto_init() {
	if File_multicast_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_multicast_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupDesc); i {
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
		file_multicast_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinGroupDesc); i {
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
		file_multicast_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuitGroupDesc); i {
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
		file_multicast_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickUserDesc); i {
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
		file_multicast_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncGroupAck); i {
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
		file_multicast_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMesageDesc); i {
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
		file_multicast_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BanTalkingMsg); i {
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
		file_multicast_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupMessage); i {
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
	file_multicast_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*GroupMessage_GroupInfo)(nil),
		(*GroupMessage_JoinGroupInfo)(nil),
		(*GroupMessage_GroupId)(nil),
		(*GroupMessage_QuitGroupInfo)(nil),
		(*GroupMessage_SyncGroupAck)(nil),
		(*GroupMessage_ChatMsg)(nil),
		(*GroupMessage_KickId)(nil),
		(*GroupMessage_BanTalking)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_multicast_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_multicast_proto_goTypes,
		DependencyIndexes: file_multicast_proto_depIdxs,
		EnumInfos:         file_multicast_proto_enumTypes,
		MessageInfos:      file_multicast_proto_msgTypes,
	}.Build()
	File_multicast_proto = out.File
	file_multicast_proto_rawDesc = nil
	file_multicast_proto_goTypes = nil
	file_multicast_proto_depIdxs = nil
}
