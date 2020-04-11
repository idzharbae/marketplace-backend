package entity

import (
	"github.com/pkg/errors"
	"time"
)

type Product struct {
	ID          int32
	ShopID      int32
	Name        string
	Description string
	Slug        string
	Quantity    int32
	PricePerKG  int32
	StockKG     float32
	PhotoURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p Product) Validate() error {
	if p.Quantity < 0 {
		return errors.New("quantity can't be negative")
	}
	if p.PricePerKG < 0 {
		return errors.New("price can't be negative")
	}
	if p.StockKG < 0 {
		return errors.New("stock can't be negative")
	}
	if p.Name == "" {
		return errors.New("name is required")
	}
	if p.Slug == "" {
		return errors.New("slug is required")
	}
	return nil
}

func (p Product) ZeroID() Product {
	p.ID = 0
	return p
}

func (p Product) AssertNoZeroID() error {
	if p.ID == 0 {
		return errors.New("id should not be 0")
	}
	return nil
}
