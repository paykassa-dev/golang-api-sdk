package priority

import "github.com/paykassa-dev/golang-api-sdk/enum"

const (
	High   = enum.TransactionPriority("high")
	Medium = enum.TransactionPriority("medium")
	Low    = enum.TransactionPriority("low")
)
