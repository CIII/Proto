package utils

import (
	"proto/internal/models"
	"reflect"
	"testing"
)

func TestGetMetricsByUserId(t *testing.T) {
	wantUserId := uint64(1111)
	otherUserId := uint64(2222)
	transaction1 := models.TransactionRecord{
		RecordType:    models.Credit,
		UnixTimestamp: 0,
		UserId:        wantUserId,
		Amount:        100,
	}

	transaction2 := models.TransactionRecord{
		RecordType:    models.Credit,
		UnixTimestamp: 0,
		UserId:        wantUserId,
		Amount:        50,
	}

	transaction3 := models.TransactionRecord{
		RecordType:    models.Debit,
		UnixTimestamp: 0,
		UserId:        wantUserId,
		Amount:        25,
	}

	transaction4 := models.TransactionRecord{
		RecordType:    models.Credit,
		UnixTimestamp: 0,
		UserId:        otherUserId,
		Amount:        500,
	}

	transaction5 := models.TransactionRecord{
		RecordType:    models.StartAutopay,
		UnixTimestamp: 0,
		UserId:        wantUserId,
	}

	transaction6 := models.TransactionRecord{
		RecordType:    models.EndAutopay,
		UnixTimestamp: 0,
		UserId:        wantUserId,
	}

	transaction7 := models.TransactionRecord{
		RecordType:    models.EndAutopay,
		UnixTimestamp: 0,
		UserId:        otherUserId,
	}

	transactions := []*models.TransactionRecord{&transaction1, &transaction2, &transaction3, &transaction4, &transaction5, &transaction6, &transaction7}
	got, err := GetMetricsByUserId(transactions, 1111)
	if err != nil {
		t.Errorf("Error running test: %v", err)
	}

	want := &models.Metrics{
		TotalCreditAmount: 150,
		TotalDebitAmount:  25,
		AutopaysStarted:   1,
		AutopaysEnded:     1,
		Balance:           125,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}
