package multicast

import (
	"github.com/ninjahome/ninja-go/cli_lib/clientMsg/unicast"
	"google.golang.org/protobuf/proto"
)

func WrapCreateGroup(nickName []string, owner, groupId, groupName string) ([]byte, error) {
	createGroup := &GroupDesc{
		GroupName:  groupName,
		GroupOwner: owner,
		NickName:   nickName,
		GroupId:    groupId,
	}

	gMsg := &GroupMessage{
		GroupMsgTyp: GroupMessageType_CreateGroupT,
		Payload: &GroupMessage_GroupInfo{
			GroupInfo: createGroup,
		},
	}

	rawData, err := proto.Marshal(gMsg)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}

func WrapJoinGroup(nickName []string, owner, groupId, groupName string, banTalking bool, newId []string) ([]byte, error) {
	groupDesc := &GroupDesc{
		GroupName:  groupName,
		GroupOwner: owner,
		NickName:   nickName,
		GroupId:    groupId,
	}

	joinGroup := &JoinGroupDesc{
		GroupInfo:  groupDesc,
		NewID:      newId,
		BanTalking: banTalking,
	}

	gMsg := &GroupMessage{
		GroupMsgTyp: GroupMessageType_JoinGroupT,
		Payload: &GroupMessage_JoinGroupInfo{
			JoinGroupInfo: joinGroup,
		},
	}

	rawData, err := proto.Marshal(gMsg)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}

func WrapQuitGroup(quitId, groupId string) ([]byte, error) {
	quitGroup := &QuitGroupDesc{
		GroupId: groupId,
		QuitId:  quitId,
	}

	gMsg := &GroupMessage{
		GroupMsgTyp: GroupMessageType_QuitGroupT,
		Payload: &GroupMessage_QuitGroupInfo{
			QuitGroupInfo: quitGroup,
		},
	}

	rawData, err := proto.Marshal(gMsg)
	if err != nil {
		return nil, err
	}

	return rawData, nil

}

func WrapKickUser(kickId, groupId string) ([]byte, error) {
	quitGroup := &QuitGroupDesc{
		GroupId: groupId,
		QuitId:  kickId,
	}

	gMsg := &GroupMessage{
		GroupMsgTyp: GroupMessageType_KickOutUserT,
		Payload: &GroupMessage_QuitGroupInfo{
			QuitGroupInfo: quitGroup,
		},
	}

	rawData, err := proto.Marshal(gMsg)
	if err != nil {
		return nil, err
	}

	return rawData, nil

}

func WrapDismisGroup(groupId string) ([]byte, error) {
	gMsg := &GroupMessage{
		GroupMsgTyp: GroupMessageType_DismisGroupT,
		Payload: &GroupMessage_GroupId{
			GroupId: groupId,
		},
	}

	rawData, err := proto.Marshal(gMsg)
	if err != nil {
		return nil, err
	}

	return rawData, nil

}

func WrapBanTalking(groupId string) ([]byte, error) {
	gMsg := &GroupMessage{
		GroupMsgTyp: GroupMessageType_BanTalkingT,
		Payload: &GroupMessage_GroupId{
			GroupId: groupId,
		},
	}

	rawData, err := proto.Marshal(gMsg)
	if err != nil {
		return nil, err
	}

	return rawData, nil

}

func WrapSyncGroupAck(nickName, memberId []string, owner, groupId, groupName string, banTalking bool) ([]byte, error) {

	groupDesc := &GroupDesc{
		GroupName:  groupName,
		GroupOwner: owner,
		NickName:   nickName,
		GroupId:    groupId,
	}

	syncGroup := &SyncGroupAck{
		GroupInfo:  groupDesc,
		BanTalking: banTalking,
		MemberId:   memberId,
	}

	gMsg := &GroupMessage{
		GroupMsgTyp: GroupMessageType_SyncGroupAckT,
		Payload: &GroupMessage_SyncGroupAck{
			SyncGroupAck: syncGroup,
		},
	}

	rawData, err := proto.Marshal(gMsg)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}

func WrapTextMesage(groupid, plainTxt string) ([]byte, error) {
	chatMessage := &unicast.ChatMessage{
		Payload: &unicast.ChatMessage_PlainTxt{
			PlainTxt: plainTxt,
		},
	}

	chatInfo := &ChatMesageDesc{
		GroupId: groupid,
		ChatMsg: chatMessage,
	}

	gMsg := &GroupMessage{
		GroupMsgTyp: GroupMessageType_ChatMessageT,
		Payload: &GroupMessage_ChatMsg{
			ChatMsg: chatInfo,
		},
	}

	rawData, err := proto.Marshal(gMsg)
	if err != nil {
		return nil, err
	}

	return rawData, nil

}

func WrapLocation(l, a float32, name, groupId string) ([]byte, error) {
	chatMessage := &unicast.ChatMessage{
		Payload: &unicast.ChatMessage_Location{
			Location: &unicast.Location{
				Latitude:  a,
				Longitude: l,
				Name:      name,
			},
		},
	}

	chatInfo := &ChatMesageDesc{
		GroupId: groupId,
		ChatMsg: chatMessage,
	}
	gMsg := &GroupMessage{
		GroupMsgTyp: GroupMessageType_ChatMessageT,
		Payload: &GroupMessage_ChatMsg{
			ChatMsg: chatInfo,
		},
	}

	rawData, err := proto.Marshal(gMsg)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}

func WrapImage(data []byte, groupId string) ([]byte, error) {
	chatMessage := &unicast.ChatMessage{
		Payload: &unicast.ChatMessage_Image{
			Image: data,
		},
	}

	chatInfo := &ChatMesageDesc{
		GroupId: groupId,
		ChatMsg: chatMessage,
	}
	gMsg := &GroupMessage{
		GroupMsgTyp: GroupMessageType_ChatMessageT,
		Payload: &GroupMessage_ChatMsg{
			ChatMsg: chatInfo,
		},
	}

	rawData, err := proto.Marshal(gMsg)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}

func WrapVoice(p []byte, l int, groupId string) ([]byte, error) {
	chatMessage := &unicast.ChatMessage{
		Payload: &unicast.ChatMessage_Voice{
			Voice: &unicast.Voice{
				Data:   p,
				Length: int32(l),
			},
		},
	}

	chatInfo := &ChatMesageDesc{
		GroupId: groupId,
		ChatMsg: chatMessage,
	}
	gMsg := &GroupMessage{
		GroupMsgTyp: GroupMessageType_ChatMessageT,
		Payload: &GroupMessage_ChatMsg{
			ChatMsg: chatInfo,
		},
	}

	rawData, err := proto.Marshal(gMsg)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}

func WrapFile(p []byte, size int, name, groupId string) ([]byte, error) {
	chatMessage := &unicast.ChatMessage{
		Payload: &unicast.ChatMessage_File{
			File: &unicast.File{
				Size: int32(size),
				Data: p,
				Name: name,
			},
		},
	}

	chatInfo := &ChatMesageDesc{
		GroupId: groupId,
		ChatMsg: chatMessage,
	}

	gMsg := &GroupMessage{
		GroupMsgTyp: GroupMessageType_ChatMessageT,
		Payload: &GroupMessage_ChatMsg{
			ChatMsg: chatInfo,
		},
	}

	rawData, err := proto.Marshal(gMsg)
	if err != nil {
		return nil, err
	}
	return rawData, nil
}
