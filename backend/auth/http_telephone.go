package main

import (
	"time"

	"github.com/McaxDev/Axolotland/backend/utils"
	unisms "github.com/apistd/uni-go-sdk/sms"
	"github.com/gin-gonic/gin"
)

func SendTelephone(c *gin.Context) {

	telephone := c.Param("number")
	authcode := utils.RandomCode(6)

	message := unisms.BuildMessage()
	message.SetTo(telephone)
	message.SetSignature(config.SMS.Signature)
	message.SetTemplateId(config.SMS.Template)
	message.SetTemplateData(map[string]string{
		"code": authcode,
		"ttl":  "10",
	})

	_, err := SMSClient.Send(message)
	if err != nil {
		c.JSON(500, utils.Resp("短信发送失败", nil))
		return
	}

	TelephoneSent[telephone] = MsgSentValue{
		Authcode: authcode,
		Expiry:   time.Now().Add(10 * time.Minute),
	}

	c.JSON(200, utils.Resp("验证码发送成功", nil))
}
