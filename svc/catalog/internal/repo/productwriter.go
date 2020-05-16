package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/catalog/util/errors"
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
	if update.PhotoURL == "" {
		update.PhotoURL = found.PhotoURL
	}
	err := pw.db.Save(&update).Error()
	if err != nil {
		return entity.Product{}, errors.WithPrefix(err, op)
	}
	return converter.ProductModelToEntity(update), nil
}

func (pw *ProductWriter) DeleteByID(productID int32) error {
	const op = "ProductWriter::Delete()"
	var found model.Product
	query := pw.db.Where("id=?", productID).First(&found)
	if err := query.Error(); err != nil {
		return errors.WithPrefix(err, op)
	}
	if query.RecordNotFound() {
		return errors.NewWithPrefix("record not found", op)
	}
	err := pw.db.Delete(&model.Product{ID: found.ID}).Error()
	if err != nil {
		return errors.WithPrefix(err, op)
	}
	return nil
}

func (pw *ProductWriter) DeleteBySlug(productSlug string) error {
	const op = "ProductWriter::Delete()"
	var found model.Product
	query := pw.db.Where("slug=?", productSlug).First(&found)
	if err := query.Error(); err != nil {
		return errors.WithPrefix(err, op)
	}
	if query.RecordNotFound() {
		return errors.NewWithPrefix("record not found", op)
	}
	err := pw.db.Delete(&model.Product{ID: found.ID}).Error()
	if err != nil {
		return errors.WithPrefix(err, op)
	}
	return nil
}
