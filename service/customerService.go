package service

import "github.com/1codingguy/go-microservice-api/banking/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

// service implementation
type DefaultCustomerService struct {
	// repo can hold anything that implements the CustomerRepository interface
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// helper func to instantiate DefaultCustomerService
func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
