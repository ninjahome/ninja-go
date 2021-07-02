package push

import (
	"crypto/tls"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
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

	ip.client.CloseIdleConnections()

	notification := &apns2.Notification{}
	notification.DeviceToken = devToken
	notification.Topic = AppBundle


	notification.Payload = payload.NewPayload().Alert(alert).Badge(1).SoundName("default")

	res,err:=ip.client.Push(notification)
	if err!=nil{
		log.Println("ios send notification error",err)
		return err
	}
	log.Println("ios send notification success",devToken,res.StatusCode, res.ApnsID, res.Reason)

	return nil
}
