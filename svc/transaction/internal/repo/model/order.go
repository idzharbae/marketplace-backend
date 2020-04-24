package model

import (
	"time"
)

type Order struct {
	ID            int64
	UserID        int64
	ShopID        int64
	TotalPrice    int64
	Status        int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
	OrderProducts []OrderProduct
}

func (o Order) TableName() string {
	return "orders"
}

func (o Order) GetProductIDs() []int64 {
	res := make([]int64, len(o.OrderProducts))
	for i, product := range o.OrderProducts {
		res[i] = product.ID
	}
	return res
}

type OrderProduct struct {
	ID        int64
	OrderID   int64
	ProductID int64
	AmountKG  float64
}

func (op OrderProduct) TableName() string {
	return "order_products"
}
