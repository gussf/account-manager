package account

import (
	"account-manager/domain/model"
	"account-manager/repository"
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

func (s *Service) CreateAccount(req model.CreateAccountRequest) (*model.Account, error) {
	err := req.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate create account request: %w", err)
	}

	createdAcc, err := s.store.CreateAccount(req)
	if err != nil {
		return nil, fmt.Errorf("create account in store: %w", err)
	}

	return createdAcc, nil
}

func (s *Service) GetAccount(accountID string) (*model.Account, error) {
	if accountID == "" {
		return nil, errors.New("invalid account id")
	}

	acc, err := s.store.GetAccount(accountID)
	if err != nil {
		return nil, fmt.Errorf("get account from store: %w", err)
	}

	return acc, nil
}
