package model

import "time"

type PaymentModel struct {
	Id            int
	CustomerId    int
	Paid          int64
	DestinationId int
	CreatedBy     string
	CreatedAt     time.Time
}
