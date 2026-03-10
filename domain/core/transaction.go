package core

import (
	"context"
	"errors"
	"time"
)

type Transaction struct {
	ID              int // Should be an UUID_V7 string, but will use int for simplicity
	AccountID       int
	OperationTypeID int
	Amount          float64
	EventDate       time.Time
}

type TransactionService interface {
	SaveTransaction(ctx context.Context, req SaveTransactionRequest) (*Transaction, error)
}

type SaveTransactionRequest struct {
	AccountID       int
	OperationTypeID int
	Amount          float64
}

func (s *SaveTransactionRequest) Validate() error {
	if s.AccountID <= 0 {
		return errors.New("invalid account id")
	}

	if s.OperationTypeID <= 0 {
		return errors.New("invalid operation type id")
	}

	if s.Amount <= 0 {
		return errors.New("amount cannot be equal or less than zero")
	}

	return nil
}
