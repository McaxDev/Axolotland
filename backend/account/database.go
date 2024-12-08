package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {

	var err error

	if DB, err = gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.DB.User,
		Config.DB.Password,
		Config.DB.Host,
		Config.DB.Port,
		Config.DB.Name,
	))); err != nil {
		return err
	}

	DB.AutoMigrate(new(User))

	return nil
}

type User struct {
	gorm.Model
	Name        string
	Password    string
	Avatar      string
	Profile     string
	Admin       bool
	Money       int
	Checkin     int64
	Setting     int64
	Email       string
	Phone       string
	QQ          string
	BedrockName string
	JavaName    string
}

type BlackList struct {
	gorm.Model
	Type   string
	Value  string
	Expiry time.Time
}
