package http

import (
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := "hello"
	respond(w, response, 200)
}
