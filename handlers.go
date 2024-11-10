package main

import (
	"net/http"
	"strconv"
	
	"github.com/labstack/echo/v4"
	
	"transactions-routines/models"
)

// CreateAccountHandler handles the creation of a new account.
func CreateAccountHandler(c echo.Context) error {
	var request models.CreateAccountRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	err := CreateAccount(request.DocumentNumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, "Success")
}

// GetAccountHandler retrieves account details by account ID.
func GetAccountHandler(c echo.Context) error {
	accountID, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid account ID"})
	}
	
	account, err := GetAccount(accountID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Account not found"})
	}
	
	return c.JSON(http.StatusOK, account)
}

// CreateTransactionHandler creates a new transaction for a specified account.
func CreateTransactionHandler(c echo.Context) error {
	var request models.CreateTxnRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	
	err := CreateTransaction(request.AccountID, request.OperationTypeID, request.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	return c.JSON(http.StatusCreated, "Success")
}
