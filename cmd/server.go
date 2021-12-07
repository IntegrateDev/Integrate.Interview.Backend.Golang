package cmd

import "Integrate.Interview.Backend.Golang/server"

func Setup(port string) {
	server.Start("127.0.0.1", port)
}
