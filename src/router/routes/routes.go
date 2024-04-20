package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represent all routes of this API
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiredAuthentication bool
}

func SetRoutes(r *mux.Router) *mux.Router {
	routes := routeUsers

	for _, userRoute := range routes {
		r.HandleFunc(userRoute.URI, userRoute.Function).Methods(userRoute.Method)
	}

	return r
}
