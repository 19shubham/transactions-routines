package models

// Account - struct will hold information for one user account.
type Account struct {
	ID             int    `json:"Account_ID"`      // ID of account, it will increment by +1 with each new account.
	DocumentNumber string `json:"Document_Number"` // DocumentNumber of account, represent no. of document.
}
