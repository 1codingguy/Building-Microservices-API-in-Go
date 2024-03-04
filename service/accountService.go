package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()

	if err != nil {
		return nil, err
	}

	// transform the details to domain
	request := domain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format(time.DateTime),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1", // active
	}

	newAccount, err := s.repo.Save(request)

	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDto()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
