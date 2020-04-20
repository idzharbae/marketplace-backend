package model

import "time"

type Payment struct {
	ID            int64
	OrderID       int64
	Amount        int64
	PaymentMethod int32
	PaymentStatus int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
