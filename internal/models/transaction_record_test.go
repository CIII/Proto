package models

import (
	"encoding/binary"
	"reflect"
	"testing"
)

func TestRawCreditDebitTransactionRecordConvertRecord(t *testing.T) {
	magicStringBytes := make([]byte, 4)
	copy(magicStringBytes, "MPS7")
	versionBytes := []byte{2}
	recordCountBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(recordCountBytes, 100)
	record := RawTransactionLogHeader{
		MagicString: [4]byte(magicStringBytes),
		Version:     [1]byte(versionBytes),
		RecordCount: [4]byte(recordCountBytes),
	}
	got := record.ConvertRecord()
	want := &TransactionLogHeader{
		MagicString: "MPS7",
		Version:     2,
		RecordCount: 100,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
