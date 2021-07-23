package multicast

import (
	"google.golang.org/protobuf/proto"
)

func WrapCreateGroup(nickName []string,owner,groupId, groupName string) ([]byte,error)  {
	createGroup:=&GroupDesc{
		GroupName: groupName,
		GroupOwner: owner,
		NickName: nickName,
		GroupId: groupId,
	}


	gMsg:=&GroupMessage{
		GroupMsgTyp: GroupMessageType_CreateGroupT,
		Payload: &GroupMessage_GroupInfo{
			GroupInfo: createGroup,
		},
	}

	rawData,err:=proto.Marshal(gMsg)
	if err!=nil{
		return nil, err
	}

	return rawData, nil
}



func WrapJoinGroup(nickName []string,owner,groupId, groupName,newId string) ([]byte,error)  {
	groupDesc:=&GroupDesc{
		GroupName: groupName,
		GroupOwner: owner,
		NickName: nickName,
		GroupId: groupId,
	}


	joinGroup:=&JoinGroupDesc{
		GroupInfo: groupDesc,
		NewID: newId,
	}

	gMsg:=&GroupMessage{
		GroupMsgTyp: GroupMessageType_JoinGroupT,
		Payload: &GroupMessage_JoinGroupInfo{
			JoinGroupInfo: joinGroup,
		},
	}

	rawData,err:=proto.Marshal(gMsg)
	if err!=nil{
		return nil, err
	}

	return rawData,nil
}

func WrapQuitGroup(quitId, groupId string) ([]byte, error)  {
	quitGroup:=&QuitGroupDesc{
		GroupId: groupId,
		QuitId: quitId,
	}

	gMsg:=&GroupMessage{
		GroupMsgTyp: GroupMessageType_QuitGroupT,
		Payload: &GroupMessage_QuitGroupInfo{
			QuitGroupInfo: quitGroup,
		},
	}

	rawData,err:=proto.Marshal(gMsg)
	if err!=nil{
		return nil, err
	}

	return rawData,nil

}


func WrapKickUser(kickId, groupId string) ([]byte, error)  {
	quitGroup:=&QuitGroupDesc{
		GroupId: groupId,
		QuitId: kickId,
	}

	gMsg:=&GroupMessage{
		GroupMsgTyp: GroupMessageType_KickOutUserT,
		Payload: &GroupMessage_QuitGroupInfo{
			QuitGroupInfo: quitGroup,
		},
	}

	rawData,err:=proto.Marshal(gMsg)
	if err!=nil{
		return nil, err
	}

	return rawData,nil

}


func WrapDismisGroup(groupId string) ([]byte, error)  {
	gMsg:=&GroupMessage{
		GroupMsgTyp: GroupMessageType_DismisGroupT,
		Payload: &GroupMessage_GroupId{
			GroupId: groupId,
		},
	}

	rawData,err:=proto.Marshal(gMsg)
	if err!=nil{
		return nil, err
	}

	return rawData,nil

}
