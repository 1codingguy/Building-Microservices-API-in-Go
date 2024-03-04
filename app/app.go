package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"banking/domain"
	"banking/service"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

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

	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDB(dbClient)
	ch := CustomerHandlers{service: service.NewService(customerRepositoryDb)}
	ah := AccountHandler{service: service.NewAccountService(accountRepositoryDb)}

	// function to register route
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	// function to start server
	// nil because relying on the default multiplexer instead of creating one our own

	server := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", server, port), router))
}

// create one connection pool, used by different repositories
func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@(%s:%s)/%s", dbUser, dbPassword, dbAddress, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
