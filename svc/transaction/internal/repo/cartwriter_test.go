package repo

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/connection/gormmock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

type cartWriterTest struct {
	ctrl *gomock.Controller
	db   *gormmock.MockGormw
	unit internal.CartWriter
}

func newCartWriterTest() *cartWriterTest {
	return &cartWriterTest{}
}

func (cw *cartWriterTest) Begin(t *testing.T) {
	cw.ctrl = gomock.NewController(t)
	cw.db = gormmock.NewMockGormw(cw.ctrl)
	cw.unit = NewCartWriter(cw.db)
}

func (cw *cartWriterTest) Finish() {
	cw.ctrl.Finish()
}

func TestCartWriter_Create(t *testing.T) {
	test := newCartWriterTest()
	t.Run("product already exists for user, return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Cart{
			ID:       1,
			Product:  entity.Product{ID: 2},
			UserID:   3,
			AmountKG: 4,
		}
		test.db.EXPECT().Where("user_id=?", req.UserID).Return(test.db)
		test.db.EXPECT().Where("product_id=?", req.Product.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))
		got, err := test.unit.Create(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Cart{}, got)
	})
	t.Run("product already exists for user, return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Cart{
			ID:       1,
			Product:  entity.Product{ID: 2},
			UserID:   3,
			AmountKG: 4,
		}
		test.db.EXPECT().Where("user_id=?", req.UserID).Return(test.db)
		test.db.EXPECT().Where("product_id=?", req.Product.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().RecordNotFound().Return(false)
		got, err := test.unit.Create(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Cart{}, got)
	})
	t.Run("error when saving, return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Cart{
			ID:       1,
			Product:  entity.Product{ID: 2},
			UserID:   3,
			AmountKG: 4,
		}
		test.db.EXPECT().Where("user_id=?", req.UserID).Return(test.db)
		test.db.EXPECT().Where("product_id=?", req.Product.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().RecordNotFound().Return(true)
		test.db.EXPECT().Save(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))
		got, err := test.unit.Create(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Cart{}, got)
	})
	t.Run("saving success, return no error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Cart{
			ID:       1,
			Product:  entity.Product{ID: 2},
			UserID:   3,
			AmountKG: 4,
		}
		test.db.EXPECT().Where("user_id=?", req.UserID).Return(test.db)
		test.db.EXPECT().Where("product_id=?", req.Product.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().RecordNotFound().Return(true)
		test.db.EXPECT().Save(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(nil)
		got, err := test.unit.Create(req)
		assert.Nil(t, err)
		assert.NotEqual(t, entity.Cart{}, got)
	})
}

func TestCartWriter_Update(t *testing.T) {
	test := newCartWriterTest()
	t.Run("error when finding cart, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Cart{
			ID:       1,
			Product:  entity.Product{ID: 21},
			UserID:   2,
			AmountKG: 3,
		}

		test.db.EXPECT().Where("id=?", req.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.Update(req)
		assert.Equal(t, entity.Cart{}, got)
		assert.NotNil(t, err)
	})
	t.Run("cart does not exists, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Cart{
			ID:       1,
			Product:  entity.Product{ID: 21},
			UserID:   2,
			AmountKG: 3,
		}

		test.db.EXPECT().Where("id=?", req.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().RecordNotFound().Return(true)

		got, err := test.unit.Update(req)
		assert.Equal(t, entity.Cart{}, got)
		assert.NotNil(t, err)
	})
	t.Run("error when saving cart, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Cart{
			ID:       1,
			Product:  entity.Product{ID: 21},
			UserID:   2,
			AmountKG: 3,
		}

		test.db.EXPECT().Where("id=?", req.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.Cart) *gormmock.MockGormw {
			*arg = model.Cart{
				ID:        1,
				ProductID: 21,
				UserID:    2,
				AmountKG:  13,
			}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().RecordNotFound().Return(false)
		test.db.EXPECT().Save(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.Update(req)
		assert.Equal(t, entity.Cart{}, got)
		assert.NotNil(t, err)
	})
	t.Run("success when saving, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Cart{
			ID:       1,
			Product:  entity.Product{ID: 21},
			UserID:   2,
			AmountKG: 3,
		}

		test.db.EXPECT().Where("id=?", req.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.Cart) *gormmock.MockGormw {
			*arg = model.Cart{
				ID:        1,
				ProductID: 21,
				UserID:    2,
				AmountKG:  13,
			}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().RecordNotFound().Return(false)
		test.db.EXPECT().Save(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.Update(req)
		assert.Equal(t, req, got)
		assert.Nil(t, err)
	})
}

func TestCartWriter_DeleteByID(t *testing.T) {
	test := newCartWriterTest()
	t.Run("error when finding cart, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)

		test.db.EXPECT().Where("id=?", req).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		err := test.unit.DeleteByID(req)
		assert.NotNil(t, err)
	})
	t.Run("record not found, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)

		test.db.EXPECT().Where("id=?", req).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().RecordNotFound().Return(true)

		err := test.unit.DeleteByID(req)
		assert.NotNil(t, err)
	})
	t.Run("error when deleting cart, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)

		test.db.EXPECT().Where("id=?", req).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.Cart) *gormmock.MockGormw {
			*arg = model.Cart{
				ID:        123,
				ProductID: 32,
				UserID:    12,
				AmountKG:  1440,
			}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().RecordNotFound().Return(false)
		test.db.EXPECT().Delete(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		err := test.unit.DeleteByID(req)
		assert.NotNil(t, err)
	})
}
