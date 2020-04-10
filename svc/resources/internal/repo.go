package internal

import "github.com/idzharbae/marketplace-backend/svc/resources/internal/entity"

//go:generate mockgen -destination=repo/repomock/filewriter_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/resources/internal FileWriter
type FileWriter interface {
	UploadFile(req entity.File) (entity.File, error)
	DeleteFile(req entity.File) error
}

//go:generate mockgen -destination=repo/repomock/ownershipwriter_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/resources/internal OwnershipWriter
type OwnershipWriter interface {
	Save(req entity.File) (entity.File, error)
	DeleteByID(id int64) error
}

//go:generate mockgen -destination=repo/repomock/ownershipreader_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/resources/internal OwnershipReader
type OwnershipReader interface {
	GetByURL(fileURL string) (entity.File, error)
}
