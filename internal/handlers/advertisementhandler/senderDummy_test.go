package advertisementhandler

import (
	"errors"
	"io"
	"time"
)

type SenderDummy struct {
	addr    string
	buf     [][]byte
	latency int
	sendErr bool
}

type SenderMaker struct{}

func (sm *SenderMaker) NewSender(URL string) (io.Writer, error) {
	return &SenderDummy{}, nil
}

func (sd *SenderDummy) Write(data []byte) (int, error) {
	time.Sleep(time.Duration(sd.latency) * time.Second)
	sd.buf = append(sd.buf, data)
	if sd.sendErr {
		return len(data), errors.New("err")
	}
	return len(data), nil
}

func (sd *SenderDummy) Close() error {
	return nil
}
func (sd *SenderDummy) SendTo() string {
	return sd.addr
}
func (sd *SenderDummy) Len() int {
	var l int
	for i := range sd.buf {
		l += len(sd.buf[i])
	}
	return l
}

func (sd *SenderDummy) RecievedMessages() int {
	return len(sd.buf)
}
