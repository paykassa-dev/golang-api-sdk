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

	request := dto.NewCheckPaymentRequest(
		"72ea64dae378eeeebc9dc8f2d79c89ac2c0bd6f6e2c4ab7b3aa08655dd2e34e4",
	)

	response := client.CheckPayment(request)

	if !response.Error {
		log.Printf("Order ID in your system: %s\n", response.Data.OrderId)
		log.Printf("Transaction Number: %s\n", response.Data.Transaction)
		log.Printf("Amount: %s\n", response.Data.Amount)
		log.Printf("System: %s\n", response.Data.System)
		log.Printf("Currency: %s\n", response.Data.Currency)
		log.Printf("Wallet: %s\n", response.Data.Address)
		log.Printf("Tag: %s\n", response.Data.Tag)
		log.Printf("Partial? %s\n", response.Data.Partial)

		// The applied amount may differ from the received amount if partial payment mode is enabled.
		// Default is 'no'.
		if response.Data.Partial == "yes" {
			log.Printf("Verify the amount: %s %s %s\n", response.Data.System, response.Data.Amount, response.Data.Currency)
		} else {
			log.Printf("You can credit this amount: %s %s %s\n", response.Data.System, response.Data.Amount, response.Data.Currency)
		}

		// your code...

		//You should confirm that you’ve credited the funds to avoid repeated IPNs.
		fmt.Printf("%s|succes", response.Data.OrderId)
	} else {
		log.Printf("Error message: %s\n", response.Message)
	}
}
