package model

import "time"

type Product struct {
	ID          int32
	ShopID      int32
	Name        string
	Slug        string
	Quantity    int32
	PricePerKG  int32
	StockKG     float32
	Description string
	PhotoURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (c Product) TableName() string {
	return "product"
}
