package main

import (
	"app/config"
	"app/server"
	"log"
)

func main() {

	srv := server.NewHttp(config.ServerAddr)

	err := srv.ListenAndServe()

	log.Fatalf("server could not be started: %v", err)
}


