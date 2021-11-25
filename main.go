package main

import (
	"net/http"

	"github.com/matheusmv/financial-planning-assistant/transaction"
)

func main() {

	http.HandleFunc("/transactions", transaction.FindAllTransactions)
	http.HandleFunc("/transactions/", transaction.FindTransactionById)
	http.HandleFunc("/transactions/create", transaction.CreateTransaction)

	http.ListenAndServe(":8080", nil)
}
