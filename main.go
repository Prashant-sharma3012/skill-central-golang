package main

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/skill-central-golang/api"
	"github.com/skill-central-golang/server"
	"github.com/skill-central-golang/utils"
)

func main() {
	fmt.Println("Starting Server")

	fmt.Println("Getting Handler")
	handler := api.GetHandler()

	fmt.Println("Loading Config")
	config := utils.LoadConfig()

	// Creating Logger
	serverConfig := &server.ServerConfig{
		Config:  config,
		Log:     logrus.New(),
		Handler: handler,
	}

	fmt.Println("Initializing Server")
	srv := server.BootstrapServer(serverConfig)

	fmt.Println("Initializing Routes")
	api.InitRoutes(srv)

	fmt.Println("Staring Server")
	srv.Start()
}
