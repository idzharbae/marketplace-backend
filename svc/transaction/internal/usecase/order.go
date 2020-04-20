package usecase

import (
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/request"
)

type Order struct {
	reader     internal.OrderReader
	writer     internal.OrderWriter
	cartReader internal.CartReader
}

func NewOrder(reader internal.OrderReader, writer internal.OrderWriter, cartReader internal.CartReader) *Order {
	return &Order{writer: writer, reader: reader, cartReader: cartReader}
}

func (o *Order) List(req request.ListOrderReq) ([]entity.Order, error) {
	return nil, nil
}

func (o *Order) Get(order entity.Order) (entity.Order, error) {
	return entity.Order{}, nil
}

func (o *Order) CreateFromCarts(req request.CheckoutReq) (entity.Order, error) {
	carts, err := o.cartReader.GetByIDs(req.CartIDs...)
	if err != nil {
		return entity.Order{}, err
	}
	productsTotalPrice := int64(0)
	for _, cart := range carts {
		productsTotalPrice += int64(float64(cart.Product.PricePerKG) * cart.Product.AmountKG)
	}
	if req.PaymentAmount != productsTotalPrice {
		return entity.Order{}, errors.New("payment amount doesn't match products total price")
	}

	order, err := o.writer.CreateFromCartsAndSubstractCustomerSaldo(req)
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}
func (o *Order) Fulfill(order entity.Order) (entity.Order, error) {
	if order.ID == 0 {
		return entity.Order{}, errors.New("orderID should not be 0")
	}
	return o.writer.UpdateOrderStatusAndAddShopSaldo(order)
}
