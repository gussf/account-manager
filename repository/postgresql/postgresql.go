package postgresql

import (
	"account-manager/config"
	"account-manager/domain/core"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func NewStore(cfg config.Config) (*Store, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DatabaseHostname, cfg.DatabasePort, cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseName,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	store := &Store{
		db: db,
	}

	err = store.runMigrations()
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return store, nil
}

func (s *Store) DB() *sql.DB {
	return s.db
}

func (s *Store) CreateAccount(ctx context.Context, req core.CreateAccountRequest) (*core.Account, error) {
	var id int
	query := "INSERT INTO accounts (document_number) VALUES ($1) RETURNING id"
	err := s.db.QueryRowContext(ctx, query, req.DocumentNumber).Scan(&id)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			switch pgErr.Code {
			case "23505":
				return nil, fmt.Errorf("%w: account document number conflict: %s", core.ErrAlreadyExists, err)
			}
		}

		return nil, fmt.Errorf("failed to exec query: %w", err)
	}

	return &core.Account{
		ID:             id,
		DocumentNumber: req.DocumentNumber,
	}, nil
}

func (s *Store) GetAccount(ctx context.Context, accountID int) (*core.Account, error) {
	var foundID int
	var foundDocNumber string
	query := "SELECT id, document_number FROM accounts WHERE id=$1"
	err := s.db.QueryRowContext(ctx, query, accountID).Scan(&foundID, &foundDocNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w: account %d not found", core.ErrNotFound, accountID)
		}

		return nil, fmt.Errorf("failed to exec query: %w", err)
	}

	return &core.Account{
		ID:             foundID,
		DocumentNumber: foundDocNumber,
	}, nil
}

func (s *Store) SaveTransaction(ctx context.Context, req core.SaveTransactionRequest) (*core.Transaction, error) {
	var createdID int
	var createdEventDate time.Time
	var createdAmount float64
	query := "INSERT INTO transactions (account_id, operation_type_id, amount, event_date) VALUES ($1, $2, $3, now()) RETURNING id, amount, event_date"
	err := s.db.QueryRowContext(ctx, query, req.AccountID, req.OperationTypeID, req.Amount).Scan(&createdID, &createdAmount, &createdEventDate)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			switch pgErr.Code {
			case "23503":
				return nil, fmt.Errorf("%w: account id or operation type id not found: %s", core.ErrNotFound, err)
			}
		}

		return nil, fmt.Errorf("failed to exec query: %w", err)
	}

	return &core.Transaction{
		ID:              createdID,
		AccountID:       req.AccountID,
		OperationTypeID: req.OperationTypeID,
		Amount:          createdAmount,
		EventDate:       createdEventDate,
	}, nil
}
