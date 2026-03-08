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
	PurchaseOperationType:                 &DebitStrategy{operation: PurchaseOperationName},
	PurchaseWithInstallmentsOperationType: &DebitStrategy{operation: PurchaseWithInstallmentsOperationName},
	WithdrawalOperationType:               &DebitStrategy{operation: WithdrawalOperationName},
	CreditVoucherOperationType:            &CreditStrategy{operation: CreditVoucherOperationName},
}

type OperationStrategy interface {
	Apply(*SaveTransactionRequest)
	Operation() string
}

func ResolveOperationStrategy(operationTypeID int) (OperationStrategy, error) {
	strategy, ok := operationStrategyByID[operationTypeID]
	if !ok {
		return nil, fmt.Errorf("operation type with id %d is not mapped", operationTypeID)
	}
	return strategy, nil
}

type DebitStrategy struct {
	operation string
}

func (d *DebitStrategy) Apply(req *SaveTransactionRequest) {
	if req.Amount > 0 {
		req.Amount *= -1
	}
}

func (d *DebitStrategy) Operation() string {
	return d.operation
}

type CreditStrategy struct {
	operation string
}

func (c *CreditStrategy) Apply(req *SaveTransactionRequest) {
	if req.Amount < 0 {
		req.Amount *= -1
	}
}

func (d *CreditStrategy) Operation() string {
	return d.operation
}
