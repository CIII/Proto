package utils

import (
	"os"
	"proto/internal/models"
	"reflect"
	"testing"
)

const FILEPATH = "../../txnlog.dat"

func TestParseHeader(t *testing.T) {
	f, err := os.Open(FILEPATH)
	if err != nil {
		t.Errorf("Failed to open test file: %v", err)
	}
	defer f.Close()

	got, err := ParseHeader(f)
	if err != nil {
		t.Errorf("Failed to parse header from test file")
	}

	want := &models.TransactionLogHeader{
		MagicString: "MPS7",
		Version:     1,
		RecordCount: 71,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}

func TestParseTransactions(t *testing.T) {
	f, err := os.Open(FILEPATH)
	if err != nil {
		t.Errorf("Failed to open test file: %v", err)
	}
	defer f.Close()

	header, err := ParseHeader(f)
	if err != nil {
		t.Errorf("Failed to parse header from test file")
	}

	got, err := ParseTransactions(f, header)
	if err != nil {
		t.Errorf("Failed to parse transactions from test file")
	}

	want := 71

	if len(got) != want {
		t.Errorf("got %d, wanted %d", len(got), want)
	}
}
