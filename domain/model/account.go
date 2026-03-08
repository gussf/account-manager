package model

import (
	"context"
	"errors"
)

type Account struct {
	Id             int // optimally should be an UUID_V7 string, but will use int for simplicity
	DocumentNumber string
}

type AccountService interface {
	CreateAccount(ctx context.Context, req CreateAccountRequest) (*Account, error)
	GetAccount(ctx context.Context, accountID int) (*Account, error)
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
