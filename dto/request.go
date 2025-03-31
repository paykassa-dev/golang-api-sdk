package dto

import (
	"net/url"

	"github.com/paykassa-dev/golang-api-sdk/enum"
	"github.com/paykassa-dev/golang-api-sdk/enum/priority"
)

// Request defines the interface for all API requests.
type Request interface {
	Normalize() url.Values
}

// BaseRequest contains common fields for requests.
type BaseRequest struct {
	Comment string
}

// SetComment sets the comment for the request.
func (r *BaseRequest) SetComment(comment string) {
	r.Comment = comment
}

// applyFields adds key-value pairs from a map to url.Values.
func applyFields(values url.Values, fields map[string]string) url.Values {
	for key, value := range fields {
		values.Set(key, value)
	}
	return values
}

// ===== CheckBalanceRequest =====

// CheckBalanceRequest represents a balance check request.
type CheckBalanceRequest struct {
	ShopId string
}

// NewCheckBalanceRequest creates a new balance check request.
func NewCheckBalanceRequest(shopId string) *CheckBalanceRequest {
	return &CheckBalanceRequest{ShopId: shopId}
}

// Normalize converts the request to url.Values.
func (r *CheckBalanceRequest) Normalize() url.Values {
	return applyFields(url.Values{}, map[string]string{
		"shop_id": r.ShopId,
	})
}

// ===== CheckPaymentRequest =====

// CheckPaymentRequest represents a payment check request.
type CheckPaymentRequest struct {
	PrivateHash string
}

// NewCheckPaymentRequest creates a new payment check request.
func NewCheckPaymentRequest(privateHash string) *CheckPaymentRequest {
	return &CheckPaymentRequest{PrivateHash: privateHash}
}

// Normalize converts the request to url.Values.
func (r *CheckPaymentRequest) Normalize() url.Values {
	return applyFields(url.Values{}, map[string]string{
		"private_hash": r.PrivateHash,
	})
}

// ===== CheckTransactionRequest =====

// CheckTransactionRequest represents a transaction check request.
type CheckTransactionRequest struct {
	PrivateHash string
}

// NewCheckTransactionRequest creates a new transaction check request.
func NewCheckTransactionRequest(privateHash string) *CheckTransactionRequest {
	return &CheckTransactionRequest{PrivateHash: privateHash}
}

// Normalize converts the request to url.Values.
func (r *CheckTransactionRequest) Normalize() url.Values {
	return applyFields(url.Values{}, map[string]string{
		"private_hash": r.PrivateHash,
	})
}

// ===== PaymentRequest =====

// PaymentRequest represents the base struct for payment-related requests.
type PaymentRequest struct {
	*BaseRequest
	OrderId  string
	Amount   string
	System   enum.System
	Currency enum.Currency
}

// newPaymentRequest creates a new base payment request.
func newPaymentRequest(orderId, amount string, system enum.System, currency enum.Currency) *PaymentRequest {
	return &PaymentRequest{
		BaseRequest: &BaseRequest{},
		OrderId:     orderId,
		Amount:      amount,
		System:      system,
		Currency:    currency,
	}
}

// normalize converts the payment request data into url.Values.
func (r *PaymentRequest) normalize() url.Values {
	return applyFields(url.Values{}, map[string]string{
		"order_id": r.OrderId,
		"amount":   r.Amount,
		"currency": string(r.Currency),
		"system":   string(r.System),
		"comment":  r.Comment,
	})
}

// ===== GenerateAddressRequest =====

// GenerateAddressRequest represents an address generation request.
type GenerateAddressRequest struct {
	*PaymentRequest
}

// NewGenerateAddressRequest creates a new address generation request.
func NewGenerateAddressRequest(orderId string, system enum.System, currency enum.Currency) *GenerateAddressRequest {
	return &GenerateAddressRequest{
		PaymentRequest: newPaymentRequest(orderId, "1.0", system, currency),
	}
}

// SetComment sets the comment and returns the current request for chaining.
func (r *GenerateAddressRequest) SetComment(comment string) *GenerateAddressRequest {
	r.PaymentRequest.SetComment(comment)
	return r
}

// Normalize converts the request to url.Values.
func (r *GenerateAddressRequest) Normalize() url.Values {
	return r.PaymentRequest.normalize()
}

// ===== GetPaymentUrlRequest =====

// GetPaymentUrlRequest represents a request to obtain a payment URL.
type GetPaymentUrlRequest struct {
	*PaymentRequest
}

// NewGetPaymentUrlRequest creates a new payment URL request.
func NewGetPaymentUrlRequest(orderId, amount string, system enum.System, currency enum.Currency) *GetPaymentUrlRequest {
	return &GetPaymentUrlRequest{
		PaymentRequest: newPaymentRequest(orderId, amount, system, currency),
	}
}

// SetComment sets the comment for the request.
func (r *GetPaymentUrlRequest) SetComment(comment string) *GetPaymentUrlRequest {
	r.PaymentRequest.SetComment(comment)
	return r
}

// Normalize converts the request to url.Values.
func (r *GetPaymentUrlRequest) Normalize() url.Values {
	return r.PaymentRequest.normalize()
}

// ===== GetTxidsOfInvoicesRequest =====

// GetTxidsOfInvoicesRequest represents a request to retrieve transaction IDs of invoices.
type GetTxidsOfInvoicesRequest struct {
	ShopId   string
	Invoices []string
}

// NewGetTxidsOfInvoicesRequest creates a new request to retrieve transaction IDs of invoices.
func NewGetTxidsOfInvoicesRequest(shopId string, invoices []string) *GetTxidsOfInvoicesRequest {
	return &GetTxidsOfInvoicesRequest{ShopId: shopId, Invoices: invoices}
}

// Normalize converts the request to url.Values.
func (r *GetTxidsOfInvoicesRequest) Normalize() url.Values {
	values := url.Values{}
	values.Set("shop_id", r.ShopId)
	for _, invoice := range r.Invoices {
		values.Add("invoices[]", invoice)
	}
	return values
}

// ===== MakePaymentRequest =====

// MakePaymentRequest represents a payment execution request.
type MakePaymentRequest struct {
	Amount   string
	System   enum.System
	Currency enum.Currency
	ShopId   string
	Number   string
	Tag      string
	Comment  string
	Priority enum.TransactionPriority
}

// NewMakePaymentRequest creates a new payment execution request.
func NewMakePaymentRequest(shopId, amount string, system enum.System, currency enum.Currency, number, comment string) *MakePaymentRequest {
	req := &MakePaymentRequest{
		Amount:   amount,
		System:   system,
		Currency: currency,
		ShopId:   shopId,
		Number:   number,
		Comment:  comment,
		Tag:      "",
		Priority: priority.High,
	}
	return req
}

// SetTag sets the tag for the request.
func (r *MakePaymentRequest) SetTag(tag string) *MakePaymentRequest {
	r.Tag = tag
	return r
}

// SetPriority sets the priority for the request.
func (r *MakePaymentRequest) SetPriority(priority enum.TransactionPriority) *MakePaymentRequest {
	r.Priority = priority
	return r
}

// Normalize converts the request to url.Values.
func (r *MakePaymentRequest) Normalize() url.Values {
	values := url.Values{}
	return applyFields(values, map[string]string{
		"amount":   r.Amount,
		"system":   string(r.System),
		"currency": string(r.Currency),
		"shop_id":  r.ShopId,
		"number":   r.Number,
		"tag":      r.Tag,
		"priority": string(r.Priority),
		"comment":  r.Comment,
	})
}
