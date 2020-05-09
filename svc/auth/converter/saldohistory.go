package converter

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/model"
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

func SaldoHistoryModelToEntity(req model.SaldoHistory) entity.SaldoHistory {
	return entity.SaldoHistory{
		ID:           req.ID,
		UserID:       req.UserID,
		SourceID:     req.SourceID,
		Description:  req.Description,
		ChangeAmount: req.ChangeAmount,
		CreatedAt:    req.CreatedAt,
		UpdatedAt:    req.UpdatedAt,
	}
}
func SaldoHistoryModelsToEntities(req []model.SaldoHistory) []entity.SaldoHistory {
	histories := make([]entity.SaldoHistory, len(req))
	for i, history := range req {
		histories[i] = SaldoHistoryModelToEntity(history)
	}
	return histories
}
func SaldoHistoryEntityToModel(req entity.SaldoHistory) model.SaldoHistory {
	return model.SaldoHistory{
		ID:           req.ID,
		UserID:       req.UserID,
		SourceID:     req.SourceID,
		Description:  req.Description,
		ChangeAmount: req.ChangeAmount,
	}
}
