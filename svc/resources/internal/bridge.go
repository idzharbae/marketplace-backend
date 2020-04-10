package internal

//go:generate mockgen -destination=bridge/bridgemock/fileio_mock.go -package=bridgemock github.com/idzharbae/marketplace-backend/svc/resources/internal FileIO
type FileIO interface {
	CreateFile(path string, data []byte) error
	DeleteFile(path string) error
}
