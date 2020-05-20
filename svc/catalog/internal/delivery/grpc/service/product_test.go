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

		req := &catalogproto.ListProductsReq{
			Pagination: &catalogproto.Pagination{
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

		req := &catalogproto.ListProductsReq{
			Pagination: &catalogproto.Pagination{
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

func TestProductService_GetProduct(t *testing.T) {
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

		req := &catalogproto.GetProductReq{
			Id: 1,
		}
		productReq := entity.Product{
			ID:   req.GetId(),
			Slug: req.GetSlug(),
		}
		productUC.EXPECT().Get(productReq).Return(entity.Product{}, errors.New("error"))

		got, err := unit.GetProduct(context.Background(), req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("uc returns no error, should return product proto", func(t *testing.T) {
		begin(t)
		defer finish()

		req := &catalogproto.GetProductReq{
			Id: 2,
		}
		productReq := entity.Product{
			ID:   req.GetId(),
			Slug: req.GetSlug(),
		}
		productUC.EXPECT().Get(productReq).Return(entity.Product{ID: req.GetId()}, nil)

		got, err := unit.GetProduct(context.Background(), req)

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

		req := &catalogproto.GetProductReq{
			Slug: "slug-1",
		}
		productReq := entity.Product{
			ID:   req.GetId(),
			Slug: req.GetSlug(),
		}
		productUC.EXPECT().Get(productReq).Return(entity.Product{}, errors.New("error"))

		got, err := unit.GetProduct(context.Background(), req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("uc returns no error, should return product proto", func(t *testing.T) {
		begin(t)
		defer finish()

		req := &catalogproto.GetProductReq{
			Slug: "slug-2",
		}
		productReq := entity.Product{
			ID:   req.GetId(),
			Slug: req.GetSlug(),
		}
		productUC.EXPECT().Get(productReq).Return(entity.Product{ID: 1, Slug: req.Slug}, nil)

		got, err := unit.GetProduct(context.Background(), req)

		assert.Nil(t, err)
		assert.NotNil(t, got)
	})
}

func TestProductService_CreateUpdateProduct(t *testing.T) {
	var (
		productUC *ucmock.MockProductUC
		ctrl      *gomock.Controller
		unit      *ProductService
		req       *catalogproto.Product
	)
	begin := func(t *testing.T) {
		req = &catalogproto.Product{
			ShopId:     1,
			Name:       "test",
			Quantity:   rand.Int31(),
			PricePerKg: rand.Int31(),
			StockKg:    rand.Float32(),
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
	)
	begin := func(t *testing.T) {
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
		req := &catalogproto.GetProductReq{Id: rand.Int31()}
		productReq := entity.Product{
			ID:   req.GetId(),
			Slug: req.GetSlug(),
		}
		productUC.EXPECT().Delete(productReq).Return(errors.New("error"))
		_, err := unit.DeleteProduct(context.Background(), req)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error should return no error", func(t *testing.T) {
		begin(t)
		defer finish()
		req := &catalogproto.GetProductReq{Id: rand.Int31()}
		productReq := entity.Product{
			ID:   req.GetId(),
			Slug: req.GetSlug(),
		}
		productUC.EXPECT().Delete(productReq).Return(nil)
		_, err := unit.DeleteProduct(context.Background(), req)
		assert.Nil(t, err)
	})
	t.Run("uc returns error should return error", func(t *testing.T) {
		begin(t)
		defer finish()
		req := &catalogproto.GetProductReq{Slug: "asdf"}
		productReq := entity.Product{
			ID:   req.GetId(),
			Slug: req.GetSlug(),
		}
		productUC.EXPECT().Delete(productReq).Return(errors.New("error"))
		_, err := unit.DeleteProduct(context.Background(), req)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error should return no error", func(t *testing.T) {
		begin(t)
		defer finish()
		req := &catalogproto.GetProductReq{Slug: "asdf"}
		productReq := entity.Product{
			ID:   req.GetId(),
			Slug: req.GetSlug(),
		}
		productUC.EXPECT().Delete(productReq).Return(nil)
		_, err := unit.DeleteProduct(context.Background(), req)
		assert.Nil(t, err)
	})
}

func TestProductService_TotalProducts(t *testing.T) {
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

		productUC.EXPECT().GetTotal(int32(123)).Return(int32(0), errors.New("error"))

		got, err := unit.TotalProducts(context.Background(), &catalogproto.TotalProductsReq{ShopId: 123})
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("uc returns no error, should return total products", func(t *testing.T) {
		begin(t)
		defer finish()

		productUC.EXPECT().GetTotal(int32(123)).Return(int32(123), nil)

		got, err := unit.TotalProducts(context.Background(), &catalogproto.TotalProductsReq{ShopId: 123})
		assert.Nil(t, err)
		assert.Equal(t, int32(123), got.GetProductCount())
	})
}
