package authorizator

import (
	"math/rand"
	"testing"
)

func FuzzCanWriteUsers(f *testing.F) {
	for i := 0; i < 10000; i++ {
		f.Add(int32(rand.Intn(int(1<<30)) | (1 << UserWrite)))
	}
	au := NewAuthrizator()
	f.Fuzz(func(t *testing.T, i int32) {
		if !au.CanWriteUsers(i) {
			t.Fatalf("wrong permission. %b != %b", i, (1 << UserWrite))
		}
	})
}

func TestCanWriteUsers(t *testing.T) {
	au := NewAuthrizator()
	if !au.CanWriteUsers(1 << UserWrite) {
		t.Fatalf("wrong permission")
	}
}

func TestCanCreateUsers(t *testing.T) {
	au := NewAuthrizator()
	if !au.CanCreateUsers(1 << UserCreate) {
		t.Fatalf("wrong permission")
	}
}

func FuzzCanCreateUsers(f *testing.F) {
	for i := 0; i < 10000; i++ {
		f.Add(int32(rand.Intn(int(1<<30)) | (1 << UserCreate)))
	}
	au := NewAuthrizator()
	f.Fuzz(func(t *testing.T, i int32) {
		if !au.CanCreateUsers(i) {
			t.Fatalf("wrong permission. %b != %b", i, (1 << UserCreate))
		}
	})
}

func TestCanReadUsers(t *testing.T) {
	au := NewAuthrizator()
	if !au.CanReadUsers(1 << UserRead) {
		t.Fatalf("wrong permission")
	}
}

func FuzzCanReadUsers(f *testing.F) {
	for i := 0; i < 10000; i++ {
		f.Add(int32(rand.Intn(int(1<<30)) | (1 << UserRead)))
	}
	au := NewAuthrizator()
	f.Fuzz(func(t *testing.T, i int32) {
		if !au.CanReadUsers(i) {
			t.Fatalf("wrong permission. %b != %b", i, (1 << UserRead))
		}
	})
}

func TestCanDeleteUsers(t *testing.T) {
	au := NewAuthrizator()
	if !au.CanRemoveUsers(1 << UserDelete) {
		t.Fatalf("wrong permission")
	}
}

func FuzzCanRemoveUsers(f *testing.F) {
	for i := 0; i < 10000; i++ {
		f.Add(int32(rand.Intn(int(1<<30)) | (1 << UserDelete)))
	}
	au := NewAuthrizator()
	f.Fuzz(func(t *testing.T, i int32) {
		if !au.CanRemoveUsers(i) {
			t.Fatalf("wrong permission. %b != %b", i, (1 << UserDelete))
		}
	})
}
