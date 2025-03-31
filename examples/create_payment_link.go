package main

import (
	"github.com/joho/godotenv"
	"github.com/paykassa-dev/golang-api-sdk/api"
	"github.com/paykassa-dev/golang-api-sdk/dto"
	"github.com/paykassa-dev/golang-api-sdk/enum/currency"
	"github.com/paykassa-dev/golang-api-sdk/enum/system"
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

	request := dto.NewGetPaymentUrlRequest(
		"1005000",
		"0.12345678",
		system.BITCOIN,
		currency.BTC,
	)

	request.SetComment("test comment") //set a comment

	response := client.GetPaymentUrl(request)

	if !response.Error {
		log.Println("Url: ", response.Data.Url)
	} else {
		log.Printf("Error message: %s\n", response.Message)
	}
}
