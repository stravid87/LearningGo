package main

import (
	"log"
	"fmt"
	"net/http"
)

const webPort = "80"

type Config struct {}

func main() {
	app := Config{}

	log.Printf("Starting Broker service on port %s\n", webPort)

	//define HTTP server 
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	//start the Server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}