package filehandler

import (
	"bytes"
	"crypto/sha1"
	_ "embed"
	"errors"
	"io"
	"runtime"
	"sync"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/logging"
)

//go:embed file.png
var dummy []byte

const miniatureSize = 64

type Filehandler struct {
	fileHandler     FileStorage
	imager          Imager
	encodedHashData map[[20]byte][]byte
	rmu             sync.RWMutex
}

func NewFilehandler(fh FileStorage, img Imager) (*Filehandler, error) {
	files, err := fh.List()
	if err != nil {
		return nil, err
	}
	return &Filehandler{
		fileHandler:     fh,
		imager:          img,
		encodedHashData: make(map[[20]byte][]byte, len(files)*2),
	}, nil
}

func (fh *Filehandler) Init() error {
	logging.Logger.Info("filehandler.Init: loading images to the cache")
	var wg sync.WaitGroup
	var cores int = runtime.NumCPU()
	if cores > 1 {
		cores /= 2
	}
	files, err := fh.fileHandler.List()
	if err != nil {
		return err
	}

	subArraysLength := len(files) / cores
	mod := len(files) % cores
	for i := 0; i < cores; i++ {
		wg.Add(1)
		if (i+1)*subArraysLength+mod == len(files) {
			go fh.convertImages(&wg, files[i*subArraysLength:])
			break
		}
		go fh.convertImages(&wg, files[i*subArraysLength:(i+1)*subArraysLength])
	}
	wg.Wait()
	return nil
}

func (fh *Filehandler) convertImages(wg *sync.WaitGroup, files []datatransferobjects.File) {
	for i := range files {
		logging.Logger.Debug("converting image", "image", files[i].Name)
		fh.fillImageData(&files[i], miniatureSize)
	}
	defer wg.Done()
}

func (fh *Filehandler) GetMiniature(name string, size int) (datatransferobjects.File, error) {
	var file datatransferobjects.File
	file.Name = name
	err := fh.fillImageData(&file, size)
	if err != nil {
		return file, err
	}
	return file, err
}

func (fh *Filehandler) Get(name string) (datatransferobjects.File, error) {
	f, _, err := fh.fileHandler.OpenForRead(name)
	if err != nil {
		return datatransferobjects.File{}, err
	}
	defer f.Close()
	var b bytes.Buffer
	s, err := io.Copy(&b, f)
	if err != nil {
		return datatransferobjects.File{}, err
	}
	img := datatransferobjects.File{
		Name: name,
		Data: b.Bytes(),
		Size: s,
	}
	return img, err
}

func (fh *Filehandler) Set(name string, data io.ReadSeekCloser) (string, error) {
	f, name, err := fh.fileHandler.NewFile(name)
	if err != nil {
		return "", err
	}
	_, err = io.CopyBuffer(f, data, nil)
	if err != nil {
		f.Close()
		fh.fileHandler.Remove(name)
	}
	f.Close()
	return name, err
}

func (fh *Filehandler) List() ([]datatransferobjects.File, error) {
	return fh.fileHandler.List()
}

func (fh *Filehandler) ListWithMiniatures(size int) ([]datatransferobjects.File, error) {
	files, err := fh.fileHandler.List()
	if err != nil {
		return nil, err
	}
	for i := range files {
		data, err := fh.seekInCache(&files[i])
		if err == nil {
			files[i].Data = data
			continue
		}
		err = fh.fillImageData(&files[i], size)
		if err != nil || errors.Is(io.EOF, err) {
			files[i].Data = nil
		}
	}
	return files, nil
}

func (fh *Filehandler) seekInCache(file *datatransferobjects.File) ([]byte, error) {
	fh.rmu.RLock()
	defer fh.rmu.RUnlock()
	f, fileSize, err := fh.fileHandler.OpenForRead(file.Name)
	if err != nil {
		logging.Logger.Error("Filehandler.fillImageData: cannot open file", "error", err)
		return nil, err
	}
	defer f.Close()
	file.Size = fileSize
	sum, err := fh.fileDataHash(f)
	if err != nil {
		return nil, err
	}
	data, ok := fh.encodedHashData[sum]
	if ok {
		return data, nil
	}
	return nil, errors.New("filehandler.seekInCache: file not found")
}

func (fh *Filehandler) fileDataHash(in io.Reader) ([20]byte, error) {
	hasher := sha1.New()
	_, err := io.CopyBuffer(hasher, in, nil)
	if err != nil {
		logging.Logger.Error("Filehandler.fillImageData: an error occured while reading from file to buffer", "error", err)
		return [20]byte{}, err
	}
	sum := [20]byte(hasher.Sum(nil))
	return sum, nil
}

func (fh *Filehandler) fillImageData(file *datatransferobjects.File, size int) error {
	var b bytes.Buffer
	f, fileSize, err := fh.fileHandler.OpenForRead(file.Name)
	if err != nil {
		logging.Logger.Error("Filehandler.fillImageData: cannot open file", "error", err)
		return err
	}
	defer f.Close()
	file.Size = fileSize
	sum, err := fh.resizeOrPlaceDummy(f, &b, size)
	if err != nil {
		logging.Logger.Debug("Filehandler.fillImageData: cannot resize image", "error", err)
		return err
	}

	file.Data = b.Bytes()
	if _, ok := fh.encodedHashData[[20]byte(sum)]; !ok && size == 64 {
		fh.rmu.Lock()
		defer fh.rmu.Unlock()
		fh.encodedHashData[[20]byte(sum)] = file.Data
	}
	return err
}

func (fh *Filehandler) resizeOrPlaceDummy(input io.Reader, output io.Writer, size int) ([20]byte, error) {
	hasher := sha1.New()
	tee := io.TeeReader(input, hasher)
	err := fh.imager.Resize(tee, output, size)
	if errors.Is(err, fh.imager.ErrFormat()) {
		hasher.Reset()
		b := bytes.NewBuffer(dummy)
		tee := io.TeeReader(b, hasher)
		err = fh.imager.Resize(tee, output, size)
	}
	return [20]byte(hasher.Sum(nil)), err
}

func (fh *Filehandler) Remove(file string) error {
	fh.rmu.Lock()
	defer fh.rmu.Unlock()
	f, _, err := fh.fileHandler.OpenForRead(file)
	if err != nil {
		return err
	}
	sum, err := fh.fileDataHash(f)
	if err != nil {
		return err
	}
	f.Close()

	delete(fh.encodedHashData, sum)
	return fh.fileHandler.Remove(file)
}

func (fh *Filehandler) Path(name string) (string, error) {
	return fh.fileHandler.Path(name)
}

func (fh *Filehandler) ReloadCache() error {
	fh.rmu.Lock()
	for k := range fh.encodedHashData {
		delete(fh.encodedHashData, k)
	}
	fh.rmu.Unlock()
	return fh.Init()
}

func (fh *Filehandler) CleanCache() error {
	var saw bool
	hashList, err := fh.fileHashList()
	if err != nil {
		return err
	}
	fh.rmu.Lock()
	defer fh.rmu.Unlock()
	for k := range fh.encodedHashData {
		for i := range hashList {
			if k == hashList[i] {
				saw = true
				break
			}
		}
		if !saw {
			delete(fh.encodedHashData, k)
		}
		saw = false
	}
	return nil
}

func (fh *Filehandler) fileHashList() ([][20]byte, error) {
	list, err := fh.fileHandler.List()
	if err != nil {
		return nil, err
	}
	hashList := make([][20]byte, 0, len(list))
	for i := range list {
		f, _, err := fh.fileHandler.OpenForRead(list[i].Name)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		hash, err := fh.fileDataHash(f)
		if err != nil {
			return nil, err
		}
		hashList = append(hashList, hash)
	}
	return hashList, nil
}

func (fh *Filehandler) SetFolder(folder string) error {
	return fh.fileHandler.SetFolder(folder)
}
