package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/converter"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
	"github.com/jinzhu/gorm"
)

type SaldoHistoryReader struct {
	db connection.Gormw
}

func NewSaldoHistoryReader(db connection.Gormw) *SaldoHistoryReader {
	return &SaldoHistoryReader{db: db}
}

func (shr *SaldoHistoryReader) ListByUserID(req request.ListSaldoHistory) ([]entity.SaldoHistory, error) {
	var saldoHistories []model.SaldoHistory
	db := shr.db.Where("user_id=?", req.UserID)
	if req.Pagination.Limit > 0 {
		db = db.Limit(req.Pagination.Limit)
	}
	if req.Pagination.Page > 0 {
		db = db.Offset(req.Pagination.OffsetFromPagination())
	}
	err := db.Find(&saldoHistories).Error()
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return converter.SaldoHistoryModelsToEntities(saldoHistories), nil
}
