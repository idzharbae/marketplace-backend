package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/usecase/ucmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type shopTest struct {
	Ctrl *gomock.Controller
	UC   *ucmock.MockShopUC
	Unit *ShopService
	Ctx  context.Context
}

func newShopTest() *shopTest {
	return &shopTest{}
}

func (st *shopTest) Begin(t *testing.T) {
	st.Ctrl = gomock.NewController(t)
	st.UC = ucmock.NewMockShopUC(st.Ctrl)
	st.Ctx = context.Background()
	st.Unit = NewShopService(st.UC)
}

func (st *shopTest) Finish() {
	st.Ctrl.Finish()
}

func (st *shopTest) GetShopExample() *catalogproto.Shop {
	return &catalogproto.Shop{
		Id:       1,
		Name:     "testname",
		Address:  "testaddr",
		Location: &catalogproto.GPS{Longitude: 123, Latitude: 321},
		Products: []*catalogproto.Product{
			{Id: 1, Name: "testprod", PricePerKg: 123, StockKg: 123, ShopId: 1},
		},
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
}

func TestNewShopService(t *testing.T) {
	test := newShopTest()
	t.Run("should return shop service with uc attached", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		assert.NotNil(t, test.Unit)
		assert.NotNil(t, test.Unit.shopUC)
	})
}

func TestShopService_ListShops(t *testing.T) {
	test := newShopTest()
	t.Run("given nil param should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.Unit.ListShops(test.Ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &catalogproto.ListShopsReq{
			Pagination: &catalogproto.Pagination{
				Page:  1,
				Limit: 10,
			},
		}

		test.UC.EXPECT().List(converter.ListShopProtoToReq(req)).Return(nil, errors.New("error"))

		got, err := test.Unit.ListShops(test.Ctx, req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
	})
	t.Run("uc returns success should return list of shops", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &catalogproto.ListShopsReq{
			Pagination: &catalogproto.Pagination{
				Page:  1,
				Limit: 10,
			},
		}

		test.UC.EXPECT().List(converter.ListShopProtoToReq(req)).Return([]entity.Shop{
			{ID: 1, Name: "test"},
			{ID: 3, Name: "teststst"},
		}, nil)

		got, err := test.Unit.ListShops(test.Ctx, req)
		assert.NotNil(t, got)
		assert.Equal(t, 2, len(got.GetShops()))
		assert.Nil(t, err)
	})
}

func TestShopService_GetShopBySlug(t *testing.T) {
	test := newShopTest()
	t.Run("given nil param should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.Unit.GetShopBySlug(test.Ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		slug := "slug"

		test.UC.EXPECT().GetBySlug(slug).Return(entity.Shop{}, errors.New("error"))

		got, err := test.Unit.GetShopBySlug(test.Ctx, &catalogproto.GetShopBySlugReq{Slug: slug})
		assert.Nil(t, got)
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
	})
	t.Run("uc returns succes should return the shop object", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		slug := "slug"

		test.UC.EXPECT().GetBySlug(slug).Return(entity.Shop{ID: 1337, Name: "test", Slug: slug}, nil)

		got, err := test.Unit.GetShopBySlug(test.Ctx, &catalogproto.GetShopBySlugReq{Slug: slug})
		assert.NotNil(t, got)
		assert.Nil(t, err)
	})
}

func TestShopService_GetShopByID(t *testing.T) {
	test := newShopTest()
	t.Run("given nil param should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.Unit.GetShopByID(test.Ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		shopID := int32(1337)

		test.UC.EXPECT().GetByID(shopID).Return(entity.Shop{}, errors.New("error"))

		got, err := test.Unit.GetShopByID(test.Ctx, &catalogproto.GetShopByPKReq{Id: shopID})
		assert.Nil(t, got)
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
	})
	t.Run("uc returns succes should return the shop object", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		shopID := int32(1337)

		test.UC.EXPECT().GetByID(shopID).Return(entity.Shop{ID: 1337, Name: "test"}, nil)

		got, err := test.Unit.GetShopByID(test.Ctx, &catalogproto.GetShopByPKReq{Id: shopID})
		assert.NotNil(t, got)
		assert.Nil(t, err)
	})
}

func TestShopService_CreateShop(t *testing.T) {
	test := newShopTest()
	t.Run("given nil param should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.Unit.CreateShop(test.Ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := test.GetShopExample()
		test.UC.EXPECT().Create(gomock.Any()).Return(entity.Shop{}, errors.New("error"))

		got, err := test.Unit.CreateShop(test.Ctx, req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
	})
	t.Run("uc returns success, should return success", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := test.GetShopExample()
		test.UC.EXPECT().Create(gomock.Any()).Return(entity.Shop{ID: 1337}, nil)

		got, err := test.Unit.CreateShop(test.Ctx, req)
		assert.NotNil(t, got)
		assert.Equal(t, int32(1337), got.Id)
		assert.Nil(t, err)
	})
}

func TestShopService_UpdateShop(t *testing.T) {
	test := newShopTest()
	t.Run("given nil param should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.Unit.UpdateShop(test.Ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := test.GetShopExample()
		test.UC.EXPECT().Update(gomock.Any()).DoAndReturn(func(shop entity.Shop) (entity.Shop, error) {
			assert.Equal(t, req.GetName(), shop.Name)
			assert.Equal(t, req.GetAddress(), shop.Address)
			assert.Equal(t, req.GetLocation().GetLatitude(), shop.Location.Latitude)
			assert.Equal(t, req.GetLocation().GetLongitude(), shop.Location.Longitude)
			return entity.Shop{}, errors.New("error")
		})

		got, err := test.Unit.UpdateShop(test.Ctx, req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
	})
	t.Run("uc returns success, should return success", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := test.GetShopExample()
		test.UC.EXPECT().Update(gomock.Any()).DoAndReturn(func(shop entity.Shop) (entity.Shop, error) {
			assert.Equal(t, req.GetName(), shop.Name)
			assert.Equal(t, req.GetAddress(), shop.Address)
			assert.Equal(t, req.GetLocation().GetLatitude(), shop.Location.Latitude)
			assert.Equal(t, req.GetLocation().GetLongitude(), shop.Location.Longitude)
			return entity.Shop{ID: 1337}, nil
		})

		got, err := test.Unit.UpdateShop(test.Ctx, req)
		assert.NotNil(t, got)
		assert.Equal(t, int32(1337), got.Id)
		assert.Nil(t, err)
	})
}

func TestShopService_DeleteShop(t *testing.T) {
	test := newShopTest()
	t.Run("given nil param should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.Unit.DeleteShop(test.Ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := int32(13)
		test.UC.EXPECT().Delete(req).Return(errors.New("error"))

		got, err := test.Unit.DeleteShop(test.Ctx, &catalogproto.PKReq{Id: req})
		assert.Nil(t, got)
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
	})
	t.Run("uc returns success, should return success", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := int32(13)
		test.UC.EXPECT().Delete(req).Return(nil)

		got, err := test.Unit.DeleteShop(test.Ctx, &catalogproto.PKReq{Id: req})
		assert.NotNil(t, got)
		assert.Nil(t, err)
	})
}
