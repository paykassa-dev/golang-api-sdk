package main

import (
	"fmt"
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

	request := dto.NewGenerateAddressRequest(
		"1005000",
		system.TON,
		currency.USDT,
	)

	request.SetComment("test comment") //set a comment

	response := client.GenerateAddress(request)

	if !response.Error {

		ext := ""
		if response.Data.IsTag {
			ext = fmt.Sprintf(" %s: %s", response.Data.TagName, response.Data.Tag)
		}

		log.Printf("Invoice ID: %s\n", response.Data.InvoiceId)
		log.Printf("Wallet: %s%s\n", response.Data.Wallet, ext)

		if testMode {
			log.Println("")

			log.Println("Test URL:", response.Data.Url)
		}
	} else {
		log.Printf("Error message: %s\n", response.Message)
	}
}
