package repo

import (
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection/gormmock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/catalog/util/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type reviewWriterTest struct {
	ctrl *gomock.Controller
	db   *gormmock.MockGormw
	unit *ReviewWriter
}

func newReviewWriterTest() *reviewWriterTest {
	return &reviewWriterTest{}
}

func (wt *reviewWriterTest) Begin(t *testing.T) {
	wt.ctrl = gomock.NewController(t)
	wt.db = gormmock.NewMockGormw(wt.ctrl)
	wt.unit = NewReviewWriter(wt.db)
}

func (wt *reviewWriterTest) Finish() {
	wt.ctrl.Finish()
}

func TestReviewWriter_Create(t *testing.T) {
	test := newReviewWriterTest()
	t.Run("db returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{
			UserID:    123,
			ProductID: 321,
			ShopID:    123,
			Title:     "asd",
			Content:   "dsa",
			PhotoURL:  "pasdasd",
			Rating:    4.3,
		}

		test.db.EXPECT().Save(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.Create(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Review{}, got)
	})
	t.Run("db returns no error, should return review entity", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{
			UserID:    123,
			ProductID: 321,
			ShopID:    123,
			Title:     "asd",
			Content:   "dsa",
			PhotoURL:  "pasdasd",
			Rating:    4.3,
		}

		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Review) *gormmock.MockGormw {
			(*arg).ID = 1337
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.Create(req)
		assert.Nil(t, err)
		assert.Equal(t, int64(1337), got.ID)
	})
}

func TestReviewWriter_Update(t *testing.T) {
	test := newReviewWriterTest()
	t.Run("error when finding review, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{
			ID:        123,
			UserID:    321,
			ProductID: 123,
			ShopID:    321,
			Title:     "asd",
			Content:   "asd",
			PhotoURL:  "asd",
			Rating:    4.9,
		}

		test.db.EXPECT().Where("id=?", req.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Review{}, got)
	})
	t.Run("requesting user is not the review writer, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{
			ID:        123,
			UserID:    321,
			ProductID: 123,
			ShopID:    321,
			Title:     "asd",
			Content:   "asd",
			PhotoURL:  "asd",
			Rating:    4.9,
		}

		test.db.EXPECT().Where("id=?", req.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.Review) *gormmock.MockGormw {
			*arg = model.Review{ID: 123, UserID: 1}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Review{}, got)
	})
	t.Run("error when saving review, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{
			ID:        123,
			UserID:    321,
			ProductID: 123,
			ShopID:    321,
			Title:     "asd",
			Content:   "asd",
			PhotoURL:  "asd",
			Rating:    4.9,
		}

		test.db.EXPECT().Where("id=?", req.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.Review) *gormmock.MockGormw {
			*arg = model.Review{ID: 123, UserID: 321}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Review{}, got)
	})
	t.Run("error when saving review, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{
			ID:        123,
			UserID:    321,
			ProductID: 123,
			ShopID:    321,
			Title:     "asd",
			Content:   "asd",
			PhotoURL:  "asd",
			Rating:    4.9,
		}

		test.db.EXPECT().Where("id=?", req.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.Review) *gormmock.MockGormw {
			*arg = model.Review{ID: 123, UserID: 321}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Review) *gormmock.MockGormw {
			*arg = converter.ReviewEntityToModel(req)
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.Update(req)
		assert.Nil(t, err)
		assert.Equal(t, req, got)
	})
}

func TestReviewWriter_Delete(t *testing.T) {
	test := newReviewWriterTest()
	t.Run("error when finding review, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{ID: 123, UserID: 321}
		test.db.EXPECT().Where("id=?", req.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))
		err := test.unit.Delete(req)
		assert.NotNil(t, err)
	})
	t.Run("requesting user is not the review writer, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{ID: 123, UserID: 321}
		test.db.EXPECT().Where("id=?", req.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.Review) *gormmock.MockGormw {
			*arg = model.Review{ID: 123, UserID: 333}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		err := test.unit.Delete(req)
		assert.NotNil(t, err)
	})
	t.Run("error when deleting, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{
			ID:        123,
			UserID:    321,
			ProductID: 123,
			ShopID:    1231,
			Title:     "asdf",
			Content:   "Asdas",
			PhotoURL:  "ASdasd",
			Rating:    4.9,
		}
		test.db.EXPECT().Where("id=?", req.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.Review) *gormmock.MockGormw {
			*arg = model.Review{ID: 123, UserID: 321}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Delete(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))
		err := test.unit.Delete(req)
		assert.NotNil(t, err)
	})
	t.Run("no error when deleting, should return nil", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.Review{
			ID:        123,
			UserID:    321,
			ProductID: 123,
			ShopID:    1231,
			Title:     "asdf",
			Content:   "Asdas",
			PhotoURL:  "ASdasd",
			Rating:    4.9,
		}
		test.db.EXPECT().Where("id=?", req.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.Review) *gormmock.MockGormw {
			*arg = model.Review{ID: 123, UserID: 321}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Delete(gomock.Any()).DoAndReturn(func(arg *model.Review) *gormmock.MockGormw {
			assert.Equal(t, model.Review{ID: req.ID}, *arg)
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		err := test.unit.Delete(req)

		assert.Nil(t, err)
	})
}
