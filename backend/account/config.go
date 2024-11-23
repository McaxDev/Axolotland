package main

import "os"

var Config struct {
	AuthAddr string
	JwtKey   string
	Port     string
	DB       DBConfig
	SSL      SSLConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type SSLConfig struct {
	Certificate string
	Private     string
}

func LoadConfig() {

	Config.AuthAddr = os.Getenv("AUTH_ADDR")
	Config.JwtKey = os.Getenv("JWT_KEY")
	Config.Port = os.Getenv("PORT")

	Config.DB.Host = os.Getenv("DB_HOST")
	Config.DB.Port = os.Getenv("DB_PORT")
	Config.DB.User = os.Getenv("DB_USER")
	Config.DB.Password = os.Getenv("DB_PASSWORD")
	Config.DB.Name = os.Getenv("DB_NAME")

	if certificate, exists := os.LookupEnv("SSL_CERTIFICATE"); exists {
		Config.SSL.Certificate = certificate
	}

	if private, exists := os.LookupEnv("SSL_PRIVATE"); exists {
		Config.SSL.Private = private
	}
}
