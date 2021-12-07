package main

import (
	"os"

	"Integrate.Interview.Backend.Golang/cmd"
)

var (
	port = os.Getenv("PORT")
)

func main() {
	if port == "" {
		port = "8084"
	}
	cmd.Setup(port)
}
