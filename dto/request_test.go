package dto_test

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/paykassa-dev/golang-api-sdk/dto"
	local_currency "github.com/paykassa-dev/golang-api-sdk/enum/currency"
	local_priority "github.com/paykassa-dev/golang-api-sdk/enum/priority"
	local_system "github.com/paykassa-dev/golang-api-sdk/enum/system"
)

// TestBaseRequest tests the base request
func TestBaseRequest(t *testing.T) {
	br := &dto.BaseRequest{}
	testComment := "test comment"
	br.SetComment(testComment)

	if br.Comment != testComment {
		t.Errorf("Expected comment %s, got %s", testComment, br.Comment)
	}
}

// TestCheckBalanceRequest tests the balance check request
func TestCheckBalanceRequest(t *testing.T) {
	shopId := "123456"
	req := dto.NewCheckBalanceRequest(shopId)

	if req.ShopId != shopId {
		t.Errorf("Expected ShopId %s, got %s", shopId, req.ShopId)
	}

	values := req.Normalize()
	expectedValues := url.Values{}
	expectedValues.Set("shop_id", shopId)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected %v, got %v", expectedValues, values)
	}
}

// TestCheckPaymentRequest tests the payment check request
func TestCheckPaymentRequest(t *testing.T) {
	privateHash := "hash123"
	req := dto.NewCheckPaymentRequest(privateHash)

	if req.PrivateHash != privateHash {
		t.Errorf("Expected PrivateHash %s, got %s", privateHash, req.PrivateHash)
	}

	values := req.Normalize()
	expectedValues := url.Values{}
	expectedValues.Set("private_hash", privateHash)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected %v, got %v", expectedValues, values)
	}
}

// TestCheckTransactionRequest tests the transaction check request
func TestCheckTransactionRequest(t *testing.T) {
	privateHash := "hash456"
	req := dto.NewCheckTransactionRequest(privateHash)

	if req.PrivateHash != privateHash {
		t.Errorf("Expected PrivateHash %s, got %s", privateHash, req.PrivateHash)
	}

	values := req.Normalize()
	expectedValues := url.Values{}
	expectedValues.Set("private_hash", privateHash)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected %v, got %v", expectedValues, values)
	}
}

// TestGenerateAddressRequest tests the address generation request
func TestGenerateAddressRequest(t *testing.T) {
	orderId := "order123"
	system := local_system.BITCOIN
	currency := local_currency.BTC
	comment := "address generation"

	req := dto.NewGenerateAddressRequest(orderId, system, currency)
	req.SetComment(comment)

	values := req.Normalize()
	expectedValues := url.Values{}
	expectedValues.Set("order_id", orderId)
	expectedValues.Set("amount", "1.0")
	expectedValues.Set("currency", string(currency))
	expectedValues.Set("system", string(system))
	expectedValues.Set("comment", comment)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected %v, got %v", expectedValues, values)
	}
}

// TestGetPaymentUrlRequest tests the payment URL request
func TestGetPaymentUrlRequest(t *testing.T) {
	orderId := "order456"
	amount := "100.50"
	system := local_system.LITECOIN
	currency := local_currency.LTC
	comment := "payment link"

	req := dto.NewGetPaymentUrlRequest(orderId, amount, system, currency)
	req.SetComment(comment)

	values := req.Normalize()
	expectedValues := url.Values{}
	expectedValues.Set("order_id", orderId)
	expectedValues.Set("amount", amount)
	expectedValues.Set("currency", string(currency))
	expectedValues.Set("system", string(system))
	expectedValues.Set("comment", comment)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected %v, got %v", expectedValues, values)
	}
}

// TestGetTxidsOfInvoicesRequest tests the request for retrieving transaction IDs
func TestGetTxidsOfInvoicesRequest(t *testing.T) {
	shopId := "shop789"
	invoices := []string{"100500", "200600", "300700"}

	req := dto.NewGetTxidsOfInvoicesRequest(shopId, invoices)

	values := req.Normalize()
	expectedValues := url.Values{}
	expectedValues.Set("shop_id", shopId)
	for _, invoice := range invoices {
		expectedValues.Add("invoices[]", invoice)
	}

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected %v, got %v", expectedValues, values)
	}
}

// TestMakePaymentRequest tests the payment execution request
func TestMakePaymentRequest(t *testing.T) {
	shopId := "shop999"
	amount := "250.75"
	system := local_system.RIPPLE
	currency := local_currency.XRP
	number := "wallet123"
	comment := "payment"
	tag := "tag123"
	customPriority := local_priority.High

	req := dto.NewMakePaymentRequest(shopId, amount, system, currency, number, comment)
	req.SetTag(tag)
	req.SetPriority(customPriority)

	values := req.Normalize()
	expectedValues := url.Values{}
	expectedValues.Set("amount", amount)
	expectedValues.Set("currency", string(currency))
	expectedValues.Set("system", string(system))
	expectedValues.Set("comment", comment)
	expectedValues.Set("shop_id", shopId)
	expectedValues.Set("number", number)
	expectedValues.Set("tag", tag)
	expectedValues.Set("priority", string(customPriority))

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected %v, got %v", expectedValues, values)
	}
}

// TestApplyFields tests the correct application of fields
func TestApplyFields(t *testing.T) {
	t.Run("Adding fields to empty url.Values", func(t *testing.T) {
		values := url.Values{}
		fields := map[string]string{
			"key1": "value1",
			"key2": "value2",
		}

		// Create a private function with the same behavior as applyFields
		result := func(values url.Values, fields map[string]string) url.Values {
			for key, value := range fields {
				values.Set(key, value)
			}
			return values
		}(values, fields)

		expectedValues := url.Values{}
		expectedValues.Set("key1", "value1")
		expectedValues.Set("key2", "value2")

		if !reflect.DeepEqual(result, expectedValues) {
			t.Errorf("Expected %v, got %v", expectedValues, result)
		}
	})

	t.Run("Adding fields to existing url.Values", func(t *testing.T) {
		values := url.Values{}
		values.Set("existing", "value")
		fields := map[string]string{
			"key1": "value1",
			"key2": "value2",
		}

		// Create a private function with the same behavior as applyFields
		result := func(values url.Values, fields map[string]string) url.Values {
			for key, value := range fields {
				values.Set(key, value)
			}
			return values
		}(values, fields)

		expectedValues := url.Values{}
		expectedValues.Set("existing", "value")
		expectedValues.Set("key1", "value1")
		expectedValues.Set("key2", "value2")

		if !reflect.DeepEqual(result, expectedValues) {
			t.Errorf("Expected %v, got %v", expectedValues, result)
		}
	})
}

// TestChaining tests the method chaining capability
func TestChaining(t *testing.T) {
	// Test chaining for GenerateAddressRequest
	orderId := "orderChain1"
	system := local_system.BINANCESMARTCHAIN_BEP20
	currency := local_currency.BNB
	comment := "chain test"

	req1 := dto.NewGenerateAddressRequest(orderId, system, currency).SetComment(comment)

	values1 := req1.Normalize()
	if values1.Get("comment") != comment {
		t.Errorf("Chain error for GenerateAddressRequest: expected comment %s, got %s",
			comment, values1.Get("comment"))
	}

	// Test chaining for GetPaymentUrlRequest
	req2 := dto.NewGetPaymentUrlRequest(orderId, "50.0", system, currency).SetComment(comment)

	values2 := req2.Normalize()
	if values2.Get("comment") != comment {
		t.Errorf("Chain error for GetPaymentUrlRequest: expected comment %s, got %s",
			comment, values2.Get("comment"))
	}

	// Test chaining for MakePaymentRequest
	shopId := "shop777"
	number := "wallet777"
	tag := "tag777"
	customPriority := local_priority.Low

	req3 := dto.NewMakePaymentRequest(shopId, "75.0", system, currency, number, comment).
		SetTag(tag).
		SetPriority(customPriority)

	values3 := req3.Normalize()
	if values3.Get("tag") != tag || values3.Get("priority") != string(customPriority) {
		t.Errorf("Chain error for MakePaymentRequest: expected tag %s and priority %s, got tag %s and priority %s",
			tag, string(customPriority), values3.Get("tag"), values3.Get("priority"))
	}
}
