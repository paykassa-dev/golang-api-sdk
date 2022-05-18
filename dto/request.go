package dto

import (
	"github.com/paykassa-dev/golang-api-sdk/enum"
	"strconv"
)

type Request interface {
	Normalize() map[string]string
}

type CheckBalanceRequest struct {
	Shop string
}

func NewCheckBalanceRequest(shop string) *CheckBalanceRequest {
	return &CheckBalanceRequest{Shop: shop}
}

type MakePaymentRequest struct {
	Shop           string
	Amount         string
	System         enum.System
	Currency       enum.Currency
	PaidCommission enum.CommissionPayer
	Number         string
	Tag            string
	Priority       enum.TransactionPriority
	Test           bool
}

func NewMakePaymentRequest(shop string, amount string, system enum.System, currency enum.Currency, number string) *MakePaymentRequest {
	return &MakePaymentRequest{Shop: shop, Amount: amount, Currency: currency, System: system, Number: number}
}

type CheckPaymentRequest struct {
	PrivateHash string
	Test        bool
}

func NewCheckPaymentRequest(privateHash string) *CheckPaymentRequest {
	return &CheckPaymentRequest{PrivateHash: privateHash}
}

type CheckTransactionRequest struct {
	PrivateHash string
	Test        bool
}

func NewCheckTransactionRequest(privateHash string) *CheckTransactionRequest {
	return &CheckTransactionRequest{PrivateHash: privateHash}
}

type GenerateAddressRequest struct {
	OrderId        string
	Amount         string
	System         enum.System
	Currency       enum.Currency
	Comment        string
	PaidCommission enum.CommissionPayer
	Test           bool
}

func NewGenerateAddressRequest(orderId string, amount string, system enum.System, currency enum.Currency) *GenerateAddressRequest {
	return &GenerateAddressRequest{OrderId: orderId, Amount: amount, System: system, Currency: currency}
}

type GetPaymentUrlRequest struct {
	OrderId        string
	Amount         string
	System         enum.System
	Currency       enum.Currency
	Comment        string
	PaidCommission enum.CommissionPayer
	Test           bool
}

func NewGetPaymentUrlRequest(orderId string, amount string, system enum.System, currency enum.Currency) *GetPaymentUrlRequest {
	return &GetPaymentUrlRequest{OrderId: orderId, Amount: amount, System: system, Currency: currency}
}

func (r GetPaymentUrlRequest) Normalize() map[string]string {
	return map[string]string{
		"order_id":        r.OrderId,
		"amount":          r.Amount,
		"currency":        string(r.Currency),
		"system":          string(r.System),
		"comment":         r.Comment,
		"paid_commission": string(r.PaidCommission),
		"phone":           "false",
		"test":            strconv.FormatBool(r.Test),
	}
}

func (r CheckBalanceRequest) Normalize() map[string]string {
	return map[string]string{
		"shop": r.Shop,
	}
}

func (r GenerateAddressRequest) Normalize() map[string]string {
	return map[string]string{
		"order_id":        r.OrderId,
		"amount":          r.Amount,
		"currency":        string(r.Currency),
		"system":          string(r.System),
		"comment":         r.Comment,
		"paid_commission": string(r.PaidCommission),
		"phone":           "false",
		"test":            strconv.FormatBool(r.Test),
	}
}

func (r CheckTransactionRequest) Normalize() map[string]string {
	return map[string]string{
		"private_hash": r.PrivateHash,
		"test":         strconv.FormatBool(r.Test),
	}
}

func (r CheckPaymentRequest) Normalize() map[string]string {
	return map[string]string{
		"private_hash": r.PrivateHash,
		"test":         strconv.FormatBool(r.Test),
	}
}

func (r MakePaymentRequest) Normalize() map[string]string {
	return map[string]string{
		"priority":        string(r.Priority),
		"tag":             r.Tag,
		"number":          r.Number,
		"paid_commission": string(r.PaidCommission),
		"system":          string(r.System),
		"currency":        string(r.Currency),
		"shop":            r.Shop,
		"amount":          r.Amount,
		"test":            strconv.FormatBool(r.Test),
	}
}
