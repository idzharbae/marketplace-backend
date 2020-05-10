package entity

import "time"

type SaldoHistory struct {
	ID           int64
	UserID       int64
	SourceID     int64
	Description  string
	ChangeAmount int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
