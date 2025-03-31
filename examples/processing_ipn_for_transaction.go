package main

import (
	"fmt"
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

	client := api.NewMerchantApi(os.Getenv("MERCHANT_ID"), os.Getenv("MERCHANT_PASSWORD"))

	testMode := false //activate a test mode

	client.SetTest(testMode)

	request := dto.NewCheckTransactionRequest(
		"607e40c901bd3df89464ea7394b0b46eb7a876c14c5bcc705da4d1703d5274c4",
	)

	response := client.CheckTransaction(request)

	if !response.Error {
		log.Printf("Order ID in your system: %s\n", response.Data.OrderId)
		log.Printf("Transaction Number: %s\n", response.Data.Transaction)
		log.Printf("Txid: %s\n", response.Data.Txid)
		log.Printf("Amount: %s\n", response.Data.Amount)
		log.Printf("Fee: %s\n", response.Data.Fee)
		log.Printf("System: %s\n", response.Data.System)
		log.Printf("Currency: %s\n", response.Data.Currency)
		log.Printf("Address from: %s\n", response.Data.AddressFrom)
		log.Printf("Address: %s\n", response.Data.Address)
		log.Printf("Tag: %s\n", response.Data.Tag)
		log.Printf("Confirmations: %d\n", response.Data.Confirmations)
		log.Printf("Required confirmations: %d\n", response.Data.RequiredConfirmations)
		log.Printf("Status: %s\n", response.Data.Status)

		// The applied amount may differ from the received amount if partial payment mode is enabled.
		// Default is 'no'.
		if response.Data.Status == "yes" {
			log.Printf("You can credit this amount: %s %s %s\n", response.Data.System, response.Data.Amount, response.Data.Currency)
			// your code...
		}

		//You should confirm that you’ve credited the funds to avoid repeated IPNs.
		fmt.Printf("%s|succes", response.Data.OrderId)
	} else {
		log.Printf("Error message: %s\n", response.Message)
	}
}
