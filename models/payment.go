package models

import "time"

type Payment struct {
	ID          int       `json:"id"`
	BillerID    string    `json:"billerId"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"paymentDate"`
	Status      string    `json:"status"`
}
