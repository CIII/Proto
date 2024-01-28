package main

import (
	"fmt"
	"log"
	"os"
	"proto/internal/utils"
)

func main() {
	f, err := os.Open("txnlog.dat")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	transactionLogHeader, err := utils.ParseHeader(f)
	if err != nil {
		log.Fatalln(err)
	}

	// we can only parse MPS7 mainframe files
	if transactionLogHeader.MagicString != "MPS7" {
		log.Fatalf("Invalid Protocol format. Unable to parse the following format: %v", transactionLogHeader.MagicString)
	}

	transactions, err := utils.ParseTransactions(f, transactionLogHeader)
	if err != nil {
		log.Fatalln(err)
	}

	id := uint64(2456938384156277127)
	metrics, err := utils.GetMetricsByUserId(transactions, id)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("total credit amount=%.2f\n", metrics.TotalCreditAmount)
	fmt.Printf("total debit amount=%.2f\n", metrics.TotalDebitAmount)
	fmt.Printf("autopays started=%v\n", metrics.AutopaysStarted)
	fmt.Printf("autopays ended=%v\n", metrics.AutopaysEnded)
	fmt.Printf("balance for user %v=%.2f\n", id, metrics.Balance)
}
