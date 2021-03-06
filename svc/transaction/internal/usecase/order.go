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
	if req.UserID != 0 {
		return o.reader.ListByUserID(req.UserID, req.Status, req.Pagination)
	}
	return o.reader.ListByShopID(req.ShopID, req.Status, req.Pagination)
}

func (o *Order) Get(order entity.Order) (entity.Order, error) {
	if order.ID == 0 {
		return entity.Order{}, errors.New("orderID is required")
	}
	if order.ShopID == 0 && order.UserID == 0 {
		return entity.Order{}, errors.New("either shopID or userID is required")
	}
	resultOrder, err := o.reader.GetByID(order.ID)
	if err != nil {
		return entity.Order{}, err
	}
	if order.UserID != 0 && order.UserID != resultOrder.UserID {
		return entity.Order{}, errors.New("user is not authorized to get this order")
	}
	if order.ShopID != 0 && order.ShopID != resultOrder.ShopID {
		return entity.Order{}, errors.New("user is not authorized to get this order")
	}
	return resultOrder, nil
}

func (o *Order) CreateFromCarts(req request.CheckoutReq) ([]entity.Order, error) {
	carts, err := o.cartReader.GetByIDs(request.Pagination{}, req.CartIDs...)
	if err != nil {
		return nil, err
	}
	if len(carts) == 0 {
		return nil, errors.New("carts not found")
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
		product.AmountKG = cartMap[product.ID].AmountKG
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
	if order.UserID == 0 {
		return entity.Order{}, errors.New("userID should not be 0")
	}
	return o.writer.UpdateOrderStatusAndAddShopSaldo(order)
}

func (o *Order) UpdateOrderStatusToOnShipment(orderID, shopID int64) (entity.Order, error) {
	return o.writer.UpdateOrderStatusToOnShipment(orderID, shopID)
}
func (o *Order) RejectOrder(orderID, shopID int64) (entity.Order, error) {
	return o.writer.UpdateOrderStatusToRejected(orderID, shopID)
}
