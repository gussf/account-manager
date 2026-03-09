package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveTransactionRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     SaveTransactionRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "success",
			req: SaveTransactionRequest{
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          15.00,
			},
			wantErr: false,
		},
		{
			name: "should fail when invalid account id",
			req: SaveTransactionRequest{
				AccountID:       -1,
				OperationTypeID: 1,
				Amount:          15.00,
			},
			wantErr: true,
			errMsg:  "invalid account id",
		},
		{
			name: "should fail when invalid amount",
			req: SaveTransactionRequest{
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          0,
			},
			wantErr: true,
			errMsg:  "amount cannot be zero",
		},
		{
			name: "should fail when invalid operation type id",
			req: SaveTransactionRequest{
				AccountID:       1,
				OperationTypeID: -1,
				Amount:          10,
			},
			wantErr: true,
			errMsg:  "invalid operation type id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if tt.wantErr {
				assert.ErrorContains(t, err, tt.errMsg)
			}
		})
	}

}
