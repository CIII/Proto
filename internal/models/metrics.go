package models

type Metrics struct {
	TotalCreditAmount float64
	TotalDebitAmount  float64
	AutopaysStarted   int64
	AutopaysEnded     int64
	Balance           float64
}
