package filestorage

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/logging"
)

type Storage struct {
	PathPrefix string
}

func NewFileStorage() *Storage {
	storage := new(Storage)
	return storage
}

func (s *Storage) SetFolder(name string) error {
	_, err := os.Stat(name)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			logging.Logger.Warn("Storage.SetFolder: folder "+name+"not exitst. ", "error", err)
			return os.Mkdir(name, 0644)
		}
	}
	s.PathPrefix = name
	return err
}

func (s *Storage) OpenForWrite(name string) (io.WriteCloser, error) {
	f, err := os.OpenFile(filepath.Join(s.PathPrefix, name), os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (s *Storage) NewFile(name string) (io.WriteCloser, string, error) {
	var i int
	var stri string
	ext := filepath.Ext(name)
	fileName, _ := strings.CutSuffix(name, ext)
	for {
		if i > 0 {
			stri = "_" + strconv.Itoa(i)
		}
		newName := fileName + stri + ext
		filePath := filepath.Join(s.PathPrefix, newName)
		_, err := os.Stat(filePath)
		if errors.Is(err, os.ErrNotExist) {
			f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				return nil, "", err
			}

			return f, newName, nil
		}
		i++
	}
}

func (s *Storage) OpenForRead(name string) (io.ReadSeekCloser, int64, error) {
	f, err := os.OpenFile(filepath.Join(s.PathPrefix, name), os.O_RDONLY, 0644)
	if err != nil {
		return nil, 0, err
	}
	fStat, err := f.Stat()
	if err != nil {
		return nil, 0, err
	}
	return f, fStat.Size(), nil
}

func (s *Storage) Remove(name string) error {
	return os.Remove(filepath.Join(s.PathPrefix, name))
}

func (s *Storage) List() ([]datatransferobjects.File, error) {
	d := os.DirFS(s.PathPrefix)
	allFiles, err := fs.Glob(d, "*")
	if err != nil {
		return nil, err
	}
	files := make([]datatransferobjects.File, 0, len(allFiles))
	for i := range allFiles {
		info, err := os.Stat(filepath.Join(s.PathPrefix, allFiles[i]))
		if err != nil {
			continue
		}
		if info.IsDir() {
			continue
		}
		files = append(files, datatransferobjects.File{Name: allFiles[i], Size: info.Size()})
	}
	return files, err
}

func (s *Storage) Path(name string) (string, error) {
	f, err := os.Stat(filepath.Join(s.PathPrefix, name))
	if err != nil {
		return "", err
	}
	if f.IsDir() {
		return "", os.ErrInvalid
	}
	return filepath.Join(s.PathPrefix, name), nil
}
