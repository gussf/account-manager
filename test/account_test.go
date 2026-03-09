package test

import (
	apihttp "account-manager/api/http"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stretchr/testify/assert"
)

func (s *IntegrationSuite) Test_CreateAccount() {
	tests := []struct {
		name               string
		req                apihttp.CreateAccountRequest
		wantStatus         int
		wantCreatedAccount bool
	}{
		{
			name: "should create account",
			req: apihttp.CreateAccountRequest{
				DocumentNumber: "existing-doc",
			},
			wantStatus:         http.StatusCreated,
			wantCreatedAccount: true,
		},
		{
			name: "should fail when document number already exists (using test above)",
			req: apihttp.CreateAccountRequest{
				DocumentNumber: "existing-doc",
			},
			wantStatus:         http.StatusConflict,
			wantCreatedAccount: false,
		},

		{
			name: "should fail when document number is empty",
			req: apihttp.CreateAccountRequest{
				DocumentNumber: "",
			},
			wantStatus:         http.StatusBadRequest,
			wantCreatedAccount: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			payload, err := json.Marshal(tt.req)
			assert.NoError(s.T(), err)

			resp, err := http.Post(s.server.URL+"/accounts", "application/json", bytes.NewReader(payload))
			assert.NoError(s.T(), err)
			assert.Equal(s.T(), tt.wantStatus, resp.StatusCode)

			if tt.wantCreatedAccount {
				var acc apihttp.CreateAccountResponse
				err = json.NewDecoder(resp.Body).Decode(&acc)
				assert.NoError(s.T(), err)

				s.AssertAccountSavedInDB(acc)
			}
		})
	}
}

func (s *IntegrationSuite) Test_GetAccount() {
	preexistingAcc := s.CreateRandomAccount()

	tests := []struct {
		name        string
		accountID   int
		wantStatus  int
		wantAccount *apihttp.GetAccountResponse
	}{
		{
			name:       "should get account",
			accountID:  preexistingAcc.ID,
			wantStatus: http.StatusOK,
			wantAccount: &apihttp.GetAccountResponse{
				ID:             preexistingAcc.ID,
				DocumentNumber: preexistingAcc.DocumentNumber,
			},
		},
		{
			name:       "should fail when account doesnt exist",
			accountID:  12394174,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "should fail when account id is invalid",
			accountID:  -1,
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			resp, err := http.Get(s.server.URL + "/accounts/" + fmt.Sprintf("%d", tt.accountID))
			assert.NoError(s.T(), err)
			assert.Equal(s.T(), tt.wantStatus, resp.StatusCode)

			if tt.wantAccount != nil {
				var acc apihttp.GetAccountResponse
				err = json.NewDecoder(resp.Body).Decode(&acc)
				assert.NoError(s.T(), err)

				assert.Equal(s.T(), preexistingAcc.ID, acc.ID)
				assert.Equal(s.T(), preexistingAcc.DocumentNumber, acc.DocumentNumber)
			}
		})
	}
}

func (s *IntegrationSuite) AssertAccountSavedInDB(acc apihttp.CreateAccountResponse) {
	var dbID int
	var dbDocNum string
	err := s.store.DB().QueryRowContext(context.Background(),
		"SELECT id, document_number FROM accounts WHERE id=$1", acc.ID).
		Scan(&dbID, &dbDocNum)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), acc.ID, dbID)
	assert.Equal(s.T(), acc.DocumentNumber, dbDocNum)
}
