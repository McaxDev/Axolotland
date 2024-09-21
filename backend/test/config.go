package main

type Config struct {
	HttpPort string
	GrpcPort string
	ImageURL string
	DB       DbConfig
}

type DbConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

var config Config
