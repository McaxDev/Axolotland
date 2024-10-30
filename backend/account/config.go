package main

import "os"

var Config struct {
	VerifyAddr string
}

func LoadConfig() {
	Config.VerifyAddr = os.Getenv("VERIFY_ADDR")
}
