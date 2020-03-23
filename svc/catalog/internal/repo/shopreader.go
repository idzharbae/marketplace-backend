package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
)

type ShopReader struct {
	db connection.Gormw
}

func NewShopReader(db connection.Gormw) *ShopReader {
	return &ShopReader{db: db}
}

func (sr *ShopReader) List(req requests.ListShop) ([]entity.Shop, error) {
	shops := make([]model.Shop, req.Pagination.Limit)
	db := sr.db.Limit(req.Pagination.Limit)
	if req.Pagination.Page > 1 {
		db = db.Offset(req.Pagination.OffsetFromPagination())
	}
	err := db.Find(&shops).Error()
	if err != nil {
		return nil, err
	}
	res := converter.ShopModelsToEntities(shops)
	return res, nil
}
func (sr *ShopReader) GetByID(shopID int32) (entity.Shop, error) {
	var shop model.Shop
	err := sr.db.Where("id=?", shopID).First(&shop).Error()
	if err != nil {
		return entity.Shop{}, err
	}
	res := converter.ShopModelToEntity(shop)
	return res, nil
}
func (sr *ShopReader) GetBySlug(slug string) (entity.Shop, error) {
	var shop model.Shop
	err := sr.db.Where("slug=?", slug).First(&shop).Error()
	if err != nil {
		return entity.Shop{}, err
	}
	res := converter.ShopModelToEntity(shop)
	return res, nil
}
