package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveOperationStrategy(t *testing.T) {
	tests := []struct {
		name            string
		operationTypeID int
		wantErr         bool
		wantOperation   string
		errMsg          string
	}{
		{
			name:            "success purchase",
			operationTypeID: PurchaseOperationType,
			wantErr:         false,
			wantOperation:   PurchaseOperationName,
		},
		{
			name:            "success purchase with installments",
			operationTypeID: PurchaseWithInstallmentsOperationType,
			wantErr:         false,
			wantOperation:   PurchaseWithInstallmentsOperationName,
		},
		{
			name:            "success withdrawal",
			operationTypeID: WithdrawalOperationType,
			wantErr:         false,
			wantOperation:   WithdrawalOperationName,
		},
		{
			name:            "success credit voucher",
			operationTypeID: CreditVoucherOperationType,
			wantErr:         false,
			wantOperation:   CreditVoucherOperationName,
		},
		{
			name:            "should fail when unknown operation type",
			operationTypeID: 100,
			wantErr:         true,
			errMsg:          "operation type with id 100 is not mapped",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strategy, err := DecideOperationStrategy(tt.operationTypeID)
			if tt.wantErr {
				assert.ErrorIs(t, err, ErrNotFound)
				assert.ErrorContains(t, err, tt.errMsg)
				assert.Nil(t, strategy)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantOperation, strategy.Operation())
		})
	}
}

func TestDebitStrategy_Apply(t *testing.T) {
	tests := []struct {
		name       string
		amount     float64
		wantAmount float64
	}{
		{
			name:       "positive amount becomes negative",
			amount:     100.0,
			wantAmount: -100.0,
		},
		{
			name:       "negative amount stays negative",
			amount:     -50.0,
			wantAmount: -50.0,
		},
	}

	strategy := &debitStrategy{operation: "TestDebit"}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &SaveTransactionRequest{Amount: tt.amount}
			strategy.Apply(req)
			assert.Equal(t, tt.wantAmount, req.Amount)
		})
	}
}

func TestCreditStrategy_Apply(t *testing.T) {
	tests := []struct {
		name       string
		amount     float64
		wantAmount float64
	}{
		{
			name:       "positive amount stays positive",
			amount:     100.0,
			wantAmount: 100.0,
		},
		{
			name:       "negative amount becomes positive",
			amount:     -50.0,
			wantAmount: 50.0,
		},
	}

	strategy := &creditStrategy{operation: "TestCredit"}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &SaveTransactionRequest{Amount: tt.amount}
			strategy.Apply(req)
			assert.Equal(t, tt.wantAmount, req.Amount)
		})
	}
}
