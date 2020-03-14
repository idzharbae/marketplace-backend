package entity

import "time"

type Shop struct {
	ID        int64
	Name      string
	Address   string
	Location  GPS
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GPS struct {
	Latitude  float64
	Longitude float64
}
