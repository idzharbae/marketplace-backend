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
	err := p.db.Order("id desc").Find(&products).Error()
	if err != nil {
		return nil, err
	}
	return converter.ProductModelsToEntities(products), nil
}

func (p *ProductReader) GetByID(productID int32) (entity.Product, error) {
	var product model.Product
	err := p.db.Where("id=?", productID).First(&product).Error()
	if err != nil {
		return entity.Product{}, err
	}
	return converter.ProductModelToEntity(product), nil
}

func (p *ProductReader) GetBySlug(slug string) (entity.Product, error) {
	var product model.Product
	err := p.db.Where("slug=?", slug).First(&product).Error()
	if err != nil {
		return entity.Product{}, err
	}
	return converter.ProductModelToEntity(product), nil
}
