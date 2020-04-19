package converter

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
)

func CartEntitiesToProtos(carts []entity.Cart) []*prototransaction.Cart {
	res := make([]*prototransaction.Cart, len(carts))
	for i, cart := range carts {
		res[i] = CartEntityToProto(cart)
	}
	return res
}

func CartEntityToProto(cart entity.Cart) *prototransaction.Cart {
	return &prototransaction.Cart{
		Id:      cart.ID,
		UserId:  cart.UserID,
		Product: ProductEntityToProto(cart.Product),
	}
}

func ProductEntityToProto(product entity.Product) *prototransaction.Product {
	return &prototransaction.Product{
		Id:         product.ID,
		ShopId:     product.ShopID,
		Name:       product.Name,
		AmountKg:   product.AmountKG,
		PricePerKg: product.PricePerKG,
		TotalPrice: product.TotalPrice,
	}
}
