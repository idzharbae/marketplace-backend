package fileio

type IOUtil struct {
}

func NewIO() *IOUtil {
	return &IOUtil{}
}

func (io *IOUtil) CreateFile(path string, data []byte) error {
	return nil
}

func (io *IOUtil) DeleteFile(path string) error {
	return nil
}
