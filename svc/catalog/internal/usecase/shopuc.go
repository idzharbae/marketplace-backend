package usecase

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
)

type Shop struct {
	internal.ShopReader
	ShopWriter internal.ShopWriter
}

func NewShop(shopReader internal.ShopReader, shopWriter internal.ShopWriter) *Shop {
	return &Shop{
		ShopReader: shopReader,
		ShopWriter: shopWriter,
	}
}

func (s *Shop) Create(shop entity.Shop) (entity.Shop, error) {
	return entity.Shop{}, nil
}
func (s *Shop) Update(shop entity.Shop) (entity.Shop, error) {
	return entity.Shop{}, nil
}
func (s *Shop) Delete(shop int32) error {
	return nil
}
