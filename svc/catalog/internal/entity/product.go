package entity

import (
	"github.com/pkg/errors"
	"time"
)

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

func (p *Product) Validate() error {
	if p.Quantity < 0 {
		return errors.New("quantity can't be negative")
	}
	if p.PricePerKG < 0 {
		return errors.New("price can't be negative")
	}
	if p.StockKG < 0 {
		return errors.New("stock can't be negative")
	}
	if p.ShopID <= 0 {
		return errors.New("shop id must be greater than 0")
	}
	return nil
}
