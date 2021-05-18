package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/hello").HandlerFunc(HelloHandler)
	return r
}
