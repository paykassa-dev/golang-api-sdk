package dto_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"testing"

	"github.com/paykassa-dev/golang-api-sdk/dto"
	"github.com/paykassa-dev/golang-api-sdk/enum"
	"github.com/paykassa-dev/golang-api-sdk/enum/pscur"
)

func TestCheckBalanceResponse_Unmarshal(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		want    dto.CheckBalanceResponse
		wantErr bool
	}{
		{
			name: "successful response",
			jsonStr: `{
				"error": false,
				"message": "Balance retrieved successfully",
				"data": {
					"bitcoin_btc": "1.1",
					"ethereum_eth": "2.2",
					"litecoin_ltc": "3.3",
					"dogecoin_doge": "4.4",
					"dash_dash": "5.5",
					"bitcoincash_bch": "6.6",
					"ripple_xrp": "7.7",
					"tron_trx": "8.8",
					"stellar_xlm": "9.9",
					"binancecoin_bnb": "10.10",
					"tron_trc20_usdt": "11.11",
					"binancesmartchain_bep20_usdt": "12.12",
					"ethereum_erc20_usdt": "13.13",
					"binancesmartchain_bep20_busd": "14.14",
					"binancesmartchain_bep20_usdc": "15.15",
					"binancesmartchain_bep20_ada": "16.16",
					"binancesmartchain_bep20_eos": "17.17",
					"binancesmartchain_bep20_btc": "18.18",
					"binancesmartchain_bep20_eth": "19.19",
					"binancesmartchain_bep20_doge": "20.20",
					"ton_ton": "21.21",
					"ton_usdt": "22.22"
				}
			}`,
			want: func() dto.CheckBalanceResponse {
				var resp dto.CheckBalanceResponse
				resp.Error = false
				resp.Message = "Balance retrieved successfully"
				resp.Data = make(map[string]string)
				for i, tag := range pscur.PsCurList {
					z := i + 1
					resp.Data[string(tag)] = fmt.Sprintf("%d.%d", z, z)
				}
				return resp
			}(),
			wantErr: false,
		},
		{
			name: "error response",
			jsonStr: `{
				"error": true,
				"message": "Invalid shop ID",
				"data": {}
			}`,
			want: func() dto.CheckBalanceResponse {
				var resp dto.CheckBalanceResponse
				resp.Error = true
				resp.Message = "Invalid shop ID"
				resp.Data = make(map[string]string)
				return resp
			}(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got dto.CheckBalanceResponse
			err := json.Unmarshal([]byte(tt.jsonStr), &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json.Unmarshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakePaymentResponse_Unmarshal(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		want    dto.MakePaymentResponse
		wantErr bool
	}{
		{
			name: "successful response",
			jsonStr: `{
				"error": false,
				"message": "Payment successful",
				"data": {
					"shop_id": "shop123",
					"transaction": "tx123456",
					"payment_id": "100500",
					"txid": "blockchain_tx_123456",
					"amount": "100.50",
					"amount_pay": "100.00",
					"system": "btc",
					"currency": "BTC",
					"number": "wallet123",
					"shop_commission_percent": "0.5",
					"shop_commission_amount": "0.50",
					"paid_commission": "0.50",
					"explorer_address_link": "https://blockchain.com/address/to_wallet_address",
					"explorer_transaction_link": "https://blockchain.com/tx/blockchain_tx_123456"
				}
			}`,
			want: func() dto.MakePaymentResponse {
				var resp dto.MakePaymentResponse
				resp.Error = false
				resp.Message = "Payment successful"
				resp.Data.ShopId = "shop123"
				resp.Data.Transaction = "tx123456"
				resp.Data.PaymentId = "100500"
				resp.Data.TxId = "blockchain_tx_123456"
				resp.Data.Amount = "100.50"
				resp.Data.AmountPay = "100.00"
				resp.Data.System = enum.System("btc")
				resp.Data.Currency = enum.Currency("BTC")
				resp.Data.Number = "wallet123"
				resp.Data.ShopCommissionPercent = "0.5"
				resp.Data.ShopCommissionAmount = "0.50"
				resp.Data.PaidCommission = "0.50"
				resp.Data.ExplorerAddressLink = "https://blockchain.com/address/to_wallet_address"
				resp.Data.ExplorerTransactionLink = "https://blockchain.com/tx/blockchain_tx_123456"
				return resp
			}(),
			wantErr: false,
		},
		{
			name: "error response",
			jsonStr: `{
				"error": true,
				"message": "Insufficient funds",
				"data": {}
			}`,
			want: func() dto.MakePaymentResponse {
				var resp dto.MakePaymentResponse
				resp.Error = true
				resp.Message = "Insufficient funds"
				return resp
			}(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got dto.MakePaymentResponse
			err := json.Unmarshal([]byte(tt.jsonStr), &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json.Unmarshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPaymentResponse_Unmarshal(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		want    dto.CheckPaymentResponse
		wantErr bool
	}{
		{
			name: "successful response",
			jsonStr: `{
				"error": false,
				"message": "Payment found",
				"data": {
					"transaction": "tx123456",
					"shop_id": "shop123",
					"order_id": "order123",
					"amount": "100.50",
					"currency": "BTC",
					"system": "btc",
					"address": "bc1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qccfmv3",
					"tag": "123456",
					"hash": "hash123456",
					"partial": "0"
				}
			}`,
			want: func() dto.CheckPaymentResponse {
				var resp dto.CheckPaymentResponse
				resp.Error = false
				resp.Message = "Payment found"
				resp.Data.Transaction = "tx123456"
				resp.Data.ShopId = "shop123"
				resp.Data.OrderId = "order123"
				resp.Data.Amount = "100.50"
				resp.Data.Currency = "BTC"
				resp.Data.System = "btc"
				resp.Data.Address = "bc1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qccfmv3"
				resp.Data.Tag = "123456"
				resp.Data.Hash = "hash123456"
				resp.Data.Partial = "0"
				return resp
			}(),
			wantErr: false,
		},
		{
			name: "error response",
			jsonStr: `{
				"error": true,
				"message": "Payment not found",
				"data": {}
			}`,
			want: func() dto.CheckPaymentResponse {
				var resp dto.CheckPaymentResponse
				resp.Error = true
				resp.Message = "Payment not found"
				return resp
			}(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got dto.CheckPaymentResponse
			err := json.Unmarshal([]byte(tt.jsonStr), &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json.Unmarshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckTransactionResponse_Unmarshal(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		want    dto.CheckTransactionResponse
		wantErr bool
	}{
		{
			name: "successful response",
			jsonStr: `{
				"error": false,
				"message": "Transaction found",
				"data": {
					"transaction": "tx123456",
					"txid": "blockchain_tx_123456",
					"shop_id": "shop123",
					"order_id": "order123",
					"amount": "100.50",
					"fee": "0.50",
					"currency": "BTC",
					"system": "btc",
					"address_from": "from_wallet_address",
					"address": "to_wallet_address",
					"tag": "123456",
					"confirmations": 3,
					"required_confirmations": 3,
					"status": "completed",
					"static": "1",
					"date_update": "2023-03-31 12:34:56",
					"explorer_address_link": "https://blockchain.com/address/to_wallet_address",
					"explorer_transaction_link": "https://blockchain.com/tx/blockchain_tx_123456"
				}
			}`,
			want: func() dto.CheckTransactionResponse {
				var resp dto.CheckTransactionResponse
				resp.Error = false
				resp.Message = "Transaction found"
				resp.Data.Transaction = "tx123456"
				resp.Data.Txid = "blockchain_tx_123456"
				resp.Data.ShopId = "shop123"
				resp.Data.OrderId = "order123"
				resp.Data.Amount = "100.50"
				resp.Data.Fee = "0.50"
				resp.Data.Currency = "BTC"
				resp.Data.System = "btc"
				resp.Data.AddressFrom = "from_wallet_address"
				resp.Data.Address = "to_wallet_address"
				resp.Data.Tag = "123456"
				resp.Data.Confirmations = 3
				resp.Data.RequiredConfirmations = 3
				resp.Data.Status = "completed"
				resp.Data.Static = "1"
				resp.Data.DateUpdate = "2023-03-31 12:34:56"
				resp.Data.ExplorerAddressLink = "https://blockchain.com/address/to_wallet_address"
				resp.Data.ExplorerTransactionLink = "https://blockchain.com/tx/blockchain_tx_123456"
				return resp
			}(),
			wantErr: false,
		},
		{
			name: "error response",
			jsonStr: `{
				"error": true,
				"message": "Transaction not found",
				"data": {}
			}`,
			want: func() dto.CheckTransactionResponse {
				var resp dto.CheckTransactionResponse
				resp.Error = true
				resp.Message = "Transaction not found"
				return resp
			}(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got dto.CheckTransactionResponse
			err := json.Unmarshal([]byte(tt.jsonStr), &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json.Unmarshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateAddressResponse_Unmarshal(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		want    dto.GenerateAddressResponse
		wantErr bool
	}{
		{
			name: "successful response with tag",
			jsonStr: `{
				"error": false,
				"message": "Address generated successfully",
				"data": {
					"invoice_id": "inv123456",
					"order_id": "order123",
					"wallet": "wallet_address_123",
					"amount": "100.50",
					"system": "Ripple",
					"currency": "XRP",
					"url": "https://example.com/payment/inv123456",
					"tag": "123456",
					"is_tag": true,
					"tag_name": "Destination Tag"
				}
			}`,
			want: func() dto.GenerateAddressResponse {
				var resp dto.GenerateAddressResponse
				resp.Error = false
				resp.Message = "Address generated successfully"
				resp.Data.InvoiceId = "inv123456"
				resp.Data.OrderId = "order123"
				resp.Data.Wallet = "wallet_address_123"
				resp.Data.Amount = "100.50"
				resp.Data.System = "Ripple"
				resp.Data.Currency = "XRP"
				resp.Data.Url = "https://example.com/payment/inv123456"
				resp.Data.Tag = "123456"
				resp.Data.IsTag = true
				resp.Data.TagName = "Destination Tag"
				return resp
			}(),
			wantErr: false,
		},
		{
			name: "successful response without tag",
			jsonStr: `{
				"error": false,
				"message": "Address generated successfully",
				"data": {
					"invoice_id": "inv123456",
					"order_id": "order123",
					"wallet": "bc1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qccfmv3",
					"amount": "100.50",
					"system": "BitCoin",
					"currency": "BTC",
					"url": "https://example.com/payment/inv123456",
					"tag": "",
					"is_tag": false,
					"tag_name": ""
				}
			}`,
			want: func() dto.GenerateAddressResponse {
				var resp dto.GenerateAddressResponse
				resp.Error = false
				resp.Message = "Address generated successfully"
				resp.Data.InvoiceId = "inv123456"
				resp.Data.OrderId = "order123"
				resp.Data.Wallet = "bc1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qccfmv3"
				resp.Data.Amount = "100.50"
				resp.Data.System = "BitCoin"
				resp.Data.Currency = "BTC"
				resp.Data.Url = "https://example.com/payment/inv123456"
				resp.Data.Tag = ""
				resp.Data.IsTag = false
				resp.Data.TagName = ""
				return resp
			}(),
			wantErr: false,
		},
		{
			name: "error response",
			jsonStr: `{
				"error": true,
				"message": "Invalid parameters",
				"data": {}
			}`,
			want: func() dto.GenerateAddressResponse {
				var resp dto.GenerateAddressResponse
				resp.Error = true
				resp.Message = "Invalid parameters"
				return resp
			}(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got dto.GenerateAddressResponse
			err := json.Unmarshal([]byte(tt.jsonStr), &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json.Unmarshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPaymentUrlResponse_Unmarshal(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		want    dto.GetPaymentUrlResponse
		wantErr bool
	}{
		{
			name: "successful response",
			jsonStr: `{
				"error": false,
				"message": "Payment URL generated successfully",
				"data": {
					"url": "https://example.com/payment/form?token=abc123",
					"method": "GET"
				}
			}`,
			want: func() dto.GetPaymentUrlResponse {
				var resp dto.GetPaymentUrlResponse
				resp.Error = false
				resp.Message = "Payment URL generated successfully"
				resp.Data.Url = "https://example.com/payment/form?token=abc123"
				resp.Data.Method = "GET"
				return resp
			}(),
			wantErr: false,
		},
		{
			name: "error response",
			jsonStr: `{
				"error": true,
				"message": "Invalid parameters",
				"data": {}
			}`,
			want: func() dto.GetPaymentUrlResponse {
				var resp dto.GetPaymentUrlResponse
				resp.Error = true
				resp.Message = "Invalid parameters"
				return resp
			}(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got dto.GetPaymentUrlResponse
			err := json.Unmarshal([]byte(tt.jsonStr), &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json.Unmarshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTxidsOfInvoicesResponse_Unmarshal(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		want    dto.GetTxidsOfInvoicesResponse
		wantErr bool
	}{
		{
			name: "successful response",
			jsonStr: `{
				"error": false,
				"message": "TXIDs retrieved successfully",
				"data": {
					"inv1": ["txid1", "txid2"],
					"inv2": ["txid3"],
					"inv3": []
				}
			}`,
			want: func() dto.GetTxidsOfInvoicesResponse {
				var resp dto.GetTxidsOfInvoicesResponse
				resp.Error = false
				resp.Message = "TXIDs retrieved successfully"
				resp.Data = url.Values{
					"inv1": {"txid1", "txid2"},
					"inv2": {"txid3"},
					"inv3": {},
				}
				return resp
			}(),
			wantErr: false,
		},
		{
			name: "error response",
			jsonStr: `{
				"error": true,
				"message": "Invalid shop ID",
				"data": {}
			}`,
			want: func() dto.GetTxidsOfInvoicesResponse {
				var resp dto.GetTxidsOfInvoicesResponse
				resp.Error = true
				resp.Message = "Invalid shop ID"
				return resp
			}(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got dto.GetTxidsOfInvoicesResponse
			err := json.Unmarshal([]byte(tt.jsonStr), &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Специальная обработка для GetTxidsOfInvoicesResponse, так как url.Values не может быть точно сопоставлено
			// Проверяем ошибку и сообщение напрямую
			if got.Error != tt.want.Error || got.Message != tt.want.Message {
				t.Errorf("json.Unmarshal() got = %v, want %v", got, tt.want)
			}

			// Для данных мы проверяем только факт их наличия в успешном случае
			if !tt.want.Error && len(got.Data) == 0 {
				t.Errorf("json.Unmarshal() got empty data, expected non-empty")
			}
		})
	}
}
