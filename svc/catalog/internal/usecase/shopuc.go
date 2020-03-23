package usecase

import (
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
)

type Shop struct {
	ShopReader internal.ShopReader
	ShopWriter internal.ShopWriter
}

func NewShop(shopReader internal.ShopReader, shopWriter internal.ShopWriter) *Shop {
	return &Shop{
		ShopReader: shopReader,
		ShopWriter: shopWriter,
	}
}

func (s *Shop) List(req requests.ListShop) ([]entity.Shop, error) {
	if err := req.Pagination.Validate(); err != nil {
		return nil, err
	}
	shops, err := s.ShopReader.List(req)
	if err != nil {
		return nil, err
	}
	return shops, nil
}

func (s *Shop) GetByID(shopID int32) (entity.Shop, error) {
	if shopID <= 0 {
		return entity.Shop{}, errors.New("shop ID should be > 0")
	}
	shop, err := s.ShopReader.GetByID(shopID)
	if err != nil {
		return entity.Shop{}, err
	}
	return shop, nil
}

func (s *Shop) GetBySlug(slug string) (entity.Shop, error) {
	if slug == "" {
		return entity.Shop{}, errors.New("slug should not be empty")
	}
	shop, err := s.ShopReader.GetBySlug(slug)
	if err != nil {
		return entity.Shop{}, err
	}
	return shop, nil
}

func (s *Shop) Create(shop entity.Shop) (entity.Shop, error) {
	err := s.validateShopInput(shop)
	if err != nil {
		return entity.Shop{}, err
	}
	shop = shop.ZeroID().ZeroProductsID()
	res, err := s.ShopWriter.Create(shop)
	if err != nil {
		return entity.Shop{}, err
	}
	return res, nil
}

func (s *Shop) Update(shop entity.Shop) (entity.Shop, error) {
	err := s.validateShopInput(shop)
	if err != nil {
		return entity.Shop{}, err
	}
	err = s.assertIDsNotZero(shop)
	if err != nil {
		return entity.Shop{}, err
	}
	res, err := s.ShopWriter.Update(shop)
	if err != nil {
		return entity.Shop{}, err
	}
	return res, nil
}
func (s *Shop) Delete(shop int32) error {
	return s.ShopWriter.Delete(shop)
}

func (s *Shop) validateShopInput(shop entity.Shop) error {
	if err := shop.Validate(); err != nil {
		return err
	}
	return shop.ValidateProducts()
}

func (s *Shop) assertIDsNotZero(shop entity.Shop) error {
	if err := shop.AssertNotZeroID(); err != nil {
		return err
	}
	return shop.AssertProductsNotZeroID()
}
