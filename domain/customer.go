package domain

import "banking/errs"

//  "domain" refers to the business domain or the problem space that the software is addressing

// business logic - what is a customer
// Customer is the domain modal/ entity here
type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

// Defining a port - an interface that define the expected interactions with external actors (like a web client, a file system, or a database)

// Repository/ interface, defines the action
// How business logic (Customer) interacts with backend
type CustomerRepository interface {
	// "repository" is more specific to data storage and retrieval abstractions
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError) // a pointer: able to use nil pointer if a customer doesn't exist
}
