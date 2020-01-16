package api

import (
	"github.com/gorilla/mux"
	"github.com/skill-central-golang/models"
	"github.com/skill-central-golang/server"
)

type Router struct {
	Root      *mux.Router
	ApiRootV1 *mux.Router // 'api/v1'
	Employee  *mux.Router // 'api/v1/employee'
}

type API struct {
	Srv        *server.Server
	BaseRouter *Router
}

func GetHandler() *mux.Router {
	router := mux.NewRouter()
	return router
}

func InitRoutes(srv *server.Server) {
	api := &API{
		Srv:        srv,
		BaseRouter: &Router{},
	}

	api.BaseRouter.Root = srv.Router
	api.BaseRouter.ApiRootV1 = srv.Router.PathPrefix(model.API_URL_SUFFIX).Subrouter()
	api.BaseRouter.Employee = api.BaseRouter.ApiRootV1.PathPrefix("/employee").Subrouter()

	api.initEmpolyee()
}
