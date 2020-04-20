package converter

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
)

func OrderEntityToProto(order entity.Order) *prototransaction.Order {
	return &prototransaction.Order{
		Id:                   order.ID,
		UserId:               order.UserID,
		TotalPrice:           order.TotalPrice,
		Products:             ProductEntitiesToProtos(order.Products),
		Status:               order.Status,
		Payment:              nil,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
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
