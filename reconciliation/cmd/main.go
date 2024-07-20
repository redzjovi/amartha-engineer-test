package main

import (
	"fmt"
	"os"
	"reconciliation/internal/repository"
	"reconciliation/internal/service"
	"time"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: go run main.go <system_transaction_directory> <statement_directory> <start_date> <end_date>")
		return
	}

	systemTransactionDir := os.Args[1]
	bankStatementDir := os.Args[2]
	startDate, err := time.Parse("2006-01-02", os.Args[3])
	if err != nil {
		fmt.Println("Error parse start date:", err)
		return
	}
	endDate, err := time.Parse("2006-01-02", os.Args[4])
	if err != nil {
		fmt.Println("Error parse end date:", err)
		return
	}

	transactions, err := repository.ReadSystemTransactions(systemTransactionDir)
	if err != nil {
		fmt.Println("Error reading transactions:", err)
		return
	}

	bankStatements, err := repository.ReadBankStatements(bankStatementDir)
	if err != nil {
		fmt.Println("Error reading bank statements:", err)
		return
	}

	result := service.ReconcileTransactions(transactions, bankStatements, startDate, endDate)

	fmt.Printf("Total Transactions: %d\n", result.TotalTransactions)
	fmt.Printf("Matched Transactions: %d\n", result.MatchedTransactions)
	fmt.Printf("Unmatched Transactions: %d\n", result.UnmatchedTransactions)
	fmt.Printf("Total Discrepancies: %f\n", result.TotalDiscrepancies)
}
