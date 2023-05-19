package dto

import (
	"github.com/paykassa-dev/golang-api-sdk/enum"
	"github.com/paykassa-dev/golang-api-sdk/enum/currency"
	"github.com/paykassa-dev/golang-api-sdk/enum/payer"
	"github.com/paykassa-dev/golang-api-sdk/enum/priority"
	"github.com/paykassa-dev/golang-api-sdk/enum/system"
	"reflect"
	"testing"
)

func TestCheckPaymentRequest_Normalize(t *testing.T) {
	type fields struct {
		privateHash string
		test        bool
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			fields: fields{
				privateHash: "851b7ac0d8bd3e31598b36e564423b4598b15e2f020c1c25fb093ab0b80ec04a",
				test:        false,
			},
			want: map[string]string{
				"private_hash": "851b7ac0d8bd3e31598b36e564423b4598b15e2f020c1c25fb093ab0b80ec04a",
				"test":         "false",
			},
		},
		{
			fields: fields{
				privateHash: "",
				test:        true,
			},
			want: map[string]string{
				"private_hash": "",
				"test":         "true",
			},
		},
		{
			fields: fields{},
			want: map[string]string{
				"private_hash": "",
				"test":         "false",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := CheckPaymentRequest{
				PrivateHash: tt.fields.privateHash,
				Test:        tt.fields.test,
			}
			if got := r.Normalize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckTransactionRequest_Normalize(t *testing.T) {
	type fields struct {
		privateHash string
		test        bool
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			fields: fields{
				privateHash: "851b7ac0d8bd3e31598b36e564423b4598b15e2f020c1c25fb093ab0b80ec04a",
				test:        false,
			},
			want: map[string]string{
				"private_hash": "851b7ac0d8bd3e31598b36e564423b4598b15e2f020c1c25fb093ab0b80ec04a",
				"test":         "false",
			},
		},
		{
			fields: fields{
				privateHash: "",
				test:        true,
			},
			want: map[string]string{
				"private_hash": "",
				"test":         "true",
			},
		},
		{
			fields: fields{},
			want: map[string]string{
				"private_hash": "",
				"test":         "false",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := CheckTransactionRequest{
				PrivateHash: tt.fields.privateHash,
				Test:        tt.fields.test,
			}
			if got := r.Normalize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckBalanceRequest_Normalize1(t *testing.T) {
	type fields struct {
		shopId string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			fields: fields{},
			want: map[string]string{
				"shop_id": "",
			},
		},
		{
			fields: fields{
				shopId: "",
			},
			want: map[string]string{
				"shop_id": "",
			},
		},
		{
			fields: fields{
				shopId: "test",
			},
			want: map[string]string{
				"shop_id": "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := CheckBalanceRequest{
				ShopId: tt.fields.shopId,
			}
			if got := r.Normalize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakePaymentRequest_Normalize(t *testing.T) {
	type fields struct {
		shopId           string
		amount         string
		currency       enum.Currency
		system         enum.System
		paidCommission enum.CommissionPayer
		number         string
		tag            string
		priority       enum.TransactionPriority
		test           bool
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			fields: fields{
				shopId:           "123",
				amount:         "0.12345678",
				currency:       currency.BTC,
				system:         system.BITCOIN,
				paidCommission: payer.Client,
				number:         "3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy",
				tag:            "549",
				priority:       priority.High,
				test:           true,
			},
			want: map[string]string{
				"shop_id":          "123",
				"amount":          	"0.12345678",
				"currency":        	"BTC",
				"system":          	"11",
				"paid_commission": 	"client",
				"number":          	"3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy",
				"tag":             	"549",
				"priority":        	"high",
				"test":            	"true",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := MakePaymentRequest{
				ShopId:           	tt.fields.shopId,
				Amount:         	tt.fields.amount,
				Currency:       	tt.fields.currency,
				System:         	tt.fields.system,
				PaidCommission: 	tt.fields.paidCommission,
				Number:         	tt.fields.number,
				Tag:            	tt.fields.tag,
				Priority:       	tt.fields.priority,
				Test:           	tt.fields.test,
			}
			if got := r.Normalize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateAddressRequest_Normalize(t *testing.T) {
	type fields struct {
		paidCommission enum.CommissionPayer
		comment        string
		currency       enum.Currency
		system         enum.System
		amount         string
		orderId        string
		test           bool
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			fields: fields{
				paidCommission: "",
				comment:        "",
				currency:       "",
				system:         "",
				amount:         "",
				orderId:        "",
				test:           false,
			},
			want: map[string]string{
				"order_id":        "",
				"amount":          "",
				"currency":        "",
				"system":          "",
				"comment":         "",
				"phone":           "false",
				"paid_commission": "",
				"test":            "false",
			},
		},
		{
			fields: fields{
				paidCommission: payer.Shop,
				comment:        "TEST",
				currency:       currency.USDT,
				system:         system.TRON_TRC20,
				amount:         "123.45",
				orderId:        "order_id",
				test:           true,
			},
			want: map[string]string{
				"order_id":        "order_id",
				"amount":          "123.45",
				"currency":        "USDT",
				"system":          "30",
				"comment":         "TEST",
				"phone":           "false",
				"paid_commission": "shop",
				"test":            "true",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := GenerateAddressRequest{
				PaidCommission: tt.fields.paidCommission,
				Comment:        tt.fields.comment,
				Currency:       tt.fields.currency,
				System:         tt.fields.system,
				Amount:         tt.fields.amount,
				OrderId:        tt.fields.orderId,
				Test:           tt.fields.test,
			}
			if got := r.Normalize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPaymentUrlRequest_Normalize(t *testing.T) {
	type fields struct {
		PaidCommission enum.CommissionPayer
		Comment        string
		Currency       enum.Currency
		System         enum.System
		Amount         string
		OrderId        string
		Test           bool
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			fields: fields{
				PaidCommission: "",
				Comment:        "",
				Currency:       "",
				System:         "",
				Amount:         "",
				OrderId:        "",
				Test:           false,
			},
			want: map[string]string{
				"order_id":        "",
				"amount":          "",
				"currency":        "",
				"system":          "",
				"comment":         "",
				"phone":           "false",
				"paid_commission": "",
				"test":            "false",
			},
		},
		{
			fields: fields{
				PaidCommission: payer.Shop,
				Comment:        "TEST",
				Currency:       currency.USDT,
				System:         system.TRON_TRC20,
				Amount:         "123.45",
				OrderId:        "order_id",
				Test:           true,
			},
			want: map[string]string{
				"order_id":        "order_id",
				"amount":          "123.45",
				"currency":        "USDT",
				"system":          "30",
				"comment":         "TEST",
				"phone":           "false",
				"paid_commission": "shop",
				"test":            "true",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := GetPaymentUrlRequest{
				PaidCommission: tt.fields.PaidCommission,
				Comment:        tt.fields.Comment,
				Currency:       tt.fields.Currency,
				System:         tt.fields.System,
				Amount:         tt.fields.Amount,
				OrderId:        tt.fields.OrderId,
				Test:           tt.fields.Test,
			}
			if got := r.Normalize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
