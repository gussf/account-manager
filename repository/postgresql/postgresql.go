package postgresql

import (
	"account-manager/domain/model"
	"context"
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateAccount(ctx context.Context, req model.CreateAccountRequest) (*model.Account, error) {
	return &model.Account{}, nil
}

func (s *Store) GetAccount(ctx context.Context, accountID string) (*model.Account, error) {
	return &model.Account{}, nil
}

func (s *Store) SaveTransaction(ctx context.Context, req model.SaveTransactionRequest) (*model.Transaction, error) {
	return &model.Transaction{}, nil
}
