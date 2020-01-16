package server

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	model "github.com/skill-central-golang/models"
)

// this is wher everything connects
type Server struct {
	Router *mux.Router
	Config *model.Config
	Log    *logrus.Logger
}

// this struct is used as a starting param to get server instance
type ServerConfig struct {
	Config  *model.Config
	Log     *logrus.Logger
	Handler *mux.Router
}

func BootstrapServer(config *ServerConfig) *Server {
	srv := &Server{
		Router: config.Handler,
		Config: config.Config,
		Log:    config.Log,
	}
	// connect DB and attach to the server struct

	return srv
}

func (s *Server) Start() {
	server := &http.Server{
		Addr:           httpAddress(s.Config.ServerSettings),
		Handler:        s.Router,
		ReadTimeout:    s.Config.ServerSettings.ReadTimeout,
		WriteTimeout:   s.Config.ServerSettings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}

func httpAddress(s *model.ServerSettings) string {
	return ":" + fmt.Sprintf("%d", s.HTTPPort)
}
