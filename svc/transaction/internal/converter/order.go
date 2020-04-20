package converter

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
	"time"
)

func OrderEntitiesToProtos(orders []entity.Order) []*prototransaction.Order {
	res := make([]*prototransaction.Order, len(orders))
	for i, order := range orders {
		res[i] = OrderEntityToProto(order)
	}
	return res
}

func OrderEntityToProto(order entity.Order) *prototransaction.Order {
	return &prototransaction.Order{
		Id:         order.ID,
		UserId:     order.UserID,
		TotalPrice: order.TotalPrice,
		Products:   ProductEntitiesToProtos(order.Products),
		Status:     order.Status,
		Payment:    nil,
	}
}

func PaymentEntityToProto(payment entity.Payment) *prototransaction.Payment {
	return &prototransaction.Payment{
		Id:            payment.ID,
		Amount:        payment.Amount,
		Status:        payment.PaymentStatus,
		PaymentMethod: payment.PaymentMethod,
		CreatedAt:     payment.CreatedAt.Unix(),
		UpdatedAt:     payment.UpdatedAt.Unix(),
	}
}

func OrderModelToEntity(order model.Order) entity.Order {
	products := make([]entity.Product, len(order.ProductID))
	for i := range order.ProductID {
		products[i] = entity.Product{ID: order.ProductID[i]}
	}
	return entity.Order{
		ID:         order.ID,
		UserID:     order.UserID,
		Products:   products,
		TotalPrice: order.TotalPrice,
		Status:     0,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}
}
