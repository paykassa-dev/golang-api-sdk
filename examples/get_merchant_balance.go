package main

import (
	"github.com/joho/godotenv"
	"github.com/paykassa-dev/golang-api-sdk/api"
	"github.com/paykassa-dev/golang-api-sdk/dto"
	"github.com/paykassa-dev/golang-api-sdk/enum/pscur"
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

	request := dto.NewCheckBalanceRequest(
		os.Getenv("MERCHANT_ID"),
	)

	response := client.CheckBalance(request)

	if !response.Error {
		for _, tag := range pscur.PsCurList {
			if balance, exist := response.Data[string(tag)]; exist {
				log.Printf("Balance %s = %s\n", tag, balance)
			}
		}

		log.Println()

		manual := pscur.BinancesmartchainBep20Btc
		log.Printf("Balance by constant value %s = %s\n", string(manual), response.Data[string(manual)])

	} else {
		log.Printf("Error message: %s\n", response.Message)
	}
}
