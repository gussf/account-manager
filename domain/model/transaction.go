package model

import (
	"errors"
	"time"
)

type Transaction struct {
	Id            string
	AccountId     string
	OperationType OperationType
	Amount        int64
	EventDate     time.Time
}

type TransactionService interface {
	SaveTransaction(SaveTransactionRequest) (*Transaction, error)
}

type SaveTransactionRequest struct {
	AccountID     string
	OperationType OperationType
	Amount        int64
}

func (s *SaveTransactionRequest) Validate() error {
	if s.AccountID == "" {
		return errors.New("invalid account id")
	}

	if s.Amount == 0 {
		return errors.New("amount cannot be zero")
	}

	// should we expect amount to be negative when op_type in (1,2,3)?

	return nil
}
