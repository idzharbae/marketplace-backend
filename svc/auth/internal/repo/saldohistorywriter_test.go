package repo

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/connection/gormmock"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/catalog/util/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type saldoHistoryWriterTest struct {
	ctrl *gomock.Controller
	db   *gormmock.MockGormw
	unit internal.SaldoHistoryWriter
	ctx  context.Context
}

func newSaldoHistoryWriterTest() *saldoHistoryWriterTest {
	return &saldoHistoryWriterTest{}
}

func (srt *saldoHistoryWriterTest) Begin(t *testing.T) {
	srt.ctrl = gomock.NewController(t)
	srt.db = gormmock.NewMockGormw(srt.ctrl)
	srt.unit = NewSaldoHistoryWriter(srt.db)
	srt.ctx = context.Background()
}

func (srt *saldoHistoryWriterTest) Finish() {
	srt.ctrl.Finish()
}

func (srt *saldoHistoryWriterTest) GetEntity() entity.SaldoHistory {
	return entity.SaldoHistory{
		ID:           1,
		UserID:       2,
		SourceID:     3,
		Description:  "asdf",
		ChangeAmount: 123,
	}
}

func TestSaldoHistoryWriter_Create(t *testing.T) {
	test := newSaldoHistoryWriterTest()
	t.Run("db returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := test.GetEntity()
		test.db.EXPECT().Save(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.Create(req)

		assert.NotNil(t, err)
		assert.Equal(t, entity.SaldoHistory{}, got)
	})
	t.Run("db returns no error, should return saldo history entity", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := test.GetEntity()
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.SaldoHistory) *gormmock.MockGormw {
			assert.Equal(t, arg.ID, int64(0))
			*arg = model.SaldoHistory{
				ID:           1,
				UserID:       req.UserID,
				SourceID:     req.SourceID,
				Description:  req.Description,
				ChangeAmount: req.ChangeAmount,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.Create(req)

		assert.Nil(t, err)
		assert.Equal(t, got.ID, int64(1))
	})
}
