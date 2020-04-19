package repo

import (
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/model"
)

type CartWriter struct {
	db connection.Gormw
}

func NewCartWriter(db connection.Gormw) *CartWriter {
	return &CartWriter{db: db}
}

func (cw *CartWriter) Create(cart entity.Cart) (entity.Cart, error) {
	var cartModel model.Cart
	queryFind := cw.db.Where("user_id=?", cart.UserID).Where("product_id=?", cart.Product.ID).First(&cartModel)
	if !queryFind.RecordNotFound() {
		return entity.Cart{}, errors.New("record already exists")
	}
	if err := queryFind.Error(); err != nil {
		return entity.Cart{}, err
	}

	cartModel = converter.CartEntityToModel(cart)
	err := cw.db.Save(&cartModel).Error()
	if err != nil {
		return entity.Cart{}, err
	}

	return converter.CartModelToEntity(cartModel), nil
}
func (cw *CartWriter) Update(cart entity.Cart) (entity.Cart, error) {
	var cartModel model.Cart
	queryFind := cw.db.Where("id=?", cart.ID).First(&cartModel)
	if queryFind.RecordNotFound() {
		return entity.Cart{}, errors.New("record not found")
	}
	if err := queryFind.Error(); err != nil {
		return entity.Cart{}, err
	}

	cartModel.AmountKG = cart.AmountKG
	err := cw.db.Save(&cartModel).Error()
	if err != nil {
		return entity.Cart{}, err
	}

	return converter.CartModelToEntity(cartModel), nil
}

func (cw *CartWriter) DeleteByID(cartID int64) error {
	var cartModel model.Cart
	queryFind := cw.db.Where("id=?", cartID).First(&cartModel)
	if queryFind.RecordNotFound() {
		return errors.New("record not found")
	}
	if err := queryFind.Error(); err != nil {
		return err
	}

	err := cw.db.Delete(&cartModel).Error()

	return err
}
