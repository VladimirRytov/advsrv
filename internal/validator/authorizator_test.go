package validator

import (
	"bytes"
	"encoding/pem"
	"os"
	"reflect"
	"testing"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/encodedecoder"
)

func TestNewWorker(t *testing.T) {
	_, err := NewWorker(encodedecoder.NewBase64Encoder())
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerateToken(t *testing.T) {
	w, err := NewWorker(encodedecoder.NewBase64Encoder())
	if err != nil {
		t.Fatal(err)
	}
	token, err := w.SignDate("Petya", 1)
	if err != nil {
		t.Fatal(err)
	}
	if false {
		priv := bytes.NewBuffer(w.privateKeyBytes())
		pub := bytes.NewBuffer(w.publicKeyBytes())
		pu := pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY",
			Bytes: pub.Bytes()})
		pr := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: priv.Bytes(),
		})
		os.WriteFile("jwt", token, 0644)
		os.WriteFile("public", pu, 0644)
		os.WriteFile("private", pr, 0644)
		t.Fatal(err)
	}
}

func TestValidateToken(t *testing.T) {
	w, err := NewWorker(encodedecoder.NewBase64Encoder())
	if err != nil {
		t.Fatal(err)
	}
	tok, err := w.SignDate("Petya", 1)
	if err != nil {
		t.Fatal(err)
	}
	err = w.Validate(tok)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFetchPayload(t *testing.T) {
	want := datatransferobjects.UserToken{Log: "Petya", Perm: 1}
	w, err := NewWorker(encodedecoder.NewBase64Encoder())
	if err != nil {
		t.Fatal(err)
	}
	tok, err := w.SignDate(want.Log, want.Perm)
	if err != nil {
		t.Fatal(err)
	}
	userPayload, err := w.FetchPayload(tok)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(userPayload, want) {
		t.Fatalf("got wrong token.")
	}
}

func BenchmarkGenerateToken(b *testing.B) {
	b.StopTimer()
	w, err := NewWorker(encodedecoder.NewBase64Encoder())
	if err != nil {
		b.Fatal(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, err := w.SignDate("Petya", 1)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkValidateToken(b *testing.B) {
	b.StopTimer()
	w, err := NewWorker(encodedecoder.NewBase64Encoder())
	if err != nil {
		b.Fatal(err)
	}
	tok, err := w.SignDate("Petya", 1)
	if err != nil {
		b.Fatal(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w.Validate(tok)
	}
}
