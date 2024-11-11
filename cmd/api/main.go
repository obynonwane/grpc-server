package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRpcPort = "50001"
)

type Config struct {
}

func main() {

	app := Config{}

	go app.gRPCListen()

	log.Printf("starting logger service on port %s\n", webPort)
	//define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	//start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
