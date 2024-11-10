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

// ConvertToType - function will convert and return int value of enum operation type to enum value.
func ConvertToType(oType int) OperationType {
	return OperationType(oType)
}

// IsValid - function will return if transaction operation type is valid or not.
func IsValid(oType int) bool {
	switch ConvertToType(oType) {
	case 1, 2, 3, 4:
		return true
	default:
		return false
	}
}
