package entity

import "time"

type Order struct {
	ID         int64
	UserID     int64
	Products   []Product
	TotalPrice int64
	Payment    Payment
	Status     int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Payment struct {
	ID            int64
	Amount        int64
	PaymentMethod int32
	PaymentStatus int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
