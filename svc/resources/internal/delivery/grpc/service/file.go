package service

import (
	"context"
	"errors"

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
	return nil, errors.New("not implemented")
}
