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

func ReadBankStatements(dirname string) (bankStatements []entity.BankStatement, err error) {
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
			var bankStatement entity.BankStatement
			bankStatement.UniqueIdentifier = record[headerMap["unique_identifier"]]
			bankStatement.Amount, _ = strconv.ParseFloat(record[headerMap["amount"]], 64)
			bankStatement.Date, _ = time.Parse("2006-01-02", record[headerMap["date"]])
			bankStatements = append(bankStatements, bankStatement)
		}

		f.Close()
	}

	return bankStatements, nil
}
