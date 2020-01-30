package main

import (
	"github.com/sirupsen/logrus"

	"github.com/skill-central-golang/api"
	"github.com/skill-central-golang/server"
	"github.com/skill-central-golang/utils"
)

func main() {

	logger := logrus.New()

	logger.Info("Starting Server")

	logger.Info("Getting Handler")
	handler := api.GetHandler()

	logger.Info("Loading Config")
	config := utils.LoadConfig()

	// Creating Logger
	serverConfig := &server.ServerConfig{
		Config:  config,
		Log:     logger,
		Handler: handler,
	}

	logger.Info("Initializing Server")
	srv := server.BootstrapServer(serverConfig)

	logger.Info("Initializing Routes")
	api.InitRoutes(srv)

	logger.Info("Staring Server")
	srv.Start()
}
