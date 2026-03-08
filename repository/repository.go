package repository

import (
	"account-manager/domain/model"
	"context"
)

type Store interface {
	CreateAccount(ctx context.Context, req model.CreateAccountRequest) (*model.Account, error)
	GetAccount(ctx context.Context, accountID int) (*model.Account, error)
	SaveTransaction(ctx context.Context, req model.SaveTransactionRequest) (*model.Transaction, error)
}
