package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/repomock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type testShop struct {
	Ctrl   *gomock.Controller
	Reader *repomock.MockShopReader
	Writer *repomock.MockShopWriter
	Unit   internal.ShopUC
}

func newTestShop() *testShop {
	return &testShop{}
}

func (ts *testShop) Begin(t *testing.T) {
	ts.Ctrl = gomock.NewController(t)
	ts.Reader = repomock.NewMockShopReader(ts.Ctrl)
	ts.Writer = repomock.NewMockShopWriter(ts.Ctrl)
	ts.Unit = NewShop(ts.Reader, ts.Writer)
}

func (ts *testShop) Finish() {
	ts.Ctrl.Finish()
}

func (ts *testShop) GetSampleShop() entity.Shop {
	return entity.Shop{
		ID:       13,
		Name:     "test",
		Address:  "asdf",
		Slug:     "asdfg",
		Location: entity.GPS{Latitude: 14.9, Longitude: 13.37},
		Products: []entity.Product{
			{
				ID:         1,
				Name:       "test",
				Slug:       "test",
				StockKG:    123,
				Quantity:   2134,
				PricePerKG: 132132,
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func TestShop_Get(t *testing.T) {
	test := newTestShop()
	t.Run("ID and slug not given, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		shop := entity.Shop{}

		got, err := test.Unit.Get(shop)
		assert.Equal(t, entity.Shop{}, got)
		assert.NotNil(t, err)
	})
	t.Run("given ID <= 0 should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		shop := entity.Shop{
			ID: -1,
		}

		got, err := test.Unit.Get(shop)
		assert.Equal(t, entity.Shop{}, got)
		assert.NotNil(t, err)
	})
	t.Run("given ID, repo returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		shop := entity.Shop{
			ID: 1337,
		}

		test.Reader.EXPECT().GetByID(shop.ID).Return(entity.Shop{}, errors.New("error"))

		got, err := test.Unit.Get(shop)
		assert.Equal(t, entity.Shop{}, got)
		assert.NotNil(t, err)
	})
	t.Run("given ID, repo returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		shop := entity.Shop{
			ID: 1337,
		}

		test.Reader.EXPECT().GetByID(shop.ID).Return(entity.Shop{ID: shop.ID}, nil)

		got, err := test.Unit.Get(shop)
		assert.Equal(t, shop.ID, got.ID)
		assert.Nil(t, err)
	})
	t.Run("given slug repo returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		shop := entity.Shop{
			Slug: "asdf",
		}

		test.Reader.EXPECT().GetBySlug(shop.Slug).Return(entity.Shop{ID: 13, Slug: shop.Slug}, errors.New("error"))

		got, err := test.Unit.Get(shop)
		assert.Equal(t, entity.Shop{}, got)
		assert.NotNil(t, err)
	})
	t.Run("given slug, repo returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		shop := entity.Shop{
			Slug: "asdf",
		}

		test.Reader.EXPECT().GetBySlug(shop.Slug).Return(entity.Shop{ID: 13, Slug: shop.Slug}, nil)

		got, err := test.Unit.Get(shop)
		assert.Equal(t, shop.Slug, got.Slug)
		assert.Nil(t, err)
	})
}

func TestShop_List(t *testing.T) {
	test := newTestShop()
	t.Run("given negative limit should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := requests.ListShop{Pagination: requests.Pagination{
			Page:  3,
			Limit: -1,
		}}
		got, err := test.Unit.List(req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("given negative page should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := requests.ListShop{Pagination: requests.Pagination{
			Page:  -2,
			Limit: 10,
		}}
		got, err := test.Unit.List(req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("repo returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := requests.ListShop{Pagination: requests.Pagination{
			Page:  2,
			Limit: 10,
		}}

		test.Reader.EXPECT().ListAll(req.Pagination).Return(nil, errors.New("error"))

		got, err := test.Unit.List(req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("repo returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := requests.ListShop{Pagination: requests.Pagination{
			Page:  2,
			Limit: 10,
		}}

		test.Reader.EXPECT().ListAll(req.Pagination).Return([]entity.Shop{
			{ID: 1337},
			{ID: 1338},
		}, nil)

		got, err := test.Unit.List(req)
		assert.NotNil(t, got)
		assert.Equal(t, 2, len(got))
		assert.Nil(t, err)
	})
}

func TestShop_Create(t *testing.T) {
	test := newTestShop()
	t.Run("given invalid input should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		var shops []entity.Shop
		validShop := test.GetSampleShop()

		shop := newShop(validShop)
		shop.Name = ""
		shops = append(shops, shop)

		shop = newShop(validShop)
		shop.Slug = ""
		shops = append(shops, shop)

		for _, item := range shops {
			got, err := test.Unit.Create(item)
			assert.Equal(t, entity.Shop{}, got)
			assert.NotNil(t, err)
		}

	})
	t.Run("given shop or products with ID should set ID to 0", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetSampleShop()

		test.Writer.EXPECT().Create(gomock.Any()).DoAndReturn(func(shop entity.Shop) (entity.Shop, error) {
			assert.Equal(t, int32(0), shop.ID)
			for _, item := range shop.Products {
				assert.Equal(t, int32(0), item.ID)
			}
			return entity.Shop{}, nil
		})

		test.Unit.Create(req)
	})
	t.Run("repo returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetSampleShop()

		test.Writer.EXPECT().Create(gomock.Any()).Return(entity.Shop{}, errors.New("error"))

		got, err := test.Unit.Create(req)
		assert.Equal(t, entity.Shop{}, got)
		assert.NotNil(t, err)
	})
	t.Run("repo returns no error, should return no error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetSampleShop()

		test.Writer.EXPECT().Create(gomock.Any()).Return(entity.Shop{ID: 1}, nil)

		got, err := test.Unit.Create(req)
		assert.NotEqual(t, entity.Shop{}, got)
		assert.Nil(t, err)
	})
}

func TestShop_Update(t *testing.T) {
	test := newTestShop()
	t.Run("given invalid input should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		var shops []entity.Shop
		validShop := test.GetSampleShop()

		shop := newShop(validShop)
		shop.Name = ""
		shops = append(shops, shop)

		shop = newShop(validShop)
		shop.Slug = ""
		shops = append(shops, shop)

		shop = newShop(validShop)
		shop.ID = 0
		shops = append(shops, shop)

		for _, item := range shops {
			got, err := test.Unit.Update(item)
			assert.Equal(t, entity.Shop{}, got)
			assert.NotNil(t, err)
		}

	})
	t.Run("IDs should not be 0", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetSampleShop()

		test.Writer.EXPECT().Update(gomock.Any()).DoAndReturn(func(shop entity.Shop) (entity.Shop, error) {
			assert.NotEqual(t, int32(0), shop.ID)
			for _, item := range shop.Products {
				assert.NotEqual(t, int32(0), item.ID)
			}
			return entity.Shop{}, nil
		})

		test.Unit.Update(req)
	})
	t.Run("repo returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetSampleShop()

		test.Writer.EXPECT().Update(gomock.Any()).Return(entity.Shop{}, errors.New("error"))

		got, err := test.Unit.Update(req)
		assert.Equal(t, entity.Shop{}, got)
		assert.NotNil(t, err)
	})
	t.Run("repo returns no error, should return no error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetSampleShop()

		test.Writer.EXPECT().Update(gomock.Any()).Return(entity.Shop{ID: 1}, nil)

		got, err := test.Unit.Update(req)
		assert.NotEqual(t, entity.Shop{}, got)
		assert.Nil(t, err)
	})
}

func TestShop_Delete(t *testing.T) {
	test := newTestShop()
	t.Run("given no slug and ID, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		err := test.Unit.Delete(entity.Shop{})
		assert.NotNil(t, err)
	})
	t.Run("given ID, repo returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Shop{ID: 1337}
		test.Writer.EXPECT().DeleteByID(req.ID).Return(errors.New("error123"))

		err := test.Unit.Delete(req)
		assert.NotNil(t, err)
		assert.Equal(t, "error123", err.Error())
	})
	t.Run("given ID, repo returns no error, should return no error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := entity.Shop{ID: 13213}
		test.Writer.EXPECT().DeleteByID(req.ID).Return(nil)

		err := test.Unit.Delete(req)
		assert.Nil(t, err)
	})
	t.Run("given slug, repo return error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Shop{Slug: "1337"}
		test.Writer.EXPECT().DeleteBySlug(req.Slug).Return(errors.New("error1233"))

		err := test.Unit.Delete(req)
		assert.NotNil(t, err)
		assert.Equal(t, "error1233", err.Error())
	})
	t.Run("given slug, repo returns no error, should return no error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := entity.Shop{Slug: "13213"}
		test.Writer.EXPECT().DeleteBySlug(req.Slug).Return(nil)

		err := test.Unit.Delete(req)
		assert.Nil(t, err)
	})
}

func newShop(shop entity.Shop) entity.Shop {
	return entity.Shop{
		ID:      shop.ID,
		Name:    shop.Name,
		Address: shop.Address,
		Slug:    shop.Slug,
		Location: entity.GPS{
			Latitude:  shop.Location.Latitude,
			Longitude: shop.Location.Longitude,
		},
		CreatedAt: shop.CreatedAt,
		UpdatedAt: shop.UpdatedAt,
	}
}
