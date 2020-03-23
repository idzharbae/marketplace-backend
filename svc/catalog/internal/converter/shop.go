package converter

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/model"
	"time"
)

func ShopEntityToProto(s entity.Shop) *catalogproto.Shop {
	return &catalogproto.Shop{
		Id:      s.ID,
		Name:    s.Name,
		Address: s.Address,
		Location: &catalogproto.GPS{
			Latitude:  float32(s.Location.Latitude),
			Longitude: float32(s.Location.Longitude),
		},
		Products:  ProductEntitiesToProtos(s.Products),
		CreatedAt: s.CreatedAt.Unix(),
		UpdatedAt: s.UpdatedAt.Unix(),
	}
}

func ShopProtoToEntity(s *catalogproto.Shop) entity.Shop {
	return entity.Shop{
		ID:      s.GetId(),
		Name:    s.GetName(),
		Address: s.GetAddress(),
		Location: entity.GPS{
			Latitude:  float64(s.GetLocation().GetLatitude()),
			Longitude: float64(s.GetLocation().GetLongitude()),
		},
		Products:  ProductProtosToEntities(s.GetProducts()),
		CreatedAt: time.Unix(s.GetCreatedAt(), 0),
		UpdatedAt: time.Unix(s.GetUpdatedAt(), 0),
	}
}

func ShopEntitiesToProtos(ss []entity.Shop) []*catalogproto.Shop {
	shops := make([]*catalogproto.Shop, len(ss))
	for i, s := range ss {
		shops[i] = ShopEntityToProto(s)
	}
	return shops
}

func ShopModelToEntity(s model.Shop) entity.Shop {
	return entity.Shop{
		ID:      int32(s.ID),
		Name:    s.Name,
		Address: s.Address,
		Slug:    s.Slug,
		Location: entity.GPS{
			Latitude:  s.Latitude,
			Longitude: s.Longitude,
		},
		Products:  ProductModelsToEntities(s.Products),
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func ShopModelsToEntities(ss []model.Shop) []entity.Shop {
	shops := make([]entity.Shop, len(ss))
	for i, item := range ss {
		shops[i] = ShopModelToEntity(item)
	}
	return shops
}
