package service

import (
	"context"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
)

type SaldoHistoryService struct {
	uc internal.SaldoHistoryUC
}

func NewSaldoHistoryService(uc internal.SaldoHistoryUC) *SaldoHistoryService {
	return &SaldoHistoryService{uc: uc}
}

func (sh *SaldoHistoryService) ListSaldoHistory(ctx context.Context, in *authproto.ListSaldoHistoryReq) (*authproto.ListSaldoHistoryResp, error) {
	return nil, nil
}

func (sh *SaldoHistoryService) CreateSaldoHistory(ctx context.Context, in *authproto.SaldoHistory) (*authproto.SaldoHistory, error) {
	return nil, nil
}
