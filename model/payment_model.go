package model

import "time"

type Payment struct {
	Id         int
	CustomerId int
	Paid       int64
	CreatedBy  string
	CreatedAt  time.Time
}
