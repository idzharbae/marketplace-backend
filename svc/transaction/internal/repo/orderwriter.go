package repo

import (
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/transaction/constants"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/request"
)

type OrderWriter struct {
	db      connection.Gormw
	auth    internal.AuthGateway
	catalog internal.CatalogGateway
}

func NewOrderWriter(db connection.Gormw, auth internal.AuthGateway, catalog internal.CatalogGateway) *OrderWriter {
	return &OrderWriter{db: db, auth: auth, catalog: catalog}
}

func (ow *OrderWriter) CreateFromCartsAndSubstractCustomerSaldo(req request.CreateOrderReq) ([]entity.Order, error) {
	err := ow.validateSaldo(req)
	if err != nil {
		return nil, err
	}

	dbTransaction := ow.db.Begin()

	ordersMap := ow.groupOrderByShopID(req)
	resultOrders, err := ow.createOrders(ordersMap, dbTransaction)
	if err != nil {
		return nil, err
	}

	err = ow.deleteCarts(req, dbTransaction)
	if err != nil {
		return nil, err
	}
	_, err = ow.auth.UpdateUserSaldo(req.UserID, -req.PaymentAmount)
	if err != nil {
		dbTransaction.Rollback()
		return nil, err
	}

	dbTransaction.Commit()
	return resultOrders, nil
}

func (ow *OrderWriter) UpdateOrderStatusAndAddShopSaldo(order entity.Order) (entity.Order, error) {
	var orderModel model.Order
	err := ow.db.Where("id=?", order.ID).First(&orderModel).Error()
	if err != nil {
		return entity.Order{}, err
	}
	if orderModel.Status != constants.OrderStatusOnShipment {
		return entity.Order{}, errors.New("order is not being shipped")
	}
	if orderModel.UserID != order.UserID {
		return entity.Order{}, errors.New("user is not authorized to fulfill this order")
	}
	order = converter.OrderModelToEntity(orderModel, model.Payment{})
	dbTransaction := ow.db.Begin()
	err = dbTransaction.Model(&model.Order{}).Where("id=?", order.ID).Update("status", constants.OrderStatusFulfilled).Error()
	if err != nil {
		return entity.Order{}, err
	}
	order.Status = constants.OrderStatusFulfilled
	_, err = ow.auth.UpdateUserSaldo(order.ShopID, order.TotalPrice)
	if err != nil {
		dbTransaction.Rollback()
		return entity.Order{}, err
	}
	dbTransaction.Commit()
	return order, nil
}

func (ow *OrderWriter) UpdateOrderStatusToOnShipment(orderID, shopID int64) (entity.Order, error) {
	return entity.Order{}, nil
}

// CreateFromCartsAndSubstractCustomerSaldo private functions
func (ow *OrderWriter) validateSaldo(req request.CreateOrderReq) error {
	user, err := ow.auth.GetUserByID(req.UserID)
	if err != nil {
		return err
	}
	if user.Saldo < req.PaymentAmount {
		return errors.New("not enough saldo")
	}
	return nil
}

func (ow *OrderWriter) groupOrderByShopID(req request.CreateOrderReq) map[int64]model.Order {
	ordersMap := make(map[int64]model.Order)
	for _, cart := range req.Carts {
		shopID := cart.Product.ShopID
		productPrice := int64(float64(cart.Product.PricePerKG) * cart.AmountKG)
		ordersMap[cart.Product.ShopID] = model.Order{
			ProductID:  append(ordersMap[shopID].ProductID, cart.Product.ID),
			UserID:     cart.UserID,
			ShopID:     cart.Product.ShopID,
			TotalPrice: ordersMap[shopID].TotalPrice + productPrice,
			Status:     constants.OrderStatusWaitingForSeller,
		}
	}
	return ordersMap
}

func (ow *OrderWriter) createOrders(ordersMap map[int64]model.Order, dbTransaction connection.Gormw) ([]entity.Order, error) {
	resultOrders := make([]entity.Order, 0, len(ordersMap))
	for _, order := range ordersMap {
		err := dbTransaction.Save(&order).Error()
		if err != nil {
			dbTransaction.Rollback()
			return nil, err
		}
		payment := model.Payment{
			OrderID:       order.ID,
			Amount:        order.TotalPrice,
			PaymentMethod: constants.PaymentTypeSaldo,
			PaymentStatus: constants.PaymentStatusPaid,
		}
		err = dbTransaction.Save(&payment).Error()
		if err != nil {
			dbTransaction.Rollback()
			return nil, err
		}
		resultOrders = append(resultOrders, converter.OrderModelToEntity(order, payment))
	}
	return resultOrders, nil
}

func (ow *OrderWriter) deleteCarts(req request.CreateOrderReq, dbTransaction connection.Gormw) error {
	carts := converter.CartEntitiesToModels(req.Carts)
	err := dbTransaction.Delete(&carts).Error()
	if err != nil {
		dbTransaction.Rollback()
		return err
	}
	return nil
}
