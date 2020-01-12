package main

import (
	"github.com/sirupsen/logrus"
	"fmt"

	"github.com/skill-central-golang/server"
	"github.com/skill-central-golang/utils"
	"github.com/skill-central-golang/models"
)

func main() {
	fmt.Println("Starting Server")

	handler := api.Gethandler();
	config := utils.LoadConfig();

	// Creating Logger
	serverConfig := &server.ServerConfig{
		Config: config,
		Log: logrus.New(),
		Handler: handler,
	}

	srv := server.BootstrapServer(serverConfig);	
	handler.InitRoutes(srv);

	srv.Start();

}
