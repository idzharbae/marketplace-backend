package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/usecase/ucmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type saldoHistoryTest struct {
	ctrl *gomock.Controller
	uc   *ucmock.MockSaldoHistoryUC
	unit *SaldoHistoryService
	ctx  context.Context
}

func newSaldoHistoryTest() *saldoHistoryTest {
	return &saldoHistoryTest{}
}

func (st *saldoHistoryTest) Begin(t *testing.T) {
	st.ctrl = gomock.NewController(t)
	st.uc = ucmock.NewMockSaldoHistoryUC(st.ctrl)
	st.unit = NewSaldoHistoryService(st.uc)
	st.ctx = context.Background()
}

func (st *saldoHistoryTest) Finish() {
	st.ctrl.Finish()
}

func (st *saldoHistoryTest) GetProto() *authproto.SaldoHistory {
	return &authproto.SaldoHistory{
		Id:           0,
		UserId:       1,
		SourceId:     2,
		Description:  "topup",
		ChangeAmount: 5000,
	}
}
func (st *saldoHistoryTest) GetEntity() entity.SaldoHistory {
	return entity.SaldoHistory{
		ID:           1,
		UserID:       1,
		SourceID:     2,
		Description:  "topup",
		ChangeAmount: 5000,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
func (st *saldoHistoryTest) GetListReq() *authproto.ListSaldoHistoryReq {
	return &authproto.ListSaldoHistoryReq{UserId: 123}
}
func (st *saldoHistoryTest) GetListResp() []entity.SaldoHistory {
	return []entity.SaldoHistory{
		{ID: 1},
		{ID: 2},
	}
}

func TestSaldoHistoryService_CreateSaldoHistory(t *testing.T) {
	test := newSaldoHistoryTest()
	t.Run("given nil params, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.CreateSaldoHistory(test.ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetProto()
		test.uc.EXPECT().Create(entity.SaldoHistory{
			ID:           req.GetId(),
			UserID:       req.GetUserId(),
			SourceID:     req.GetSourceId(),
			Description:  req.GetDescription(),
			ChangeAmount: req.GetChangeAmount(),
		}).Return(entity.SaldoHistory{}, errors.New("error"))

		got, err := test.unit.CreateSaldoHistory(test.ctx, test.GetProto())
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc no error, should return saldo history", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetProto()
		resp := test.GetEntity()
		test.uc.EXPECT().Create(entity.SaldoHistory{
			ID:           req.GetId(),
			UserID:       req.GetUserId(),
			SourceID:     req.GetSourceId(),
			Description:  req.GetDescription(),
			ChangeAmount: req.GetChangeAmount(),
		}).Return(resp, nil)

		got, err := test.unit.CreateSaldoHistory(test.ctx, test.GetProto())
		assert.Nil(t, err)
		assert.Equal(t, got.GetId(), resp.ID)
		assert.Equal(t, got.GetCreatedAt(), resp.CreatedAt.Unix())
		assert.Equal(t, got.GetUpdatedAt(), resp.UpdatedAt.Unix())
	})
}

func TestSaldoHistoryService_ListSaldoHistory(t *testing.T) {
	test := newSaldoHistoryTest()
	t.Run("given nil params, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.ListSaldoHistory(test.ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetListReq()
		test.uc.EXPECT().List(request.ListSaldoHistory{
			UserID: req.GetUserId(),
		}).Return(nil, errors.New("error"))

		got, err := test.unit.ListSaldoHistory(test.ctx, req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc no error, should return saldo history", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetListReq()
		resp := test.GetListResp()
		test.uc.EXPECT().List(request.ListSaldoHistory{
			UserID: req.GetUserId(),
		}).Return(resp, nil)

		got, err := test.unit.ListSaldoHistory(test.ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, len(got.GetSaldoHistories()), len(resp))
	})
}
