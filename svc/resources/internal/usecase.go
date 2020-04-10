package internal

import "github.com/idzharbae/marketplace-backend/svc/resources/internal/entity"

//go:generate mockgen -destination=usecase/ucmock/fileuc_mock.go -package=ucmock github.com/idzharbae/marketplace-backend/svc/resources/internal FileUC
type FileUC interface {
	UploadFile(req entity.File) (entity.File, error)
	DeleteFile(req entity.File) error
}
