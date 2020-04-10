package entity

import (
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/repo/db/model"
)

type File struct {
	ID        int64
	Name      string
	Extension string
	Type      string
	Data      []byte
	URL       string
	OwnerID   int64
}

func (f File) ToFileOwnershipModel() model.FileOwnership {
	return model.FileOwnership{
		ID:       f.ID,
		FilePath: "/" + f.Type + "/" + f.Name + "." + f.Extension,
		FileURL:  f.URL,
		OwnerID:  f.OwnerID,
	}
}
