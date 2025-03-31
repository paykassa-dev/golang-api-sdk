package main

import (
	"github.com/joho/godotenv"
	"github.com/paykassa-dev/golang-api-sdk/api"
	"github.com/paykassa-dev/golang-api-sdk/dto"
	"github.com/paykassa-dev/golang-api-sdk/enum/currency"
	"github.com/paykassa-dev/golang-api-sdk/enum/priority"
	"github.com/paykassa-dev/golang-api-sdk/enum/system"
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

	request := dto.NewMakePaymentRequest(
		os.Getenv("MERCHANT_ID"),
		"0.01",
		system.TON,
		currency.USDT,
		"EQBkmfJlIjxZqB8xndrlDX05gGLKFPy84PKT4NRcmWy0PCaL",
		"test_transfer",
	)

	request.SetTag("100500")             //Ripple, Stellar, TON
	request.SetPriority(priority.Medium) //Bitcoin, Litecoin, Dogecoin, Dash, BitcoinCash

	response := client.MakePayment(request)

	if !response.Error {
		log.Printf(
			"We have sent the %s %s %s to %s.The txid is %s",
			response.Data.System,
			response.Data.AmountPay,
			response.Data.Currency,
			response.Data.Number,
			response.Data.TxId,
		)

		if response.Data.PaymentId == response.Data.TxId {
			log.Printf("Status txid is not ready! Safe the invoice id %s to get the txid later", response.Data.Transaction)
		}

		log.Printf("Deducted amount is %s", response.Data.Amount)
		log.Printf("Address link is %s", response.Data.ExplorerAddressLink)
		log.Printf("Transaction link is %s", response.Data.ExplorerTransactionLink)

	} else {
		log.Printf("Error message: %s\n", response.Message)
	}
}
