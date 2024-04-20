package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate will return router with setted routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.SetRoutes(r)
}
