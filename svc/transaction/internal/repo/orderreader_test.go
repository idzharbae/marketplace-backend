package repo

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/constants"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/gateway/gatewaymock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/connection/gormmock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/request"
	"github.com/stretchr/testify/assert"
	"testing"
)

type orderReaderTest struct {
	ctrl    *gomock.Controller
	db      *gormmock.MockGormw
	catalog *gatewaymock.MockCatalogGateway
	unit    internal.OrderReader
}

func newOrderReaderTest() *orderReaderTest {
	return &orderReaderTest{}
}

func (ct *orderReaderTest) Begin(t *testing.T) {
	ct.ctrl = gomock.NewController(t)
	ct.db = gormmock.NewMockGormw(ct.ctrl)
	ct.catalog = gatewaymock.NewMockCatalogGateway(ct.ctrl)
	ct.unit = NewOrderReader(ct.db, ct.catalog)
}

func (ct *orderReaderTest) Finish() {
	ct.ctrl.Finish()
}

func TestOrderReader_ListByShopID(t *testing.T) {
	testList("shop_id", t)
}
func TestOrderReader_ListByUserID(t *testing.T) {
	testList("user_id", t)
}

func testList(arg string, t *testing.T) {
	test := newOrderReaderTest()
	t.Run("error fetching order, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)
		test.db.EXPECT().Preload("OrderProducts").Return(test.db)
		test.db.EXPECT().Where(arg+"=?", req).Return(test.db)
		test.db.EXPECT().Find(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		var got []entity.Order
		var err error
		if arg == "shop_id" {
			got, err = test.unit.ListByShopID(req, 0, request.Pagination{})
		} else {
			got, err = test.unit.ListByUserID(req, 0, request.Pagination{})
		}

		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("error fetching payment, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)
		orderModelList := test.GetOrderList()
		test.db.EXPECT().Preload("OrderProducts").Return(test.db)
		test.db.EXPECT().Where(arg+"=?", req).Return(test.db)
		test.db.EXPECT().Find(gomock.Any()).DoAndReturn(func(arg *[]model.Order) *gormmock.MockGormw {
			*arg = orderModelList
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Where("order_id=?", orderModelList[0].ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		var got []entity.Order
		var err error
		if arg == "shop_id" {
			got, err = test.unit.ListByShopID(req, 0, request.Pagination{})
		} else {
			got, err = test.unit.ListByUserID(req, 0, request.Pagination{})
		}
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("error fetching products, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)
		orderModelList := test.GetOrderList()
		test.db.EXPECT().Preload("OrderProducts").Return(test.db)
		test.db.EXPECT().Where(arg+"=?", req).Return(test.db)
		test.db.EXPECT().Find(gomock.Any()).DoAndReturn(func(arg *[]model.Order) *gormmock.MockGormw {
			*arg = orderModelList
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Where("order_id=?", orderModelList[0].ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.Payment) *gormmock.MockGormw {
			*arg = model.Payment{
				ID:            1,
				OrderID:       req,
				Amount:        123,
				PaymentMethod: 1,
				PaymentStatus: 2,
			}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.catalog.EXPECT().GetProductsByID([]int64(orderModelList[0].GetProductIDs())).Return(nil, errors.New("error"))

		var got []entity.Order
		var err error
		if arg == "shop_id" {
			got, err = test.unit.ListByShopID(req, 0, request.Pagination{})
		} else {
			got, err = test.unit.ListByUserID(req, 0, request.Pagination{})
		}
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("success fetching products, should return order data", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)
		orderModelList := test.GetOrderList()
		test.db.EXPECT().Preload("OrderProducts").Return(test.db)
		test.db.EXPECT().Where(arg+"=?", req).Return(test.db)
		test.db.EXPECT().Find(gomock.Any()).DoAndReturn(func(arg *[]model.Order) *gormmock.MockGormw {
			*arg = orderModelList
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Where("order_id=?", orderModelList[0].ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.Payment) *gormmock.MockGormw {
			*arg = model.Payment{
				ID:            1,
				OrderID:       req,
				Amount:        123,
				PaymentMethod: 1,
				PaymentStatus: 2,
			}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.catalog.EXPECT().GetProductsByID([]int64(orderModelList[0].GetProductIDs())).Return([]entity.Product{
			{ID: 1},
			{ID: 2},
			{ID: 3},
		}, nil)

		var got []entity.Order
		var err error
		if arg == "shop_id" {
			got, err = test.unit.ListByShopID(req, 0, request.Pagination{})
		} else {
			got, err = test.unit.ListByUserID(req, 0, request.Pagination{})
		}
		assert.Nil(t, err)
		assert.NotNil(t, got)
		assert.Equal(t, 1, len(got))
	})
	t.Run("given status param should filter by status", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)
		orderModelList := test.GetOrderList()
		test.db.EXPECT().Preload("OrderProducts").Return(test.db)
		test.db.EXPECT().Where(arg+"=?", req).Return(test.db)
		test.db.EXPECT().Where("status=?", int32(constants.OrderStatusOnShipment)).Return(test.db)
		test.db.EXPECT().Find(gomock.Any()).DoAndReturn(func(arg *[]model.Order) *gormmock.MockGormw {
			*arg = orderModelList
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Where("order_id=?", orderModelList[0].ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.Payment) *gormmock.MockGormw {
			*arg = model.Payment{
				ID:            1,
				OrderID:       req,
				Amount:        123,
				PaymentMethod: 1,
				PaymentStatus: 2,
			}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.catalog.EXPECT().GetProductsByID([]int64(orderModelList[0].GetProductIDs())).Return([]entity.Product{
			{ID: 1},
			{ID: 2},
			{ID: 3},
		}, nil)

		var got []entity.Order
		var err error
		if arg == "shop_id" {
			got, err = test.unit.ListByShopID(req, constants.OrderStatusOnShipment, request.Pagination{})
		} else {
			got, err = test.unit.ListByUserID(req, constants.OrderStatusOnShipment, request.Pagination{})
		}
		assert.Nil(t, err)
		assert.NotNil(t, got)
		assert.Equal(t, 1, len(got))
	})
}

func TestOrderReader_GetByID(t *testing.T) {
	test := newOrderReaderTest()
	t.Run("error fetching order, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)
		test.db.EXPECT().Where("id=?", req).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.GetByID(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Order{}, got)
	})
	t.Run("error fetching payment, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)
		test.db.EXPECT().Where("id=?", req).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.Order) *gormmock.MockGormw {
			*arg = model.Order{
				ID:            req,
				OrderProducts: []model.OrderProduct{{ID: 1}, {ID: 2}, {ID: 3}},
				UserID:        1,
				ShopID:        2,
				TotalPrice:    3,
				Status:        4,
			}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Where("order_id=?", req).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.GetByID(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Order{}, got)
	})
	t.Run("error fetching product, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)
		test.db.EXPECT().Where("id=?", req).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.Order) *gormmock.MockGormw {
			*arg = model.Order{
				ID:            req,
				OrderProducts: []model.OrderProduct{{ID: 1, ProductID: 1}, {ID: 2, ProductID: 2}, {ID: 3, ProductID: 3}},
				UserID:        1,
				ShopID:        2,
				TotalPrice:    3,
				Status:        4,
			}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Where("order_id=?", req).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.Payment) *gormmock.MockGormw {
			*arg = model.Payment{
				ID:            1,
				OrderID:       req,
				Amount:        1000,
				PaymentMethod: 1,
				PaymentStatus: 1,
			}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.catalog.EXPECT().GetProductsByID([]int64{1, 2, 3}).Return(nil, errors.New("error"))

		got, err := test.unit.GetByID(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Order{}, got)
	})
	t.Run("error fetching product, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)
		test.db.EXPECT().Where("id=?", req).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.Order) *gormmock.MockGormw {
			*arg = model.Order{
				ID:            req,
				OrderProducts: []model.OrderProduct{{ID: 1, ProductID: 1}, {ID: 2, ProductID: 2}, {ID: 3, ProductID: 3}},
				UserID:        1,
				ShopID:        2,
				TotalPrice:    3,
				Status:        4,
			}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Where("order_id=?", req).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.Payment) *gormmock.MockGormw {
			*arg = model.Payment{
				ID:            1,
				OrderID:       req,
				Amount:        1000,
				PaymentMethod: 1,
				PaymentStatus: 1,
			}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.catalog.EXPECT().GetProductsByID([]int64{1, 2, 3}).Return([]entity.Product{
			{ID: 1},
			{ID: 2},
			{ID: 3},
		}, nil)

		got, err := test.unit.GetByID(req)
		assert.Nil(t, err)
		assert.Equal(t, entity.Order{
			ID: req,
			Products: []entity.Product{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
			UserID:     1,
			ShopID:     2,
			TotalPrice: 3,
			Status:     4,
			Payment: entity.Payment{
				ID:            1,
				Amount:        1000,
				PaymentMethod: 1,
				PaymentStatus: 1,
			},
		}, got)
	})
}

func (ct *orderReaderTest) GetOrderList() []model.Order {
	return []model.Order{
		{
			ID:            1,
			OrderProducts: []model.OrderProduct{{ID: 1}, {ID: 2}, {ID: 3}},
			UserID:        2,
			ShopID:        3,
			TotalPrice:    4,
			Status:        5,
		},
	}
}
