package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/usecase/ucmock"
	"github.com/stretchr/testify/assert"
	"math/rand"
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

func TestProductService_CreateUpdateProduct(t *testing.T) {
	var (
		productUC *ucmock.MockProductUC
		ctrl      *gomock.Controller
		unit      *ProductService
		req       *marketplaceproto.Product
	)
	begin := func(t *testing.T) {
		req = &marketplaceproto.Product{
			ShopID:     1,
			Name:       "test",
			Quantity:   rand.Int31(),
			PricePerKG: rand.Int31(),
			StockKG:    rand.Float32(),
			Slug:       "slugname",
		}
		ctrl = gomock.NewController(t)
		productUC = ucmock.NewMockProductUC(ctrl)
		unit = NewProductService(productUC)
	}
	finish := func() {
		ctrl.Finish()
	}
	t.Run("[create] uc returns error, should return error", func(t *testing.T) {
		begin(t)
		defer finish()
		productUC.EXPECT().Create(converter.ProductProtoToEntity(req)).Return(entity.Product{}, errors.New("error"))

		got, err := unit.CreateProduct(context.Background(), req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("[create] uc returns no error, should return saved product entity", func(t *testing.T) {
		begin(t)
		defer finish()
		productUC.EXPECT().Create(converter.ProductProtoToEntity(req)).Return(converter.ProductProtoToEntity(req), nil)

		got, err := unit.CreateProduct(context.Background(), req)
		assert.Nil(t, err)
		assert.NotNil(t, got)
		assert.Equal(t, req.Name, got.Name)
	})
	t.Run("[update] uc returns error, should return error", func(t *testing.T) {
		begin(t)
		defer finish()
		productUC.EXPECT().Update(converter.ProductProtoToEntity(req)).Return(entity.Product{}, errors.New("error"))

		got, err := unit.UpdateProduct(context.Background(), req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("[update] uc returns no error, should return saved product entity", func(t *testing.T) {
		begin(t)
		defer finish()
		productUC.EXPECT().Update(converter.ProductProtoToEntity(req)).Return(converter.ProductProtoToEntity(req), nil)

		got, err := unit.UpdateProduct(context.Background(), req)
		assert.Nil(t, err)
		assert.NotNil(t, got)
		assert.Equal(t, req.Name, got.Name)
	})
}

func TestProductService_DeleteProduct(t *testing.T) {
	var (
		productUC *ucmock.MockProductUC
		ctrl      *gomock.Controller
		unit      *ProductService
		req       *catalogproto.ProductPKReq
	)
	begin := func(t *testing.T) {
		req = &catalogproto.ProductPKReq{Id: rand.Int31()}
		ctrl = gomock.NewController(t)
		productUC = ucmock.NewMockProductUC(ctrl)
		unit = NewProductService(productUC)
	}
	finish := func() {
		ctrl.Finish()
	}
	t.Run("uc returns error should return error", func(t *testing.T) {
		begin(t)
		defer finish()
		productUC.EXPECT().Delete(req.GetId()).Return(errors.New("error"))
		_, err := unit.DeleteProduct(context.Background(), req)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error should return no error", func(t *testing.T) {
		begin(t)
		defer finish()
		productUC.EXPECT().Delete(req.GetId()).Return(nil)
		_, err := unit.DeleteProduct(context.Background(), req)
		assert.Nil(t, err)
	})
}
