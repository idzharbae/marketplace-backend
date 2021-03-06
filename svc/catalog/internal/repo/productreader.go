package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
	"github.com/lib/pq"
)

type ProductReader struct {
	db connection.Gormw
}

func NewProductReader(db connection.Gormw) *ProductReader {
	return &ProductReader{db: db}
}

func (p *ProductReader) ListAll(req requests.ListProduct) ([]entity.Product, error) {
	var products []model.Product
	db := p.db.Preload("Reviews")

	if req.Category != "" {
		db = db.Where("category=?", req.Category)
	}
	if req.Search != "" {
		db = db.Where("name ilike ?", "%"+req.Search+"%")
	}
	if req.ShopIDs != nil {
		db = db.Where("shop_id=ANY(?)", req.ShopIDs)
	}
	db = db.Order(getOrder(req.OrderBy, req.OrderType))
	if req.Pagination.Limit > 0 {
		db = db.Limit(req.Pagination.Limit)
	}
	if req.Pagination.Page > 0 {
		db = db.Offset(req.Pagination.OffsetFromPagination())
	}
	err := db.Find(&products).Error()
	if err != nil {
		return nil, err
	}
	return converter.ProductModelsToEntities(products), nil
}

func (p *ProductReader) ListByShopID(shopID int64, pagination requests.Pagination) ([]entity.Product, error) {
	var products []model.Product
	db := p.db.Where("shop_id=?", shopID).Order("id desc")
	db = applyPagination(pagination, db)
	err := db.Find(&products).Error()
	if err != nil {
		return nil, err
	}
	return converter.ProductModelsToEntities(products), nil
}

func (p *ProductReader) ListByIDs(productIDs pq.Int64Array) ([]entity.Product, error) {
	var products []model.Product
	db := p.db.Where("id=ANY(?)", productIDs).Order("id desc")
	err := db.Find(&products).Error()
	if err != nil {
		return nil, err
	}
	return converter.ProductModelsToEntities(products), nil
}

func (p *ProductReader) GetByID(productID int32) (entity.Product, error) {
	var product model.Product
	err := p.db.Preload("Reviews").Where("id=?", productID).First(&product).Error()
	if err != nil {
		return entity.Product{}, err
	}
	return converter.ProductModelToEntity(product), nil
}

func (p *ProductReader) GetBySlug(slug string) (entity.Product, error) {
	var product model.Product
	err := p.db.Preload("Reviews").Where("slug=?", slug).First(&product).Error()
	if err != nil {
		return entity.Product{}, err
	}
	return converter.ProductModelToEntity(product), nil
}

func (p *ProductReader) GetByShopID(shopID int32) ([]entity.Product, error) {
	var product []model.Product
	err := p.db.Preload("Reviews").Where("shop_id=?", shopID).First(&product).Error()
	if err != nil {
		return nil, err
	}
	return converter.ProductModelsToEntities(product), nil
}

func (p *ProductReader) GetTotalByShopID(shopID int32) (int32, error) {
	count := int32(0)
	err := p.db.Model(&model.Product{}).Where("shop_id=?", shopID).Count(&count).Error()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func getOrder(orderBy, orderType string) string {
	if orderBy != "id" && orderBy != "name" && orderBy != "created_at" && orderBy != "updated_at" &&
		orderBy != "price_per_kg" && orderBy != "stock_kg" && orderBy != "quantity" {
		orderBy = "id"
	}
	if orderType != "asc" && orderType != "desc" {
		orderType = "desc"
	}
	return orderBy + " " + orderType
}
