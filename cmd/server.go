package cmd

import "Integrate.Interview.Backend.Golang/server"

var ip = "127.0.0.1"

func Setup(port string) {

	/*
		Dependency Injection here to expost lower layers
	*/

	server.Start(ip, port)
}
