package transaction

import (
	"account-manager/domain/core"
	"account-manager/repository"
	"context"
	"fmt"
	"log/slog"
)

type Service struct {
	store repository.Store
}

func NewService(store repository.Store) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) SaveTransaction(ctx context.Context, req core.SaveTransactionRequest) (*core.Transaction, error) {
	err := req.Validate()
	if err != nil {
		return nil, fmt.Errorf("%w: save transaction request: %s", core.ErrValidation, err)
	}

	operationStrategy, err := core.ResolveOperationStrategy(req.OperationTypeID)
	if err != nil {
		return nil, fmt.Errorf("resolve operation strategy for transaction: %w", err)
	}

	slog.Info("applying operation strategy", "operation", operationStrategy.Operation())
	operationStrategy.Apply(&req)

	createdTx, err := s.store.SaveTransaction(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("save transaction to store: %w", err)
	}

	return createdTx, nil
}
