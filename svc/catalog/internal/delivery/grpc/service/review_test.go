package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/usecase/ucmock"
	"github.com/stretchr/testify/assert"
)

type reviewTest struct {
	ctrl *gomock.Controller
	uc   *ucmock.MockReviewUC
	unit *ReviewService
	ctx  context.Context
}

func newReviewTest() *reviewTest {
	return &reviewTest{}
}

func (rt *reviewTest) Begin(t *testing.T) {
	rt.ctrl = gomock.NewController(t)
	rt.uc = ucmock.NewMockReviewUC(rt.ctrl)
	rt.unit = NewReviewService(rt.uc)
	rt.ctx = context.Background()
}

func (rt *reviewTest) Finish() {
	rt.ctrl.Finish()
}

func TestReviewService_ListReviews(t *testing.T) {
	test := newReviewTest()
	t.Run("given nil params should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.ListReviews(test.ctx, nil)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := &catalogproto.ListReviewsReq{
			ProductId: 1,
			ShopId:    2,
			Pagination: &catalogproto.Pagination{
				Page:  123,
				Limit: 1213,
			},
		}
		test.uc.EXPECT().List(requests.ListReview{
			ShopID:    req.GetShopId(),
			ProductID: req.GetProductId(),
			Pagination: requests.Pagination{
				Page:  int(req.GetPagination().GetPage()),
				Limit: int(req.GetPagination().GetLimit()),
			},
		}).Return(nil, errors.New("error"))

		got, err := test.unit.ListReviews(test.ctx, req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := &catalogproto.ListReviewsReq{
			ProductId: 1,
			ShopId:    2,
			Pagination: &catalogproto.Pagination{
				Page:  123,
				Limit: 1213,
			},
		}
		test.uc.EXPECT().List(requests.ListReview{
			ShopID:    req.GetShopId(),
			ProductID: req.GetProductId(),
			Pagination: requests.Pagination{
				Page:  int(req.GetPagination().GetPage()),
				Limit: int(req.GetPagination().GetLimit()),
			},
		}).Return([]entity.Review{{ID: 1}, {ID: 2}}, nil)

		got, err := test.unit.ListReviews(test.ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(got.GetReviews()))
	})
}

func TestReviewService_GetReview(t *testing.T) {
	test := newReviewTest()
	t.Run("given nil param should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.GetReview(test.ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &catalogproto.GetReviewReq{
			ReviewId: 123,
		}
		test.uc.EXPECT().Get(req.GetReviewId()).Return(entity.Review{}, errors.New("error"))
		got, err := test.unit.GetReview(test.ctx, req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error, should no return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &catalogproto.GetReviewReq{
			ReviewId: 123,
		}

		resp := entity.Review{
			ID:        123,
			UserID:    123,
			ProductID: 321,
			ShopID:    312,
			Title:     "asdf",
			Content:   "dsas",
			PhotoURL:  "asdsadll",
			Rating:    3.2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		test.uc.EXPECT().Get(req.GetReviewId()).Return(resp, nil)

		got, err := test.unit.GetReview(test.ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, resp.ID, got.GetId())
		assert.Equal(t, resp.UserID, got.GetUserId())
		assert.Equal(t, resp.Title, got.GetTitle())
		assert.Equal(t, resp.CreatedAt.Unix(), got.GetCreatedAt())
	})
}

func TestReviewService_CreateReview(t *testing.T) {
	test := newReviewTest()
	t.Run("given nil param should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.CreateReview(test.ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &catalogproto.Review{
			UserId:    123,
			ProductId: 321,
			ShopId:    123,
			Title:     "asdf",
			Content:   "sdsd",
			PhotoUrl:  "asdasd",
			Rating:    3.0,
		}
		test.uc.EXPECT().Create(converter.ReviewProtoToEntity(req)).Return(entity.Review{}, errors.New("error"))
		got, err := test.unit.CreateReview(test.ctx, req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error, should no return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &catalogproto.Review{
			UserId:    123,
			ProductId: 321,
			ShopId:    123,
			Title:     "asdf",
			Content:   "sdsd",
			PhotoUrl:  "asdasd",
			Rating:    3.0,
		}
		test.uc.EXPECT().Create(converter.ReviewProtoToEntity(req)).Return(entity.Review{}, nil)
		got, err := test.unit.CreateReview(test.ctx, req)
		assert.NotNil(t, got)
		assert.Nil(t, err)
	})
}
func TestReviewService_UpdateReview(t *testing.T) {
	test := newReviewTest()
	t.Run("given nil param should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.UpdateReview(test.ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &catalogproto.Review{
			Id:        123,
			UserId:    123,
			ProductId: 321,
			ShopId:    123,
			Title:     "asdf",
			Content:   "sdsd",
			PhotoUrl:  "asdasd",
			Rating:    3.0,
		}
		test.uc.EXPECT().Update(converter.ReviewProtoToEntity(req)).Return(entity.Review{}, errors.New("error"))
		got, err := test.unit.UpdateReview(test.ctx, req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error, should no return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &catalogproto.Review{
			Id:        123,
			UserId:    123,
			ProductId: 321,
			ShopId:    123,
			Title:     "asdf",
			Content:   "sdsd",
			PhotoUrl:  "asdasd",
			Rating:    3.0,
		}
		test.uc.EXPECT().Update(converter.ReviewProtoToEntity(req)).Return(entity.Review{}, nil)
		got, err := test.unit.UpdateReview(test.ctx, req)
		assert.NotNil(t, got)
		assert.Nil(t, err)
	})
}
func TestReviewService_DeleteReview(t *testing.T) {
	test := newReviewTest()
	t.Run("given nil param should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.DeleteReview(test.ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &catalogproto.Review{
			Id: 123,
		}
		test.uc.EXPECT().Delete(entity.Review{ID: req.GetId()}).Return(errors.New("error"))
		got, err := test.unit.DeleteReview(test.ctx, req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error, should no return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &catalogproto.Review{
			Id: 123,
		}
		test.uc.EXPECT().Delete(entity.Review{ID: req.GetId()}).Return(nil)
		got, err := test.unit.DeleteReview(test.ctx, req)
		assert.Nil(t, err)
		assert.NotNil(t, got)
	})
}
