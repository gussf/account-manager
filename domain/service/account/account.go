package account

import (
	"account-manager/domain/model"
	"account-manager/repository"
	"context"
	"errors"
	"fmt"
)

type Service struct {
	store repository.Store
}

func NewService(store repository.Store) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) CreateAccount(ctx context.Context, req model.CreateAccountRequest) (*model.Account, error) {
	err := req.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate create account request: %w", err)
	}

	createdAcc, err := s.store.CreateAccount(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("create account in store: %w", err)
	}

	return createdAcc, nil
}

func (s *Service) GetAccount(ctx context.Context, accountID int) (*model.Account, error) {
	if accountID <= 0 {
		return nil, errors.New("invalid account id")
	}

	acc, err := s.store.GetAccount(ctx, accountID)
	if err != nil {
		return nil, fmt.Errorf("get account from store: %w", err)
	}

	return acc, nil
}
