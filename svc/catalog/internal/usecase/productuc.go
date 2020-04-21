package usecase

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
	"github.com/idzharbae/marketplace-backend/svc/catalog/util/errors"
)

type Product struct {
	ProductReader internal.ProductReader
	ProductWriter internal.ProductWriter
}

func NewProduct(productReader internal.ProductReader, productWriter internal.ProductWriter) *Product {
	return &Product{ProductReader: productReader, ProductWriter: productWriter}
}

func (p *Product) List(req requests.ListProduct) ([]entity.Product, error) {
	if len(req.ProductIDs) > 0 {
		return p.ProductReader.ListByIDs(req.ProductIDs)
	}
	return p.ProductReader.ListAll(req)
}

func (p *Product) Get(product entity.Product) (entity.Product, error) {
	if product.ID != 0 {
		got, err := p.ProductReader.GetByID(product.ID)
		if err != nil {
			return entity.Product{}, err
		}
		return got, nil
	}
	if product.Slug == "" {
		return entity.Product{}, errors.New("either product ID or slug should not be empty")
	}
	got, err := p.ProductReader.GetBySlug(product.Slug)
	if err != nil {
		return entity.Product{}, err
	}
	return got, nil
}

func (p *Product) Create(product entity.Product) (entity.Product, error) {
	const op = "Product::Create()"
	if err := product.Validate(); err != nil {
		return entity.Product{}, errors.WithPrefix(err, op)
	}
	product.ID = 0
	if product.ShopID == 0 {
		return entity.Product{}, errors.NewWithPrefix("shop id should not be 0", op)
	}
	return p.ProductWriter.Create(product)
}
func (p *Product) Update(product entity.Product) (entity.Product, error) {
	const op = "Product::Update()"
	if err := product.Validate(); err != nil {
		return entity.Product{}, errors.WithPrefix(err, op)
	}
	if product.ID == 0 {
		return entity.Product{}, errors.NewWithPrefix("product ID can't be 0", op)
	}
	if product.ShopID == 0 {
		return entity.Product{}, errors.NewWithPrefix("shop id should not be 0", op)
	}
	return p.ProductWriter.Update(product)
}
func (p *Product) Delete(product entity.Product) error {
	if product.ID != 0 {
		return p.ProductWriter.DeleteByID(product.ID)
	}
	return p.ProductWriter.DeleteBySlug(product.Slug)
}
