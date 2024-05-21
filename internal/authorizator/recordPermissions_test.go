package authorizator

import (
	"math/rand"
	"testing"
)

func TestCanWriteRecords(t *testing.T) {
	au := NewAuthrizator()
	if !au.CanWriteRecords(1 << RecordWrite) {
		t.Fatalf("wrong permission")
	}
}

func FuzzCanWriteRecords(f *testing.F) {
	for i := 0; i < 10000; i++ {
		f.Add(int32(rand.Intn(int(1<<30)) | (1 << RecordWrite)))
	}
	au := NewAuthrizator()
	f.Fuzz(func(t *testing.T, i int32) {
		if !au.CanWriteRecords(i) {
			t.Fatalf("wrong permission. %b != %b", i, (1 << RecordWrite))
		}
	})
}

func TestCanCreateRecords(t *testing.T) {
	au := NewAuthrizator()
	if !au.CanCreateRecords(1 << RecordCreate) {
		t.Fatalf("wrong permission")
	}
}

func FuzzCanCreateRecords(f *testing.F) {
	for i := 0; i < 10000; i++ {
		f.Add(int32(rand.Intn(int(1<<30)) | (1 << RecordCreate)))
	}
	au := NewAuthrizator()
	f.Fuzz(func(t *testing.T, i int32) {
		if !au.CanCreateRecords(i) {
			t.Fatalf("wrong permission. %b != %b", i, (1 << RecordCreate))
		}
	})
}

func TestCanReadRecords(t *testing.T) {
	au := NewAuthrizator()
	if !au.CanReadRecords(1 << RecordRead) {
		t.Fatalf("wrong permission")
	}
}

func FuzzCanReadRecords(f *testing.F) {
	for i := 0; i < 10000; i++ {
		f.Add(int32(rand.Intn(int(1<<30)) | (1 << RecordRead)))
	}
	au := NewAuthrizator()
	f.Fuzz(func(t *testing.T, i int32) {
		if !au.CanReadRecords(i) {
			t.Fatalf("wrong permission. %b != %b", i, (1 << RecordRead))
		}
	})
}

func TestCanDeleteRecords(t *testing.T) {
	au := NewAuthrizator()
	if !au.CanRemoveRecords(1 << RecordDelete) {
		t.Fatalf("wrong permission")
	}
}

func FuzzCanRemoveRecords(f *testing.F) {
	for i := 0; i < 10000; i++ {
		f.Add(int32(rand.Intn(int(1<<30)) | (1 << RecordDelete)))
	}
	au := NewAuthrizator()
	f.Fuzz(func(t *testing.T, i int32) {
		if !au.CanRemoveRecords(i) {
			t.Fatalf("wrong permission. %b != %b", i, (1 << RecordDelete))
		}
	})
}
