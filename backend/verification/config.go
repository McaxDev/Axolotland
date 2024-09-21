package main

import (
	"os"
)

type Config struct {
	GrpcPort  string
	HttpPort  string
	GeoServer string
	SMS       SmsConfig
	SMTP      SmtpConfig
}

type SmsConfig struct {
	ID        string
	Secret    string
	Signature string
	Template  string
}

type SmtpConfig struct {
	Server   string
	Port     string
	Mail     string
	Password string
}

var config Config

func GetConfig() {
	var exist bool
	config.GrpcPort, exist = os.LookupEnv("GRPC_PORT")
	if !exist {
		config.GrpcPort = "50051"
	}
	config.HttpPort, exist = os.LookupEnv("HTTP_PORT")
	if !exist {
		config.HttpPort = "8080"
	}
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
