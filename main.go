package main

import (
	httpapi "account-manager/api/http"
	"account-manager/config"
	"account-manager/domain/service/account"
	"account-manager/domain/service/transaction"
	"account-manager/repository/postgresql"
	"context"
	"database/sql"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	pgStore := postgresql.NewStore(&sql.DB{}) // todo: config
	accSvc := account.NewService(pgStore)
	txSvc := transaction.NewService(pgStore)

	server := httpapi.NewServer(cfg, accSvc, txSvc)
	err = server.Start(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
