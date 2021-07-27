package iosLib

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/ninjahome/ninja-go/cli_lib/clientMsg/multicast"
	"github.com/ninjahome/ninja-go/cli_lib/clientMsg/unicast"
	"github.com/ninjahome/ninja-go/cli_lib/utils"
	"github.com/ninjahome/ninja-go/common"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"github.com/ninjahome/ninja-go/service/client"
	"github.com/ninjahome/ninja-go/service/websocket"
	"github.com/ninjahome/ninja-go/wallet"
	"github.com/polydawn/refmt/json"
	"google.golang.org/protobuf/proto"
)

func NewWallet(auth string) string {
	key := wallet.NewKey()
	_inst.key = key
	return key.StoreString(auth)
}

func ActiveAddress() string {
	if _inst.key == nil {
		return ""
	}

	return _inst.key.Address.String()
}

type IosAPP struct {
	key       *wallet.Key
	unicast   UnicastCallBack
	multicast MulticastCallBack
	wsEnd     string
	websocket *client.WSClient
	unreadSeq int64
}

func (a IosAPP) OnlineSuccess() {
	if err := _inst.websocket.PullUnreadMsg(a.unreadSeq); err != nil {
		fmt.Println(err.Error())
	}
}

func (a IosAPP) ImmediateMessage(msg *pbs.WSCryptoMsg) error {
	if msg == nil {
		return errors.New("msg is nil")
	}
	return a.unicastMsg(msg)
}

func (a IosAPP) ImmediateGMessage(msg *pbs.WSCryptoGroupMsg) error {
	if msg == nil {
		return errors.New("msg is nil")
	}

	var to []string
	for i := 0; i < len(msg.To); i++ {
		to = append(to, msg.To[i].MemberId)
	}

	return a.multicastMsg(to, msg)
}

func (a IosAPP) WebSocketClosed() {
	a.unicast.WebSocketClosed()
}

//func (i AndroidAPP) UnreadMsg(ack *pbs.WSUnreadAck) error {
//	payload := ack.Payload
//
//	for j := 0; j < len(payload); j++ {
//		if err := i.unicastMsg(payload[j]); err != nil {
//			//TODO::notify app to know the failure
//			continue
//		}
//	}
//	return nil
//}

func (i IosAPP) unicastMsg(msg *pbs.WSCryptoMsg) error {

	chatMessage := &unicast.ChatMessage{}
	if err := proto.Unmarshal(msg.PayLoad, chatMessage); err != nil {
		return err
	}
	switch chatMessage.Payload.(type) {

	case *unicast.ChatMessage_PlainTxt:

		rawData := chatMessage.Payload.(*unicast.ChatMessage_PlainTxt)

		return i.unicast.TextMessage(msg.From,
			msg.To,
			rawData.PlainTxt,
			msg.UnixTime)

	case *unicast.ChatMessage_Image:

		rawData := chatMessage.Payload.(*unicast.ChatMessage_Image)

		return i.unicast.ImageMessage(msg.From,
			msg.To,
			rawData.Image,
			msg.UnixTime)

	case *unicast.ChatMessage_Voice:

		voiceMessage := chatMessage.Payload.(*unicast.ChatMessage_Voice).Voice

		return i.unicast.VoiceMessage(msg.From,
			msg.To,
			voiceMessage.Data,
			int(voiceMessage.Length),
			msg.UnixTime)

	case *unicast.ChatMessage_Location:

		locationMessage := chatMessage.Payload.(*unicast.ChatMessage_Location).Location

		return i.unicast.LocationMessage(msg.From,
			msg.To,
			locationMessage.Longitude,
			locationMessage.Latitude,
			locationMessage.Name,
			msg.UnixTime)
	case *unicast.ChatMessage_SyncGroupId:
		grouId := chatMessage.Payload.(*unicast.ChatMessage_SyncGroupId).SyncGroupId

		groupInfo := i.multicast.SyncGroup(grouId)

		if groupInfo == "" {
			fmt.Println("no group in cell phone")
			return nil
		}

		gi := &GroupInfo{}
		if err := json.Unmarshal([]byte(groupInfo), gi); err != nil {
			fmt.Println(err)
			return nil
		}

		rawData, err := multicast.WrapSyncGroupAck(gi.NickName, gi.MemberId, gi.OwnerId, gi.GroupId, gi.GroupName)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if err := i.websocket.Write(msg.From, rawData); err != nil {
			fmt.Println(err)
			return nil
		}
		return nil
	default:
		return errors.New("msg not recognize")
	}

}

func UnmarshalGoByte(s string) []byte {
	b, e := base64.StdEncoding.DecodeString(s)
	if e != nil {
		return nil
	}
	return b
}

var _inst = &IosAPP{unreadSeq: 0}

type UnicastCallBack interface {
	VoiceMessage(from, to string, payload []byte, length int, time int64) error
	ImageMessage(from, to string, payload []byte, time int64) error
	LocationMessage(from, to string, l, a float32, name string, time int64) error
	TextMessage(from, to string, payload string, time int64) error
	WebSocketClosed()
}

func ConfigApp(addr string, unicast UnicastCallBack, multicast MulticastCallBack) {

	if addr == "" {
		addr = client.RandomBootNode()
	}
	//fmt.Println("======>", addr)
	_inst.wsEnd = addr
	_inst.unicast = unicast
	_inst.multicast = multicast
}

func ActiveWallet(cipherTxt, auth string, devtoken string) error {

	key, err := wallet.LoadKeyFromJsonStr(cipherTxt, auth)
	if err != nil {
		return err
	}
	_inst.key = key
	ws, err := client.NewWSClient(devtoken, _inst.wsEnd, websocket.DevTypeAndroid, key, _inst) //202.182.101.145//167.179.78.33//127.0.0.1//
	if err != nil {
		return err
	}
	_inst.websocket = ws

	return ws.Online()
}

func WalletIsOpen() bool {
	return _inst.key != nil && _inst.key.IsOpen()
}

func WSIsOnline() bool {
	return _inst.websocket != nil && _inst.websocket.IsOnline
}

func WSOnline() error {
	if WSIsOnline() {
		return nil
	}

	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")
	}

	return _inst.websocket.Online()
}

func WSOffline() {
	if _inst.websocket == nil {
		fmt.Println("nil, no need to offline")
		return
	}

	_inst.websocket.ShutDown()
}

func WriteMessage(to string, plainTxt string) error {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")
	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}
	rawData, err := unicast.WrapPlainTxt(plainTxt)
	if err != nil {
		return err
	}
	return _inst.websocket.Write(to, rawData)
}

func WriteLocationMessage(to string, longitude, latitude float32, name string) error {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")
	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}

	rawData, err := unicast.WrapLocation(longitude, latitude, name)
	if err != nil {
		return err
	}

	return _inst.websocket.Write(to, rawData)
}

func WriteImageMessage(to string, payload []byte) error {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")
	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}

	rawData, err := unicast.WrapImage(payload)
	if err != nil {
		return err
	}
	return _inst.websocket.Write(to, rawData)
}

func WriteVoiceMessage(to string, payload []byte, len int) error {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")
	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}

	rawData, err := unicast.WrapVoice(payload, len)
	if err != nil {
		return err
	}

	return _inst.websocket.Write(to, rawData)
}

func SyncGroup(to string, groupId string) error {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")
	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}

	rawData, err := unicast.WrapSyncGroup(groupId)
	if err != nil {
		return err
	}

	return _inst.websocket.Write(to, rawData)
}

func IsValidNinjaAddr(addr string) bool {
	_, err := common.HexToAddress(addr)
	if err != nil {
		return false
	}
	return true
}

func IconIndex(id string, mod int) int64 {
	return int64(utils.ID2IconIdx(id, mod))
}
