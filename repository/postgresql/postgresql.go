package postgresql

import (
	"account-manager/config"
	"account-manager/domain/model"
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
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

func (s *Store) CreateAccount(ctx context.Context, req model.CreateAccountRequest) (*model.Account, error) {
	var id int
	query := "INSERT INTO accounts (document_number) VALUES ($1) RETURNING id"
	err := s.db.QueryRowContext(ctx, query, req.DocumentNumber).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to exec query: %w", err)
	}

	return &model.Account{
		Id:             id,
		DocumentNumber: req.DocumentNumber,
	}, nil
}

func (s *Store) GetAccount(ctx context.Context, accountID int) (*model.Account, error) {
	var foundID int
	var foundDocNumber string
	query := "SELECT id, document_number FROM accounts WHERE id=$1"
	err := s.db.QueryRowContext(ctx, query, accountID).Scan(&foundID, &foundDocNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to exec query: %w", err)
	}

	return &model.Account{
		Id:             foundID,
		DocumentNumber: foundDocNumber,
	}, nil
}

func (s *Store) SaveTransaction(ctx context.Context, req model.SaveTransactionRequest) (*model.Transaction, error) {
	var id int
	var eventDate time.Time
	query := "INSERT INTO transactions (account_id, operation_type_id, amount, event_date) VALUES ($1, $2, $3, now()) RETURNING id, event_date"
	err := s.db.QueryRowContext(ctx, query, req.AccountID, req.OperationType, req.Amount).Scan(&id, &eventDate)
	if err != nil {
		return nil, fmt.Errorf("failed to exec query: %w", err)
	}

	return &model.Transaction{
		Id:            id,
		AccountId:     req.AccountID,
		OperationType: req.OperationType,
		Amount:        req.Amount,
		EventDate:     eventDate,
	}, nil
}
