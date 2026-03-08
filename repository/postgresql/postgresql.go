package postgresql

import (
	"account-manager/domain/model"
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

func (s *Store) CreateAccount(model.CreateAccountRequest) (*model.Account, error) {
	return &model.Account{}, nil
}

func (s *Store) GetAccount(accountID string) (*model.Account, error) {
	return &model.Account{}, nil
}

func (s *Store) SaveTransaction(model.SaveTransactionRequest) (*model.Transaction, error) {
	return &model.Transaction{}, nil
}
