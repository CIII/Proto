package models

import "encoding/binary"

type TransactionLogHeader struct {
	MagicString string
	Version     uint32
	RecordCount uint32
}

type RawTransactionLogHeader struct {
	MagicString [4]byte
	Version     [1]byte
	RecordCount [4]byte
}

func (raw *RawTransactionLogHeader) ConvertRecord() *TransactionLogHeader {
	return &TransactionLogHeader{
		MagicString: string(raw.MagicString[:]),
		Version:     uint32(raw.Version[0]),
		RecordCount: binary.BigEndian.Uint32(raw.RecordCount[:]),
	}
}
