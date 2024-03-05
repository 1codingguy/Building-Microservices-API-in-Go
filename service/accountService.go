package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
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

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	// transaction validation
	err := req.Validate()

	if err != nil {
		return nil, err
	}

	// server side validation for checking the available balance in the account
	// actually checking data stored in sql
	if req.IsWithdrawal() {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}

		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}

	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(time.DateTime),
	}

	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}

	// convert Transaction object into TransactionResponse
	response := transaction.ToDto()
	return &response, nil
}
