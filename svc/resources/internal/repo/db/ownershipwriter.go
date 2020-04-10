package db

import (
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/repo/db/connection"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/repo/db/model"
)

type OwnershipWriter struct {
	db connection.Gormw
}

func NewOwnershipWriter(db connection.Gormw) *OwnershipWriter {
	return &OwnershipWriter{db: db}
}

func (ow *OwnershipWriter) Save(req entity.File) (entity.File, error) {
	fileModel := req.ToFileOwnershipModel()
	query := ow.db.Save(&fileModel)
	if err := query.Error(); err != nil {
		return entity.File{}, err
	}
	return req, nil
}

func (ow *OwnershipWriter) DeleteByID(id int64) error {
	var file model.FileOwnership
	query := ow.db.Where("id=?", id).First(&file)
	if query.RecordNotFound() {
		return errors.New("file does not exists")
	}
	if err := query.Error(); err != nil {
		return err
	}
	err := ow.db.Delete(&file).Error()
	return err
}
