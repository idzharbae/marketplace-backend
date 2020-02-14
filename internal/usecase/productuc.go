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

func (p *ProductUC) ListProducts(req requests.ListProduct) ([]entity.Product, error) {
	return p.ProductReader.ListProducts(req)
}
