package repo

import (
	"github.com/idzharbae/marketplace-backend/internal/converter"
	"github.com/idzharbae/marketplace-backend/internal/entity"
	"github.com/idzharbae/marketplace-backend/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/internal/repo/model"
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
	const op = "ProductWriter::Update()"
	var found model.Product
	query := pw.db.Where(model.Product{
		ID: req.ID,
	}).First(&found)
	if err := query.Error(); err != nil {
		return entity.Product{}, errors.WithPrefix(err, op)
	}
	if query.RecordNotFound() {
		return entity.Product{}, errors.NewWithPrefix("product does not exist", op)
	}
	update := converter.ProductEntityToModel(req)
	err := pw.db.Save(&update).Error()
	if err != nil {
		return entity.Product{}, errors.WithPrefix(err, op)
	}
	return converter.ProductModelToEntity(update), nil
}

func (pw *ProductWriter) Delete(productID int32) error {
	const op = "ProductWriter::Delete()"
	var found model.Product
	query := pw.db.Where("id=?", productID).First(&found)
	if err := query.Error(); err != nil {
		return errors.WithPrefix(err, op)
	}
	if query.RecordNotFound() {
		return errors.NewWithPrefix("record not found", op)
	}
	err := pw.db.Delete(&model.Product{ID: productID}).Error()
	if err != nil {
		return errors.WithPrefix(err, op)
	}
	return nil
}
