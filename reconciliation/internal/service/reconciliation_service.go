package service

import (
	"fmt"
	"math"
	"reconciliation/internal/entity"
	"time"
)

type ReconciliationResult struct {
	TotalTransactions           int
	MatchedTransactions         int
	UnmatchedTransactions       int
	MissingInBankStatements     []entity.Transaction
	MissingInSystemTransactions []entity.BankStatement
	TotalDiscrepancies          float64
}

func ReconcileTransactions(transactions []entity.Transaction, bankStatements []entity.BankStatement, startDate, endDate time.Time) ReconciliationResult {
	// Filter transactions based on date range
	filteredTransactions := filterTransactionsByDate(transactions, startDate, endDate)
	filteredBankStatements := filterBankStatementsByDate(bankStatements, startDate, endDate)

	// Perform reconciliation
	matchedTransactions, unmatchedTransactions, missingInBankStatements, missingInSystemTransactions, totalDiscrepancies := matchTransactions(filteredTransactions, filteredBankStatements)

	return ReconciliationResult{
		TotalTransactions:           len(filteredTransactions),
		MatchedTransactions:         len(matchedTransactions),
		UnmatchedTransactions:       len(unmatchedTransactions),
		MissingInBankStatements:     missingInBankStatements,
		MissingInSystemTransactions: missingInSystemTransactions,
		TotalDiscrepancies:          totalDiscrepancies,
	}
}

func filterTransactionsByDate(transactions []entity.Transaction, startDate, endDate time.Time) []entity.Transaction {
	var filtered []entity.Transaction
	for _, transaction := range transactions {
		if transaction.TransactionTime.After(startDate) && transaction.TransactionTime.Before(endDate) {
			filtered = append(filtered, transaction)
		}
	}
	return filtered
}

func filterBankStatementsByDate(bankStatements []entity.BankStatement, startDate, endDate time.Time) []entity.BankStatement {
	var filtered []entity.BankStatement
	for _, bankStatement := range bankStatements {
		if bankStatement.Date.After(startDate) && bankStatement.Date.Before(endDate) {
			filtered = append(filtered, bankStatement)
		}
	}
	return filtered
}

func matchTransactions(transactions []entity.Transaction, bankStatements []entity.BankStatement) (matchedTransactions, unmatchedTransactions []entity.Transaction, missingInBankStatements []entity.Transaction, missingInSystemTransactions []entity.BankStatement, totalDiscrepancies float64) {
	mapTransaction := make(map[string]entity.Transaction)
	mapTransactionByID := make(map[string]entity.Transaction)
	for _, transaction := range transactions {
		var amount float64
		if transaction.Type == entity.CREDIT {
			amount = transaction.Amount
		} else if transaction.Type == entity.DEBIT {
			amount = -transaction.Amount
		}
		key := fmt.Sprintf("%s-%v-%s", transaction.TrxID, math.Abs(amount), transaction.TransactionTime.Format("2006-01-02"))
		mapTransaction[key] = transaction
		mapTransactionByID[transaction.TrxID] = transaction
	}

	mapBankStatement := make(map[string]entity.BankStatement)
	mapBankStatementByID := make(map[string]entity.BankStatement)
	for _, bankStatement := range bankStatements {
		key := fmt.Sprintf("%s-%v-%s", bankStatement.UniqueIdentifier, math.Abs(bankStatement.Amount), bankStatement.Date.Format("2006-01-02"))
		mapBankStatement[key] = bankStatement
		mapBankStatementByID[bankStatement.UniqueIdentifier] = bankStatement
	}
	for key, transaction := range mapTransaction {
		if _, exists := mapBankStatement[key]; exists {
			matchedTransactions = append(matchedTransactions, transaction)
			delete(mapBankStatement, key)
		} else {
			unmatchedTransactions = append(unmatchedTransactions, transaction)
			missingInBankStatements = append(missingInBankStatements, transaction)
		}
	}

	for _, bankStatement := range mapBankStatement {
		_, exists := mapTransactionByID[bankStatement.UniqueIdentifier]
		if !exists {
			missingInSystemTransactions = append(missingInSystemTransactions, bankStatement)
		}
	}

	for _, transaction := range unmatchedTransactions {
		bankStatement, exists := mapBankStatementByID[transaction.TrxID]
		if exists {
			totalDiscrepancies += transaction.GetRealAmount() - bankStatement.Amount
		}
	}

	return
}
