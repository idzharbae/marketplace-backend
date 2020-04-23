package usecase

import (
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
)

type Cart struct {
	CartReader     internal.CartReader
	CartWriter     internal.CartWriter
	CatalogGateway internal.CatalogGateway
}

func NewCart(reader internal.CartReader, writer internal.CartWriter, gateway internal.CatalogGateway) *Cart {
	return &Cart{
		CartReader:     reader,
		CartWriter:     writer,
		CatalogGateway: gateway,
	}
}

func (c *Cart) List(userID int64) ([]entity.Cart, error) {
	if userID == 0 {
		return nil, errors.New("user ID is required")
	}
	carts, err := c.CartReader.ListByUserID(userID)
	if err != nil {
		return nil, err
	}

	for i, cart := range carts {
		res, err := c.CatalogGateway.GetProductByID(cart.Product.ID)
		if err != nil {
			return nil, err
		}
		carts[i].Product = res
		carts[i].Product.AmountKG = carts[i].AmountKG
		carts[i].Product.TotalPrice = int64(carts[i].AmountKG * float64(carts[i].Product.PricePerKG))
	}

	return carts, nil
}
func (c *Cart) Add(cart entity.Cart) (entity.Cart, error) {
	if cart.AmountKG <= 0 {
		return entity.Cart{}, errors.New("amount cant be <= 0")
	}
	_, err := c.CatalogGateway.GetProductByID(cart.Product.ID)
	if err != nil {
		return entity.Cart{}, errors.New("error fetching product: " + err.Error())
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
	if cart.UserID == 0 {
		return entity.Cart{}, errors.New("userID is required")
	}

	res, err := c.CartWriter.Update(cart)
	if err != nil {
		return entity.Cart{}, err
	}
	res.Product.AmountKG = res.AmountKG
	return res, nil
}
func (c *Cart) Remove(cartID, userID int64) error {
	if cartID == 0 {
		return errors.New("cart ID is required")
	}
	if userID == 0 {
		return errors.New("userID is required")
	}
	return c.CartWriter.DeleteByID(cartID, userID)
}
