package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/resources/constants"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/repo/repomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type fileTest struct {
	ctrl       *gomock.Controller
	mockrepo   *repomock.MockFileWriter
	mockdb     *repomock.MockOwnershipWriter
	mockreader *repomock.MockOwnershipReader
	unit       internal.FileUC
}

func newFileTest() *fileTest {
	return &fileTest{}
}

func (ft *fileTest) Begin(t *testing.T) {
	ft.ctrl = gomock.NewController(t)
	ft.mockrepo = repomock.NewMockFileWriter(ft.ctrl)
	ft.mockdb = repomock.NewMockOwnershipWriter(ft.ctrl)
	ft.mockreader = repomock.NewMockOwnershipReader(ft.ctrl)
	ft.unit = NewFile(ft.mockrepo, ft.mockdb, ft.mockreader)
}

func (ft *fileTest) Finish() {
	ft.ctrl.Finish()
}

func TestFile_UploadFile(t *testing.T) {
	test := newFileTest()
	t.Run("missing ownerID should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.File{
			Extension: "jpg",
			Data:      []byte("GIF89a"),
		}
		got, err := test.unit.UploadFile(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.File{}, got)
	})
	t.Run("size too big, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.File{
			OwnerID:   123,
			Extension: "jpg",
			Data:      append([]byte("GIF89a"), make([]byte, constants.MaxFileSize)...),
		}
		got, err := test.unit.UploadFile(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.File{}, got)
	})
	t.Run("invalid extension, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.File{
			OwnerID:   123,
			Extension: "php",
			Data:      []byte("GIF89a"),
		}
		got, err := test.unit.UploadFile(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.File{}, got)
	})
	t.Run("invalid mimetype, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.File{
			OwnerID:   123,
			Extension: "jpeg",
			Data:      []byte("<?php system('ls'); ?>"),
		}
		got, err := test.unit.UploadFile(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.File{}, got)
	})
	t.Run("repo returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.File{
			OwnerID:   123,
			Extension: "jpeg",
			Data:      []byte{0xff, 0xd8, 0xff},
		}
		test.mockrepo.EXPECT().UploadFile(req).Return(entity.File{}, errors.New("error"))

		got, err := test.unit.UploadFile(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.File{}, got)
	})
	t.Run("db returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.File{
			OwnerID:   123,
			Extension: "jpeg",
			Data:      []byte{0xff, 0xd8, 0xff},
		}
		resp := entity.File{
			URL:       "asdf",
			Name:      "asdf",
			Extension: req.Extension,
			OwnerID:   req.OwnerID,
		}
		test.mockrepo.EXPECT().UploadFile(req).Return(resp, nil)
		test.mockdb.EXPECT().Save(resp).Return(entity.File{}, errors.New("error"))

		got, err := test.unit.UploadFile(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.File{}, got)
	})
	t.Run("db returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.File{
			OwnerID:   123,
			Extension: "jpeg",
			Data:      []byte{0xff, 0xd8, 0xff},
		}
		resp := entity.File{
			URL:       "asdf",
			Name:      "asdf",
			Extension: req.Extension,
			OwnerID:   req.OwnerID,
		}
		test.mockrepo.EXPECT().UploadFile(req).Return(resp, nil)
		test.mockdb.EXPECT().Save(resp).Return(entity.File{}, nil)

		got, err := test.unit.UploadFile(req)
		assert.Nil(t, err)
		assert.Equal(t, resp, got)
	})
}

func TestFile_DeleteFile(t *testing.T) {
	test := newFileTest()
	t.Run("db returns error when querying file, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.File{
			URL:     "asdf",
			OwnerID: 1337,
		}
		test.mockreader.EXPECT().GetByURL(req.URL).Return(entity.File{}, errors.New("error"))

		err := test.unit.DeleteFile(req)
		assert.NotNil(t, err)
	})
	t.Run("user does not own the file, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.File{
			URL:     "asdf",
			OwnerID: 1337,
		}
		test.mockreader.EXPECT().GetByURL(req.URL).Return(entity.File{
			Name:      "asdf",
			Extension: "Asdf",
			Type:      "asdff",
			Data:      nil,
			URL:       "asdf",
			OwnerID:   12312312,
		}, nil)

		err := test.unit.DeleteFile(req)
		assert.NotNil(t, err)
	})
	t.Run("db returns error when deleting ownership, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.File{
			URL:     "asdf",
			OwnerID: 1337,
		}
		resp := entity.File{
			ID:        1233,
			Name:      "asdf",
			Extension: "Asdf",
			Type:      "asdff",
			Data:      nil,
			URL:       "asdf",
			OwnerID:   req.OwnerID,
		}

		test.mockreader.EXPECT().GetByURL(req.URL).Return(resp, nil)
		test.mockdb.EXPECT().DeleteByID(resp.ID).Return(errors.New("error"))

		err := test.unit.DeleteFile(req)
		assert.NotNil(t, err)
	})
	t.Run("storage returns error when deleting file, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.File{
			URL:     "asdf",
			OwnerID: 1337,
		}
		resp := entity.File{
			ID:        1233,
			Name:      "asdf",
			Extension: "Asdf",
			Type:      "asdff",
			Data:      nil,
			URL:       "asdf",
			OwnerID:   req.OwnerID,
		}

		test.mockreader.EXPECT().GetByURL(req.URL).Return(resp, nil)
		test.mockdb.EXPECT().DeleteByID(resp.ID).Return(nil)
		test.mockrepo.EXPECT().DeleteFile(resp).Return(errors.New("error"))

		err := test.unit.DeleteFile(req)
		assert.NotNil(t, err)
	})
	t.Run("no error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.File{
			URL:     "asdf",
			OwnerID: 1337,
		}
		resp := entity.File{
			ID:        1233,
			Name:      "asdf",
			Extension: "Asdf",
			Type:      "asdff",
			Data:      nil,
			URL:       "asdf",
			OwnerID:   req.OwnerID,
		}

		test.mockreader.EXPECT().GetByURL(req.URL).Return(resp, nil)
		test.mockdb.EXPECT().DeleteByID(resp.ID).Return(nil)
		test.mockrepo.EXPECT().DeleteFile(resp).Return(nil)

		err := test.unit.DeleteFile(req)
		assert.Nil(t, err)
	})
}
