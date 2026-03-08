package transaction

import (
	"account-manager/domain/model"
	"account-manager/repository"
	"context"
	"fmt"
)

type Service struct {
	store repository.Store
}

func NewService(store repository.Store) *Service {
	return &Service{}
}

func (s *Service) SaveTransaction(ctx context.Context, req model.SaveTransactionRequest) (*model.Transaction, error) {
	err := req.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate save transaction request: %w", err)
	}

	createdTx, err := s.store.SaveTransaction(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("save transaction to store: %w", err)
	}

	return createdTx, nil
}
