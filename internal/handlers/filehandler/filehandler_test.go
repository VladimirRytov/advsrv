package filehandler

import (
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/filestorage"
	"github.com/VladimirRytov/advsrv/internal/logging"
	"github.com/VladimirRytov/advsrv/internal/minimizer"
)

func TestInit(t *testing.T) {
	CreateLogger()
	fstor := filestorage.NewFileStorage()

	fh, err := NewFilehandler(fstor, &minimizer.Imager{})
	if err != nil {
		t.Fatal(err)
	}

	err = fh.Init()
	if err == nil {
		t.Fatal(err, len(fh.encodedHashData))
	}
}

func TestLoadImage(t *testing.T) {
	fstor := filestorage.NewFileStorage()

	fh, err := NewFilehandler(fstor, &minimizer.Imager{})
	if err != nil {
		t.Fatal(err)
	}
	file := &datatransferobjects.File{
		Name: "_DSC0021.jpg",
	}
	err = fh.fillImageData(file, 128)
	if err != nil {
		t.Fatal(err)
	}
	f, err := os.OpenFile("got.jpeg", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.Write(file.Data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestListWithMiniatures(t *testing.T) {
	fstor := filestorage.NewFileStorage()

	err := os.Mkdir("D:/testImages/deploy", 0750)
	if err != nil {
		t.Fatal(err)
	}
	fh, err := NewFilehandler(fstor, &minimizer.Imager{})
	if err != nil {
		t.Fatal(err)
	}
	files, err := fh.ListWithMiniatures(128)
	if err != nil {
		t.Fatal(err)
	}
	for i := range files {
		f, err := os.OpenFile(filepath.Join("D:/testImages/deploy", files[i].Name), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			t.Fatal(err)
		}
		_, err = f.Write(files[i].Data)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelDebug}, false, os.Stderr)
}
