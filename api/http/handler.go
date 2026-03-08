package http

import (
	"account-manager/domain/core"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
)

func (s *Server) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("received create account request")

	var req CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("failed to decode create account request body", "error", err, "status", http.StatusBadRequest)
		sendResponse(w, http.StatusBadRequest, errorMessage("invalid request body: "+err.Error()))
		return
	}

	domainReq := req.toDomain()
	createdAcc, err := s.AccountService.CreateAccount(r.Context(), domainReq)
	if err != nil {
		var status int
		var msg any
		switch {
		case errors.Is(err, core.ErrValidation):
			status = http.StatusBadRequest
			msg = errorMessage(err.Error())
		case errors.Is(err, core.ErrAlreadyExists):
			status = http.StatusConflict
			msg = errorMessage("account already exists")
		default:
			status = http.StatusInternalServerError
			msg = errorMessage("error creating account")
		}

		slog.Error("failed to create account", "error", err, "status", status)
		sendResponse(w, status, msg)
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
		var status int
		var msg any
		switch {
		case errors.Is(err, core.ErrValidation):
			status = http.StatusBadRequest
			msg = errorMessage(err.Error())
		case errors.Is(err, core.ErrNotFound):
			status = http.StatusNotFound
			msg = errorMessage("account not found")
		default:
			status = http.StatusInternalServerError
			msg = errorMessage("error getting account")
		}

		slog.Error("failed to get account", "error", err, "status", status)
		sendResponse(w, status, msg)
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

	var req SaveTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("failed to decode save transaction request body", "error", err, "status", http.StatusBadRequest)
		sendResponse(w, http.StatusBadRequest, errorMessage("invalid request body: "+err.Error()))
		return
	}

	domainReq := req.toDomain()
	createdTx, err := s.TransactionService.SaveTransaction(r.Context(), domainReq)
	if err != nil {
		var status int
		var msg any
		switch {
		case errors.Is(err, core.ErrValidation):
			status = http.StatusBadRequest
			msg = errorMessage(err.Error())
		case errors.Is(err, core.ErrNotFound):
			status = http.StatusNotFound
			msg = errorMessage("account id or operation type id not found")
		default:
			status = http.StatusInternalServerError
			msg = errorMessage("error saving transaction")
		}

		slog.Error("failed to save transaction", "error", err, "status", status)
		sendResponse(w, status, msg)
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
