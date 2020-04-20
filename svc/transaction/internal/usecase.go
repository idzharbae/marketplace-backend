package internal

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/request"
)

//go:generate mockgen -destination=usecase/ucmock/cartuc_mock.go -package=ucmock github.com/idzharbae/marketplace-backend/svc/transaction/internal CartUC
type CartUC interface {
	List(userID int64) ([]entity.Cart, error)
	Add(cart entity.Cart) (entity.Cart, error)
	Update(cart entity.Cart) (entity.Cart, error)
	Remove(cartID int64) error
}

//go:generate mockgen -destination=usecase/ucmock/orderuc_mock.go -package=ucmock github.com/idzharbae/marketplace-backend/svc/transaction/internal OrderUC
type OrderUC interface {
	List(req request.ListOrderReq) ([]entity.Order, error)
	Get(order entity.Order) (entity.Order, error)

	Create(order entity.Order) (entity.Order, error)
	Update(order entity.Order) (entity.Order, error)
	Delete(order entity.Order) error
}
