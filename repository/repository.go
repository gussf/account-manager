package repository

import "account-manager/domain/model"

type Store interface {
	CreateAccount(model.CreateAccountRequest) (*model.Account, error)
	GetAccount(accountID string) (*model.Account, error)
	SaveTransaction(model.SaveTransactionRequest) (*model.Transaction, error)
}
