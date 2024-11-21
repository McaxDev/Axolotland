package main

var Servers []Server

type Server struct {
	Name    string
	Comment string
	Game    string
	Path    struct {
		Server string
		World  string
	}
	Backup struct {
		Enable    bool
		Frequency string
		Path      string
		Limit     int
	}
	AllowedCommands []string
}
