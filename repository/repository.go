package repository

import (
	"account-manager/domain/core"
	"context"
)

type Store interface {
	CreateAccount(ctx context.Context, req core.CreateAccountRequest) (*core.Account, error)
	GetAccount(ctx context.Context, accountID int) (*core.Account, error)
	SaveTransaction(ctx context.Context, req core.SaveTransactionRequest) (*core.Transaction, error)
}
