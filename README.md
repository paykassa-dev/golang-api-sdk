# Paykassa SCI & API

## Installation

```
go get github.com/paykassa-dev/golang-api-sdk
```

## Payment API

### Initialize Client

```go
import "github.com/paykassa-dev/golang-api-sdk/api"

client := api.NewPaymentApi("apiId", "apiKey")
```

### Check Balance

```go
import "github.com/paykassa-dev/golang-api-sdk/dto"

request := dto.NewCheckBalanceRequest("shopId")
response := client.CheckBalance(request)

if !response.Error {
    fmt.Println(response.Data.BitcoinBtc)
    fmt.Println(response.Data.LitecoinLtc)
}
```

### Make Payment

```go
import (
    "github.com/paykassa-dev/golang-api-sdk/dto"
    "github.com/paykassa-dev/golang-api-sdk/enum/currency"
    "github.com/paykassa-dev/golang-api-sdk/enum/system"
)

request := dto.NewMakePaymentRequest(
    "15743",
    "0.12345678",
    system.BITCOIN,
    currency.BTC,
    "3LaKdUrPfVyZeEVYpZei3HwjqQj5AHHTCE",
)

response := client.MakePayment(request)

if !response.Error {
    fmt.Println(response.Data.Transaction)
    fmt.Println(response.Data.PaidCommission)
}
```

## Merchant API

### Initialize Client

```go
import "github.com/paykassa-dev/golang-api-sdk/api"

client = api.MerchantApi("sci_id", "sci_key")
```

### Check Payment (IPN)

```go
import "github.com/paykassa-dev/golang-api-sdk/dto"

request := dto.NewCheckPaymentRequest("privateHash")
response := client.CheckPayment(request)

if !response.Error {
    fmt.Println(response.Data.Amount)
    fmt.Println(response.Data.Transaction)
}
```

### Check Transaction (IPN)

```go
import "github.com/paykassa-dev/golang-api-sdk/dto"

request := dto.NewCheckPaymentRequest("privateHash")
response := client.CheckPayment(request)

if !response.Error {
    fmt.Println(response.Data.AddressFrom)
    fmt.Println(response.Data.Confmations)
}
```

### Generate Address

```go
import (
    "github.com/paykassa-dev/golang-api-sdk/dto"
    "github.com/paykassa-dev/golang-api-sdk/enum/currency"
    "github.com/paykassa-dev/golang-api-sdk/enum/system"
)

request := dto.NewGenerateAddressRequest(
    "orderId",
    "0.12345678",
    system.RIPPLE,
    currency.XRP,
)

response := client.GenerateAddress(request)

if !response.Error {
    fmt.Println(response.Data.Invoice)
    fmt.Println(response.Data.Wallet)
    fmt.Println(response.Data.Tag)
}
```

### Get Payment Url

```go
import (
    "github.com/paykassa-dev/golang-api-sdk/dto"
    "github.com/paykassa-dev/golang-api-sdk/enum/currency"
    "github.com/paykassa-dev/golang-api-sdk/enum/system"
)

request := dto.NewGetPaymentUrlRequest(
    "orderId",
    "0.12345678",
    system.PERFECTMONEY,
    currency.USD,
)

response := client.GetPaymentUrl(request)

if !response.Error {
    fmt.Println(response.Data.Url)
    fmt.Println(response.Data.Method)
}
```

## References
- [Devs Documentation](https://paykassa.pro/en/developers)
- [API Documentation](https://paykassa.pro/docs/)
