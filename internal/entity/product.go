package entity

import "time"

type Product struct {
	ID int32
	ShopID int32
	Name string
	Quantity int32
	PricePerKG int32
	StockKG float32
	CreatedAt time.Time
	UpdatedAt time.Time
}