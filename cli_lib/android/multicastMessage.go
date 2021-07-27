package androidlib

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/ninjahome/ninja-go/cli_lib/clientMsg/multicast"
	"github.com/ninjahome/ninja-go/cli_lib/clientMsg/unicast"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
)

type GroupInfo struct {
	GroupId   string   `json:"group_id"`
	GroupName string   `json:"group_name"`
	OwnerId   string   `json:"owner_id"`
	MemberId  []string `json:"member_id"`
	NickName  []string `json:"nick_name"`
}

type MulticastCallBack interface {
	CreateGroup(groupId, groupName, owner string, memberId, memberNickName []string) error
	JoinGroup(from, groupId, groupName, owner string, memberId, memberNickName []string, newId []string) error
	KickOutUser(from, groupId, kickId string) error
	QuitGroup(from, groupId, quitId string) error
	DismisGroup(groupId string) error
	SyncGroup(groupId string) string
	//same as CreateGroup
	SyncGroupAck(groupId, groupName, owner string, memberId, memberNickName []string) error
	VoiceMessage(from, groupId string, payload []byte, length int, time int64) error
	ImageMessage(from, groupId string, payload []byte, time int64) error
	LocationMessage(from, groupId string, l, a float32, name string, time int64) error
	TextMessage(from, groupId string, payload string, time int64) error
}

func (i AndroidAPP) multicastMsg(to []string, msg *pbs.WSCryptoGroupMsg) error {
	groupMessage := &multicast.GroupMessage{}
	if err := proto.Unmarshal(msg.PayLoad, groupMessage); err != nil {
		return err
	}

	switch groupMessage.GroupMsgTyp {
	case multicast.GroupMessageType_CreateGroupT:
		groupInfo := groupMessage.Payload.(*multicast.GroupMessage_GroupInfo)
		groupDesc := groupInfo.GroupInfo

		return i.multicast.CreateGroup(groupDesc.GroupId,
			groupDesc.GroupName,
			groupDesc.GroupOwner,
			to,
			groupDesc.NickName)

	case multicast.GroupMessageType_JoinGroupT:
		jInfo := groupMessage.Payload.(*multicast.GroupMessage_JoinGroupInfo)
		joinGroup := jInfo.JoinGroupInfo

		return i.multicast.JoinGroup(msg.From,
			joinGroup.GroupInfo.GroupId,
			joinGroup.GroupInfo.GroupName,
			joinGroup.GroupInfo.GroupOwner,
			to,
			joinGroup.GroupInfo.NickName,
			joinGroup.NewID,
		)

	case multicast.GroupMessageType_QuitGroupT:
		quitInfo := groupMessage.Payload.(*multicast.GroupMessage_QuitGroupInfo)
		quitGroup := quitInfo.QuitGroupInfo

		return i.multicast.QuitGroup(msg.From,quitGroup.GroupId, quitGroup.QuitId)

	case multicast.GroupMessageType_KickOutUserT:
		kickInfo := groupMessage.Payload.(*multicast.GroupMessage_QuitGroupInfo)
		kickGroup := kickInfo.QuitGroupInfo

		return i.multicast.KickOutUser(msg.From,kickGroup.GroupId, kickGroup.QuitId)

	case multicast.GroupMessageType_SyncGroupAckT:
		syncGroupAck := groupMessage.Payload.(*multicast.GroupMessage_SyncGroupAck)
		syncGroup := syncGroupAck.SyncGroupAck

		return i.multicast.SyncGroupAck(syncGroup.GroupInfo.GroupId,
			syncGroup.GroupInfo.GroupName,
			syncGroup.GroupInfo.GroupOwner,
			syncGroup.MemberId,
			syncGroup.GroupInfo.NickName)

	case multicast.GroupMessageType_ChatMessageT:
		chatMessage := groupMessage.Payload.(*multicast.GroupMessage_ChatMsg)
		chatMsg := chatMessage.ChatMsg

		return i.multicastChatMsg(msg.From, chatMsg, msg.UnixTime)

	case multicast.GroupMessageType_DismisGroupT:
		dismisGroup := groupMessage.Payload.(*multicast.GroupMessage_GroupId)

		return i.multicast.DismisGroup(dismisGroup.GroupId)

	}

	return nil
}

func (i AndroidAPP) multicastChatMsg(from string, msg *multicast.ChatMesageDesc, ts int64) error {
	switch msg.ChatMsg.Payload.(type) {

	case *unicast.ChatMessage_PlainTxt:

		rawData := msg.ChatMsg.Payload.(*unicast.ChatMessage_PlainTxt)

		return i.multicast.TextMessage(from,
			msg.GroupId,
			rawData.PlainTxt,
			ts)

	case *unicast.ChatMessage_Image:

		rawData := msg.ChatMsg.Payload.(*unicast.ChatMessage_Image)

		return i.multicast.ImageMessage(from,
			msg.GroupId,
			rawData.Image,
			ts)

	case *unicast.ChatMessage_Voice:

		voiceMessage := msg.ChatMsg.Payload.(*unicast.ChatMessage_Voice).Voice

		return i.multicast.VoiceMessage(from,
			msg.GroupId,
			voiceMessage.Data,
			int(voiceMessage.Length),
			ts)

	case *unicast.ChatMessage_Location:

		locationMessage := msg.ChatMsg.Payload.(*unicast.ChatMessage_Location).Location

		return i.multicast.LocationMessage(from,
			msg.GroupId,
			locationMessage.Longitude,
			locationMessage.Latitude,
			locationMessage.Name,
			ts)
	default:
		return errors.New("msg not recognize")
	}

}

func CreateGroup(to, nickname []string, groupId, groupName string) error {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")

	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}

	owner := _inst.websocket.Address()

	rawData, err := multicast.WrapCreateGroup(nickname, owner, groupId, groupName)
	if err != nil {
		return err
	}

	err = _inst.websocket.GWrite(to, rawData)
	if err != nil {
		return err
	}

	return nil
}

func JoinGroup(to, nickName []string, groupId, groupName, groupOwner, newId []string) error {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")

	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	rawData, err := multicast.WrapJoinGroup(nickName, groupOwner, groupId, groupName, newId)
	if err != nil {
		return err
	}

	err = _inst.websocket.GWrite(to, rawData)
	if err != nil {
		return err
	}

	return nil
}

func QuitGroup(to []string, groupId string) error {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")

	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	quitId := _inst.websocket.Address()

	rawData, err := multicast.WrapQuitGroup(quitId, groupId)
	if err != nil {
		return err
	}

	err = _inst.websocket.GWrite(to, rawData)
	if err != nil {
		return err
	}

	return nil
}

func KickOutUser(to []string, groupId, owner, kickUserId string) error {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")

	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}

	localUser := _inst.websocket.Address()
	if localUser != owner {
		return fmt.Errorf("only owner can kick out group member")
	}

	rawData, err := multicast.WrapKickUser(kickUserId, groupId)
	if err != nil {
		return err
	}

	err = _inst.websocket.GWrite(to, rawData)
	if err != nil {
		return err
	}

	return nil
}

func DismisGroup(to []string, owner, groupId string) error {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")

	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}

	localUser := _inst.websocket.Address()
	if localUser != owner {
		return fmt.Errorf("only owner can dismis group")
	}

	rawData, err := multicast.WrapDismisGroup(groupId)
	if err != nil {
		return err
	}

	err = _inst.websocket.GWrite(to, rawData)
	if err != nil {
		return err
	}

	return nil
}

func WriteGroupMessage(to []string, groupId, plainTxt string) error {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")
	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}

	rawData, err := multicast.WrapTextMesage(groupId, plainTxt)

	if err != nil {
		return err
	}

	err = _inst.websocket.GWrite(to, rawData)
	if err != nil {
		return err
	}

	return nil
}

func WriteLocationGroupMessage(to []string, longitude, latitude float32, name, groupId string) error {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")
	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}

	rawData, err := multicast.WrapLocation(longitude, latitude, name, groupId)

	if err != nil {
		return err
	}

	err = _inst.websocket.GWrite(to, rawData)
	if err != nil {
		return err
	}

	return nil
}

func WriteImageGroupMessage(to []string, payload []byte, groupId string) error {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")
	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}

	rawData, err := multicast.WrapImage(payload, groupId)

	if err != nil {
		return err
	}

	err = _inst.websocket.GWrite(to, rawData)
	if err != nil {
		return err
	}

	return nil
}

func WriteVoiceGroupMessage(to []string, payload []byte, length int, groupId string) error {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")
	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}

	rawData, err := multicast.WrapVoice(payload, length, groupId)

	if err != nil {
		return err
	}

	err = _inst.websocket.GWrite(to, rawData)
	if err != nil {
		return err
	}

	return nil
}

func NewGroupId() string {
	buf := make([]byte, 32)
	for {
		if n, err := rand.Read(buf); err != nil || n != len(buf) {
			continue
		}
		break
	}

	return base64.StdEncoding.EncodeToString(buf)
}


//type GroupInfo struct {
//	GroupId   string   `json:"group_id"`
//	GroupName string   `json:"group_name"`
//	OwnerId   string   `json:"owner_id"`
//	MemberId  []string `json:"member_id"`
//	NickName  []string `json:"nick_name"`
//}


func GroupInfo2Str(groupId, groupName, owner string, memberIds, nickNames []string) string {
	gi:=&GroupInfo{
		GroupId: groupId,
		GroupName: groupName,
		OwnerId: owner,
		MemberId: memberIds,
		NickName: nickNames,
	}

	j,_:=json.Marshal(*gi)

	return string(j)
}