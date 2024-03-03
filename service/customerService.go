package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]*dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

// service implementation
// "implementing" an interface in Go means that a type (like a struct) has all the methods that the interface describes.

type Service struct {
	// Whatever repo ends up being, it must know how to FindAll customers.
	repo domain.CustomerRepository
}

func (s Service) GetAllCustomer(status string) ([]*dto.CustomerResponse, *errs.AppError) {
	// This bit has nothing to do with CustomerRepositoryInterface.
	// GetAllCustomer() is just a wrapper to call FindAll() in repo

	// Service layer should be responsible translating the status into code for sql query

	c, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	var response []*dto.CustomerResponse

	for _, customer := range c {
		r := customer.ToDto()
		response = append(response, &r)
	}
	return response, nil
}

func (s Service) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

// helper func to instantiate Service
func NewService(repo domain.CustomerRepository) Service {
	// This function takes a repo, and create a Service struct with it.
	// It means injecting a repo and creating a new struct.
	// That is why "Service uses an instance of CustomerRepositoryInterface"
	return Service{repo}
}
