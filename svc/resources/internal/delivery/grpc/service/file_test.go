package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/usecase/ucmock"
	"github.com/idzharbae/marketplace-backend/svc/resources/protoresources"
	"github.com/stretchr/testify/assert"
)

type fileServiceTest struct {
	ctrl   *gomock.Controller
	mockuc *ucmock.MockFileUC
	unit   *FileService
}

func newFileServiceTest() *fileServiceTest {
	return &fileServiceTest{}
}

func (fst *fileServiceTest) Begin(t *testing.T) {
	fst.ctrl = gomock.NewController(t)
	fst.mockuc = ucmock.NewMockFileUC(fst.ctrl)
	fst.unit = NewFileService(fst.mockuc)
}

func (fst *fileServiceTest) Finish() {
	fst.ctrl.Finish()
}

func TestFileService_UploadPhoto(t *testing.T) {
	test := newFileServiceTest()
	t.Run("given nil param should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.UploadPhoto(context.Background(), nil)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &protoresources.UploadPhotoReq{
			OwnerId: 123,
			File:    []byte("GIF"),
			FileExt: "jpg",
		}

		test.mockuc.EXPECT().UploadFile(entity.File{
			OwnerID:   req.GetOwnerId(),
			Extension: req.GetFileExt(),
			Data:      req.GetFile(),
		}).Return(entity.File{}, errors.New("error"))

		got, err := test.unit.UploadPhoto(context.Background(), req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &protoresources.UploadPhotoReq{
			OwnerId: 123,
			File:    []byte("GIF"),
			FileExt: "jpg",
		}
		res := entity.File{
			URL: "127.0.0.1:1441/img/asdf.jpg",
		}

		test.mockuc.EXPECT().UploadFile(entity.File{
			OwnerID:   req.GetOwnerId(),
			Extension: req.GetFileExt(),
			Data:      req.GetFile(),
		}).Return(res, nil)

		got, err := test.unit.UploadPhoto(context.Background(), req)
		assert.NotNil(t, got)
		assert.Equal(t, res.URL, got.GetFileUrl())
		assert.Nil(t, err)
	})
}

func TestFileService_DeletePhoto(t *testing.T) {
	test := newFileServiceTest()
	t.Run("given nil params should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.DeletePhoto(context.Background(), nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &protoresources.DeletePhotoReq{
			UserId:  123,
			FileUrl: "asdfg",
		}
		test.mockuc.EXPECT().DeleteFile(entity.File{
			URL:     req.GetFileUrl(),
			OwnerID: req.GetUserId(),
		}).Return(errors.New("error"))

		got, err := test.unit.DeletePhoto(context.Background(), req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &protoresources.DeletePhotoReq{
			FileUrl: "asdfg",
			UserId:  123,
		}
		test.mockuc.EXPECT().DeleteFile(entity.File{
			URL:     req.GetFileUrl(),
			OwnerID: req.GetUserId(),
		}).Return(nil)

		got, err := test.unit.DeletePhoto(context.Background(), req)
		assert.NotNil(t, got)
		assert.Nil(t, err)
	})
}
