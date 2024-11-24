package main

import (
	"os"
)

var config struct {
	GrpcPort  string
	HttpPort  string
	GeoServer string
	SMS       struct {
		ID        string
		Secret    string
		Signature string
		Template  string
	}
	SMTP struct {
		Server   string
		Port     string
		Mail     string
		Password string
	}
}

func LoadConfig() {
	config.GrpcPort = os.Getenv("GRPC_PORT")
	config.HttpPort = os.Getenv("HTTP_PORT")
	config.GeoServer = os.Getenv("GEO_SERVER")
	config.SMS.ID = os.Getenv("SMS_ID")
	config.SMS.Secret = os.Getenv("SMS_SECRET")
	config.SMS.Signature = os.Getenv("SMS_SIGNATURE")
	config.SMS.Template = os.Getenv("SMS_TEMPLATE")
	config.SMTP.Server = os.Getenv("SMTP_SERVER")
	config.SMTP.Port = os.Getenv("SMTP_PORT")
	config.SMTP.Mail = os.Getenv("SMTP_MAIL")
	config.SMTP.Password = os.Getenv("SMTP_PASSWORD")
}
