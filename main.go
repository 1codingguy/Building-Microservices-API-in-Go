package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func main() {
	// user standard http library instead of external libraries

	// function to register route
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// function to start server
	// nil because relying on the default multiplexer instead of creating one our own
	err := http.ListenAndServe("localhost:9000", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func greet(w http.ResponseWriter, r *http.Request) {
	// send the response back to the client
	// Fprintf takes the "hello world" string, write it to w
	fmt.Fprint(w, "hello world")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{
			Name:    "John",
			City:    "Phuket",
			ZipCode: "83000",
		},
		{
			Name:    "Peter",
			City:    "Bangkok",
			ZipCode: "10000",
		},
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
