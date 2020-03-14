package usecase

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/util/errors"
)

type Product struct {
	internal.ProductReader
	ProductWriter internal.ProductWriter
}

func NewProduct(productReader internal.ProductReader, productWriter internal.ProductWriter) *Product {
	return &Product{ProductReader: productReader, ProductWriter: productWriter}
}

func (p *Product) Create(product entity.Product) (entity.Product, error) {
	const op = "Product::Create()"
	if err := product.Validate(); err != nil {
		return entity.Product{}, errors.WithPrefix(err, op)
	}
	product.ID = 0
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
	return p.ProductWriter.Update(product)
}
func (p *Product) Delete(productID int32) error {
	return p.ProductWriter.Delete(productID)
}
