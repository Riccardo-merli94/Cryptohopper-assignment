package server

import (
	"app/config"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func NewHttp(serverAddress string) *http.Server {
	r := mux.NewRouter()

	config.Routes(r)

	server := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
	}
	return server
}
