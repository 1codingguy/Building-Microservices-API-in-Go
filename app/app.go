package app

import (
	"log"
	"net/http"
)

func Start() {
	// uses standard http library instead of external libraries

	// function to register route
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// function to start server
	// nil because relying on the default multiplexer instead of creating one our own
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}
