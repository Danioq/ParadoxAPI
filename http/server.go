package http

import "net/http"

func GetServer(listen string) *http.Server {
	router := getRouter()
	return &http.Server{
		Addr:    listen,
		Handler: router,
	}

}
