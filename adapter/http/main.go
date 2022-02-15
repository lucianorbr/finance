package http

import (
	"net/http"

	"github.com/lucianorbr/finance/adapter/http/actuator"
	"github.com/lucianorbr/finance/adapter/http/transaction"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)

	http.HandleFunc("/health", actuator.Health)

	http.ListenAndServe(":8080", nil)
}
