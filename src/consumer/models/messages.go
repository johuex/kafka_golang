package models

import "time"

type TransactionType int

const (
	Payment TransactionType = iota
	Authorization
	Capture
	Void
	Refund
	RecurringPayment
)

func (t TransactionType) String() string {
	switch t {
	case Payment:
		return "Payment"
	case Authorization:
		return "Authorization"
	case Capture:
		return "Capture"
	case Void:
		return "Void"
	case Refund:
		return "Refund"
	case RecurringPayment:
		return "RecurringPayment" // will be deprecated on Consumer
	}
	return "Undefined" // will be deprecated on Consumer
}

var Locations = [5]string{
	"LED",
	"MOW",
	"KHV",
	"KZN",
	"LPP", // will be deprecated on Consumer
}

type Transaction struct {
	TransactionType TransactionType `json:"transaction_type"`
	Location        string          `json:"location"`
	TransactionId   int64           `json:"transaction_id"`
	AccountNumber   int64           `json:"account_number"`
	Amount          int             `json:"amount"` // allow only <= 100k
	Time            time.Time       `json:"time"`
}
