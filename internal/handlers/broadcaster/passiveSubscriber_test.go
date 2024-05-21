package broadcaster

import (
	"slices"
	"testing"
	"time"
)

func TestNewSubscriber(t *testing.T) {
	CreateLogger()
	sd := &SenderDummy{}
	sub := NewpassiveSubscriber(10, 1, "a", sd)
	sub.AppendData([]byte("Hello"))
	time.Sleep(1 * time.Second)
	if !slices.Equal(sd.buf[0], []byte("Hello")) {
		t.Fatalf("Data must be equal.\ngot %v\nwant: %v", sd.buf[0], []byte("Hello"))
	}
}

func TestTimeout(t *testing.T) {
	CreateLogger()
	sd := &SenderDummy{}
	sd.sendErr = true
	sub := NewpassiveSubscriber(10, 1, "a", sd)
	testCase := []byte("Hello world")
	for i := 0; i < 5; i++ {
		sub.AppendData(testCase)
	}
	time.Sleep(5 * time.Second)
	if !sub.Closed() {
		t.Fatal("subscriber must be closed")
	}
}

func TestRunningCancel(t *testing.T) {
	sd := &SenderDummy{}
	sd.latency = 5
	sub := NewpassiveSubscriber(10, 1, "a", sd)
	testCase := [][]byte{[]byte("Hello world"), []byte("!")}
	for i := range testCase {
		sub.AppendData(testCase[i])
	}
	sub.Close()
	time.Sleep(3 * time.Second)
	if !sub.Closed() {
		t.Fatal("subscriber must be closed")
	}
}

func TestEstablished(t *testing.T) {
	CreateLogger()
	sd := &SenderDummy{}
	sd.sendErr = true
	sub := NewpassiveSubscriber(10, 3, "a", sd)
	testCase := []byte("Hello world")
	sub.AppendData(testCase)
	time.Sleep(1 * time.Second)
	sd.sendErr = false
	time.Sleep(1 * time.Second)
	if sub.Closed() {
		t.Fatal("subscriber must be open")
	}
}

func TestOverflowBuffer(t *testing.T) {
	CreateLogger()
	sd := &SenderDummy{}
	sub := NewpassiveSubscriber(1, 2, "a", sd)
	testCase := [][]byte{[]byte("Hello"), []byte("World"), []byte("!")}
	for i := range testCase {
		sub.AppendData(testCase[i])
	}
	time.Sleep(1 * time.Second)
	if !sub.Closed() {
		t.Fatal("subscriber must be closed")
	}
}
