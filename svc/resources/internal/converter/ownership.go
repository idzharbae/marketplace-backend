package converter

import (
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/repo/db/model"
)

func OwnershipModelToEntity(f model.FileOwnership) entity.File {
	return entity.File{
		ID:        f.ID,
		Name:      f.Name(),
		Extension: f.Ext(),
		Type:      f.Type(),
		URL:       f.FileURL,
		OwnerID:   f.OwnerID,
	}
}
