package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	AuthenticationRequired bool
}

func Configure(r *mux.Router) *mux.Router {
	apiRoutes := accountsRoute

	for _, apiRoute := range apiRoutes {
		r.HandleFunc(apiRoute.URI, apiRoute.Function).Methods(apiRoute.Method)
	}

	return r
}
