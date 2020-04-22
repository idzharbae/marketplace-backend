package converter

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/model"
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

func ProductEntitiesToProtos(products []entity.Product) []*prototransaction.Product {
	res := make([]*prototransaction.Product, len(products))
	for i, product := range products {
		res[i] = ProductEntityToProto(product)
	}
	return res
}

func ProductEntityToProto(product entity.Product) *prototransaction.Product {
	return &prototransaction.Product{
		Id:         product.ID,
		ShopId:     product.ShopID,
		Name:       product.Name,
		AmountKg:   product.AmountKG,
		PricePerKg: product.PricePerKG,
		TotalPrice: product.TotalPrice,
		PhotoUrl:   product.PhotoURL,
	}
}

func CartModelsToEntities(carts []model.Cart) []entity.Cart {
	res := make([]entity.Cart, len(carts))
	for i, item := range carts {
		res[i] = CartModelToEntity(item)
	}
	return res
}

func CartModelToEntity(cart model.Cart) entity.Cart {
	return entity.Cart{
		ID:       cart.ID,
		Product:  entity.Product{ID: cart.ProductID},
		UserID:   cart.UserID,
		AmountKG: cart.AmountKG,
	}
}
func CartEntityToModel(cart entity.Cart) model.Cart {
	return model.Cart{
		ID:        cart.ID,
		ProductID: cart.Product.ID,
		UserID:    cart.UserID,
		AmountKG:  cart.AmountKG,
	}
}

func CartEntitiesToModels(carts []entity.Cart) []model.Cart {
	res := make([]model.Cart, len(carts))
	for i, cart := range carts {
		res[i] = CartEntityToModel(cart)
	}
	return res
}
