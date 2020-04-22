package gateway

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/gateway/connection/connectionmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type catalogTest struct {
	ctrl   *gomock.Controller
	client *connectionmock.MockCatalog
	unit   internal.CatalogGateway
}

func newCatalogTest() *catalogTest {
	return &catalogTest{}
}

func (at *catalogTest) Begin(t *testing.T) {
	at.ctrl = gomock.NewController(t)
	at.client = connectionmock.NewMockCatalog(at.ctrl)
	at.unit = NewCatalog(at.client)
}

func (at *catalogTest) Finish() {
	at.ctrl.Finish()
}

func TestCatalog_GetProductByID(t *testing.T) {
	test := newCatalogTest()
	t.Run("client return error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int32(123)
		test.client.EXPECT().GetProduct(context.Background(), &catalogproto.GetProductReq{Id: req}).Return(nil, errors.New("error"))
		got, err := test.unit.GetProductByID(int64(req))
		assert.NotNil(t, err)
		assert.Equal(t, entity.Product{}, got)
	})
	t.Run("client return no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int32(123)
		test.client.EXPECT().GetProduct(context.Background(), &catalogproto.GetProductReq{Id: req}).Return(&catalogproto.Product{
			Id:         123,
			ShopId:     12,
			Name:       "asdf",
			PhotoUrl:   "qweqwe",
			PricePerKg: 12345,
		}, nil)
		got, err := test.unit.GetProductByID(int64(req))
		assert.Nil(t, err)
		assert.NotEqual(t, entity.Product{}, got)
	})
}

func TestCatalog_GetProductsByID(t *testing.T) {
	test := newCatalogTest()
	t.Run("client return error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := []int64{1, 2, 3}
		test.client.EXPECT().ListProducts(context.Background(), &catalogproto.ListProductsReq{ProductIds: req}).Return(nil, errors.New("error"))
		got, err := test.unit.GetProductsByID(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("client return no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := []int64{1, 2, 3}
		test.client.EXPECT().ListProducts(context.Background(), &catalogproto.ListProductsReq{ProductIds: req}).Return(&catalogproto.ListProductsResp{
			Products: []*catalogproto.Product{
				{Id: 1, ShopId: 2, PricePerKg: 3, Name: "123", PhotoUrl: "123"},
				{Id: 2, ShopId: 2, PricePerKg: 3, Name: "123", PhotoUrl: "123"},
				{Id: 3, ShopId: 2, PricePerKg: 3, Name: "123", PhotoUrl: "123"},
			},
		}, nil)
		got, err := test.unit.GetProductsByID(req)
		assert.Nil(t, err)
		assert.NotNil(t, got)
	})
}
