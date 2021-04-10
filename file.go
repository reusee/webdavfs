package webdavfs

import (
	"net/http"

	"golang.org/x/net/webdav"
)

type File struct {
	http.File
}

var _ webdav.File = new(File)

func (f *File) Write(_ []byte) (int, error) {
	return 0, ErrNotSupported
}
