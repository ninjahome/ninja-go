package push

import (
	"crypto/tls"
	"github.com/polydawn/refmt/json"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"log"
)

const(
	AppBundle = "com.hop.ninja.light"
)

type IOSPush struct {
	cert tls.Certificate
	client *apns2.Client
}

type ApsContent struct {
	Alert string `json:"alert"`
	Sound string `json:"sound,omitempty"`
	Badge int    `json:"badge"`
}

type PayloadContent struct {
	Aps *ApsContent `json:"aps"`
}

func NewIOSPush(certfile string) *IOSPush {
	cert,err:=certificate.FromP12File(certfile,"")
	if err!=nil{
		return nil
	}
	client := apns2.NewClient(cert).Production()

	return &IOSPush{
		client: client,
		cert: cert,
	}
}

func (ip *IOSPush)IOSPushMessage(alert string, devToken string) error  {
	notification := &apns2.Notification{}
	notification.DeviceToken = devToken
	notification.Topic = AppBundle

	payload:=&PayloadContent{
		Aps: &ApsContent{},
	}

	payload.Aps.Alert = alert
	payload.Aps.Badge = 1

	j,_:=json.Marshal(payload)

	notification.Payload = j

	res,err:=ip.client.Push(notification)
	if err!=nil{
		log.Println("ios send notification error",err)
		return err
	}
	log.Println("ios send notification success",devToken,res.StatusCode, res.ApnsID, res.Reason)

	return nil
}
