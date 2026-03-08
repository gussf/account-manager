package model

type OperationType int

const (
	NormalPurchase OperationType = iota
	PurchaseWithInstallments
	Withdrawal
	CreditVoucher
)
