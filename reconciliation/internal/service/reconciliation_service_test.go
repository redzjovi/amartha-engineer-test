package service

import (
	"reconciliation/internal/entity"
	"reflect"
	"testing"
	"time"
)

func TestReconcileTransactions(t *testing.T) {
	type args struct {
		transactions   []entity.Transaction
		bankStatements []entity.BankStatement
		startDate      time.Time
		endDate        time.Time
	}
	tests := []struct {
		name string
		args args
		want ReconciliationResult
	}{
		{
			name: "Match",
			args: args{
				transactions: []entity.Transaction{
					{
						TrxID:           "1",
						Amount:          100.00,
						Type:            "DEBIT",
						TransactionTime: time.Date(2024, 1, 1, 15, 0, 0, 0, time.UTC),
					},
				},
				bankStatements: []entity.BankStatement{
					{
						UniqueIdentifier: "1",
						Amount:           -100.00,
						Date:             time.Date(2024, 1, 1, 16, 0, 0, 0, time.UTC),
					},
				},
				startDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				endDate:   time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC),
			},
			want: ReconciliationResult{
				TotalTransactions:           1,
				MatchedTransactions:         1,
				UnmatchedTransactions:       0,
				MissingInBankStatements:     nil,
				MissingInSystemTransactions: nil,
				TotalDiscrepancies:          0,
			},
		},
		{
			name: "Unmatch",
			args: args{
				transactions: []entity.Transaction{
					{
						TrxID:           "1",
						Amount:          100.00,
						Type:            "DEBIT",
						TransactionTime: time.Date(2024, 1, 1, 15, 0, 0, 0, time.UTC),
					},
				},
				bankStatements: []entity.BankStatement{
					{
						UniqueIdentifier: "1",
						Amount:           -200.00,
						Date:             time.Date(2024, 1, 2, 16, 0, 0, 0, time.UTC),
					},
				},
				startDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				endDate:   time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC),
			},
			want: ReconciliationResult{
				TotalTransactions:     1,
				MatchedTransactions:   0,
				UnmatchedTransactions: 1,
				MissingInBankStatements: []entity.Transaction{
					{
						TrxID:           "1",
						Amount:          100.00,
						Type:            "DEBIT",
						TransactionTime: time.Date(2024, 1, 1, 15, 0, 0, 0, time.UTC),
					},
				},
				MissingInSystemTransactions: nil,
				TotalDiscrepancies:          100,
			},
		},
		{
			name: "Missing in system transaction",
			args: args{
				transactions: []entity.Transaction{},
				bankStatements: []entity.BankStatement{
					{
						UniqueIdentifier: "1",
						Amount:           -100.00,
						Date:             time.Date(2024, 1, 2, 16, 0, 0, 0, time.UTC),
					},
				},
				startDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				endDate:   time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC),
			},
			want: ReconciliationResult{
				TotalTransactions:       0,
				MatchedTransactions:     0,
				UnmatchedTransactions:   0,
				MissingInBankStatements: nil,
				MissingInSystemTransactions: []entity.BankStatement{
					{
						UniqueIdentifier: "1",
						Amount:           -100.00,
						Date:             time.Date(2024, 1, 2, 16, 0, 0, 0, time.UTC),
					},
				},
				TotalDiscrepancies: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReconcileTransactions(tt.args.transactions, tt.args.bankStatements, tt.args.startDate, tt.args.endDate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReconcileTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}
