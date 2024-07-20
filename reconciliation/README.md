# Transaction Reconciliation Service

## Overview

This service reconciles transactions from the system against corresponding transactions in bank statements.

## Data Model

- **Transaction**
  - `trxID`: Unique identifier for the transaction (string)
  - `amount`: Transaction amount (decimal)
  - `type`: Transaction type (enum: DEBIT, CREDIT)
  - `transactionTime`: Date and time of the transaction (datetime)

- **Bank Statement**
  - `unique_identifier`: Unique identifier for the transaction in the bank statement (string)
  - `amount`: Transaction amount (decimal) (can be negative for debits)
  - `date`: Date of the transaction (date)

## Usage

### Running the service

```bash
go run cmd/main.go ./example/system-transaction/ ./example/bank-statement/ 2024-01-01 2024-06-3
```

### Run unit test

```bash
go test ./...
```
