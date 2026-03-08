package main

import (
	httpapi "account-manager/api/http"
	"account-manager/domain/service/account"
	"account-manager/domain/service/transaction"
	"account-manager/repository/postgresql"
	"context"
	"database/sql"
	"log"
)

func main() {
	pgStore := postgresql.NewStore(&sql.DB{}) // todo: config
	accSvc := account.NewService(pgStore)
	txSvc := transaction.NewService(pgStore)

	server := httpapi.NewServer(accSvc, txSvc)
	err := server.Start(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
