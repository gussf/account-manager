package account

import (
	"account-manager/domain/core"
	"account-manager/repository"
	"context"
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

func (s *Service) CreateAccount(ctx context.Context, req core.CreateAccountRequest) (*core.Account, error) {
	err := req.Validate()
	if err != nil {
		return nil, fmt.Errorf("%w: create account request: %s", core.ErrValidation, err)
	}

	createdAcc, err := s.store.CreateAccount(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("create account in store: %w", err)
	}

	return createdAcc, nil
}

func (s *Service) GetAccount(ctx context.Context, accountID int) (*core.Account, error) {
	if accountID <= 0 {
		return nil, fmt.Errorf("%w: invalid account id", core.ErrValidation)
	}

	acc, err := s.store.GetAccount(ctx, accountID)
	if err != nil {
		return nil, fmt.Errorf("get account from store: %w", err)
	}

	return acc, nil
}
