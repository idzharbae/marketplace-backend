package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/repomock"
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
