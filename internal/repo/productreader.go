package repo

import (
	"github.com/idzharbae/marketplace-backend/internal/converter"
	"github.com/idzharbae/marketplace-backend/internal/entity"
	"github.com/idzharbae/marketplace-backend/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/internal/requests"
)

type ProductReader struct {
	db connection.Gormw
}

func NewProductReader(db connection.Gormw) *ProductReader {
	return &ProductReader{db: db}
}

func (p *ProductReader) List(req requests.ListProduct) ([]entity.Product, error) {
	var products []model.Product
	p.db.Order("id desc").Find(&products)
	return converter.ProductModelsToEntities(products), nil
}

func (p *ProductReader) GetByID(productID int32) (entity.Product, error) {
	return entity.Product{}, nil
}

func (p *ProductReader) GetBySlug(slug string) (entity.Product, error) {
	return entity.Product{}, nil
}
