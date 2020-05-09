package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/connection"
)

type SaldoHistoryWriter struct {
	db connection.Gormw
}

func NewSaldoHistoryWriter(db connection.Gormw) *SaldoHistoryWriter {
	return &SaldoHistoryWriter{db: db}
}

func (shw *SaldoHistoryWriter) Create(history entity.SaldoHistory) (entity.SaldoHistory, error) {
	return entity.SaldoHistory{}, nil
}
