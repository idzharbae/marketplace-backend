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

func TestProductReader_List(t *testing.T) {
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
		req := requests.ListProduct{Pagination: requests.Pagination{Limit: 10, Page: 1}}

		db.EXPECT().Order("id desc").Return(db)
		db.EXPECT().Find(gomock.Any()).Return(db)
		db.EXPECT().Error().Return(errors.New("error"))

		got, err := unit.List(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("db returns no error, should return entity converted from model", func(t *testing.T) {
		begin(t)
		defer finish()
		req := requests.ListProduct{Pagination: requests.Pagination{Limit: 10, Page: 1}}

		db.EXPECT().Order("id desc").Return(db)
		db.EXPECT().Find(gomock.Any()).DoAndReturn(func(out *[]model.Product, where ...interface{}) *gormmock.MockGormw {
			*out = []model.Product{
				{ID: 10},
				{ID: 13},
			}
			return db
		})
		db.EXPECT().Error().Return(nil)

		got, err := unit.List(req)
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
