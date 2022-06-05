package main

import (
	"os"

	"github.com/gocopper/copper"
)

func main() {
	var app = copper.New()

	server, err := InitServer(app)
	if err != nil {
		app.Logger.Error("Failed to init server", err)
		os.Exit(1)
	}

	app.Start(server)
}
