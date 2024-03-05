package domain

import (
	"banking/dto"
	"banking/errs"
)

type Account struct {
	AccountId   string  `db:"account_id" json:"account_id"`
	CustomerId  string  `db:"customer_id" json:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
	FindBy(string) (*Account, *errs.AppError)
}

func (r Account) CanWithdraw(amount float64) bool {
	return r.Amount >= amount
}
