package service

import (
	"context"
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/converter"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
)

type SaldoHistoryService struct {
	uc internal.SaldoHistoryUC
}

func NewSaldoHistoryService(uc internal.SaldoHistoryUC) *SaldoHistoryService {
	return &SaldoHistoryService{uc: uc}
}

func (sh *SaldoHistoryService) ListSaldoHistory(ctx context.Context, in *authproto.ListSaldoHistoryReq) (*authproto.ListSaldoHistoryResp, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	res, err := sh.uc.List(in.GetUserId())
	if err != nil {
		return nil, err
	}
	return &authproto.ListSaldoHistoryResp{
		SaldoHistories: converter.SaldoHistoryEntitiesToProtos(res),
	}, nil
}

func (sh *SaldoHistoryService) CreateSaldoHistory(ctx context.Context, in *authproto.SaldoHistory) (*authproto.SaldoHistory, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	res, err := sh.uc.Create(converter.SaldoHistoryProtoToEntity(in))
	if err != nil {
		return nil, err
	}
	return converter.SaldoHistoryEntityToProto(res), nil
}
