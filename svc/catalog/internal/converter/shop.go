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
			Latitude:  s.Location.Latitude,
			Longitude: s.Location.Longitude,
		},
		Products:  ProductEntitiesToProtos(s.Products),
		CreatedAt: s.CreatedAt.Unix(),
		UpdatedAt: s.UpdatedAt.Unix(),
		PhotoUrl:  s.PhotoURL,
	}
}

func ShopProtoToEntity(s *catalogproto.Shop) entity.Shop {
	return entity.Shop{
		ID:      s.GetId(),
		Slug:    s.GetSlug(),
		Name:    s.GetName(),
		Address: s.GetAddress(),
		Location: entity.GPS{
			Latitude:  s.GetLocation().GetLatitude(),
			Longitude: s.GetLocation().GetLongitude(),
		},
		PhotoURL:  s.GetPhotoUrl(),
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
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
		PhotoURL:  s.PhotoURL,
	}
}

func ShopModelsToEntities(ss []model.Shop) []entity.Shop {
	shops := make([]entity.Shop, len(ss))
	for i, item := range ss {
		shops[i] = ShopModelToEntity(item)
	}
	return shops
}

func ShopEntityToModel(s entity.Shop) model.Shop {
	return model.Shop{
		ID:        int64(s.ID),
		Name:      s.Name,
		Address:   s.Address,
		Longitude: s.Location.Longitude,
		Latitude:  s.Location.Latitude,
		Slug:      s.Slug,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
		PhotoURL:  s.PhotoURL,
	}
}
