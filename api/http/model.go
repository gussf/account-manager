package http

import (
	"account-manager/domain/core"
	"time"
)

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

func (c CreateAccountRequest) toDomain() core.CreateAccountRequest {
	return core.CreateAccountRequest{
		DocumentNumber: c.DocumentNumber,
	}
}

type CreateAccountResponse struct {
	ID             int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

type GetAccountResponse struct {
	ID             int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

type SaveTransactionRequest struct {
	AccountID       int     `json:"account_id"`
	OperationTypeID int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

func (s SaveTransactionRequest) toDomain() core.SaveTransactionRequest {
	return core.SaveTransactionRequest{
		AccountID:       s.AccountID,
		OperationTypeID: s.OperationTypeID,
		Amount:          s.Amount,
	}
}

type SaveTransactionResponse struct {
	ID              int       `json:"id"`
	OperationTypeID int       `json:"operation_type_id"`
	Amount          float64   `json:"amount"`
	EventDate       time.Time `json:"event_date,omitempty"`
}

type ErrorMessage struct {
	Error string `json:"error"`
}
