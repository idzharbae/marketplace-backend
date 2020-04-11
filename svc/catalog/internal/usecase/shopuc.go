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
	shops, err := s.ShopReader.ListAll(req.Pagination)
	if err != nil {
		return nil, err
	}
	return shops, nil
}

func (s *Shop) Get(shop entity.Shop) (entity.Shop, error) {
	if shop.ID != 0 {
		if shop.ID <= 0 {
			return entity.Shop{}, errors.New("shop ID should be > 0")
		}
		shop, err := s.ShopReader.GetByID(shop.ID)
		if err != nil {
			return entity.Shop{}, err
		}
		return shop, nil
	}
	if shop.Slug == "" {
		return entity.Shop{}, errors.New("either shop slug or ID should not be empty")
	}
	shop, err := s.ShopReader.GetBySlug(shop.Slug)
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
	shop = shop.ZeroProductsID()
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
func (s *Shop) Delete(shop entity.Shop) error {
	if shop.ID != 0 {
		return s.ShopWriter.DeleteByID(shop.ID)
	}
	if shop.Slug == "" {
		return errors.New("either id or slug must be provided")
	}
	return s.ShopWriter.DeleteBySlug(shop.Slug)
}

func (s *Shop) validateShopInput(shop entity.Shop) error {
	return shop.Validate()
}

func (s *Shop) assertIDsNotZero(shop entity.Shop) error {
	return shop.AssertNotZeroID()
}
