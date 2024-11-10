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
func CreateAccount(documentNumber string) *models.Account {
	accountMutex.Lock()
	defer accountMutex.Unlock()
	
	account := &models.Account{
		ID:             accountIDCounter,
		DocumentNumber: documentNumber,
	}
	accountsData[accountIDCounter] = account
	accountIDCounter = accountIDCounter + 1
	log.Printf("Account created: %+v\n", account)
	return account
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
