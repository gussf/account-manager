package model

import (
	"context"
	"errors"
)

type Account struct {
	Id             string
	DocumentNumber string
}

type AccountService interface {
	CreateAccount(ctx context.Context, req CreateAccountRequest) (*Account, error)
	GetAccount(ctx context.Context, accountID string) (*Account, error)
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
