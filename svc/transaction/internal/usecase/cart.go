package usecase

import (
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
)

type Cart struct {
	CartReader internal.CartReader
	CartWriter internal.CartWriter
}

func NewCart(reader internal.CartReader, writer internal.CartWriter) *Cart {
	return &Cart{
		CartReader: reader,
		CartWriter: writer,
	}
}

func (c *Cart) List(userID int64) ([]entity.Cart, error) {
	if userID == 0 {
		return nil, errors.New("user ID is required")
	}
	return c.CartReader.ListByUserID(userID)
}
func (c *Cart) Add(cart entity.Cart) (entity.Cart, error) {
	if cart.AmountKG <= 0 {
		return entity.Cart{}, errors.New("amount cant be <= 0")
	}
	cart.ID = 0
	return c.CartWriter.Create(cart)
}
func (c *Cart) Update(cart entity.Cart) (entity.Cart, error) {
	if cart.ID == 0 {
		return entity.Cart{}, errors.New("cart ID is required")
	}
	if cart.AmountKG <= 0 {
		return entity.Cart{}, errors.New("amount cant be <= 0")
	}
	return c.CartWriter.Update(cart)
}
func (c *Cart) Remove(cartID int64) error {
	if cartID == 0 {
		return errors.New("cart ID is required")
	}
	return c.CartWriter.DeleteByID(cartID)
}
