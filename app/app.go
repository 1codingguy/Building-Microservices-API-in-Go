package app

import (
	"log"
	"net/http"
)

func Start() {

	// create our own multiplexer
	mux := http.NewServeMux()

	// function to register route
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	// function to start server
	// nil because relying on the default multiplexer instead of creating one our own
	log.Fatal(http.ListenAndServe("localhost:9000", mux))
}
