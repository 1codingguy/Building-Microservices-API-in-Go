package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"banking/service"

	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	// service is a type that has a GetAllCustomer() method
	service service.CustomerService
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["customer_id"]

	customer, err := ch.service.GetCustomer(id)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Add("Content-Type", "application/json") // must be first
	w.WriteHeader(code)                                // must be before sending the json
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
