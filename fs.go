package webdavfs

import (
	"context"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/webdav"
)

type FS struct {
	httpFS http.FileSystem
	iofs   fs.FS
}

var _ webdav.FileSystem = new(FS)

func New(iofs fs.FS) *FS {
	httpFS := http.FS(iofs)
	return &FS{
		httpFS: httpFS,
		iofs:   iofs,
	}
}

func (_ *FS) Mkdir(ctx context.Context, name string, perm os.FileMode) error {
	return ErrNotSupported
}

func (_ *FS) RemoveAll(ctx context.Context, name string) error {
	return ErrNotSupported
}

func (_ *FS) Rename(ctx context.Context, oldName string, newName string) error {
	return ErrNotSupported
}

func convertName(name string) string {
	if name == "/" || name == "" {
		return "."
	}
	name = strings.TrimLeft(name, "/")  // trim all prefix /
	name = strings.TrimRight(name, "/") // trim all suffix /
	return name
}

func (f *FS) OpenFile(ctx context.Context, name string, flag int, perm os.FileMode) (webdav.File, error) {
	name = convertName(name)
	httpFile, err := f.httpFS.Open(name)
	if err != nil {
		return nil, err
	}
	return &File{
		File: httpFile,
	}, nil
}

func (f *FS) Stat(ctx context.Context, name string) (os.FileInfo, error) {
	name = convertName(name)
	info, err := fs.Stat(f.iofs, name)
	if err != nil {
		return nil, err
	}
	return info, nil
}
