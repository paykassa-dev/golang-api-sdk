package dto

import (
	"github.com/paykassa-dev/golang-api-sdk/enum"
	"net/url"
)

type CheckBalanceResponse struct {
	Error   bool              `json:"error"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

type MakePaymentResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    struct {
		ShopId                  string        `json:"shop_id"`
		Transaction             string        `json:"transaction"`
		PaymentId               string        `json:"payment_id"`
		TxId                    string        `json:"txid"`
		Amount                  string        `json:"amount"`
		AmountPay               string        `json:"amount_pay"`
		System                  enum.System   `json:"system"`
		Currency                enum.Currency `json:"currency"`
		Number                  string        `json:"number"`
		ShopCommissionPercent   string        `json:"shop_commission_percent"`
		ShopCommissionAmount    string        `json:"shop_commission_amount"`
		PaidCommission          string        `json:"paid_commission"`
		ExplorerAddressLink     string        `json:"explorer_address_link"`
		ExplorerTransactionLink string        `json:"explorer_transaction_link"`
	} `json:"data"`
}

type CheckPaymentResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    struct {
		Transaction string `json:"transaction"`
		ShopId      string `json:"shop_id"`
		OrderId     string `json:"order_id"`
		Amount      string `json:"amount"`
		Currency    string `json:"currency"`
		System      string `json:"system"`
		Address     string `json:"address"`
		Tag         string `json:"tag"`
		Hash        string `json:"hash"`
		Partial     string `json:"partial"`
	} `json:"data"`
}

type CheckTransactionResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    struct {
		Transaction             string `json:"transaction"`
		Txid                    string `json:"txid"`
		ShopId                  string `json:"shop_id"`
		OrderId                 string `json:"order_id"`
		Amount                  string `json:"amount"`
		Fee                     string `json:"fee"`
		Currency                string `json:"currency"`
		System                  string `json:"system"`
		AddressFrom             string `json:"address_from"`
		Address                 string `json:"address"`
		Tag                     string `json:"tag"`
		Confirmations           int    `json:"confirmations"`
		RequiredConfirmations   int    `json:"required_confirmations"`
		Status                  string `json:"status"`
		Static                  string `json:"static"`
		DateUpdate              string `json:"date_update"`
		ExplorerAddressLink     string `json:"explorer_address_link"`
		ExplorerTransactionLink string `json:"explorer_transaction_link"`
	} `json:"data"`
}

type GenerateAddressResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    struct {
		InvoiceId string `json:"invoice_id"`
		OrderId   string `json:"order_id"`
		Wallet    string `json:"wallet"`
		Amount    string `json:"amount"`
		System    string `json:"system"`
		Currency  string `json:"currency"`
		Url       string `json:"url"`
		Tag       string `json:"tag"`
		IsTag     bool   `json:"is_tag"`
		TagName   string `json:"tag_name"`
	} `json:"data"`
}

type GetPaymentUrlResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    struct {
		Url    string `json:"url"`
		Method string `json:"method"`
	} `json:"data"`
}

type GetTxidsOfInvoicesResponse struct {
	Error   bool       `json:"error"`
	Message string     `json:"message"`
	Data    url.Values `json:"data"`
}
