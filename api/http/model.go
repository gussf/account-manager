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
	Id             int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

type GetAccountResponse struct {
	Id             int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

type SaveTransactionRequest struct {
	AccountId     int                 `json:"account_id"`
	OperationType model.OperationType `json:"operation_type"`
	Amount        int64               `json:"amount"`
}

func (s SaveTransactionRequest) toDomain() model.SaveTransactionRequest {
	return model.SaveTransactionRequest{
		AccountID:     s.AccountId,
		OperationType: s.OperationType,
		Amount:        s.Amount,
	}
}

type SaveTransactionResponse struct {
	Id            int                 `json:"id"`
	OperationType model.OperationType `json:"operation_type"`
	Amount        int64               `json:"amount"`
	EventDate     time.Time           `json:"event_date,omitempty"`
}
