package internal

import "github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"

type CartReader interface {
	ListByUserID(userID int64) ([]entity.Cart, error)
}

type CartWriter interface {
	Create(cart entity.Cart) (entity.Cart, error)
	Update(cart entity.Cart) (entity.Cart, error)
	DeleteByID(cartID int64) error
}
