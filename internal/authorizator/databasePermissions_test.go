package authorizator

import (
	"math/rand"
	"testing"
)

func TestCanWriteDatabases(t *testing.T) {
	au := NewAuthrizator()
	if !au.CanWriteDatabases(1 << DatabaseWrite) {
		t.Fatalf("wrong permission")
	}
}

func FuzzCanWriteDatabases(f *testing.F) {
	for i := 0; i < 10000; i++ {
		f.Add(int32(rand.Intn(int(1<<30)) | (1 << DatabaseWrite)))
	}
	au := NewAuthrizator()
	f.Fuzz(func(t *testing.T, i int32) {
		if !au.CanWriteDatabases(i) {
			t.Fatalf("wrong permission. %b != %b", i, (1 << DatabaseWrite))
		}
	})
}

func TestSetCanWriteDatabases(t *testing.T) {
	au := NewAuthrizator()
	if !au.CanWriteDatabases(au.SetCanWriteDatabases()) {
		t.Fatalf("wrong permission")
	}
}

func TestCanCreateDatabases(t *testing.T) {
	au := NewAuthrizator()
	if !au.CanCreateDatabases(1 << DatabaseCreate) {
		t.Fatalf("wrong permission")
	}
}

func FuzzCanCreateDatabases(f *testing.F) {
	for i := 0; i < 10000; i++ {
		f.Add(int32(rand.Intn(int(1<<30)) | (1 << DatabaseCreate)))
	}
	au := NewAuthrizator()
	f.Fuzz(func(t *testing.T, i int32) {
		if !au.CanCreateDatabases(i) {
			t.Fatalf("wrong permission. %b != %b", i, (1 << DatabaseCreate))
		}
	})
}

func TestCanReadDatabases(t *testing.T) {
	au := NewAuthrizator()
	if !au.CanReadDatabases(1 << DatabaseRead) {
		t.Fatalf("wrong permission")
	}
}

func FuzzCanReadDatabases(f *testing.F) {
	for i := 0; i < 10000; i++ {
		f.Add(int32(rand.Intn(int(1<<30)) | (1 << DatabaseRead)))
	}
	au := NewAuthrizator()
	f.Fuzz(func(t *testing.T, i int32) {
		if !au.CanReadDatabases(i) {
			t.Fatalf("wrong permission. %b != %b", i, (1 << DatabaseRead))
		}
	})
}

func TestCanDeleteDatabases(t *testing.T) {
	au := NewAuthrizator()
	if !au.CanRemoveDatabases(1 << DatabaseDelete) {
		t.Fatalf("wrong permission")
	}
}

func FuzzCanRemoveDatabases(f *testing.F) {
	for i := 0; i < 10000; i++ {
		f.Add(int32(rand.Intn(int(1<<30)) | (1 << DatabaseDelete)))
	}
	au := NewAuthrizator()
	f.Fuzz(func(t *testing.T, i int32) {
		if !au.CanRemoveDatabases(i) {
			t.Fatalf("wrong permission. %b != %b", i, (1 << DatabaseDelete))
		}
	})
}
