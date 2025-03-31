# Paykassa SCI & API

## Installation

```bash
go get github.com/paykassa-dev/golang-api-sdk
```




**To run examples:**
```bash
cp ./.env.example ./.env
```

### Get a deposit address
```golang
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

```

**Example:**
```bash
go run ./examples/create_address.go
```

### Check an IPN of a transaction
```golang
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

```

**Example:**
```bash
go run ./examples/processing_ipn_for_transaction.go
```

### Get a payment link(create an order)
```golang
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

```

**Example:**
```bash
go run ./examples/create_payment_link.go
```

### Check an IPN of an order
```golang
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

```

**Example:**
```bash
go run ./examples/processing_ipn_for_order.go
```

### Send money
```golang
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

```

**Example:**
```bash
go run ./examples/send_money.go
```

### Get a merchant balance
```golang
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

```

**Example:**
```bash
go run ./examples/get_merchant_balance.go
```

### Get a txid of a payment
```golang
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

```

**Example:**
```bash
go run ./examples/get_payment_txid.go
```


## References
- [Devs Documentation](https://paykassa.pro/en/developers)
- [API Documentation](https://paykassa.pro/docs/)

## Contributing
If during your work with this wrapper you encounter a bug or have a suggestion to help improve it for others, you are welcome to open a Github issue on this repository and it will be reviewed by one of our development team members. The Paykassa.pro bug bounty does not cover this wrapper.

## License
MIT - see LICENSE