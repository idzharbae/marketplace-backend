package converter

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/model"
	"time"
)

func ProductModelToEntity(p model.Product) entity.Product {
	return entity.Product{
		ID:         p.ID,
		ShopID:     p.ShopID,
		Name:       p.Name,
		Slug:       p.Slug,
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

func ProductEntityToModel(p entity.Product) model.Product {
	return model.Product{
		ID:         p.ID,
		ShopID:     p.ShopID,
		Name:       p.Name,
		Slug:       p.Slug,
		Quantity:   p.Quantity,
		PricePerKG: p.PricePerKG,
		StockKG:    p.StockKG,
	}
}

func ProductEntityToProto(p entity.Product) *catalogproto.Product {
	return &catalogproto.Product{
		Id:         p.ID,
		ShopId:     p.ShopID,
		Name:       p.Name,
		Quantity:   p.Quantity,
		Slug:       p.Slug,
		PricePerKg: p.PricePerKG,
		StockKg:    p.StockKG,
		CreatedAt:  p.CreatedAt.Unix(),
		UpdatedAt:  p.UpdatedAt.Unix(),
	}
}

func ProductEntitiesToProtos(ps []entity.Product) []*catalogproto.Product {
	products := make([]*catalogproto.Product, len(ps))
	for i, p := range ps {
		products[i] = ProductEntityToProto(p)
	}
	return products
}

func ProductProtoToEntity(p *catalogproto.Product) entity.Product {
	return entity.Product{
		ID:         p.GetId(),
		ShopID:     p.GetShopId(),
		Name:       p.GetName(),
		Slug:       p.GetSlug(),
		Quantity:   p.GetQuantity(),
		PricePerKG: p.GetPricePerKg(),
		StockKG:    p.GetStockKg(),
		CreatedAt:  time.Unix(p.GetCreatedAt(), 0),
		UpdatedAt:  time.Unix(p.GetUpdatedAt(), 0),
	}
}
