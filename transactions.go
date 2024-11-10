package main

import (
	"errors"
	"log"
	"sync"
	"time"
	
	"transactions-routines/models"
)

// Global variables for in-memory storage of transaction data and transaction mutex.
var (
	transactionsData     = make(map[int]*models.Transaction)
	transactionIDCounter = 1
	transactionMutex     = &sync.Mutex{}
)

// CreateTransaction creates n stores transaction associated with a specific account and operation type and adds the
// transaction details in txnData
func CreateTransaction(accountId, operationTypeId int, amt float64) error {
	transactionMutex.Lock()
	defer transactionMutex.Unlock()
	
	// Checking if the account exists
	if _, ok := accountsData[accountId]; !ok {
		return errors.New("account not found")
	}
	
	// Checking if operationTypeId is valid or not.
	if ok := models.IsValid(operationTypeId); !ok {
		return errors.New("invalid operation type")
	}
	
	transactionsData[transactionIDCounter] = &models.Transaction{
		ID:              transactionIDCounter,
		AccountID:       accountId,
		OperationTypeID: operationTypeId,
		Amount:          amt,
		EventDate:       time.Now(),
	}
	transactionIDCounter = transactionIDCounter + 1
	log.Print("Transaction created")
	return nil
}
