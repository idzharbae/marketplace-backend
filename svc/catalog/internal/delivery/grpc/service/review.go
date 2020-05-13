package service

import (
	"context"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal"
)

type ReviewService struct {
	ReviewUC internal.ReviewUC
}

func NewReviewService(reviewUC internal.ReviewUC) *ReviewService {
	return &ReviewService{ReviewUC: reviewUC}
}

func (r *ReviewService) ListReviews(context.Context, *catalogproto.ListReviewsReq) (*catalogproto.ListReviewsResp, error) {
	return nil, nil
}
func (r *ReviewService) GetReview(context.Context, *catalogproto.GetReviewReq) (*catalogproto.Review, error) {
	return nil, nil
}
func (r *ReviewService) CreateReview(context.Context, *catalogproto.Review) (*catalogproto.Review, error) {
	return nil, nil
}
func (r *ReviewService) UpdateReview(context.Context, *catalogproto.Review) (*catalogproto.Review, error) {
	return nil, nil
}
func (r *ReviewService) DeleteReview(context.Context, *catalogproto.GetReviewReq) (*catalogproto.Empty, error) {
	return nil, nil
}
