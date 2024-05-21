package authorizator

import (
	"math/rand"
	"testing"
)

func TestCanWriteFiles(t *testing.T) {
	au := NewAuthrizator()
	if !au.CanWriteFiles(1 << FileWrite) {
		t.Fatalf("wrong permission")
	}
}

func FuzzCanWriteFiles(f *testing.F) {
	for i := 0; i < 10000; i++ {
		f.Add(int32(rand.Intn(int(1<<30)) | (1 << FileWrite)))
	}
	au := NewAuthrizator()
	f.Fuzz(func(t *testing.T, i int32) {
		if !au.CanWriteFiles(i) {
			t.Fatalf("wrong permission. %b != %b", i, (1 << FileWrite))
		}
	})
}

func TestCanCreateFiles(t *testing.T) {
	au := NewAuthrizator()
	if !au.CanCreateFiles(1 << FileCreate) {
		t.Fatalf("wrong permission")
	}
}

func FuzzCanCreateFiles(f *testing.F) {
	for i := 0; i < 10000; i++ {
		f.Add(int32(rand.Intn(int(1<<30)) | (1 << FileCreate)))
	}
	au := NewAuthrizator()
	f.Fuzz(func(t *testing.T, i int32) {
		if !au.CanCreateFiles(i) {
			t.Fatalf("wrong permission. %b != %b", i, (1 << FileCreate))
		}
	})
}

func TestCanReadFiles(t *testing.T) {
	au := NewAuthrizator()
	if !au.CanReadFiles(1 << FileRead) {
		t.Fatalf("wrong permission")
	}
}

func FuzzCanReadFiles(f *testing.F) {
	for i := 0; i < 10000; i++ {
		f.Add(int32(rand.Intn(int(1<<30)) | (1 << FileRead)))
	}
	au := NewAuthrizator()
	f.Fuzz(func(t *testing.T, i int32) {
		if !au.CanReadFiles(i) {
			t.Fatalf("wrong permission. %b != %b", i, (1 << FileRead))
		}
	})
}

func TestCanDeleteFiles(t *testing.T) {
	au := NewAuthrizator()
	if !au.CanRemoveFiles(1 << FileDelete) {
		t.Fatalf("wrong permission")
	}
}

func FuzzCanRemoveFiles(f *testing.F) {
	for i := 0; i < 10000; i++ {
		f.Add(int32(rand.Intn(int(1<<30)) | (1 << FileDelete)))
	}
	au := NewAuthrizator()
	f.Fuzz(func(t *testing.T, i int32) {
		if !au.CanRemoveFiles(i) {
			t.Fatalf("wrong permission. %b != %b", i, (1 << FileDelete))
		}
	})
}
