package androidlib

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/ninjahome/ninja-go/cli_lib/chat_msg"
	"github.com/ninjahome/ninja-go/cli_lib/utils"
	"github.com/ninjahome/ninja-go/common"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"github.com/ninjahome/ninja-go/service/client"
	"github.com/ninjahome/ninja-go/service/websocket"
	"github.com/ninjahome/ninja-go/wallet"
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

type AndroidAPP struct {
	key       *wallet.Key
	cb        AppCallBack
	wsEnd     string
	websocket *client.WSClient
	unreadSeq int64
}

func (a AndroidAPP) OnlineSuccess() {
	if err := _inst.websocket.PullUnreadMsg(a.unreadSeq); err != nil {
		fmt.Println(err.Error())
	}
}

func (a AndroidAPP) ImmediateMessage(msg *pbs.WSCryptoMsg) error {
	if msg == nil {
		return errors.New("msg is nil")
	}
	return a.callback(msg)
}

func (a AndroidAPP) WebSocketClosed() {
	a.cb.WebSocketClosed()
}

func (i AndroidAPP) UnreadMsg(ack *pbs.WSUnreadAck) error {
	payload := ack.Payload

	for j := 0; j < len(payload); j++ {
		if err := i.callback(payload[j]); err != nil {
			//TODO::notify app to know the failure
			continue
		}
	}
	return nil
}

func (i AndroidAPP) callback(msg *pbs.WSCryptoMsg) error {

	chatMessage := &chat_msg.ChatMessage{}
	if err := proto.Unmarshal(msg.PayLoad, chatMessage); err != nil {
		return err
	}
	switch chatMessage.Payload.(type) {

	case *chat_msg.ChatMessage_PlainTxt:

		rawData := chatMessage.Payload.(*chat_msg.ChatMessage_PlainTxt)

		return i.cb.TextMessage(msg.From,
			msg.To,
			rawData.PlainTxt,
			msg.UnixTime)

	case *chat_msg.ChatMessage_Image:

		rawData := chatMessage.Payload.(*chat_msg.ChatMessage_Image)

		return i.cb.ImageMessage(msg.From,
			msg.To,
			rawData.Image,
			msg.UnixTime)

	case *chat_msg.ChatMessage_Voice:

		voiceMessage := chatMessage.Payload.(*chat_msg.ChatMessage_Voice).Voice

		return i.cb.VoiceMessage(msg.From,
			msg.To,
			voiceMessage.Data,
			int(voiceMessage.Length),
			msg.UnixTime)

	case *chat_msg.ChatMessage_Location:

		locationMessage := chatMessage.Payload.(*chat_msg.ChatMessage_Location).Location

		return i.cb.LocationMessage(msg.From,
			msg.To,
			locationMessage.Longitude,
			locationMessage.Latitude,
			locationMessage.Name,
			msg.UnixTime)
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

var _inst = &AndroidAPP{unreadSeq: 0}

type AppCallBack interface {
	VoiceMessage(from, to string, payload []byte, length int, time int64) error
	ImageMessage(from, to string, payload []byte, time int64) error
	LocationMessage(from, to string, l, a float32, name string, time int64) error
	TextMessage(from, to string, payload string, time int64) error
	WebSocketClosed()
}

func ConfigApp(addr string, callback AppCallBack) {

	if addr == "" {
		addr = client.RandomBootNode()
	}
	//fmt.Println("======>", addr)
	_inst.wsEnd = addr
	_inst.cb = callback
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
	rawData, err := chat_msg.WrapPlainTxt(plainTxt)
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

	rawData, err := chat_msg.WrapLocation(longitude, latitude, name)
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

	rawData, err := chat_msg.WrapImage(payload)
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

	rawData, err := chat_msg.WrapVoice(payload, len)
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

func IconIndex(id string,mod int) int64 {
	return int64(utils.ID2IconIdx(id,mod))
}