package repo

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection/gormmock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type shopReaderTest struct {
	Ctrl *gomock.Controller
	DB   *gormmock.MockGormw
	Unit *ShopReader
}

func newShopReaderTest() *shopReaderTest {
	return &shopReaderTest{}
}

func (srt *shopReaderTest) Begin(t *testing.T) {
	srt.Ctrl = gomock.NewController(t)
	srt.DB = gormmock.NewMockGormw(srt.Ctrl)
	srt.Unit = NewShopReader(srt.DB)
}

func (srt *shopReaderTest) Finish() {
	srt.Ctrl.Finish()
}

func TestShopReader_GetByID(t *testing.T) {
	test := newShopReaderTest()
	t.Run("db returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		shopID := int32(1337)

		test.DB.EXPECT().Where("id=?", shopID).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(errors.New("error"))

		got, err := test.Unit.GetByID(shopID)
		assert.Equal(t, entity.Shop{}, got)
		assert.NotNil(t, err)
	})
	t.Run("db returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		shopID := int32(1337)
		result := model.Shop{
			ID:        1337,
			Name:      "asdf",
			Address:   "asdf",
			Slug:      "sdddsd",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		test.DB.EXPECT().Where("id=?", shopID).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).DoAndReturn(func(res *model.Shop) *gormmock.MockGormw {
			*res = result
			return test.DB
		})
		test.DB.EXPECT().Error().Return(nil)

		got, err := test.Unit.GetByID(shopID)
		assert.Equal(t, converter.ShopModelToEntity(result), got)
		assert.Nil(t, err)
	})
}

func TestShopReader_GetBySlug(t *testing.T) {
	test := newShopReaderTest()
	t.Run("db returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		slug := "slug"

		test.DB.EXPECT().Where("slug=?", slug).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(errors.New("error"))

		got, err := test.Unit.GetBySlug(slug)
		assert.Equal(t, entity.Shop{}, got)
		assert.NotNil(t, err)
	})
	t.Run("db returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		slug := "slug"
		result := model.Shop{
			ID:        1337,
			Name:      "asdf",
			Address:   "asdf",
			Slug:      slug,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		test.DB.EXPECT().Where("slug=?", slug).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).DoAndReturn(func(res *model.Shop) *gormmock.MockGormw {
			*res = result
			return test.DB
		})
		test.DB.EXPECT().Error().Return(nil)

		got, err := test.Unit.GetBySlug(slug)
		assert.Equal(t, converter.ShopModelToEntity(result), got)
		assert.Nil(t, err)
	})
}

func TestShopReader_List(t *testing.T) {
	test := newShopReaderTest()
	t.Run("given page = 1 should not exec offset", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := requests.ListShop{Pagination: requests.Pagination{
			Page:  1,
			Limit: 10,
		}}
		test.DB.EXPECT().Limit(req.Pagination.Limit).Return(test.DB)
		test.DB.EXPECT().Find(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(nil)
		test.Unit.ListAll(req.Pagination)
	})
	t.Run("given page > 1 should exec offset", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := requests.ListShop{Pagination: requests.Pagination{
			Page:  1337,
			Limit: 10,
		}}
		test.DB.EXPECT().Limit(req.Pagination.Limit).Return(test.DB)
		test.DB.EXPECT().Offset((req.Pagination.Page - 1) * req.Pagination.Limit).Return(test.DB)
		test.DB.EXPECT().Find(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(nil)
		test.Unit.ListAll(req.Pagination)
	})
	t.Run("db returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := requests.ListShop{Pagination: requests.Pagination{
			Page:  3,
			Limit: 10,
		}}

		test.DB.EXPECT().Limit(req.Pagination.Limit).Return(test.DB)
		test.DB.EXPECT().Offset((req.Pagination.Page - 1) * req.Pagination.Limit).Return(test.DB)
		test.DB.EXPECT().Find(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(errors.New("error"))

		got, err := test.Unit.ListAll(req.Pagination)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("db returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := requests.ListShop{Pagination: requests.Pagination{
			Page:  2,
			Limit: 10,
		}}

		test.DB.EXPECT().Limit(req.Pagination.Limit).Return(test.DB)
		test.DB.EXPECT().Offset((req.Pagination.Page - 1) * req.Pagination.Limit).Return(test.DB)
		test.DB.EXPECT().Find(gomock.Any()).DoAndReturn(func(shops *[]model.Shop) *gormmock.MockGormw {
			for i := 0; i < req.Pagination.Limit; i++ {
				(*shops)[i] = model.Shop{
					ID:        int64(i + req.Pagination.Page),
					Name:      "test",
					Address:   "test",
					Longitude: 1,
					Latitude:  1,
					Slug:      "sad",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
			}
			return test.DB
		})
		test.DB.EXPECT().Error().Return(nil)

		got, err := test.Unit.ListAll(req.Pagination)
		assert.NotNil(t, got)
		assert.Nil(t, err)
		assert.Equal(t, req.Pagination.Limit, len(got))
	})
}
