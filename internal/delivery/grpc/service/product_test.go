package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/internal/entity"
	"github.com/idzharbae/marketplace-backend/internal/requests"
	"github.com/idzharbae/marketplace-backend/internal/usecase/ucmock"
	"github.com/idzharbae/marketplace-backend/marketplaceproto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductService_ListProducts(t *testing.T) {
	var (
		productUC *ucmock.MockProductUC
		ctrl      *gomock.Controller
		unit      *ProductService
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		productUC = ucmock.NewMockProductUC(ctrl)
		unit = NewProductService(productUC)
	}
	finish := func() {
		ctrl.Finish()
	}
	t.Run("uc returns error, should return error", func(t *testing.T) {
		begin(t)
		defer finish()

		req := &marketplaceproto.ListProductsReq{
			Pagination: &marketplaceproto.Pagination{
				Page:  1,
				Limit: 10,
			},
		}
		ucReq := requests.ListProduct{Pagination: requests.Pagination{Limit: 10, Page: 1}}
		productUC.EXPECT().List(ucReq).Return(nil, errors.New("error"))
		got, err := unit.ListProducts(context.Background(), req)

		assert.NotNil(t, err)
		assert.Nil(t, got)
	})

	t.Run("uc returns no error, should return product protos", func(t *testing.T) {
		begin(t)
		defer finish()

		req := &marketplaceproto.ListProductsReq{
			Pagination: &marketplaceproto.Pagination{
				Page:  1,
				Limit: 10,
			},
		}
		res := []entity.Product{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
		}
		ucReq := requests.ListProduct{Pagination: requests.Pagination{Limit: 10, Page: 1}}
		productUC.EXPECT().List(ucReq).Return(res, nil)
		got, err := unit.ListProducts(context.Background(), req)

		assert.Nil(t, err)
		assert.NotNil(t, got)
	})
}

func TestProductService_GetProductByID(t *testing.T) {
	var (
		productUC *ucmock.MockProductUC
		ctrl      *gomock.Controller
		unit      *ProductService
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		productUC = ucmock.NewMockProductUC(ctrl)
		unit = NewProductService(productUC)
	}
	finish := func() {
		ctrl.Finish()
	}
	t.Run("uc returns error, should return error", func(t *testing.T) {
		begin(t)
		defer finish()

		req := &marketplaceproto.GetProductByIDReq{
			ID: 1,
		}
		productUC.EXPECT().GetByID(req.ID).Return(entity.Product{}, errors.New("error"))

		got, err := unit.GetProductByID(context.Background(), req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("uc returns no error, should return product proto", func(t *testing.T) {
		begin(t)
		defer finish()

		req := &marketplaceproto.GetProductByIDReq{
			ID: 2,
		}
		productUC.EXPECT().GetByID(req.ID).Return(entity.Product{ID: req.ID}, nil)

		got, err := unit.GetProductByID(context.Background(), req)

		assert.Nil(t, err)
		assert.NotNil(t, got)
	})
}

func TestProductService_GetProductBySlug(t *testing.T) {
	var (
		productUC *ucmock.MockProductUC
		ctrl      *gomock.Controller
		unit      *ProductService
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		productUC = ucmock.NewMockProductUC(ctrl)
		unit = NewProductService(productUC)
	}
	finish := func() {
		ctrl.Finish()
	}
	t.Run("uc returns error, should return error", func(t *testing.T) {
		begin(t)
		defer finish()

		req := &marketplaceproto.GetProductBySlugReq{
			Slug: "slug-1",
		}
		productUC.EXPECT().GetBySlug(req.Slug).Return(entity.Product{}, errors.New("error"))

		got, err := unit.GetProductBySlug(context.Background(), req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("uc returns no error, should return product proto", func(t *testing.T) {
		begin(t)
		defer finish()

		req := &marketplaceproto.GetProductBySlugReq{
			Slug: "slug-2",
		}
		productUC.EXPECT().GetBySlug(req.Slug).Return(entity.Product{ID: 1, Slug: req.Slug}, nil)

		got, err := unit.GetProductBySlug(context.Background(), req)

		assert.Nil(t, err)
		assert.NotNil(t, got)
	})
}
