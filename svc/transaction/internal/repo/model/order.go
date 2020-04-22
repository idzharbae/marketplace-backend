package model

import (
	"github.com/lib/pq"
	"time"
)

type Order struct {
	ID         int64
	ProductID  pq.Int64Array
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
