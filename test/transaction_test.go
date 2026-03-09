package test

import (
	apihttp "account-manager/api/http"
	"account-manager/domain/core"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/stretchr/testify/assert"
)

func (s *IntegrationSuite) Test_SaveTransaction() {
	testAccount := s.CreateRandomAccount()

	tests := []struct {
		name                 string
		req                  apihttp.SaveTransactionRequest
		body                 string
		wantStatus           int
		wantSavedTransaction bool
	}{
		{
			name: "should save purchase transaction and register with negative amount",
			req: apihttp.SaveTransactionRequest{
				AccountID:       testAccount.ID,
				OperationTypeID: core.PurchaseOperationType,
				Amount:          50.0,
			},
			wantStatus:           http.StatusCreated,
			wantSavedTransaction: true,
		},
		{
			name: "should save purchase with installments transaction with installments and register with negative amount",
			req: apihttp.SaveTransactionRequest{
				AccountID:       testAccount.ID,
				OperationTypeID: core.PurchaseWithInstallmentsOperationType,
				Amount:          30.45,
			},
			wantStatus:           http.StatusCreated,
			wantSavedTransaction: true,
		},
		{
			name: "should save withdrawal transaction and register with negative amount",
			req: apihttp.SaveTransactionRequest{
				AccountID:       testAccount.ID,
				OperationTypeID: core.WithdrawalOperationType,
				Amount:          20.3,
			},
			wantStatus:           http.StatusCreated,
			wantSavedTransaction: true,
		},
		{
			name: "should save credit voucher transaction and register with positive amount",
			req: apihttp.SaveTransactionRequest{
				AccountID:       testAccount.ID,
				OperationTypeID: core.CreditVoucherOperationType,
				Amount:          10.50,
			},
			wantStatus:           http.StatusCreated,
			wantSavedTransaction: true,
		},
		{
			name: "should fail when account does not exist",
			req: apihttp.SaveTransactionRequest{
				AccountID:       12345,
				OperationTypeID: core.CreditVoucherOperationType,
				Amount:          1.00,
			},
			wantStatus: http.StatusNotFound,
		},
		{
			name: "should fail when operation type does not exist",
			req: apihttp.SaveTransactionRequest{
				AccountID:       testAccount.ID,
				OperationTypeID: 135,
				Amount:          1.00,
			},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			payload, err := json.Marshal(tt.req)
			assert.NoError(s.T(), err)

			resp, err := http.Post(s.server.URL+"/transactions", "application/json", bytes.NewReader(payload))
			assert.NoError(s.T(), err)
			assert.Equal(s.T(), tt.wantStatus, resp.StatusCode)

			if tt.wantSavedTransaction {
				var tx apihttp.SaveTransactionResponse
				err = json.NewDecoder(resp.Body).Decode(&tx)
				assert.NoError(s.T(), err)

				s.AssertTransactionSavedInDB(tx)
			}
		})
	}
}

func (s *IntegrationSuite) AssertTransactionSavedInDB(tx apihttp.SaveTransactionResponse) {
	var dbID, dbAccID, dbOpTypeID int
	var dbAmount float64
	var dbEventDate time.Time
	err := s.store.DB().QueryRowContext(context.Background(),
		"SELECT id, account_id, operation_type_id, amount, event_date FROM transactions WHERE id=$1", tx.ID).
		Scan(&dbID, &dbAccID, &dbOpTypeID, &dbAmount, &dbEventDate)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), tx.ID, dbID)
	assert.Equal(s.T(), tx.AccountID, dbAccID)
	assert.Equal(s.T(), tx.OperationTypeID, dbOpTypeID)
	assert.Equal(s.T(), tx.Amount, dbAmount)
	assert.True(s.T(), tx.EventDate.Equal(dbEventDate))
}

func floatPtr(f float64) *float64 {
	return &f
}
