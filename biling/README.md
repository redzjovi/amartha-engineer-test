# Billing

## Getting Started

### Run unit test

```bash
go test -v ./test/
```

## System Design Overview

### Entities

1. Loan:

    - `ID`: Unique identifier.
    - `StartAt`.
    - `EndAt`.
    - `Principal`: Principal amount.
    - `InterestRate`: Annual interest rate.
    - `OutstandingAmount`: Outnstanding amount.

2. Loan Payment:

    - `ID`: Unique identifier.
    - `LoanID`.
    - `StartAt`.
    - `EndAt`.
    - `Amount`: Amount to be paid each week.
    - `PaidAt`: timestamp flag indicating if the payment has been made.

### Repository

1. LoanRepository:

    - `Create`: create the loan in the repository.
    - `FindByID`: Retrieves a loan by its ID.
    - `Save`: Saves or updates the loan in the repository.

2. LoanPaymentRepository:

    - `Create`: create the loan payment in the repository.
    - `ListOutstandingByLoanID`: Retrieves a list outstanding loan payment by ID.
    - `Save`: Saves or updates the loan payment in the repository.

Implementation: Uses an in-memory map for storage, suitable for simple use cases. For a production system, you might replace this with a database.

### Service

1. LoanService:

    - `GetOutstanding`: Retrive remaining outstanding amount.
    - `IsDelinquent`: Determines if the borrower is delinquent based on missed payments.
    - `MakePayment`: Processes a payment, marking the corresponding Payment as paid if valid and substract outstanding amount.

Implementation: The service methods use the repository to fetch loan data, apply business rules, and update the repository accordingly.
