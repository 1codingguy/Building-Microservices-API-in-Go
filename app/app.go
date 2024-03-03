package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"banking/domain"
	"banking/service"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func Start() {

	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	server := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", server, port), router))
}
