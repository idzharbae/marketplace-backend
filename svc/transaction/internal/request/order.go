package request

import "github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"

type ListOrderReq struct {
	UserID int64
	ShopID int64
	Status int32
}

type CheckoutReq struct {
	UserID        int64
	CartIDs       []int64
	PaymentAmount int64
	ProductIDs    []int64
}

type CreateOrderReq struct {
	UserID        int64
	Carts         []entity.Cart
	PaymentAmount int64
}
