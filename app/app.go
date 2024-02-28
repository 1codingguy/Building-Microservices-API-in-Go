package app

import (
	"log"
	"net/http"

	"banking/domain"
	"banking/service"
	"github.com/gorilla/mux"
)

func Start() {

	// create our own multiplexer
	router := mux.NewRouter()

	// wiring
	// service.NewService() - service refers to the package, not Service struct
	// CustomerRepositoryStub implements CustomerRepository interface since it has a FindAll()
	// ch := CustomerHandlers{service: service.NewService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewService(domain.NewCustomerRepositoryDb())}

	// function to register route
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// function to start server
	// nil because relying on the default multiplexer instead of creating one our own
	log.Fatal(http.ListenAndServe("localhost:9000", router))
}
