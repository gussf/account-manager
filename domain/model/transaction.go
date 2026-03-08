package model

import (
	"context"
	"errors"
	"time"
)

type Transaction struct {
	ID            int // Should be an UUID_V7 string, but will use int for simplicity
	AccountID     int
	OperationType OperationType
	Amount        float64
	EventDate     time.Time
}

type TransactionService interface {
	SaveTransaction(ctx context.Context, req SaveTransactionRequest) (*Transaction, error)
}

type SaveTransactionRequest struct {
	AccountID     int
	OperationType OperationType
	Amount        float64
}

func (s *SaveTransactionRequest) Validate() error {
	if s.AccountID <= 0 {
		return errors.New("invalid account id")
	}

	if s.Amount == 0 {
		return errors.New("amount cannot be zero")
	}

	return nil
}
