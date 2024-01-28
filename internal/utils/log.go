package utils

import (
	"encoding/binary"
	"fmt"
	"os"
	"proto/internal/models"
)

func ParseHeader(f *os.File) (*models.TransactionLogHeader, error) {
	rawHeader := make([]models.RawTransactionLogHeader, 1)
	if err := binary.Read(f, binary.BigEndian, &rawHeader); err != nil {
		return nil, err
	}

	return rawHeader[0].ConvertRecord(), nil
}

func ParseTransactions(f *os.File, transactionLogHeader *models.TransactionLogHeader) ([]*models.TransactionRecord, error) {
	var transactions []*models.TransactionRecord
	for i := 0; i < int(transactionLogHeader.RecordCount); i++ {
		rawRecordType := make([]byte, 1)
		if err := binary.Read(f, binary.BigEndian, &rawRecordType); err != nil {
			return nil, err
		}

		recordType := models.RECORD_TYPES[int(rawRecordType[0])]
		switch recordType {
		case models.Credit, models.Debit:
			rawTransaction := make([]models.RawCreditDebitTransactionRecord, 1)
			if err := binary.Read(f, binary.BigEndian, &rawTransaction); err != nil {
				return nil, err
			}
			transactions = append(transactions, rawTransaction[0].ConvertRecord(recordType))
		case models.StartAutopay, models.EndAutopay:
			rawTransaction := make([]models.RawAutoPayTransactionRecord, 1)
			if err := binary.Read(f, binary.BigEndian, &rawTransaction); err != nil {
				return nil, err
			}
			transactions = append(transactions, rawTransaction[0].ConvertRecord(recordType))
		default:
			return nil, fmt.Errorf("Unrecognized Record Type")
		}
	}
	return transactions, nil
}
