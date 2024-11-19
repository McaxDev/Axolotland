package config

var Servers struct {
	Servers []struct {
		Name   string
		Game   string
		Path   string
		Backup string
	}
	Options []struct {
		Frequency string
		Limit     int
	}
}
