package main

import "os"

var Config struct {
	Port string
}

func GetConfig() {
	Config.Port = os.Getenv("PORT")
}
