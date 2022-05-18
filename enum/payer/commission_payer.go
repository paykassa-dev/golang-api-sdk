package payer

import "github.com/paykassa-dev/golang-api-sdk/enum"

const (
	Shop   = enum.CommissionPayer("shop")
	Client = enum.CommissionPayer("client")
)
