package models

import (
	"time"
)

// Transaction - struct will hold information for one transaction.
type Transaction struct {
	ID              string    `json:"Transaction_ID"`   // ID - id of transaction. it will be 1,2..
	AccountID       int       `json:"Account_ID"`       // AccountID - id of account for which txn is done.
	OperationTypeID int       `json:"OperationType_ID"` // OperationTypeID - type of operation done in this txn.
	Amount          float64   `json:"Amount"`           // Amount - amt involved in txn can be negative as well.
	EventDate       time.Time `json:"EventDate"`        // EventDate - date time stamp of txn record.
}
