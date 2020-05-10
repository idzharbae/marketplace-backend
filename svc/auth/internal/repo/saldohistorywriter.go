package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/converter"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/connection"
)

type SaldoHistoryWriter struct {
	db connection.Gormw
}

func NewSaldoHistoryWriter(db connection.Gormw) *SaldoHistoryWriter {
	return &SaldoHistoryWriter{db: db}
}

func (shw *SaldoHistoryWriter) Create(req entity.SaldoHistory) (entity.SaldoHistory, error) {
	history := converter.SaldoHistoryEntityToModel(req)
	history.ID = 0
	err := shw.db.Save(&history).Error()
	if err != nil {
		return entity.SaldoHistory{}, err
	}
	return converter.SaldoHistoryModelToEntity(history), nil
}
