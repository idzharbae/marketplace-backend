package repo

import (
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection/gormmock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/catalog/util/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type reviewReaderTest struct {
	ctrl *gomock.Controller
	db   *gormmock.MockGormw
	unit *ReviewReader
}

func newReviewReaderTest() *reviewReaderTest {
	return &reviewReaderTest{}
}

func (rt *reviewReaderTest) Begin(t *testing.T) {
	rt.ctrl = gomock.NewController(t)
	rt.db = gormmock.NewMockGormw(rt.ctrl)
	rt.unit = NewReviewReader(rt.db)
}

func (rt *reviewReaderTest) Finish() {
	rt.ctrl.Finish()
}

func TestReviewReader_GetByID(t *testing.T) {
	test := newReviewReaderTest()
	t.Run("db returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{ID: 123}
		test.db.EXPECT().Where("id=?", req.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.GetByID(req.ID)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Review{}, got)
	})
	t.Run("db returns no error, should return review entity", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{ID: 123}
		resp := model.Review{
			ID:        123,
			UserID:    4,
			ProductID: 5,
			ShopID:    6,
			Title:     "7",
			Content:   "8",
			PhotoURL:  "9",
			Rating:    4.4,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		test.db.EXPECT().Where("id=?", req.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.Review) *gormmock.MockGormw {
			*arg = resp
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.GetByID(req.ID)
		assert.Nil(t, err)
		assert.Equal(t, resp.Rating, got.Rating)
		assert.Equal(t, resp.Content, got.Content)
	})
}

func TestReviewReader_ListByProductID(t *testing.T) {
	test := newReviewReaderTest()
	t.Run("db returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)
		test.db.EXPECT().Where("product_id=?", req).Return(test.db)
		test.db.EXPECT().Find(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.ListByProductID(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("db returns no error, should return review entity", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)
		test.db.EXPECT().Where("product_id=?", req).Return(test.db)
		test.db.EXPECT().Find(gomock.Any()).DoAndReturn(func(arg *[]model.Review) *gormmock.MockGormw {
			*arg = []model.Review{
				{ID: 123},
				{ID: 321},
			}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.ListByProductID(req)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(got))
	})
}

func TestReviewReader_ListByShopID(t *testing.T) {
	test := newReviewReaderTest()
	t.Run("db returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)
		test.db.EXPECT().Where("shop_id=?", req).Return(test.db)
		test.db.EXPECT().Find(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.ListByShopID(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("db returns no error, should return review entity", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)
		test.db.EXPECT().Where("shop_id=?", req).Return(test.db)
		test.db.EXPECT().Find(gomock.Any()).DoAndReturn(func(arg *[]model.Review) *gormmock.MockGormw {
			*arg = []model.Review{
				{ID: 123},
				{ID: 321},
			}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.ListByShopID(req)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(got))
	})
}
