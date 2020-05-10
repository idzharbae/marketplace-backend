package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/converter"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/model"
	"github.com/jinzhu/gorm"
)

type SaldoHistoryReader struct {
	db connection.Gormw
}

func NewSaldoHistoryReader(db connection.Gormw) *SaldoHistoryReader {
	return &SaldoHistoryReader{db: db}
}

func (shr *SaldoHistoryReader) ListByUserID(userID int64) ([]entity.SaldoHistory, error) {
	var saldoHistories []model.SaldoHistory
	err := shr.db.Where("user_id=?", userID).Find(&saldoHistories).Error()
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return converter.SaldoHistoryModelsToEntities(saldoHistories), nil
}
