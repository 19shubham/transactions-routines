package main

import (
	"log"
	
	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Echo instance
	e := echo.New()
	
	// Routes for account and transaction management
	e.POST("/accounts", CreateAccountHandler)         // Create account
	e.GET("/accounts/:accountId", GetAccountHandler)  // Retrieve account by ID
	e.POST("/transactions", CreateTransactionHandler) // Create transaction
	
	// Start the server on port 8080
	log.Println("Starting server on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
