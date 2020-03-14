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
	Unit   *ProductUC
}

func NewProductUCTest() *ProductUCTest {
	return &ProductUCTest{}
}

func (p *ProductUCTest) Begin(t *testing.T) {
	p.Ctrl = gomock.NewController(t)
	p.Reader = repomock.NewMockProductReader(p.Ctrl)
	p.Writer = repomock.NewMockProductWriter(p.Ctrl)
	p.Unit = NewProductUC(p.Reader, p.Writer)
}

func (p *ProductUCTest) Finish() {
	p.Ctrl.Finish()
}

func TestProductUC_List(t *testing.T) {
	test := NewProductUCTest()
	t.Run("reader returns error, should return error", func(t *testing.T) {
		// preparations
		test.Begin(t)
		defer test.Finish()

		// input and output
		req := requests.ListProduct{Pagination: requests.Pagination{
			Page:  1,
			Limit: 10,
		}}
		// expects
		test.Reader.EXPECT().List(req).Return(nil, errors.New("error"))

		// assertions
		got, err := test.Unit.List(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("reader returns no error, should return entity slice", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := requests.ListProduct{Pagination: requests.Pagination{
			Page:  1,
			Limit: 10,
		}}
		res := []entity.Product{
			{ID: 1}, {ID: 2},
		}
		test.Reader.EXPECT().List(req).Return(res, nil)

		got, err := test.Unit.List(req)
		assert.Nil(t, err)
		assert.NotNil(t, got)
	})
}

func TestProductUC_GetByID(t *testing.T) {
	test := NewProductUCTest()
	t.Run("reader returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		id := int32(1)
		test.Reader.EXPECT().GetByID(id).Return(entity.Product{}, errors.New("error"))

		got, err := test.Unit.GetByID(id)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Product{}, got)
	})
	t.Run("reader returns no error, should return product entity with the same id", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		id := int32(3)
		test.Reader.EXPECT().GetByID(id).Return(entity.Product{ID: id}, nil)

		got, err := test.Unit.GetByID(id)
		assert.Nil(t, err)
		assert.Equal(t, id, got.ID)
	})
}

func TestProductUC_GetBySlug(t *testing.T) {
	test := NewProductUCTest()
	t.Run("reader returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		slug := "slug-1"
		test.Reader.EXPECT().GetBySlug(slug).Return(entity.Product{}, errors.New("error"))

		got, err := test.Unit.GetBySlug(slug)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Product{}, got)
	})
	t.Run("reader returns no error, should return product entity with the same slug", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		slug := "slug-3412"
		test.Reader.EXPECT().GetBySlug(slug).Return(entity.Product{ID: 1, Slug: slug}, nil)

		got, err := test.Unit.GetBySlug(slug)
		assert.Nil(t, err)
		assert.Equal(t, slug, got.Slug)
	})
}

func TestProductUC_Create(t *testing.T) {
	test := NewProductUCTest()
	t.Run("given negative quantity, price, or stock should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		reqs := []entity.Product{
			{ID: 1, Quantity: -1},
			{ID: 1, PricePerKG: -1},
			{ID: 1, StockKG: -1},
		}
		for _, req := range reqs {
			got, err := test.Unit.Create(req)

			assert.NotNil(t, err)
			assert.Equal(t, entity.Product{}, got)
		}
	})
	t.Run("repo returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Product{ID: 1}
		test.Writer.EXPECT().Create(req).Return(entity.Product{}, errors.New("error"))

		got, err := test.Unit.Create(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Product{}, got)
	})
	t.Run("repo returns no error, shoul return corresponding entity", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Product{ID: 1}
		test.Writer.EXPECT().Create(req).Return(entity.Product{ID: 3}, nil)

		got, err := test.Unit.Create(req)

		assert.Nil(t, err)
		assert.NotEqual(t, entity.Product{}, got)
	})
}
