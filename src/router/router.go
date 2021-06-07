package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

func Make() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
