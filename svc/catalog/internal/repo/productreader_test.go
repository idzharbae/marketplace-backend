package repo

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection/gormmock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductReader_ListAll(t *testing.T) {
	var (
		db   *gormmock.MockGormw
		ctrl *gomock.Controller
		unit *ProductReader
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		db = gormmock.NewMockGormw(ctrl)
		unit = NewProductReader(db)
	}
	finish := func() {
		ctrl.Finish()
	}
	// filter flow:
	// 1. filter category
	// 2. filter search
	// 3. filter province (use list of shop_id that are located in the province)
	// 4. order by
	// 5. pagination
	t.Run("all filter empty, should only execute find and order by id desc", func(t *testing.T) {
		begin(t)
		defer finish()
		req := requests.ListProduct{}

		db.EXPECT().Order("id desc").Return(db)
		db.EXPECT().Find(gomock.Any()).Return(db)
		db.EXPECT().Error().Return(nil)

		unit.ListAll(req)
	})
	t.Run("given only category, should filter category and order by id desc", func(t *testing.T) {
		begin(t)
		defer finish()
		req := requests.ListProduct{Category: "test"}

		db.EXPECT().Where("category=?", req.Category).Return(db)
		db.EXPECT().Order("id desc").Return(db)
		db.EXPECT().Find(gomock.Any()).Return(db)
		db.EXPECT().Error().Return(nil)

		unit.ListAll(req)
	})
	t.Run("given only search, should filter search and order by id desc", func(t *testing.T) {
		begin(t)
		defer finish()
		req := requests.ListProduct{Search: "asdfg"}

		db.EXPECT().Where("name ilike ?", "%"+req.Search+"%").Return(db)
		db.EXPECT().Order("id desc").Return(db)
		db.EXPECT().Find(gomock.Any()).Return(db)
		db.EXPECT().Error().Return(nil)

		unit.ListAll(req)
	})
	t.Run("given only province, should filter province and order by id desc", func(t *testing.T) {
		begin(t)
		defer finish()
		req := requests.ListProduct{ShopIDs: []int64{1, 2, 3}}

		db.EXPECT().Where("shop_id=ANY(?)", req.ShopIDs).Return(db)
		db.EXPECT().Order("id desc").Return(db)
		db.EXPECT().Find(gomock.Any()).Return(db)
		db.EXPECT().Error().Return(nil)

		unit.ListAll(req)
	})
	t.Run("given order param should order by the given param", func(t *testing.T) {
		begin(t)
		defer finish()
		req := requests.ListProduct{OrderBy: "created_at", OrderType: "asc"}

		db.EXPECT().Order(req.OrderBy + " " + req.OrderType).Return(db)
		db.EXPECT().Find(gomock.Any()).Return(db)
		db.EXPECT().Error().Return(nil)

		unit.ListAll(req)
	})
	t.Run("given invalid order param should order by id desc", func(t *testing.T) {
		begin(t)
		defer finish()
		req := requests.ListProduct{OrderBy: "(case if 1=1 then created_at else updated_at end)", OrderType: "desc"}

		db.EXPECT().Order("id desc").Return(db)
		db.EXPECT().Find(gomock.Any()).Return(db)
		db.EXPECT().Error().Return(nil)

		unit.ListAll(req)
	})
	t.Run("given limit, should apply limit filter filter", func(t *testing.T) {
		begin(t)
		defer finish()
		req := requests.ListProduct{Pagination: requests.Pagination{Limit: 10, Page: 0}}

		db.EXPECT().Order("id desc").Return(db)
		db.EXPECT().Limit(req.Pagination.Limit).Return(db)
		db.EXPECT().Find(gomock.Any()).Return(db)
		db.EXPECT().Error().Return(nil)

		unit.ListAll(req)
	})
	t.Run("given page, should apply offset filter", func(t *testing.T) {
		begin(t)
		defer finish()
		req := requests.ListProduct{Pagination: requests.Pagination{Limit: 0, Page: 10}}

		db.EXPECT().Order("id desc").Return(db)
		db.EXPECT().Offset(req.Pagination.OffsetFromPagination()).Return(db)
		db.EXPECT().Find(gomock.Any()).Return(db)
		db.EXPECT().Error().Return(nil)

		unit.ListAll(req)
	})
	t.Run("db returns error, should return error", func(t *testing.T) {
		begin(t)
		defer finish()
		req := requests.ListProduct{
			Category:   "asdf",
			Search:     "asldkls",
			ShopIDs:    []int64{1, 2, 3},
			OrderBy:    "stock_kg",
			OrderType:  "asc",
			Pagination: requests.Pagination{Limit: 10, Page: 2},
		}

		db.EXPECT().Where("category=?", req.Category).Return(db)
		db.EXPECT().Where("name ilike ?", "%"+req.Search+"%").Return(db)
		db.EXPECT().Where("shop_id=ANY(?)", req.ShopIDs).Return(db)
		db.EXPECT().Order("stock_kg asc").Return(db)
		db.EXPECT().Limit(req.Pagination.Limit).Return(db)
		db.EXPECT().Offset(req.Pagination.OffsetFromPagination()).Return(db)
		db.EXPECT().Find(gomock.Any()).DoAndReturn(func(out *[]model.Product, where ...interface{}) *gormmock.MockGormw {
			*out = []model.Product{
				{ID: 10},
				{ID: 13},
			}
			return db
		})
		db.EXPECT().Error().Return(errors.New("error"))

		got, err := unit.ListAll(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("db returns no error, should return entity converted from model", func(t *testing.T) {
		begin(t)
		defer finish()
		req := requests.ListProduct{
			Category:   "asdf",
			Search:     "asldkls",
			ShopIDs:    []int64{1, 2, 3},
			OrderBy:    "stock_kg",
			OrderType:  "asc",
			Pagination: requests.Pagination{Limit: 10, Page: 2},
		}

		db.EXPECT().Where("category=?", req.Category).Return(db)
		db.EXPECT().Where("name ilike ?", "%"+req.Search+"%").Return(db)
		db.EXPECT().Where("shop_id=ANY(?)", req.ShopIDs).Return(db)
		db.EXPECT().Order("stock_kg asc").Return(db)
		db.EXPECT().Limit(req.Pagination.Limit).Return(db)
		db.EXPECT().Offset(req.Pagination.OffsetFromPagination()).Return(db)
		db.EXPECT().Find(gomock.Any()).DoAndReturn(func(out *[]model.Product, where ...interface{}) *gormmock.MockGormw {
			*out = []model.Product{
				{ID: 10},
				{ID: 13},
			}
			return db
		})
		db.EXPECT().Error().Return(nil)

		got, err := unit.ListAll(req)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(got))
		assert.Equal(t, int32(10), got[0].ID)
		assert.Equal(t, int32(13), got[1].ID)
	})
}
func TestProductReader_GetByID(t *testing.T) {
	var (
		db   *gormmock.MockGormw
		ctrl *gomock.Controller
		unit *ProductReader
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		db = gormmock.NewMockGormw(ctrl)
		unit = NewProductReader(db)
	}
	finish := func() {
		ctrl.Finish()
	}
	t.Run("db returns error, should return error", func(t *testing.T) {
		begin(t)
		defer finish()
		id := int32(4)

		db.EXPECT().Where("id=?", id).Return(db)
		db.EXPECT().First(gomock.Any()).Return(db)
		db.EXPECT().Error().Return(errors.New("error"))

		got, err := unit.GetByID(id)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Product{}, got)
	})
	t.Run("db returns no error, should return product entity", func(t *testing.T) {
		begin(t)
		defer finish()
		id := int32(4)

		db.EXPECT().Where("id=?", id).Return(db)
		db.EXPECT().First(gomock.Any()).DoAndReturn(func(out *model.Product, where ...interface{}) *gormmock.MockGormw {
			*out = model.Product{ID: 4}
			return db
		})
		db.EXPECT().Error().Return(nil)

		got, err := unit.GetByID(id)
		assert.Nil(t, err)
		assert.Equal(t, id, got.ID)
	})
}

func TestProductReader_GetBySlug(t *testing.T) {
	var (
		db   *gormmock.MockGormw
		ctrl *gomock.Controller
		unit *ProductReader
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		db = gormmock.NewMockGormw(ctrl)
		unit = NewProductReader(db)
	}
	finish := func() {
		ctrl.Finish()
	}
	t.Run("db returns error, should return error", func(t *testing.T) {
		begin(t)
		defer finish()
		slug := "slug-1"

		db.EXPECT().Where("slug=?", slug).Return(db)
		db.EXPECT().First(gomock.Any()).Return(db)
		db.EXPECT().Error().Return(errors.New("error"))

		got, err := unit.GetBySlug(slug)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Product{}, got)
	})
	t.Run("db returns no error, should return product entity", func(t *testing.T) {
		begin(t)
		defer finish()
		slug := "slug-13"

		db.EXPECT().Where("slug=?", slug).Return(db)
		db.EXPECT().First(gomock.Any()).DoAndReturn(func(out *model.Product, where ...interface{}) *gormmock.MockGormw {
			*out = model.Product{ID: 1, Slug: slug}
			return db
		})
		db.EXPECT().Error().Return(nil)

		got, err := unit.GetBySlug(slug)
		assert.Nil(t, err)
		assert.Equal(t, slug, got.Slug)
	})
}

func TestProductReader_GetByShopID(t *testing.T) {
	var (
		db   *gormmock.MockGormw
		ctrl *gomock.Controller
		unit *ProductReader
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		db = gormmock.NewMockGormw(ctrl)
		unit = NewProductReader(db)
	}
	finish := func() {
		ctrl.Finish()
	}
	t.Run("db returns error, should return error", func(t *testing.T) {
		begin(t)
		defer finish()
		shopID := int32(1337)

		db.EXPECT().Where("shop_id=?", shopID).Return(db)
		db.EXPECT().First(gomock.Any()).Return(db)
		db.EXPECT().Error().Return(errors.New("error"))

		got, err := unit.GetByShopID(shopID)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("db returns no error, should return product entity", func(t *testing.T) {
		begin(t)
		defer finish()
		shopID := int32(1337)

		db.EXPECT().Where("shop_id=?", shopID).Return(db)
		db.EXPECT().First(gomock.Any()).Return(db)
		db.EXPECT().Error().Return(nil)

		got, err := unit.GetByShopID(shopID)
		assert.Nil(t, err)
		assert.NotNil(t, got)
	})
}
