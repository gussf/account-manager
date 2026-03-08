package main

import (
	httpapi "account-manager/api/http"
	"account-manager/config"
	"account-manager/domain/service/account"
	"account-manager/domain/service/transaction"
	"account-manager/repository/postgresql"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("init: failed to load configuration: %s", err)
	}

	pgStore, err := postgresql.NewStore(*cfg)
	if err != nil {
		log.Fatalf("init: failed to start postgres store: %s", err)
	}

	accSvc := account.NewService(pgStore)
	txSvc := transaction.NewService(pgStore)

	server := httpapi.NewServer(*cfg, accSvc, txSvc)
	err = server.Start(ctx)
	if err != nil {
		log.Fatalf("init: failed to start http server: %s", err)
	}
}
