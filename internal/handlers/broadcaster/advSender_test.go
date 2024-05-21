package broadcaster

import (
	"log/slog"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/VladimirRytov/advsrv/internal/encodedecoder"
	"github.com/VladimirRytov/advsrv/internal/logging"
	"github.com/VladimirRytov/advsrv/internal/validator"
)

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
}

func TestSendData(t *testing.T) {
	CreateLogger()
	validator, err := validator.NewWorker(encodedecoder.NewBase64Encoder())
	if err != nil {
		t.Fatal(err)
	}
	sm := &SenderMaker{}
	br := NewSubHandler(validator, sm, encodedecoder.NewBase64Encoder())
	for i := 0; i < 10; i++ {
		br.NewPassiveSub(strconv.Itoa(i), "i")
	}
	d := []byte("Hello World!")
	err = br.SendData(d, "test", 1)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	for i := range sm.senders {
		if sm.senders[i].Len() == 0 {
			t.Fatal("all recievers must recieve data")
		}
	}
}

func TestSeveralData(t *testing.T) {
	CreateLogger()
	validator, err := validator.NewWorker(encodedecoder.NewBase64Encoder())
	if err != nil {
		t.Fatal(err)
	}
	sm := &SenderMaker{}
	br := NewSubHandler(validator, sm, encodedecoder.NewBase64Encoder())
	for i := 0; i < 10; i++ {
		br.NewPassiveSub(strconv.Itoa(i), "i")
	}
	d := []byte("Hello World!")
	for i := range d {
		err = br.SendData([]byte{d[i]}, "test", 1)
		if err != nil {
			t.Fatal(err)
		}
	}
	time.Sleep(1 * time.Second)
	for i := range sm.senders {
		if sm.senders[i].RecievedMessages() != len(d) {
			t.Fatal("all recievers must recieve the same number of messages")
		}
	}
}

func BenchmarkSendData(b *testing.B) {
	CreateLogger()
	validator, err := validator.NewWorker(encodedecoder.NewBase64Encoder())
	if err != nil {
		b.Fatal(err)
	}
	sm := &SenderMaker{}
	br := NewSubHandler(validator, sm, encodedecoder.NewBase64Encoder())
	for i := 0; i < 1000; i++ {
		br.NewPassiveSub(strconv.Itoa(i), "i")
	}
	d := []byte("Hello World!")
	for i := 0; i < b.N; i++ {
		err = br.SendData(d, "test", 1)
		if err != nil {
			b.Fatal(err)
		}
	}
}
