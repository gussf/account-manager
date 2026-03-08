package model

import "errors"

type Account struct {
	Id             string `json:"id"`
	DocumentNumber string `json:"document_number"`
}

type AccountService interface {
	CreateAccount(CreateAccountRequest) (*Account, error)
	GetAccount(accountID string) (*Account, error)
}

type CreateAccountRequest struct {
	DocumentNumber string
}

func (c *CreateAccountRequest) Validate() error {
	if c.DocumentNumber == "" {
		return errors.New("invalid document number")
	}

	return nil
}
