package domain

//  "domain" refers to the business domain or the problem space that the software is addressing

// business logic - what is a customer
// Customer is the domain modal/ entity here
type Customer struct {
	Id          string
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string
	Status      string
}

// Defining a port - an interface that define the expected interactions with external actors (like a web client, a file system, or a database)

// Repository/ interface, defines the action
// How business logic (Customer) interacts with backend
type CustomerRepository interface {
	// "repository" is more specific to data storage and retrieval abstractions
	FindAll() ([]Customer, error)
}
