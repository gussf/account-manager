package main

import (
	httpapi "account-manager/api/http"
	"account-manager/config"
	"account-manager/domain/service/account"
	"account-manager/domain/service/transaction"
	"account-manager/repository/postgresql"
	"context"
	"log/slog"
	"os"
)

func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.Load()
	if err != nil {
		slog.Error("init: failed to load configuration", "error", err)
		os.Exit(1)
	}

	pgStore, err := postgresql.NewStore(*cfg)
	if err != nil {
		slog.Error("init: failed to start postgres store", "error", err)
		os.Exit(1)
	}

	accSvc := account.NewService(pgStore)
	txSvc := transaction.NewService(pgStore)

	server := httpapi.NewServer(*cfg, accSvc, txSvc)
	err = server.Start(ctx)
	if err != nil {
		slog.Error("init: failed to start http server", "error", err)
		os.Exit(1)
	}
}
