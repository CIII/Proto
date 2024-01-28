package models

import (
	"encoding/binary"
	"math"
)

const (
	Credit       = "Credit"
	Debit        = "Debit"
	StartAutopay = "StartAutopay"
	EndAutopay   = "EndAutopay"
)

var RECORD_TYPES = map[int]string{
	0: Credit,
	1: Debit,
	2: StartAutopay,
	3: EndAutopay,
}

type TransactionRecord struct {
	RecordType    string
	UnixTimestamp uint32
	UserId        uint64
	Amount        float64
}

type RawCreditDebitTransactionRecord struct {
	UnixTimestamp [4]byte
	UserId        [8]byte
	Amount        [8]byte
}

type RawAutoPayTransactionRecord struct {
	UnixTimestamp [4]byte
	UserId        [8]byte
}

func (raw *RawCreditDebitTransactionRecord) ConvertRecord(recordType string) *TransactionRecord {

	record := TransactionRecord{
		RecordType:    recordType,
		UnixTimestamp: binary.BigEndian.Uint32(raw.UnixTimestamp[:]),
		UserId:        binary.BigEndian.Uint64(raw.UserId[:]),
	}

	if recordType == Credit || recordType == Debit {
		record.Amount = math.Float64frombits(binary.BigEndian.Uint64(raw.Amount[:]))
	}

	return &record
}

func (raw *RawAutoPayTransactionRecord) ConvertRecord(recordType string) *TransactionRecord {
	return &TransactionRecord{
		RecordType:    recordType,
		UnixTimestamp: binary.BigEndian.Uint32(raw.UnixTimestamp[:]),
		UserId:        binary.BigEndian.Uint64(raw.UserId[:]),
	}
}
