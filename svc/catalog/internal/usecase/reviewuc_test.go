package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/constant"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/repomock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type reviewTest struct {
	ctrl   *gomock.Controller
	reader *repomock.MockReviewReader
	writer *repomock.MockReviewWriter
	unit   *Review
}

func newReviewTest() *reviewTest {
	return &reviewTest{}
}

func (rt *reviewTest) Begin(t *testing.T) {
	rt.ctrl = gomock.NewController(t)
	rt.reader = repomock.NewMockReviewReader(rt.ctrl)
	rt.writer = repomock.NewMockReviewWriter(rt.ctrl)
	rt.unit = NewReview(rt.reader, rt.writer)
}

func (rt *reviewTest) Finish() {
	rt.ctrl.Finish()
}

func TestReview_Create(t *testing.T) {
	test := newReviewTest()
	t.Run("rating less than 0, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{
			ID:        12,
			UserID:    12,
			ProductID: 12,
			ShopID:    12,
			Title:     "Asdf",
			Content:   "asdf",
			PhotoURL:  "asdf",
			Rating:    -1,
		}
		got, err := test.unit.Create(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Review{}, got)
	})
	t.Run("rating more than maximum rating, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{
			ID:        12,
			UserID:    12,
			ProductID: 12,
			ShopID:    12,
			Title:     "Asdf",
			Content:   "asdf",
			PhotoURL:  "asdf",
			Rating:    constant.MaxRatingValue + 0.1,
		}
		got, err := test.unit.Create(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Review{}, got)
	})
	t.Run("reader returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{
			ID:        123,
			UserID:    32,
			ProductID: 21,
			ShopID:    421,
			Title:     "asdf",
			Content:   "asd",
			PhotoURL:  "asd",
			Rating:    4.2,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}
		test.writer.EXPECT().Create(gomock.Any()).DoAndReturn(func(arg entity.Review) (entity.Review, error) {
			assert.Equal(t, int64(0), arg.ID)
			assert.Equal(t, req.UserID, arg.UserID)
			assert.Equal(t, req.Title, arg.Title)
			return entity.Review{}, errors.New("error")
		})

		got, err := test.unit.Create(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Review{}, got)
	})
	t.Run("reader returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{
			ID:        123,
			UserID:    32,
			ProductID: 21,
			ShopID:    421,
			Title:     "asdf",
			Content:   "asd",
			PhotoURL:  "asd",
			Rating:    4.2,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}
		test.writer.EXPECT().Create(gomock.Any()).DoAndReturn(func(arg entity.Review) (entity.Review, error) {
			assert.Equal(t, int64(0), arg.ID)
			assert.Equal(t, req.UserID, arg.UserID)
			assert.Equal(t, req.Title, arg.Title)
			arg = req
			arg.ID = 123
			return arg, nil
		})

		got, err := test.unit.Create(req)
		assert.Nil(t, err)
		assert.Equal(t, int64(123), got.ID)
	})
}
func TestReview_Update(t *testing.T) {
	test := newReviewTest()
	t.Run("rating less than 0, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{
			ID:        12,
			UserID:    12,
			ProductID: 12,
			ShopID:    12,
			Title:     "Asdf",
			Content:   "asdf",
			PhotoURL:  "asdf",
			Rating:    -1,
		}
		got, err := test.unit.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Review{}, got)
	})
	t.Run("rating more than maximum rating, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{
			ID:        12,
			UserID:    12,
			ProductID: 12,
			ShopID:    12,
			Title:     "Asdf",
			Content:   "asdf",
			PhotoURL:  "asdf",
			Rating:    constant.MaxRatingValue + 0.1,
		}
		got, err := test.unit.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Review{}, got)
	})
	t.Run("reader returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{
			ID:        123,
			UserID:    32,
			ProductID: 21,
			ShopID:    421,
			Title:     "asdf",
			Content:   "asd",
			PhotoURL:  "asd",
			Rating:    4.2,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}
		test.writer.EXPECT().Update(req).Return(entity.Review{}, errors.New("error"))

		got, err := test.unit.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Review{}, got)
	})
	t.Run("reader returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{
			ID:        123,
			UserID:    32,
			ProductID: 21,
			ShopID:    421,
			Title:     "asdf",
			Content:   "asd",
			PhotoURL:  "asd",
			Rating:    4.2,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}
		test.writer.EXPECT().Update(req).Return(entity.Review{ID: 321}, nil)

		got, err := test.unit.Update(req)
		assert.Nil(t, err)
		assert.Equal(t, int64(321), got.ID)
	})
}

func TestReview_List(t *testing.T) {
	test := newReviewTest()
	t.Run("given product id, list by product id", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := requests.ListReview{
			ProductID: 123,
		}
		test.reader.EXPECT().ListByProductID(req.ProductID, req.Pagination)
		test.unit.List(req)
	})
	t.Run("given shop id, list by shop id", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := requests.ListReview{
			ShopID: 123,
		}
		test.reader.EXPECT().ListByShopID(req.ShopID, req.Pagination)
		test.unit.List(req)
	})
}
