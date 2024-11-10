package main

import (
	"fmt"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

// TestCreateAccount ensures that the CreateAccount function creates a new account.
func TestCreateAccount(t *testing.T) {
	err := CreateAccount("12345678900")
	if err != nil {
		fmt.Println("error in creating account")
		return
	}
	fmt.Println("success")
}

// TestCreateAccountHandler ensures that the GetAccount function fetches the given account.
func TestGetAccount(t *testing.T) {
	err := CreateAccount("12345678900")
	if err != nil {
		fmt.Println("error in creating account")
		return
	}
	account, err := GetAccount(1)
	if err != nil {
		fmt.Println("error in fetching account")
		return
	}
	fmt.Println("success")
	assert.Equal(t, account.DocumentNumber, "12345678900")
}
