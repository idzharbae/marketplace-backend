package gateway

import (
	"context"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/gateway/connection"
)

type Catalog struct {
	conn connection.Catalog
}

func NewCatalog(conn connection.Catalog) *Catalog {
	return &Catalog{conn: conn}
}

func (c *Catalog) GetProductByID(productID int64) (entity.Product, error) {
	res, err := c.conn.GetProduct(context.Background(), &catalogproto.GetProductReq{Id: int32(productID)})
	if err != nil {
		return entity.Product{}, err
	}
	return entity.Product{
		ID:         int64(res.GetId()),
		ShopID:     int64(res.GetShopId()),
		Name:       res.GetName(),
		PricePerKG: res.GetPricePerKg(),
		PhotoURL:   res.GetPhotoUrl(),
	}, nil
}
func (c *Catalog) GetProductsByID(productIDs []int64) ([]entity.Product, error) {
	res, err := c.conn.ListProducts(context.Background(), &catalogproto.ListProductsReq{ProductIds: productIDs})
	if err != nil {
		return nil, err
	}
	products := make([]entity.Product, len(res.GetProducts()))
	for i, product := range res.GetProducts() {
		products[i] = entity.Product{
			ID:         int64(product.GetId()),
			ShopID:     int64(product.GetShopId()),
			Name:       product.GetName(),
			PricePerKG: product.GetPricePerKg(),
			PhotoURL:   product.GetPhotoUrl(),
		}
	}
	return products, nil
}
