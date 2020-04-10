package db

import (
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/repo/db/connection"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/repo/db/model"
)

type OwnershipReader struct {
	db connection.Gormw
}

func NewOwnershipReader(db connection.Gormw) *OwnershipReader {
	return &OwnershipReader{db: db}
}

func (reader *OwnershipReader) GetByURL(fileURL string) (entity.File, error) {
	var fileOwnership model.FileOwnership
	query := reader.db.Where("file_url=?", fileURL).First(&fileOwnership)
	if query.RecordNotFound() {
		return entity.File{}, errors.New("file doesn't exists")
	}
	if err := query.Error(); err != nil {
		return entity.File{}, err
	}
	return converter.OwnershipModelToEntity(fileOwnership), nil
}
