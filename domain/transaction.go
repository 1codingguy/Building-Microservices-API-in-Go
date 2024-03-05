package domain

import "banking/dto"

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionId   string
	AccountId       string
	Amount          float64
	TransactionType string
	TransactionDate string
	UpdatedBalance  float64
}

func (t Transaction) IsWithdrawal() bool {
	// used in accountRepositoryDb
	return t.TransactionType == WITHDRAWAL
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		UpdatedBalance:  t.UpdatedBalance,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
