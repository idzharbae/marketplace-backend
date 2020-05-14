package entity

import (
	"errors"
	"fmt"
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

func (r Review) ValidateRating() error {
	if r.Rating < 0 {
		return errors.New("rating can't be negative")
	}
	if r.Rating > constant.MaxRatingValue {
		return errors.New(fmt.Sprintf("rating can't be more than %.1f", constant.MaxRatingValue))
	}
	return nil
}
