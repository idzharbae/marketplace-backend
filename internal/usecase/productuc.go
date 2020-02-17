package usecase

import (
	"github.com/idzharbae/marketplace-backend/internal"
	"github.com/idzharbae/marketplace-backend/internal/entity"
	"github.com/idzharbae/marketplace-backend/internal/requests"
)

type ProductUC struct {
	ProductReader internal.ProductReader
}

func NewProductUC(productReader internal.ProductReader) *ProductUC {
	return &ProductUC{ProductReader: productReader}
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
	return entity.Product{}, nil
}
func (p *ProductUC) Update(product entity.Product) (entity.Product, error) {
	return entity.Product{}, nil
}
func (p *ProductUC) Delete(productID int32) error {
	return nil
}
