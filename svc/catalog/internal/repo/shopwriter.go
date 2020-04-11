package repo

import (
	"fmt"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/catalog/util/errors"
)

type ShopWriter struct {
	db connection.Gormw
}

func NewShopWriter(db connection.Gormw) *ShopWriter {
	return &ShopWriter{db: db}
}

func (sw *ShopWriter) Create(req entity.Shop) (entity.Shop, error) {
	const op = "ShopWriter::Create()"
	shopModel := converter.ShopEntityToModel(req)
	query := sw.db.Save(&shopModel)
	if err := query.Error(); err != nil {
		return entity.Shop{}, errors.WithPrefix(err, op)
	}
	shopEntity := converter.ShopModelToEntity(shopModel)
	return shopEntity, nil
}
func (sw *ShopWriter) Update(shop entity.Shop) (entity.Shop, error) {
	const op = "ShopWriter::Update()"
	var shopModel model.Shop
	query := sw.db.Where("id=?", shop.ID).First(&shopModel)
	if err := query.Error(); err != nil {
		return entity.Shop{}, errors.WithPrefix(err, op)
	}
	if query.RecordNotFound() {
		return entity.Shop{}, errors.NewWithPrefix(fmt.Sprintf("record not found with id %d", shop.ID), op)
	}
	shopModel = converter.ShopEntityToModel(shop)
	err := sw.db.Save(&shopModel).Error()
	if err != nil {
		return entity.Shop{}, errors.WithPrefix(err, op)
	}
	return converter.ShopModelToEntity(shopModel), nil
}
func (sw *ShopWriter) DeleteByID(shopID int32) error {
	const op = "ShopWriter::Update()"
	var shopModel model.Shop
	query := sw.db.Where("id=?", shopID).First(&shopModel)
	if err := query.Error(); err != nil {
		return errors.WithPrefix(err, op)
	}
	if query.RecordNotFound() {
		return errors.NewWithPrefix(fmt.Sprintf("record not found with id %d", shopID), op)
	}
	err := sw.db.Delete(&shopModel).Error()
	if err != nil {
		return errors.WithPrefix(err, op)
	}
	return nil
}
func (sw *ShopWriter) DeleteBySlug(shopSlug string) error {
	const op = "ShopWriter::Update()"
	var shopModel model.Shop
	query := sw.db.Where("slug=?", shopSlug).First(&shopModel)
	if err := query.Error(); err != nil {
		return errors.WithPrefix(err, op)
	}
	if query.RecordNotFound() {
		return errors.NewWithPrefix(fmt.Sprintf("record not found with id %s", shopSlug), op)
	}
	err := sw.db.Delete(&shopModel).Error()
	if err != nil {
		return errors.WithPrefix(err, op)
	}
	return nil
}
