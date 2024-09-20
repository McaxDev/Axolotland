package main

import "os"

type Config struct {
	SMS SmsConfig
}

type SmsConfig struct {
	ID        string
	Secret    string
	Signature string
	Template  string
}

var config Config

func GetConfig() {
	config.SMS.ID = os.Getenv("SMS_ID")
	config.SMS.Secret = os.Getenv("SMS_SECRET")
	config.SMS.Signature = os.Getenv("SMS_SIGNATURE")
	config.SMS.Template = os.Getenv("SMS_TEMPLATE")
}
