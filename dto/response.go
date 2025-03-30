package dto

import "github.com/paykassa-dev/golang-api-sdk/enum"

type CheckBalanceResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    struct {
		BitcoinBtc                 string `json:"bitcoin_btc"`
		EthereumEth                string `json:"ethereum_eth"`
		LitecoinLtc                string `json:"litecoin_ltc"`
		DogecoinDoge               string `json:"dogecoin_doge"`
		DashDash                   string `json:"dash_dash"`
		BitcoincashBch             string `json:"bitcoincash_bch"`
		RippleXrp                  string `json:"ripple_xrp"`
		TronTrx                    string `json:"tron_trx"`
		StellarXlm                 string `json:"stellar_xlm"`
		BinancecoinBnb             string `json:"binancecoin_bnb"`
		TronTrc20Usdt              string `json:"tron_trc20_usdt"`
		BinancesmartchainBep20Usdt string `json:"binancesmartchain_bep20_usdt"`
		EthereumErc20Usdt          string `json:"ethereum_erc20_usdt"`
		BinancesmartchainBep20Busd string `json:"binancesmartchain_bep20_busd"`
		BinancesmartchainBep20Usdc string `json:"binancesmartchain_bep20_usdc"`
		BinancesmartchainBep20Ada  string `json:"binancesmartchain_bep20_ada"`
		BinancesmartchainBep20Eos  string `json:"binancesmartchain_bep20_eos"`
		BinancesmartchainBep20Btc  string `json:"binancesmartchain_bep20_btc"`
		BinancesmartchainBep20Eth  string `json:"binancesmartchain_bep20_eth"`
		BinancesmartchainBep20Doge string `json:"binancesmartchain_bep20_doge"`
		TonTon                     string `json:"ton_ton"`
		TonUsdt                    string `json:"ton_usdt"`
	} `json:"data"`
}

type MakePaymentResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    struct {
		ShopId                string        `json:"shop_id"`
		Transaction           string        `json:"transaction"`
		TxId                  string        `json:"txid"`
		Amount                string        `json:"amount"`
		AmountPay             string        `json:"amount_pay"`
		System                enum.System   `json:"system"`
		Currency              enum.Currency `json:"currency"`
		Number                string        `json:"number"`
		ShopCommissionPercent string        `json:"shop_commission_percent"`
		ShopCommissionAmount  string        `json:"shop_commission_amount"`
		PaidCommission        string        `json:"paid_commission"`
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
		Confirmations           string `json:"confirmations"`
		RequiredConfirmations   string `json:"required_confirmations"`
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
