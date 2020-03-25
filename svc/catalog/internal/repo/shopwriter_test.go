package repo

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection/gormmock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type shopWriterTest struct {
	Ctrl *gomock.Controller
	DB   *gormmock.MockGormw
	Unit *ShopWriter
}

func newShopWriterTest() *shopWriterTest {
	return &shopWriterTest{}
}

func (srt *shopWriterTest) Begin(t *testing.T) {
	srt.Ctrl = gomock.NewController(t)
	srt.DB = gormmock.NewMockGormw(srt.Ctrl)
	srt.Unit = NewShopWriter(srt.DB)
}

func (srt *shopWriterTest) Finish() {
	srt.Ctrl.Finish()
}

func TestShopWriter_Create(t *testing.T) {
	test := newShopWriterTest()
	t.Run("db return error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetShop()

		test.DB.EXPECT().Save(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(errors.New("error"))

		got, err := test.Unit.Create(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Shop{}, got)
	})
	t.Run("db return no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetShop()

		test.DB.EXPECT().Save(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(errors.New("error"))

		got, err := test.Unit.Create(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Shop{}, got)
	})
}

func TestShopWriter_Update(t *testing.T) {
	test := newShopWriterTest()
	t.Run("db error when finding, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetShop()

		test.DB.EXPECT().Where("id=?", req.ID).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(errors.New("error"))

		got, err := test.Unit.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Shop{}, got)
	})
	t.Run("shop not found, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetShop()

		test.DB.EXPECT().Where("id=?", req.ID).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(nil)
		test.DB.EXPECT().RecordNotFound().Return(true)

		got, err := test.Unit.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Shop{}, got)
	})
	t.Run("db return error when saving, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetShop()

		test.DB.EXPECT().Where("id=?", req.ID).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(nil)
		test.DB.EXPECT().RecordNotFound().Return(false)
		test.DB.EXPECT().Save(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(errors.New("error"))

		got, err := test.Unit.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Shop{}, got)
	})
	t.Run("db return no error , should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetShop()

		test.DB.EXPECT().Where("id=?", req.ID).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(nil)
		test.DB.EXPECT().RecordNotFound().Return(false)
		test.DB.EXPECT().Save(gomock.Any()).DoAndReturn(func(shop *model.Shop) *gormmock.MockGormw {
			*shop = model.Shop{
				ID:        213,
				Name:      shop.Name,
				Address:   shop.Address,
				Longitude: shop.Longitude,
				Latitude:  shop.Latitude,
				Slug:      shop.Slug,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			return test.DB
		})
		test.DB.EXPECT().Error().Return(nil)

		got, err := test.Unit.Update(req)
		assert.Nil(t, err)
		assert.NotEqual(t, entity.Shop{}, got)
	})
}

func (wt *shopWriterTest) GetShop() entity.Shop {
	return entity.Shop{
		ID:      123,
		Name:    "asdasd",
		Address: "asdasd",
		Slug:    "asdasdasd",
		Location: entity.GPS{
			Latitude:  123,
			Longitude: 123,
		},
		Products: nil,
	}
}
