package http

import (
	"account-manager/config"
	"account-manager/domain/core"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	handler            http.Handler
	AccountService     core.AccountService
	TransactionService core.TransactionService
	requestTimeout     time.Duration
	port               int
}

func NewServer(cfg config.Config, accSvc core.AccountService, txSvc core.TransactionService) *Server {
	sv := &Server{
		AccountService:     accSvc,
		TransactionService: txSvc,
		requestTimeout:     cfg.HTTPServerRequestTimeout,
		port:               cfg.HTTPServerPort,
	}

	mux := http.NewServeMux()
	sv.registerRoutes(mux)
	sv.handler = http.TimeoutHandler(mux, sv.requestTimeout, "request timeout")

	return sv
}

func (s *Server) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /accounts", s.CreateAccountHandler)
	mux.HandleFunc("GET /accounts/{accountId}", s.GetAccountHandler)
	mux.HandleFunc("POST /transactions", s.SaveTransactionHandler)
}

func (s *Server) Start(ctx context.Context) error {
	addr := fmt.Sprintf(":%d", s.port)
	log.Printf("server running on %s\n", addr)
	err := http.ListenAndServe(addr, s.handler)
	if err != nil {
		return err
	}

	return nil
}
