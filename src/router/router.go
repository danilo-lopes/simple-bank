package router

import (
	"github.com/danilo-lopes/simple-bank/src/router/routes"
	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.Configure(r)
}
