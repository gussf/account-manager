package http

import (
	"fmt"
	"net/http"
)

func (s *Server) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	// todo: create struct for response
	fmt.Fprint(w, "Account created")
}

func (s *Server) GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("accountId")

	acc, err := s.AccountService.GetAccount(id)
	if err != nil {
		// todo: check if err is validation for badrequest
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to fetch: " + err.Error()))
		return
	}

	// todo: create struct for response
	fmt.Fprintf(w, "Fetching account: %s", acc.Id)
}

func (s *Server) SaveTransactionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	// todo: create struct for response
	fmt.Fprint(w, "Transaction processed")
}
