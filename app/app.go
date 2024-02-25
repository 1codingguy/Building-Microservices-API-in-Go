package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// create our own multiplexer
	router := mux.NewRouter()

	// function to register route
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// function to start server
	// nil because relying on the default multiplexer instead of creating one our own
	log.Fatal(http.ListenAndServe("localhost:9000", router))
}
