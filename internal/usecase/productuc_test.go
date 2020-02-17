package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/internal/entity"
	"github.com/idzharbae/marketplace-backend/internal/repo/repomock"
	"github.com/idzharbae/marketplace-backend/internal/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductUC_List(t *testing.T) {
	var (
		repo *repomock.MockProductReader
		ctrl *gomock.Controller
		unit *ProductUC
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		repo = repomock.NewMockProductReader(ctrl)
		unit = NewProductUC(repo)
	}
	finish := func() {
		ctrl.Finish()
	}
	t.Run("repo returns error, should return error", func(t *testing.T) {
		begin(t)
		defer finish()
		req := requests.ListProduct{Pagination: requests.Pagination{
			Page:  1,
			Limit: 10,
		}}
		repo.EXPECT().List(req).Return(nil, errors.New("error"))

		got, err := unit.List(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("repo returns no error, should return entity slice", func(t *testing.T) {
		begin(t)
		defer finish()

		req := requests.ListProduct{Pagination: requests.Pagination{
			Page:  1,
			Limit: 10,
		}}
		res := []entity.Product{
			{ID: 1}, {ID: 2},
		}
		repo.EXPECT().List(req).Return(res, nil)

		got, err := unit.List(req)
		assert.Nil(t, err)
		assert.NotNil(t, got)
	})
}

func TestProductUC_GetByID(t *testing.T) {
	var (
		repo *repomock.MockProductReader
		ctrl *gomock.Controller
		unit *ProductUC
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		repo = repomock.NewMockProductReader(ctrl)
		unit = NewProductUC(repo)
	}
	finish := func() {
		ctrl.Finish()
	}
	t.Run("repo returns error, should return error", func(t *testing.T) {
		begin(t)
		defer finish()
		id := int32(1)
		repo.EXPECT().GetByID(id).Return(entity.Product{}, errors.New("error"))

		got, err := unit.GetByID(id)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Product{}, got)
	})
	t.Run("repo returns no error, should return product entity with the same id", func(t *testing.T) {
		begin(t)
		defer finish()
		id := int32(3)
		repo.EXPECT().GetByID(id).Return(entity.Product{ID: id}, nil)

		got, err := unit.GetByID(id)
		assert.Nil(t, err)
		assert.Equal(t, id, got.ID)
	})
}

func TestProductUC_GetBySlug(t *testing.T) {
	var (
		repo *repomock.MockProductReader
		ctrl *gomock.Controller
		unit *ProductUC
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		repo = repomock.NewMockProductReader(ctrl)
		unit = NewProductUC(repo)
	}
	finish := func() {
		ctrl.Finish()
	}
	t.Run("repo returns error, should return error", func(t *testing.T) {
		begin(t)
		defer finish()
		slug := "slug-1"
		repo.EXPECT().GetBySlug(slug).Return(entity.Product{}, errors.New("error"))

		got, err := unit.GetBySlug(slug)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Product{}, got)
	})
	t.Run("repo returns no error, should return product entity with the same slug", func(t *testing.T) {
		begin(t)
		defer finish()
		slug := "slug-3412"
		repo.EXPECT().GetBySlug(slug).Return(entity.Product{ID: 1, Slug: slug}, nil)

		got, err := unit.GetBySlug(slug)
		assert.Nil(t, err)
		assert.Equal(t, slug, got.Slug)
	})
}
