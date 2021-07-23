package androidlib

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/ninjahome/ninja-go/cli_lib/clientMsg/multicast"
)

type GroupMember struct {
	MemberId string `json:"member_id"`
	NickName string `json:"nick_name"`
}

type GroupInfo struct {
	GroupId string `json:"group_id"`
	GroupName string `json:"group_name"`
	OwnerId   string `json:"owner_id"`
	Members   []*GroupMember `json:"members"`
}

type MulticastCallBack interface {
	CreateGroup(groupId, groupName, owner string, memberId, memberNickName []string) error
	JoinGroup(from,groupId, groupName, owner string, memberId,memberNickName []string, newId string) error
	KickOutUser(groupId, kickId string)
	QuitGroup(groupId, quitId string) error
	DismisGroup(groupId string) error
	SyncGroup(groupId string) string
}


func CreateGroup(to,nickname []string, groupId,groupName string) error  {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")

	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}

	owner:=_inst.websocket.Address()

	rawData,err:=multicast.WrapCreateGroup(nickname,owner,groupId,groupName)
	if err!=nil{
		return err
	}

	err = _inst.websocket.GWrite(to,rawData)
	if err!=nil{
		return err
	}

	return nil
}

func JoinGroup(to,nickName []string, groupId, groupName,groupOwner, newId string) error  {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")

	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			fmt.Println( err.Error())
			return err
		}
	}

	rawData,err:=multicast.WrapJoinGroup(nickName,groupOwner,groupId,groupName,newId)
	if err!=nil{
		return err
	}

	err = _inst.websocket.GWrite(to,rawData)
	if err!=nil{
		return err
	}

	return nil
}

func QuitGroup(to []string, groupId string) error  {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")

	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			fmt.Println( err.Error())
			return err
		}
	}

	quitId:=_inst.websocket.Address()

	rawData,err:=multicast.WrapQuitGroup(quitId, groupId)
	if err!=nil{
		return err
	}

	err = _inst.websocket.GWrite(to,rawData)
	if err!=nil{
		return err
	}

	return nil
}

func KickOutUser(to []string, groupId, owner,kickUserId string) error {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")

	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}

	localUser:=_inst.websocket.Address()
	if localUser != owner{
		return fmt.Errorf("only owner can kick out group member")
	}

	rawData,err:=multicast.WrapKickUser(kickUserId, groupId)
	if err!=nil{
		return err
	}

	err = _inst.websocket.GWrite(to,rawData)
	if err!=nil{
		return err
	}

	return nil
}

func DismisGroup(to []string, owner, groupId string) error  {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")

	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}

	localUser:=_inst.websocket.Address()
	if localUser != owner{
		return fmt.Errorf("only owner can dismis group")
	}

	rawData,err:=multicast.WrapDismisGroup(groupId)
	if err!=nil{
		return err
	}

	err = _inst.websocket.GWrite(to,rawData)
	if err!=nil{
		return err
	}

	return nil
}




func WriteGroupMessage(to []string, plainTxt string) error  {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")
	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}




}


func WriteLocationGroupMessage(to []string, longitude, latitude float32, name string) error  {

}


func WriteImageGroupMessage(to string, payload []byte) error{

}

func WriteVoiceGroupMessage(to string, payload []byte, len int) error{

}

func NewGroupId() string  {
	buf:=make([]byte,32)
	for{
		if n,err:=rand.Read(buf);err!=nil || n!=len(buf){
			continue
		}
		break
	}

	return base64.StdEncoding.EncodeToString(buf)
}
