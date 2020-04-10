package service

import (
	"context"
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/entity"

	"github.com/idzharbae/marketplace-backend/svc/resources/internal"
	"github.com/idzharbae/marketplace-backend/svc/resources/protoresources"
)

type FileService struct {
	FileUC internal.FileUC
}

func NewFileService(fileUC internal.FileUC) *FileService {
	return &FileService{FileUC: fileUC}
}

func (f *FileService) UploadPhoto(ctx context.Context, in *protoresources.UploadPhotoReq) (*protoresources.UploadPhotoResp, error) {
	if in == nil {
		return nil, errors.New("param should not be nil")
	}
	res, err := f.FileUC.UploadFile(entity.File{
		OwnerID:   in.GetOwnerId(),
		Extension: in.GetFileExt(),
		Data:      in.GetFile(),
	})
	if err != nil {
		return nil, err
	}
	return &protoresources.UploadPhotoResp{
		FileUrl: res.URL,
	}, nil
}

func (f *FileService) DeletePhoto(ctx context.Context, in *protoresources.DeletePhotoReq) (*protoresources.DeletePhotoResp, error) {
	if in == nil {
		return nil, errors.New("parameters should not be nil")
	}
	err := f.FileUC.DeleteFile(entity.File{URL: in.GetFileUrl(), OwnerID: in.GetUserId()})
	if err != nil {
		return nil, err
	}
	return &protoresources.DeletePhotoResp{Success: true}, nil
}
