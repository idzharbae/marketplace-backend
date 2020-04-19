package usecase

import (
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
	return nil, nil
}
func (c *Cart) Add(cart entity.Cart) (entity.Cart, error) {
	return entity.Cart{}, nil
}
func (c *Cart) Update(cart entity.Cart) (entity.Cart, error) {
	return entity.Cart{}, nil
}
func (c *Cart) Remove(cartID int64) error {
	return nil
}
