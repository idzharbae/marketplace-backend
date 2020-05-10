package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/repomock"
	"testing"
)

type saldoHistoryTest struct {
	ctrl   *gomock.Controller
	reader *repomock.MockSaldoHistoryReader
	writer *repomock.MockSaldoHistoryWriter
	unit   *SaldoHistory
	ctx    context.Context
}

func newSaldoHistoryTest() *saldoHistoryTest {
	return &saldoHistoryTest{}
}

func (st *saldoHistoryTest) Begin(t *testing.T) {
	st.ctrl = gomock.NewController(t)
	st.reader = repomock.NewMockSaldoHistoryReader(st.ctrl)
	st.writer = repomock.NewMockSaldoHistoryWriter(st.ctrl)
	st.unit = NewSaldoHistory(st.reader, st.writer)
	st.ctx = context.Background()
}

func (st *saldoHistoryTest) Finish() {
	st.ctrl.Finish()
}
