package push

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"log"
	"strconv"
	"time"
)

var hostUmengPush = "http://msg.umeng.com"
var uploadPath = "/upload"
var postPath = "/api/send"

var appKeyAndroid = "60a723d4d827ab124e923f62"
var masterSecreptAndroid = "1h3foewyyxhclfpalpprugwazbcvvdcw"

var pushProductionMode = "true"

var pushAndroidActityChat = "com.ninjahome.ninja.ui.activity.splash.SplashActivity"

type UmengAdroid struct {
	Appkey    string `json:"appkey"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`

	DeviceTokens    string          `json:"device_tokens"`
	Production_mode string          `json:"production_mode"`
	Payload         *PayloadAndroid `json:"payload"`
	Description     string          `json:"description"`
}

type PayloadAndroid struct {
	Display_type string            `json:"display_type"`
	Body         *BodyAndroid      `json:"body"`
	Extral       map[string]string `json:"extra"`
}
type BodyAndroid struct {
	Ticker     string `json:"ticker"`
	Title      string `json:"title"`
	Text       string `json:"text"`
	After_open string `json:"after_open"`
	Activity   string `json:"activity"`
}

func AndroidMessagePush(title string, deviceToken string, extralDatas map[string]string) {
	if deviceToken == "" {
		log.Println("device Token must not null")
		return
	}

	body := &BodyAndroid{}
	// 必填 通知栏提示文字
	body.Ticker = title
	// 必填 通知标题
	body.Title = title
	// 必填 通知文字描述
	body.Text = title
	// 打开Android端的Activity
	body.After_open = "go_activity"

	payLoad := &PayloadAndroid{}
	payLoad.Display_type = "notification"
	payLoad.Body = body
	/*
	   额外携带的信息
	*/
	payLoad.Extral = extralDatas

	messageAndroid := UmengAdroid{}
	messageAndroid.Appkey = appKeyAndroid

	//打开聊天
	//body.Activity = pushAndroidActityChat

	messageAndroid.Type = "unicast"
	messageAndroid.DeviceTokens = deviceToken

	// 打开聊天
	body.Activity = pushAndroidActityChat

	timeInt64 := time.Now().Unix()
	timestamp := strconv.FormatInt(timeInt64, 10)
	messageAndroid.Timestamp = timestamp
	messageAndroid.Production_mode = pushProductionMode
	messageAndroid.Payload = payLoad
	messageAndroid.Description = title

	postBody, _ := json.Marshal(messageAndroid)
	url := hostUmengPush + postPath

	// MD5加密
	sign := Md5("POST" + url + string(postBody) + masterSecreptAndroid)
	url = url + "?sign=" + sign

	req := httplib.Post(url)
	req.JSONBody(messageAndroid)

	byteData, err := req.Bytes()
	if err != nil {
		fmt.Println("android err back,", nil, err)
	} else {
		strData := string(byteData)
		fmt.Println("android back,", strData)
		fmt.Println("android send,", messageAndroid)
	}
}

//md5加密
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	sStr := hex.EncodeToString(h.Sum(nil))
	return sStr
}
