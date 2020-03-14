package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection"
)

type ShopWriter struct {
	db connection.Gormw
}

func NewShopWriter(db connection.Gormw) *ShopWriter {
	return &ShopWriter{db: db}
}

func (sw *ShopWriter) Create(shop entity.Shop) (entity.Shop, error) {
	return entity.Shop{}, nil
}
func (sw *ShopWriter) Update(shop entity.Shop) (entity.Shop, error) {
	return entity.Shop{}, nil
}
func (sw *ShopWriter) Delete(shop int32) error {
	return nil
}
