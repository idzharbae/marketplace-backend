package model

import (
	"time"
)

type Shop struct {
	ID        int64
	Name      string
	Address   string
	Longitude float64
	Latitude  float64
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Products  []Product
}
