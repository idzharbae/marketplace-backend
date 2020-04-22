package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/util/errors"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/model"
)

type OrderReader struct {
	db      connection.Gormw
	catalog internal.CatalogGateway
}

func NewOrderReader(db connection.Gormw, catalog internal.CatalogGateway) *OrderReader {
	return &OrderReader{db: db, catalog: catalog}
}

func (or *OrderReader) ListByUserID(userID int64) ([]entity.Order, error) {
	var orders []model.Order
	err := or.db.Where("user_id=?", userID).Find(&orders).Error()
	if err != nil {
		return nil, errors.WithPrefix(err, "error fetching order")
	}
	resultOrders, err := or.addPaymentAndProductsToOrders(orders)
	if err != nil {
		return nil, err
	}

	return resultOrders, nil
}

func (or *OrderReader) ListByShopID(shopID int64) ([]entity.Order, error) {
	var orders []model.Order
	err := or.db.Where("shop_id=?", shopID).Find(&orders).Error()
	if err != nil {
		return nil, errors.WithPrefix(err, "error fetching order")
	}
	resultOrders, err := or.addPaymentAndProductsToOrders(orders)
	if err != nil {
		return nil, err
	}

	return resultOrders, nil
}
func (or *OrderReader) GetByID(orderID int64) (entity.Order, error) {
	var order model.Order
	err := or.db.Where("id=?", orderID).First(&order).Error()
	if err != nil {
		return entity.Order{}, errors.WithPrefix(err, "error fetching order")
	}

	payment, err := or.getPaymentByOrderID(orderID)
	if err != nil {
		return entity.Order{}, err
	}

	products, err := or.catalog.GetProductsByID(order.ProductID)
	if err != nil {
		return entity.Order{}, errors.WithPrefix(err, "error fetching order products")
	}
	resultOrder := converter.OrderModelToEntity(order, payment)
	resultOrder.Products = products

	return resultOrder, nil
}

func (or *OrderReader) addPaymentAndProductsToOrders(orders []model.Order) ([]entity.Order, error) {
	resultOrders := make([]entity.Order, len(orders))
	for i, order := range orders {
		payment, err := or.getPaymentByOrderID(order.ID)
		if err != nil {
			return nil, errors.WithPrefix(err, "error fetching payment")
		}
		products, err := or.catalog.GetProductsByID(order.ProductID)
		if err != nil {
			return nil, errors.WithPrefix(err, "error fetching products")
		}

		orderEntity := converter.OrderModelToEntity(order, payment)
		orderEntity.Products = products
		resultOrders[i] = orderEntity
	}
	return resultOrders, nil
}

func (or *OrderReader) getPaymentByOrderID(orderID int64) (model.Payment, error) {
	var payment model.Payment
	err := or.db.Where("order_id=?", orderID).First(&payment).Error()
	if err != nil {
		return model.Payment{}, errors.WithPrefix(err, "error fetching payment")
	}
	return payment, nil
}
