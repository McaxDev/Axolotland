package config

var servers []struct {
	Name   string
	Game   string
	Path   string
	Backup struct {
		Frequency string
		Limit     int
	}
}
