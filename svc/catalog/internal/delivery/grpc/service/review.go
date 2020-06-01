package service

import (
	"context"
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
)

type ReviewService struct {
	ReviewUC internal.ReviewUC
}

func NewReviewService(reviewUC internal.ReviewUC) *ReviewService {
	return &ReviewService{ReviewUC: reviewUC}
}

func (r *ReviewService) ListReviews(ctx context.Context, in *catalogproto.ListReviewsReq) (*catalogproto.ListReviewsResp, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	res, err := r.ReviewUC.List(requests.ListReview{
		ShopID:    in.GetShopId(),
		ProductID: in.GetProductId(),
		Pagination: requests.Pagination{
			Page:  int(in.GetPagination().GetPage()),
			Limit: int(in.GetPagination().GetLimit()),
		},
	})
	if err != nil {
		return nil, err
	}
	return &catalogproto.ListReviewsResp{
		Reviews: converter.ReviewEntitiesToProtos(res),
	}, nil
}
func (r *ReviewService) GetReview(ctx context.Context, in *catalogproto.GetReviewReq) (*catalogproto.Review, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	res, err := r.ReviewUC.Get(requests.GetReview{
		ReviewID:   in.GetReviewId(),
		CustomerID: in.GetCustomerId(),
		ProductID:  in.GetProductId(),
	})
	if err != nil {
		return nil, err
	}
	return converter.ReviewEntityToProto(res), nil
}
func (r *ReviewService) CreateReview(ctx context.Context, in *catalogproto.Review) (*catalogproto.Review, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	res, err := r.ReviewUC.Create(converter.ReviewProtoToEntity(in))
	if err != nil {
		return nil, err
	}
	return converter.ReviewEntityToProto(res), nil
}
func (r *ReviewService) UpdateReview(ctx context.Context, in *catalogproto.Review) (*catalogproto.Review, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	res, err := r.ReviewUC.Update(converter.ReviewProtoToEntity(in))
	if err != nil {
		return nil, err
	}
	return converter.ReviewEntityToProto(res), nil
}
func (r *ReviewService) DeleteReview(ctx context.Context, in *catalogproto.Review) (*catalogproto.Empty, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	err := r.ReviewUC.Delete(entity.Review{ID: in.GetId(), UserID: in.GetUserId()})
	if err != nil {
		return nil, err
	}
	return &catalogproto.Empty{}, nil
}
