package entity

import "time"

type Order struct {
	ID         int64
	UserID     int64
	ProductIDs []int64
	TotalPrice int64
	Payment    Payment
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Payment struct {
	Amount        int64
	PaymentMethod int32
	PaymentStatus int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
