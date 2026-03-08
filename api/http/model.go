package http

import (
	"account-manager/domain/model"
	"time"
)

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

func (c CreateAccountRequest) toDomain() model.CreateAccountRequest {
	return model.CreateAccountRequest{
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
	AccountID     int                 `json:"account_id"`
	OperationType model.OperationType `json:"operation_type"`
	Amount        float64             `json:"amount"`
}

func (s SaveTransactionRequest) toDomain() model.SaveTransactionRequest {
	return model.SaveTransactionRequest{
		AccountID:     s.AccountID,
		OperationType: s.OperationType,
		Amount:        s.Amount,
	}
}

type SaveTransactionResponse struct {
	ID            int                 `json:"id"`
	OperationType model.OperationType `json:"operation_type"`
	Amount        float64             `json:"amount"`
	EventDate     time.Time           `json:"event_date,omitempty"`
}
