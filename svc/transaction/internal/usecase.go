package internal

import "github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"

type CartUC interface {
	List(userID int64) ([]entity.Cart, error)
	Add(cart entity.Cart) (entity.Cart, error)
	Update(cart entity.Cart) (entity.Cart, error)
	Remove(cartID int64) error
}
