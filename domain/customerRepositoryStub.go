package domain

// placeholder for a mock?
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			Id:          "1",
			Name:        "John",
			City:        "Phuket",
			ZipCode:     "83000",
			DateOfBirth: "2000-01-01",
			Status:      "active",
		},
		{
			Id:          "2",
			Name:        "Peter",
			City:        "Bangkok",
			ZipCode:     "10000",
			DateOfBirth: "1999-01-01",
			Status:      "active",
		},
	}
	return CustomerRepositoryStub{customers: customers}
}
