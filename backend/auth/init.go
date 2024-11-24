package main

import (
	"time"

	"github.com/McaxDev/Axolotland/backend/auth/rpc"
	unisms "github.com/apistd/uni-go-sdk/sms"
	"github.com/go-resty/resty/v2"
)

type MsgSent = map[string]MsgSentValue

type MsgSentValue struct {
	Expiry   time.Time
	Authcode string
}

type RPCServer struct {
	rpc.UnimplementedAuthServer
}

var (
	EmailSent     MsgSent
	TelephoneSent MsgSent
	QQSent        MsgSent
	RestyClient   *resty.Client
	SMSClient     *unisms.UniSMSClient
)

func Init() {
	LoadConfig()
	EmailSent = make(MsgSent)
	TelephoneSent = make(MsgSent)
	QQSent = make(MsgSent)
	RestyClient = resty.New().SetTimeout(5 * time.Second)
	SMSClient = unisms.NewClient(config.SMS.ID, config.SMS.Secret)

}
