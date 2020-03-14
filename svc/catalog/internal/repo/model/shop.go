package model

import (
	"github.com/tokopedia/intools/src/b2b-eproc-new/models"
	"time"
)

type Shop struct {
	ID        int64
	Name      string
	Address   string
	Longitude float64
	Latitude  float64
	CreatedAt time.Time
	UpdatedAt time.Time
	Products  []models.Product
}
