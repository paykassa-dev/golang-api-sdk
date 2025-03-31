package pscur

import "github.com/paykassa-dev/golang-api-sdk/enum"

const (
	BitcoinBtc                 = enum.PsCur("bitcoin_btc")
	EthereumEth                = enum.PsCur("ethereum_eth")
	LitecoinLtc                = enum.PsCur("litecoin_ltc")
	DogecoinDoge               = enum.PsCur("dogecoin_doge")
	DashDash                   = enum.PsCur("dash_dash")
	BitcoincashBch             = enum.PsCur("bitcoincash_bch")
	RippleXrp                  = enum.PsCur("ripple_xrp")
	TronTrx                    = enum.PsCur("tron_trx")
	StellarXlm                 = enum.PsCur("stellar_xlm")
	BinancecoinBnb             = enum.PsCur("binancecoin_bnb")
	TronTrc20Usdt              = enum.PsCur("tron_trc20_usdt")
	BinancesmartchainBep20Usdt = enum.PsCur("binancesmartchain_bep20_usdt")
	EthereumErc20Usdt          = enum.PsCur("ethereum_erc20_usdt")
	BinancesmartchainBep20Busd = enum.PsCur("binancesmartchain_bep20_busd")
	BinancesmartchainBep20Usdc = enum.PsCur("binancesmartchain_bep20_usdc")
	BinancesmartchainBep20Ada  = enum.PsCur("binancesmartchain_bep20_ada")
	BinancesmartchainBep20Eos  = enum.PsCur("binancesmartchain_bep20_eos")
	BinancesmartchainBep20Btc  = enum.PsCur("binancesmartchain_bep20_btc")
	BinancesmartchainBep20Eth  = enum.PsCur("binancesmartchain_bep20_eth")
	BinancesmartchainBep20Doge = enum.PsCur("binancesmartchain_bep20_doge")
	TonTon                     = enum.PsCur("ton_ton")
	TonUsdt                    = enum.PsCur("ton_usdt")
)

// PsCurList contains all supported payment system currencies
var PsCurList = []enum.PsCur{
	BitcoinBtc,
	EthereumEth,
	LitecoinLtc,
	DogecoinDoge,
	DashDash,
	BitcoincashBch,
	RippleXrp,
	TronTrx,
	StellarXlm,
	BinancecoinBnb,
	TronTrc20Usdt,
	BinancesmartchainBep20Usdt,
	EthereumErc20Usdt,
	BinancesmartchainBep20Busd,
	BinancesmartchainBep20Usdc,
	BinancesmartchainBep20Ada,
	BinancesmartchainBep20Eos,
	BinancesmartchainBep20Btc,
	BinancesmartchainBep20Eth,
	BinancesmartchainBep20Doge,
	TonTon,
	TonUsdt,
}
