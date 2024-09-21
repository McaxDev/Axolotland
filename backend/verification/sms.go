package main

import (
	"time"

	"github.com/McaxDev/Axolotland/backend/utils"
	unisms "github.com/apistd/uni-go-sdk/sms"
)

var smsSent = make(map[string]smsSentValue)

type smsSentValue struct {
	AuthCode string
	Expiry   time.Time
}

func SendSMS(telephone string) error {
	authcode := utils.RandomCode(6)
	client := unisms.NewClient(config.SMS.ID, config.SMS.Secret)

	message := unisms.BuildMessage()
	message.SetTo(telephone)
	message.SetSignature(config.SMS.Signature)
	message.SetTemplateId(config.SMS.Template)
	message.SetTemplateData(map[string]string{
		"code": authcode,
		"ttl":  "10",
	})

	_, err := client.Send(message)
	if err != nil {
		return err
	}

	smsSent[telephone] = smsSentValue{
		AuthCode: authcode,
		Expiry:   time.Now().Add(10 * time.Minute),
	}

	return nil
}
