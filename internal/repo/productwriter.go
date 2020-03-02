package repo

import (
	"github.com/idzharbae/marketplace-backend/internal/converter"
	"github.com/idzharbae/marketplace-backend/internal/entity"
	"github.com/idzharbae/marketplace-backend/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/util/errors"
)

type ProductWriter struct {
	db connection.Gormw
}

func NewProductWriter(db connection.Gormw) *ProductWriter {
	return &ProductWriter{db: db}
}

func (pw *ProductWriter) Create(req entity.Product) (entity.Product, error) {
	const op = "ProductWriter::Create()"
	productModel := converter.ProductEntityToModel(req)
	productModel.ID = 0
	query := pw.db.Save(&productModel)
	if err := query.Error(); err != nil {
		return entity.Product{}, errors.WithPrefix(err, op)
	}
	productEntity := converter.ProductModelToEntity(productModel)
	return productEntity, nil
}

func (pw *ProductWriter) Update(req entity.Product) (entity.Product, error) {
	return entity.Product{}, nil
}

func (pw *ProductWriter) Delete(productID int32) error {
	return nil
}
