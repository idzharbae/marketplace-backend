package converter

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
)

func SaldoHistoryProtoToEntity(in *authproto.SaldoHistory) entity.SaldoHistory {
	return entity.SaldoHistory{
		ID:           in.GetId(),
		UserID:       in.GetUserId(),
		SourceID:     in.GetSourceId(),
		Description:  in.GetDescription(),
		ChangeAmount: in.GetChangeAmount(),
	}
}

func SaldoHistoryEntityToProto(req entity.SaldoHistory) *authproto.SaldoHistory {
	return &authproto.SaldoHistory{
		Id:           req.ID,
		UserId:       req.UserID,
		SourceId:     req.SourceID,
		Description:  req.Description,
		ChangeAmount: req.ChangeAmount,
		CreatedAt:    req.CreatedAt.Unix(),
		UpdatedAt:    req.UpdatedAt.Unix(),
	}
}

func SaldoHistoryEntitiesToProtos(req []entity.SaldoHistory) []*authproto.SaldoHistory {
	histories := make([]*authproto.SaldoHistory, len(req))
	for i, history := range req {
		histories[i] = SaldoHistoryEntityToProto(history)
	}
	return histories
}
