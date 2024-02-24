package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipCode"`
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

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
