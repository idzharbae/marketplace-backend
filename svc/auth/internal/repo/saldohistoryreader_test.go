package repo

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/connection/gormmock"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/model"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

type saldoHistoryReaderTest struct {
	ctrl *gomock.Controller
	db   *gormmock.MockGormw
	unit internal.SaldoHistoryReader
	ctx  context.Context
}

func newSaldoHistoryReaderTest() *saldoHistoryReaderTest {
	return &saldoHistoryReaderTest{}
}

func (srt *saldoHistoryReaderTest) Begin(t *testing.T) {
	srt.ctrl = gomock.NewController(t)
	srt.db = gormmock.NewMockGormw(srt.ctrl)
	srt.unit = NewSaldoHistoryReader(srt.db)
	srt.ctx = context.Background()
}

func (srt *saldoHistoryReaderTest) Finish() {
	srt.ctrl.Finish()
}

func TestSaldoHistoryReader_ListByUserID(t *testing.T) {
	test := newSaldoHistoryReaderTest()
	t.Run("db returns error, should return error", func(t *testing.T) {
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
	t.Run("records not found, should return empty data", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)

		test.db.EXPECT().Where("user_id=?", req).Return(test.db)
		test.db.EXPECT().Find(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(gorm.ErrRecordNotFound)

		got, err := test.unit.ListByUserID(req)
		assert.Nil(t, got)
		assert.Nil(t, err)
	})
	t.Run("records found, should return saldo history data", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)

		test.db.EXPECT().Where("user_id=?", req).Return(test.db)
		test.db.EXPECT().Find(gomock.Any()).DoAndReturn(func(arg *[]model.SaldoHistory) *gormmock.MockGormw {
			*arg = []model.SaldoHistory{
				{ID: 1},
				{ID: 2},
			}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.ListByUserID(req)
		assert.Equal(t, len(got), 2)
		assert.Nil(t, err)
	})
}
