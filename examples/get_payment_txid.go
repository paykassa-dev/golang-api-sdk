package main

import (
	"github.com/joho/godotenv"
	"github.com/paykassa-dev/golang-api-sdk/api"
	"github.com/paykassa-dev/golang-api-sdk/dto"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't fetch .env")
	}

	client := api.NewPaymentApi(os.Getenv("API_ID"), os.Getenv("API_PASSWORD"))

	testMode := false //activate a test mode

	client.SetTest(testMode)

	invoices := []string{
		"37433236",
		"32531999", //for test not found
		"37433238",
		"37433220",
		"37433196",
	}
	request := dto.NewGetTxidsOfInvoicesRequest(os.Getenv("MERCHANT_ID"), invoices)

	response := client.GetTxidsOfInvoices(request)

	if !response.Error {
		for _, invoiceID := range invoices {
			txids, ok := response.Data[invoiceID]
			if !ok || len(txids) == 0 {
				log.Printf("Invoice is not found: %s\n", invoiceID)
				continue
			}

			for _, txid := range txids {
				log.Printf("Invoice %s: %s\n", invoiceID, txid)
			}
		}
	} else {
		log.Printf("Error message: %s\n", response.Message)
	}
}
