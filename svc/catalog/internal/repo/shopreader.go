package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
)

type ShopReader struct {
	db connection.Gormw
}

func NewShopReader(db connection.Gormw) *ShopReader {
	return &ShopReader{db: db}
}

func (sr *ShopReader) List(req requests.ListShop) ([]entity.Shop, error) {
	return nil, nil
}
func (sr *ShopReader) GetByID(shopID int32) (entity.Shop, error) {
	return entity.Shop{}, nil
}
func (sr *ShopReader) GetBySlug(slug string) (entity.Shop, error) {
	return entity.Shop{}, nil
}
