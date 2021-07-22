package androidlib

import "fmt"



type MulticastCallBack interface {
	CreateGroup(groupId, groupName, owner string, memberId []string, memberNickName []string) error
	JoinGroup(from,groupId, groupName, owner string, memberId []string,memberNickName []string,newId string) error
	QuitGroup(groupId, quitId string) error
	DismisGroup(groupId string) error
	SyncGroup(groupId string)
}


func CreateGroup(to []string, groupName string) string  {

}

func JoinGroup(to []string, groupId string, groupName string) error  {

}

func QuitGroup(to []string, groupId string) error  {

}

func DismisGroup(to []string, groupId string) error  {

}

func SyncGroupKey(to []string, groupId string) error  {

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
