package usecase

import (
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/request"
)

type Order struct {
	reader         internal.OrderReader
	writer         internal.OrderWriter
	cartReader     internal.CartReader
	catalogGateway internal.CatalogGateway
}

func NewOrder(reader internal.OrderReader, writer internal.OrderWriter, cartReader internal.CartReader, gateway internal.CatalogGateway) *Order {
	return &Order{writer: writer, reader: reader, cartReader: cartReader, catalogGateway: gateway}
}

func (o *Order) List(req request.ListOrderReq) ([]entity.Order, error) {
	return nil, nil
}

func (o *Order) Get(order entity.Order) (entity.Order, error) {
	return entity.Order{}, nil
}

func (o *Order) CreateFromCarts(req request.CheckoutReq) ([]entity.Order, error) {
	carts, err := o.cartReader.GetByIDs(req.CartIDs...)
	if err != nil {
		return nil, err
	}
	productsTotalPrice := int64(0)
	cartMap := make(map[int64]*entity.Cart, len(carts))
	productIDs := make([]int64, len(carts))
	for i, cart := range carts {
		if cart.UserID != req.UserID {
			return nil, errors.New("one of the cart item is not owned by user")
		}
		productIDs[i] = cart.Product.ID
		cartMap[cart.Product.ID] = &carts[i]
	}

	products, err := o.catalogGateway.GetProductsByID(productIDs)
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		productsTotalPrice += int64(cartMap[product.ID].AmountKG * float64(product.PricePerKG))
		cartMap[product.ID].Product = product
	}

	if req.PaymentAmount != productsTotalPrice {
		return nil, errors.New("payment amount doesn't match products total price")
	}

	order, err := o.writer.CreateFromCartsAndSubstractCustomerSaldo(request.CreateOrderReq{
		UserID:        req.UserID,
		PaymentAmount: req.PaymentAmount,
		Carts:         carts,
	})
	if err != nil {
		return nil, err
	}
	return order, nil
}
func (o *Order) Fulfill(order entity.Order) (entity.Order, error) {
	if order.ID == 0 {
		return entity.Order{}, errors.New("orderID should not be 0")
	}
	return o.writer.UpdateOrderStatusAndAddShopSaldo(order)
}
