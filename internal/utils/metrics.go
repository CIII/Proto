package utils

import (
	"fmt"
	"proto/internal/models"
)

func GetMetricsByUserId(transactions []*models.TransactionRecord, userId uint64) (*models.Metrics, error) {
	var (
		totalCredit, totalDebit              = 0.0, 0.0
		autopaysStarted, autopaysEnded int64 = 0, 0
	)
	for _, transaction := range transactions {
		if transaction.UserId == userId {
			switch transaction.RecordType {
			case models.Credit:
				totalCredit += transaction.Amount
			case models.Debit:
				totalDebit += transaction.Amount
			case models.StartAutopay:
				autopaysStarted += 1
			case models.EndAutopay:
				autopaysEnded += 1
			default:
				return nil, fmt.Errorf("Unrecognized Record Type")
			}
		}
	}

	metrics := models.Metrics{
		TotalCreditAmount: totalCredit,
		TotalDebitAmount:  totalDebit,
		AutopaysStarted:   autopaysStarted,
		AutopaysEnded:     autopaysEnded,
		Balance:           (totalCredit - totalDebit),
	}

	return &metrics, nil
}
