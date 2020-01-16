package api

import "net/http"

func (api *API) initEmpolyee() {
	api.BaseRouter.Employee.HandleFunc("", api.createEmployee).Methods("POST")
	api.BaseRouter.Employee.HandleFunc("/all", api.getEmployees).Methods("GET")
}

func (a *API) createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Work In Progress"))
}

func (a *API) getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Work In Progress"))
}
