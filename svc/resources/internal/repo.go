package internal

import "github.com/idzharbae/marketplace-backend/svc/resources/internal/entity"

type FileWriter interface {
	UploadFile(req entity.File) (entity.File, error)
	DeleteFile(req entity.File) error
}
