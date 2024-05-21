package filestorage

import (
	"slices"
	"testing"
)

func TestOpenForWrite(t *testing.T) {
	fileName := "test"
	files := &Storage{}
	f, err := files.OpenForWrite(fileName)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	f.Write([]byte("Hello"))
	f.Write([]byte("World"))

}

func TestOpenForRead(t *testing.T) {
	fileName := "test"
	files := &Storage{}
	f, _, err := files.OpenForRead(fileName)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	testCase := []byte("HelloWorld")
	var b []byte = make([]byte, len(testCase))
	f.Read(b)
	if !slices.Equal(b, testCase) {
		t.Fatalf("want %v,got %v", testCase, b)
	}
}

func TestList(t *testing.T) {
	files := &Storage{PathPrefix: "."}
	f, err := files.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(f) != 3 {
		t.Fatal(err)
	}
}
