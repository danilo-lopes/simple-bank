package models

import "time"

type Transaction struct {
	Amount   int       `json:"amount,required"`
	Merchant string    `json:"merchant,required"`
	Time     time.Time `json:"localDateTime,omitempty"`
}
