package http

import (
	domain "account-manager/domain/model"
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	mux                *http.ServeMux
	AccountService     domain.AccountService
	TransactionService domain.TransactionService
}

func NewServer(accSvc domain.AccountService, txSvc domain.TransactionService) *Server {
	sv := &Server{
		mux:                http.NewServeMux(),
		AccountService:     accSvc,
		TransactionService: txSvc,
	}

	sv.registerRoutes()

	return sv
}

func (s *Server) registerRoutes() {
	s.mux.HandleFunc("POST /accounts", s.CreateAccountHandler)
	s.mux.HandleFunc("GET /accounts/{accountId}", s.GetAccountHandler)
	s.mux.HandleFunc("POST /transactions", s.SaveTransactionHandler)
}

func (s *Server) Start(ctx context.Context) error {
	// todo: config
	fmt.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", s.mux)
	if err != nil {
		return err
	}

	return nil
}
