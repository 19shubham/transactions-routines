package models

import (
	"errors"
)

type OperationType int

const (
	NormalPurchase      OperationType = 1
	InstallmentPurchase OperationType = 2
	Withdrawal          OperationType = 3
	CreditVoucher       OperationType = 4
)

// Int - method will convert and return int value of enum operation type.
func (ot OperationType) Int() int {
	return int(ot)
}

// GetDescription - method will return description from int val of operation type and error if none of the type matches.
func (ot OperationType) GetDescription() (string, error) {
	switch ot {
	case 1:
		return "Normal Purchase", nil
	case 2:
		return "Purchase with installments", nil
	case 3:
		return "Withdrawal", nil
	case 4:
		return "Credit Voucher", nil
	default:
		return "", errors.New("invalid operation type")
	}
}
