package model

import "time"

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

func (r Review) TableName() string {
	return "review"
}
