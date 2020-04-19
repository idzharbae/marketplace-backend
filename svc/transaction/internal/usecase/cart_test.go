package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/repomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type cartTest struct {
	ctrl   *gomock.Controller
	reader *repomock.MockCartReader
	writer *repomock.MockCartWriter
	unit   internal.CartUC
}

func newCartTest() *cartTest {
	return &cartTest{}
}

func (ct *cartTest) Begin(t *testing.T) {
	ct.ctrl = gomock.NewController(t)
	ct.reader = repomock.NewMockCartReader(ct.ctrl)
	ct.writer = repomock.NewMockCartWriter(ct.ctrl)
	ct.unit = NewCart(ct.reader, ct.writer)
}
func (ct *cartTest) Finish() {
	ct.ctrl.Finish()
}

func TestCart_List(t *testing.T) {
	test := newCartTest()
	t.Run("given id == 0 should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.List(0)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("repo returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)
		test.reader.EXPECT().ListByUserID(req).Return(nil, errors.New("error"))
		got, err := test.unit.List(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
}

func TestCart_Add(t *testing.T) {
	test := newCartTest()
	t.Run("given amount <= 0 should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Cart{
			ID:       1337,
			Product:  entity.Product{ID: 123},
			UserID:   124,
			AmountKG: -203,
		}
		got, err := test.unit.Add(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Cart{}, got)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Cart{
			ID:       1337,
			Product:  entity.Product{ID: 123},
			UserID:   124,
			AmountKG: 124,
		}
		test.writer.EXPECT().Create(entity.Cart{
			ID:       0,
			Product:  req.Product,
			UserID:   req.UserID,
			AmountKG: req.AmountKG,
		}).Return(entity.Cart{}, errors.New("error"))
		got, err := test.unit.Add(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Cart{}, got)
	})
}

func TestCart_Update(t *testing.T) {
	test := newCartTest()
	t.Run("given id == 0 should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Cart{
			ID:       0,
			Product:  entity.Product{ID: 123},
			UserID:   124,
			AmountKG: 203,
		}
		got, err := test.unit.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Cart{}, got)
	})
	t.Run("given amount <= 0 should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Cart{
			Product:  entity.Product{ID: 123},
			UserID:   124,
			AmountKG: -203,
		}
		got, err := test.unit.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Cart{}, got)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Cart{
			ID:       123,
			Product:  entity.Product{ID: 123},
			UserID:   124,
			AmountKG: 124,
		}
		test.writer.EXPECT().Update(req).Return(entity.Cart{}, errors.New("error"))
		got, err := test.unit.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Cart{}, got)
	})
}

func TestCart_Remove(t *testing.T) {
	test := newCartTest()
	t.Run("given id == 0 should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(0)
		err := test.unit.Remove(req)
		assert.NotNil(t, err)
	})
	t.Run("repo returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)
		test.writer.EXPECT().DeleteByID(req).Return(errors.New("error"))
		err := test.unit.Remove(req)
		assert.NotNil(t, err)
	})
}
