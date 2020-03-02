package usecase

import (
	"github.com/idzharbae/marketplace-backend/internal"
	"github.com/idzharbae/marketplace-backend/internal/constant"
	"github.com/idzharbae/marketplace-backend/internal/entity"
	"github.com/idzharbae/marketplace-backend/internal/requests"
	"github.com/idzharbae/marketplace-backend/util/errors"
)

type ProductUC struct {
	ProductReader internal.ProductReader
	ProductWriter internal.ProductWriter
}

func NewProductUC(productReader internal.ProductReader, productWriter internal.ProductWriter) *ProductUC {
	return &ProductUC{ProductReader: productReader, ProductWriter: productWriter}
}

func (p *ProductUC) List(req requests.ListProduct) ([]entity.Product, error) {
	return p.ProductReader.List(req)
}

func (p *ProductUC) GetByID(productID int32) (entity.Product, error) {
	return p.ProductReader.GetByID(productID)
}
func (p *ProductUC) GetBySlug(slug string) (entity.Product, error) {
	return p.ProductReader.GetBySlug(slug)
}

func (p *ProductUC) Create(product entity.Product) (entity.Product, error) {
	const op = "ProductUC::Create()"
	if !product.Validate() {
		return entity.Product{}, errors.NewWithPrefix(constant.ErrInvalidInput, op)
	}
	return p.ProductWriter.Create(product)
}
func (p *ProductUC) Update(product entity.Product) (entity.Product, error) {
	const op = "ProductUC::Update()"
	if !product.Validate() {
		return entity.Product{}, errors.NewWithPrefix(constant.ErrInvalidInput, op)
	}
	return p.ProductWriter.Update(product)
}
func (p *ProductUC) Delete(productID int32) error {
	return p.ProductWriter.Delete(productID)
}
