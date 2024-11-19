package config

import (
	"encoding/json"

	"github.com/McaxDev/Axolotland/backend/utils"
)

var Servers []Server

type Server struct {
	Name    string
	Comment string
	Game    string
	Path    string
	Backup  struct {
		enable    bool
		Frequency string
		Limit     int
	}
}

func LoadServers() error {

	serversByte, err := utils.ReadFile(ServersPath)
	if err != nil {
		return err
	}

	return json.Unmarshal(serversByte, &Servers)
}
