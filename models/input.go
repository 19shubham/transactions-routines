package models

// CreateAccountRequest - it will contain information of creating a document.
type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number"` // DocumentNumber - no. of document for account
}

// CreateTxnRequest - it will contain information of creation a txn.
type CreateTxnRequest struct {
	AccountID       int     `json:"account_id"`        // AccountID - id of account which txn belongs to.
	OperationTypeID int     `json:"operation_type_id"` // OperationTypeID - id of type of operation like credit or debit.
	Amount          float64 `json:"amount"`            // Amount - amt of txn happened n needed to be saved.
}
