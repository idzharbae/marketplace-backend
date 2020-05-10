package usecase

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
)

type SaldoHistory struct {
	saldoReader internal.SaldoHistoryReader
	saldoWriter internal.SaldoHistoryWriter
}

func NewSaldoHistory(saldoReader internal.SaldoHistoryReader, saldoWriter internal.SaldoHistoryWriter) *SaldoHistory {
	return &SaldoHistory{
		saldoReader: saldoReader,
		saldoWriter: saldoWriter,
	}
}

func (sh *SaldoHistory) Create(req entity.SaldoHistory) (entity.SaldoHistory, error) {
	return sh.saldoWriter.Create(req)
}

func (sh *SaldoHistory) List(req request.ListSaldoHistory) ([]entity.SaldoHistory, error) {
	return sh.saldoReader.ListByUserID(req)
}
