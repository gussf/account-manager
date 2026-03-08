package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

func (s *Server) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read create account request body: %s\n", err)
		sendResponse(w, http.StatusBadRequest, "failed to read request body")
		return
	}

	var req CreateAccountRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("failed to unmarshal create account request body: %s\n", err)
		sendResponse(w, http.StatusBadRequest, "failed to unmarshal request body: "+err.Error())
		return
	}

	domainReq := req.toDomain()
	createdAcc, err := s.AccountService.CreateAccount(r.Context(), domainReq)
	if err != nil {
		log.Printf("failed to create account: %s\n", err)
		// todo: check if err is validation for badrequest, check if conflict
		sendResponse(w, http.StatusInternalServerError, "error creating account")
		return
	}

	respBody := CreateAccountResponse{
		Id:             createdAcc.Id,
		DocumentNumber: createdAcc.DocumentNumber,
	}

	sendResponse(w, http.StatusCreated, respBody)
}

func (s *Server) GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("accountId"))
	if err != nil {
		log.Printf("invalid account id in request: %s\n", err)
		sendResponse(w, http.StatusBadRequest, "account id query parameter should be a number")
		return
	}

	acc, err := s.AccountService.GetAccount(r.Context(), id)
	if err != nil {
		// todo: check if err is validation for badrequest, or notfound
		log.Printf("failed to get account: %s\n", err)
		sendResponse(w, http.StatusInternalServerError, "error getting account")
		return
	}

	respBody := GetAccountResponse{
		Id:             acc.Id,
		DocumentNumber: acc.DocumentNumber,
	}

	sendResponse(w, http.StatusOK, respBody)
}

func (s *Server) SaveTransactionHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read save transaction request body: %s\n", err)
		sendResponse(w, http.StatusBadRequest, "failed to read request body")
		return
	}

	var req SaveTransactionRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("failed to unmarshal save transaction body: %s\n", err)
		sendResponse(w, http.StatusBadRequest, "failed to unmarshal request body: "+err.Error())
		return
	}

	domainReq := req.toDomain()
	createdTx, err := s.TransactionService.SaveTransaction(r.Context(), domainReq)
	if err != nil {
		log.Printf("failed to save transaction: %s\n", err)
		// todo: check if err is validation for badrequest, check if conflict
		sendResponse(w, http.StatusInternalServerError, "error saving transaction")
		return
	}

	respBody := SaveTransactionResponse{
		Id:            createdTx.Id,
		OperationType: createdTx.OperationType,
		Amount:        createdTx.Amount,
		EventDate:     createdTx.EventDate,
	}

	sendResponse(w, http.StatusCreated, respBody)
}

func sendResponse(w http.ResponseWriter, status int, payload any) {
	w.WriteHeader(status)
	resp, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(resp)
}
