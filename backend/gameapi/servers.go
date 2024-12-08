package main

import "gorm.io/gorm"

var Servers []Server

type Server struct {
	gorm.Model
	Name       string
	Comment    string
	Game       string
	ServerPath string
	Backup     struct {
		Enable     bool
		Frequency  string
		Limit      int
		WorldPath  string
		BackupPath string
	} `gorm:"type:json"`
	AllowedCmds []string `gorm:"type:json"`
}
