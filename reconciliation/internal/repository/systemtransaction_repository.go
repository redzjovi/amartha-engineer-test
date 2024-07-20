package repository

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"reconciliation/internal/entity"
)

func ReadSystemTransactions(dirname string) (transactions []entity.Transaction, err error) {
	files, err := os.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".csv") {
			continue
		}

		filePath := filepath.Join(dirname, file.Name())

		f, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}

		r := csv.NewReader(f)
		header, err := r.Read()
		if err != nil {
			return nil, err
		}
		headerMap := make(map[string]int)
		for i, col := range header {
			headerMap[col] = i
		}

		records, err := r.ReadAll()
		if err != nil {
			return nil, err
		}

		for _, record := range records {
			var transaction entity.Transaction
			transaction.TrxID = record[headerMap["trxID"]]
			transaction.Amount, _ = strconv.ParseFloat(record[headerMap["amount"]], 64)
			transaction.Type = entity.TransactionType(record[headerMap["type"]])
			transaction.TransactionTime, _ = time.Parse(time.RFC3339, record[headerMap["transactionTime"]])
			transactions = append(transactions, transaction)
		}

		f.Close()
	}

	return transactions, nil
}
