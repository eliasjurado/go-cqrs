package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "8080"

type Config struct {
}

func main() {
	app := Config{}
	log.Printf("Starting broker service on port %+v\n", port)

	//define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}

	//start Server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
