package model

import "time"

type Order struct {
	ID         int64
	ProductID  []int64
	UserID     int64
	ShopID     int64
	TotalPrice int64
	Status     int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (o Order) TableName() string {
	return "orders"
}
