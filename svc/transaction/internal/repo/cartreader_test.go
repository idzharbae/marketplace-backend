package repo

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/connection/gormmock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/model"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

type cartReaderTest struct {
	ctrl *gomock.Controller
	db   *gormmock.MockGormw
	unit internal.CartReader
}

func newCartReaderTest() *cartReaderTest {
	return &cartReaderTest{}
}

func (ct *cartReaderTest) Begin(t *testing.T) {
	ct.ctrl = gomock.NewController(t)
	ct.db = gormmock.NewMockGormw(ct.ctrl)
	ct.unit = NewCartReader(ct.db)
}

func (ct *cartReaderTest) Finish() {
	ct.ctrl.Finish()
}

func TestCartReader_ListByUserID(t *testing.T) {
	test := newCartReaderTest()
	t.Run("db return error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)

		test.db.EXPECT().Where("user_id=?", req).Return(test.db)
		test.db.EXPECT().Find(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.ListByUserID(req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("db returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)

		test.db.EXPECT().Where("user_id=?", req).Return(test.db)
		test.db.EXPECT().Find(gomock.Any()).DoAndReturn(func(arg *[]model.Cart) *gormmock.MockGormw {
			*arg = []model.Cart{
				{
					ID:        123,
					ProductID: 456,
					UserID:    789,
					AmountKG:  123,
				},
			}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.ListByUserID(req)
		assert.NotNil(t, got)
		assert.Nil(t, err)
	})
}

func TestCartReader_GetByIDs(t *testing.T) {
	test := newCartReaderTest()
	t.Run("error when finding, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := []int64{1, 2, 3}

		test.db.EXPECT().Where("id=ANY(?)", gomock.Any()).DoAndReturn(func(query string, arg pq.Int64Array) *gormmock.MockGormw {
			arr := pq.Int64Array(req)
			assert.Equal(t, arr, arg)
			return test.db
		})
		test.db.EXPECT().Find(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.GetByIDs(req...)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := []int64{1, 2, 3}

		test.db.EXPECT().Where("id=ANY(?)", gomock.Any()).DoAndReturn(func(query string, arg pq.Int64Array) *gormmock.MockGormw {
			arr := pq.Int64Array(req)
			assert.Equal(t, arr, arg)
			return test.db
		})
		test.db.EXPECT().Find(gomock.Any()).DoAndReturn(func(arg *[]model.Cart) *gormmock.MockGormw {
			*arg = []model.Cart{
				{ID: 123},
			}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.GetByIDs(req...)
		assert.NotNil(t, got)
		assert.Nil(t, err)
	})
}
