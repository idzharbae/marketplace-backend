package converter

import (
	"github.com/idzharbae/marketplace-backend/internal/entity"
	"github.com/idzharbae/marketplace-backend/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/marketplaceproto"
)

func ProductModelToEntity(p model.Product) entity.Product {
	return entity.Product{
		ID:         p.ID,
		ShopID:     p.ShopID,
		Name:       p.Name,
		Quantity:   p.Quantity,
		PricePerKG: p.PricePerKG,
		StockKG:    p.StockKG,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}

func ProductModelsToEntities(ps []model.Product) []entity.Product {
	products := make([]entity.Product, len(ps))
	for i, p := range ps {
		products[i] = ProductModelToEntity(p)
	}
	return products
}

func ProductEntityToProto(p entity.Product) *marketplaceproto.Product {
	return &marketplaceproto.Product{
		ID:         p.ID,
		ShopID:     p.ShopID,
		Name:       p.Name,
		Quantity:   p.Quantity,
		PricePerKG: p.PricePerKG,
		StockKG:    p.StockKG,
		CreatedAt:  p.CreatedAt.Unix(),
		UpdatedAt:  p.UpdatedAt.Unix(),
	}
}

func ProductEntitiesToProtos(ps []entity.Product) []*marketplaceproto.Product {
	products := make([]*marketplaceproto.Product, len(ps))
	for i, p := range ps {
		products[i] = ProductEntityToProto(p)
	}
	return products
}
