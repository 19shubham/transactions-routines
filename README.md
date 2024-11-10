# transactions-routines
This project is a RESTful API built with Go and the Echo framework. The API provides basic functionality for managing 
user accounts and financial transactions in an in-memory store. The application is containerized with Docker.

## Features
### Account Management:
* Create a new account
* Retrieve account information
### Transaction Management:
* Create a financial transaction with various operation types, including purchase, installment purchase, withdrawal, and credit voucher.

## Prerequisites:
* Go 1.18+
* Docker

## Getting Started
1. Clone the Repository
   >git clone $repo-url$

   >cd $repo-name$"

2. Build and Run with Docker

   Build the Docker Image
   >docker build -t transactions-routines .

   Run the Docker Container
   > docker run -p 8080:8080 transactions-routines

   The API will be accessible at http://localhost:8080.

## API Documentation
1. Create Account - 
   #### Endpoint: POST /accounts
   Description: Creates a new account with a unique document number.
   Request Body:
   {
   "document_number": "12345678900"
   }
   Response:
   Success/Error 

2. Get Account by ID
   #### Endpoint: GET /accounts/:accountId
   Description: Retrieves the account information for the specified account Id.
   Param- Id is auto id of account for which information needs to be fetched.
   Response:
   {
   "account_id": 1,
   "document_number": "12345678900"
   }

3. Create Transaction
   #### Endpoint: POST /transactions
   Description: Creates a financial transaction for a specific account.
   Request Body:
   {
   "account_id": 1,
   "operation_type_id": 1,
   "amount": 100.0
   }
   Response:
   Success/Error