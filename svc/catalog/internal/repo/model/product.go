package model

import (
	"time"
)

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
	Category    string
	Reviews     []Review `gorm:"ForeignKey:ProductID"`
}

func (p Product) TableName() string {
	return "product"
}

func (p Product) TotalReviews() int32 {
	return int32(len(p.Reviews))
}
func (p Product) AverageRating() float32 {
	if p.TotalReviews() == 0 {
		return 0
	}
	ratings := float64(0)
	for _, review := range p.Reviews {
		ratings += review.Rating
	}
	return float32(ratings) / float32(p.TotalReviews())
}
