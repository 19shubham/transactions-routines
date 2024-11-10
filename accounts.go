package main

import (
	"errors"
	"log"
	"sync"
	
	"transactions-routines/models"
)

var (
	accountsData     = make(map[int]*models.Account)
	accountIDCounter = 1
	accountMutex     = &sync.Mutex{}
)

// CreateAccount creates a new account with the specified document number.
func CreateAccount(documentNumber string) error {
	accountMutex.Lock()
	defer accountMutex.Unlock()
	accountsData[accountIDCounter] = &models.Account{
		ID:             accountIDCounter,
		DocumentNumber: documentNumber,
	}
	accountIDCounter = accountIDCounter + 1
	log.Print("Account created successfully")
	return nil
}

// GetAccount retrieves an account by its ID. Returns an error if the account is not found.
func GetAccount(accountID int) (*models.Account, error) {
	accountMutex.Lock()
	defer accountMutex.Unlock()
	
	account, exists := accountsData[accountID]
	if !exists {
		return nil, errors.New("account not found")
	}
	return account, nil
}
