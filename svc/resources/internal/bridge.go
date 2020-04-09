package internal

type FileIO interface {
	CreateFile(path string, data []byte) error
	DeleteFile(path string) error
}
