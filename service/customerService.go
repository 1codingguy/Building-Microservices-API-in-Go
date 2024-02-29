package service

import (
	"banking/domain"
	"banking/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(id string) (*domain.Customer, *errs.AppError)
}

// service implementation
// "implementing" an interface in Go means that a type (like a struct) has all the methods that the interface describes.

type Service struct {
	// Whatever repo ends up being, it must know how to FindAll customers.
	repo domain.CustomerRepository
}

func (s Service) GetAllCustomer() ([]domain.Customer, error) {
	// This bit has nothing to do with CustomerRepositoryInterface.
	// GetAllCustomer() is just a wrapper to call FindAll() in repo
	return s.repo.FindAll()
}

func (s Service) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

// helper func to instantiate Service
func NewService(repo domain.CustomerRepository) Service {
	// This function takes a repo, and create a Service struct with it.
	// It means injecting a repo and creating a new struct.
	// That is why "Service uses an instance of CustomerRepositoryInterface"
	return Service{repo}
}
