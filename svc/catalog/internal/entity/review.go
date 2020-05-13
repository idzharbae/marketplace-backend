package entity

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/constant"
	"time"
)

type Review struct {
	ID        int64
	UserID    int64
	ProductID int64
	ShopID    int64
	Title     string
	Content   string
	PhotoURL  string
	Rating    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r Review) ValidateRating() bool {
	return r.Rating <= constant.MaxRatingValue
}
