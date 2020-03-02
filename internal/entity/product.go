package entity

import "time"

type Product struct {
	ID         int32
	ShopID     int32
	Name       string
	Slug       string
	Quantity   int32
	PricePerKG int32
	StockKG    float32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (p *Product) Validate() bool {
	return p.Quantity >= 0 && p.PricePerKG >= 0 && p.StockKG >= 0
}
