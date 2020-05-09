package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/connection"
)

type SaldoHistoryReader struct {
	db connection.Gormw
}

func NewSaldoHistoryReader(db connection.Gormw) *SaldoHistoryReader {
	return &SaldoHistoryReader{db: db}
}

func (shr *SaldoHistoryReader) ListByUserID(userID int64) ([]entity.SaldoHistory, error) {
	return nil, nil
}
