package filehandler

import (
	"io"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

type FileStorage interface {
	Path(string) (string, error)
	SetFolder(string) error
	OpenForRead(name string) (io.ReadSeekCloser, int64, error)
	NewFile(string) (io.WriteCloser, string, error)
	OpenForWrite(name string) (io.WriteCloser, error)
	List() ([]datatransferobjects.File, error)
	Remove(name string) error
}

type Imager interface {
	Resize(in io.Reader, out io.Writer, lostgestLen int) error
	ErrFormat() error
}
