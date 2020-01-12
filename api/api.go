package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/skill-central-golang/server"
)

type Router struct {
	Root *mux.Router
}

type API struct {
	Srv        *server.Server
	BaseRouter *Router
}

func GetHandler() http.Handler {
	router := mux.NewRouter()
	return router
}

func (r *Router) InitRoutes(srv *server.Server) {
	api := &API{
		Srv:        srv,
		BaseRouter: r,
	}

	api.initEmpolyee()
}
