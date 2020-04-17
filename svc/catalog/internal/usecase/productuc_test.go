package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/repomock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ProductUCTest struct {
	Reader *repomock.MockProductReader
	Writer *repomock.MockProductWriter
	Ctrl   *gomock.Controller
	Unit   *Product
}

func NewProductUCTest() *ProductUCTest {
	return &ProductUCTest{}
}

func (p *ProductUCTest) Begin(t *testing.T) {
	p.Ctrl = gomock.NewController(t)
	p.Reader = repomock.NewMockProductReader(p.Ctrl)
	p.Writer = repomock.NewMockProductWriter(p.Ctrl)
	p.Unit = NewProduct(p.Reader, p.Writer)
}

func (p *ProductUCTest) Finish() {
	p.Ctrl.Finish()
}

func TestProductUC_List(t *testing.T) {
	test := NewProductUCTest()
	t.Run("no shopID, reader returns error, should return error", func(t *testing.T) {
		// preparations
		test.Begin(t)
		defer test.Finish()

		// input and output
		req := requests.ListProduct{Pagination: requests.Pagination{
			Page:  1,
			Limit: 10,
		}}
		// expects
		test.Reader.EXPECT().ListAll(req).Return(nil, errors.New("error"))

		// assertions
		got, err := test.Unit.List(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("no shopID, reader returns no error, should return entity slice", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := requests.ListProduct{Pagination: requests.Pagination{
			Page:  1,
			Limit: 10,
		}}
		res := []entity.Product{
			{ID: 1}, {ID: 2},
		}
		test.Reader.EXPECT().ListAll(req).Return(res, nil)

		got, err := test.Unit.List(req)
		assert.Nil(t, err)
		assert.NotNil(t, got)
	})
	t.Run("given shopID, reader returns error, should return error", func(t *testing.T) {
		// preparations
		test.Begin(t)
		defer test.Finish()

		// input and output
		req := requests.ListProduct{
			ShopID: 1337,
			Pagination: requests.Pagination{
				Page:  1,
				Limit: 10,
			}}
		// expects
		test.Reader.EXPECT().ListByShopID(req.ShopID, req.Pagination).Return(nil, errors.New("error"))

		// assertions
		got, err := test.Unit.List(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("given shopID, reader returns no error, should return entity slice", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := requests.ListProduct{
			ShopID: 1337,
			Pagination: requests.Pagination{
				Page:  1,
				Limit: 10,
			}}
		res := []entity.Product{
			{ID: 1}, {ID: 2},
		}
		test.Reader.EXPECT().ListByShopID(req.ShopID, req.Pagination).Return(res, nil)

		got, err := test.Unit.List(req)
		assert.Nil(t, err)
		assert.NotNil(t, got)
	})
}

func TestProductUC_Get(t *testing.T) {
	test := NewProductUCTest()
	t.Run("given no id and slug, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		product := entity.Product{}

		got, err := test.Unit.Get(product)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Product{}, got)
	})
	// ID
	t.Run("reader returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		product := entity.Product{
			ID: 1,
		}
		test.Reader.EXPECT().GetByID(product.ID).Return(entity.Product{}, errors.New("error"))

		got, err := test.Unit.Get(product)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Product{}, got)
	})
	t.Run("reader returns no error, should return product entity with the same id", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		product := entity.Product{
			ID: 3,
		}
		test.Reader.EXPECT().GetByID(product.ID).Return(entity.Product{ID: product.ID}, nil)

		got, err := test.Unit.Get(product)
		assert.Nil(t, err)
		assert.Equal(t, product.ID, got.ID)
	})
	// Slug
	t.Run("reader returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		product := entity.Product{
			Slug: "slug-5432",
		}
		test.Reader.EXPECT().GetBySlug(product.Slug).Return(entity.Product{}, errors.New("error"))

		got, err := test.Unit.Get(product)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Product{}, got)
	})
	t.Run("reader returns no error, should return product entity with the same slug", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		product := entity.Product{
			Slug: "slug-5432",
		}
		test.Reader.EXPECT().GetBySlug(product.Slug).Return(entity.Product{ID: 1, Slug: product.Slug}, nil)

		got, err := test.Unit.Get(product)
		assert.Nil(t, err)
		assert.Equal(t, product.Slug, got.Slug)
	})
}
