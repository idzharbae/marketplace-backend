package converter

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/model"
	"time"
)

func ReviewEntityToProto(review entity.Review) *catalogproto.Review {
	return &catalogproto.Review{
		Id:        review.ID,
		UserId:    review.UserID,
		ProductId: review.ProductID,
		ShopId:    review.ShopID,
		Title:     review.Title,
		Content:   review.Content,
		PhotoUrl:  review.PhotoURL,
		Rating:    review.Rating,
		CreatedAt: review.CreatedAt.Unix(),
		UpdatedAt: review.UpdatedAt.Unix(),
	}
}
func ReviewEntitiesToProtos(reviews []entity.Review) []*catalogproto.Review {
	protos := make([]*catalogproto.Review, len(reviews))
	for i, review := range reviews {
		protos[i] = ReviewEntityToProto(review)
	}
	return protos
}

func ReviewProtoToEntity(in *catalogproto.Review) entity.Review {
	return entity.Review{
		ID:        in.GetId(),
		UserID:    in.GetUserId(),
		ProductID: in.GetProductId(),
		ShopID:    in.GetShopId(),
		Title:     in.GetTitle(),
		Content:   in.GetContent(),
		PhotoURL:  in.GetPhotoUrl(),
		Rating:    in.GetRating(),
		CreatedAt: time.Unix(in.GetCreatedAt(), 0),
		UpdatedAt: time.Unix(in.GetUpdatedAt(), 0),
	}
}

func ReviewModelToEntity(mod model.Review) entity.Review {
	return entity.Review{
		ID:        mod.ID,
		UserID:    mod.UserID,
		ProductID: mod.ProductID,
		ShopID:    mod.ShopID,
		Title:     mod.Title,
		Content:   mod.Content,
		PhotoURL:  mod.PhotoURL,
		Rating:    mod.Rating,
		CreatedAt: mod.CreatedAt,
		UpdatedAt: mod.UpdatedAt,
	}
}

func ReviewModelsToEntities(mods []model.Review) []entity.Review {
	reviews := make([]entity.Review, len(mods))
	for i, review := range mods {
		reviews[i] = ReviewModelToEntity(review)
	}
	return reviews
}
