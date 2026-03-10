package core

import "fmt"

const (
	PurchaseOperationType                 = 1
	PurchaseWithInstallmentsOperationType = 2
	WithdrawalOperationType               = 3
	CreditVoucherOperationType            = 4
)

const (
	PurchaseOperationName                 = "Purchase"
	PurchaseWithInstallmentsOperationName = "PurchaseWithInstallments"
	WithdrawalOperationName               = "Withdrawal"
	CreditVoucherOperationName            = "CreditVoucher"
)

var operationStrategyByID = map[int]OperationStrategy{
	PurchaseOperationType:                 &debitStrategy{operation: PurchaseOperationName},
	PurchaseWithInstallmentsOperationType: &debitStrategy{operation: PurchaseWithInstallmentsOperationName},
	WithdrawalOperationType:               &debitStrategy{operation: WithdrawalOperationName},
	CreditVoucherOperationType:            &creditStrategy{operation: CreditVoucherOperationName},
}

type OperationStrategy interface {
	Apply(*SaveTransactionRequest)
	Operation() string
}

func DecideOperationStrategy(operationTypeID int) (OperationStrategy, error) {
	strategy, ok := operationStrategyByID[operationTypeID]
	if !ok {
		return nil, fmt.Errorf("%w: operation type with id %d is not mapped", ErrNotFound, operationTypeID)
	}
	return strategy, nil
}

type debitStrategy struct {
	operation string
}

func (d *debitStrategy) Apply(req *SaveTransactionRequest) {
	if req.Amount > 0 {
		req.Amount *= -1
	}
}

func (d *debitStrategy) Operation() string {
	return d.operation
}

type creditStrategy struct {
	operation string
}

func (c *creditStrategy) Apply(req *SaveTransactionRequest) {
	if req.Amount < 0 {
		req.Amount *= -1
	}
}

func (d *creditStrategy) Operation() string {
	return d.operation
}
