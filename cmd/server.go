package cmd

import "Integrate.Interview.Backend.Golang/server"

func Setup() {
	server.Start("127.0.0.1", "8084")
}
