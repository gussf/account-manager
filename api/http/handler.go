package http

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"
)

func (s *Server) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("received create account request")
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("failed to read create account request body", "error", err, "status", http.StatusBadRequest)
		sendResponse(w, http.StatusBadRequest, errorMessage("failed to read request body"))
		return
	}

	var req CreateAccountRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		slog.Error("failed to unmarshal create account request body", "error", err, "status", http.StatusBadRequest)
		sendResponse(w, http.StatusBadRequest, errorMessage("failed to unmarshal request body: "+err.Error()))
		return
	}

	domainReq := req.toDomain()
	createdAcc, err := s.AccountService.CreateAccount(r.Context(), domainReq)
	if err != nil {
		slog.Error("failed to create account", "error", err, "status", http.StatusInternalServerError)
		// todo: check if err is validation for badrequest, check if conflict
		sendResponse(w, http.StatusInternalServerError, errorMessage("error creating account"))
		return
	}

	respBody := CreateAccountResponse{
		ID:             createdAcc.ID,
		DocumentNumber: createdAcc.DocumentNumber,
	}

	slog.Info("create account response", "body", respBody, "status", http.StatusCreated)
	sendResponse(w, http.StatusCreated, respBody)
}

func (s *Server) GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("received get account request")
	id, err := strconv.Atoi(r.PathValue("accountId"))
	if err != nil {
		slog.Error("invalid account id in request", "error", err, "status", http.StatusBadRequest)
		sendResponse(w, http.StatusBadRequest, errorMessage("account id query parameter should be a number"))
		return
	}

	acc, err := s.AccountService.GetAccount(r.Context(), id)
	if err != nil {
		// todo: check if err is validation for badrequest, or notfound
		slog.Error("failed to get account", "error", err, "status", http.StatusInternalServerError)
		sendResponse(w, http.StatusInternalServerError, errorMessage("error getting account"))
		return
	}

	respBody := GetAccountResponse{
		ID:             acc.ID,
		DocumentNumber: acc.DocumentNumber,
	}

	slog.Info("get account response", "body", respBody, "status", http.StatusOK)
	sendResponse(w, http.StatusOK, respBody)
}

func (s *Server) SaveTransactionHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("received save transaction request")
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("failed to read save transaction request body", "error", err, "status", http.StatusBadRequest)
		sendResponse(w, http.StatusBadRequest, errorMessage("failed to read request body"))
		return
	}

	var req SaveTransactionRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		slog.Error("failed to unmarshal save transaction body", "error", err, "status", http.StatusBadRequest)
		sendResponse(w, http.StatusBadRequest, errorMessage("failed to unmarshal request body: "+err.Error()))
		return
	}

	domainReq := req.toDomain()
	createdTx, err := s.TransactionService.SaveTransaction(r.Context(), domainReq)
	if err != nil {
		slog.Error("failed to save transaction", "error", err, "status", http.StatusInternalServerError)
		// todo: check if err is validation for badrequest, check if conflict
		sendResponse(w, http.StatusInternalServerError, errorMessage("error saving transaction"))
		return
	}

	respBody := SaveTransactionResponse{
		ID:              createdTx.ID,
		OperationTypeID: createdTx.OperationTypeID,
		Amount:          createdTx.Amount,
		EventDate:       createdTx.EventDate,
	}

	slog.Info("save transaction response", "body", respBody, "status", http.StatusCreated)
	sendResponse(w, http.StatusCreated, respBody)
}

func sendResponse(w http.ResponseWriter, status int, payload any) {
	resp, err := json.Marshal(payload)
	if err != nil {
		slog.Error("failed to marshal response", "error", err, "status", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(resp)
}

func errorMessage(msg string) ErrorMessage {
	return ErrorMessage{
		Error: msg,
	}
}
