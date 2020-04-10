package entity

import (
	"errors"
	"fmt"
	"time"
)

type Shop struct {
	ID        int32
	Name      string
	Address   string
	Slug      string
	PhotoURL  string
	Location  GPS
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GPS struct {
	Latitude  float64
	Longitude float64
}

func (s Shop) Validate() error {
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.Slug == "" {
		return errors.New("slug is required")
	}
	if s.Address == "" {
		return errors.New("address is required")
	}
	return nil
}

func (s Shop) ValidateProducts() error {
	for i, item := range s.Products {
		if err := item.Validate(); err != nil {
			return errors.New(fmt.Sprintf("product #%d: %s", i, err.Error()))
		}
	}
	return nil
}

func (s Shop) ZeroID() Shop {
	s.ID = 0
	return s
}

func (s Shop) ZeroProductsID() Shop {
	for i, item := range s.Products {
		item.ID = 0
		s.Products[i] = item
	}
	return s
}

func (s Shop) AssertNotZeroID() error {
	if s.ID == 0 {
		return errors.New("ID should not be 0")
	}
	return nil
}

func (s Shop) AssertProductsNotZeroID() error {
	for i, item := range s.Products {
		if err := item.AssertNoZeroID(); err != nil {
			return errors.New(fmt.Sprintf("product #%d: %s", i, err.Error()))
		}
	}
	return nil
}
