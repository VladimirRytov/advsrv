package broadcaster

import (
	"io"
	"time"
)

type PassiveSubscriber struct {
	sendingData bool
	closed      bool

	close             chan struct{}
	buffer            chan []byte
	buffcap           int
	out               io.Writer
	timeoutMultiplier int
	addr              string
}

func NewpassiveSubscriber(buffLen, timeoutMult int, addr string, sender io.Writer) *PassiveSubscriber {
	sub := &PassiveSubscriber{
		buffer:            make(chan []byte, buffLen),
		close:             make(chan struct{}),
		buffcap:           buffLen,
		out:               sender,
		timeoutMultiplier: timeoutMult,
		addr:              addr,
	}
	sub.startSendingData()
	return sub
}

func (sub *PassiveSubscriber) Buffer() chan []byte {
	return sub.buffer
}

func (sub *PassiveSubscriber) AppendData(data []byte) {
	if sub.closed {
		return
	}
	if len(sub.buffer) == sub.buffcap {
		sub.Close()
		return
	}
	sub.buffer <- data
}

func (sub *PassiveSubscriber) Adress() string {
	return sub.addr
}

func (sub *PassiveSubscriber) Close() {
	if !sub.closed {
		sub.closed = true
		sub.close <- struct{}{}
	}
}

func (sub *PassiveSubscriber) Closed() bool {
	return sub.closed
}

func (sub *PassiveSubscriber) startSendingData() {
	if sub.sendingData {
		return
	}
	sub.sendingData = true
	go func() {
		for {
			select {
			case v := <-sub.buffer:
				if err := sub.sendData(v); err != nil {
					sub.Close()
				}
			case <-sub.close:
				close(sub.buffer)
				for range sub.buffer {
					if len(sub.buffer) == 0 {
						return
					}
				}
			}
		}
	}()
}

func (sub *PassiveSubscriber) sendData(data []byte) error {
	_, err := sub.out.Write(data)
	if err != nil {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Duration(sub.timeoutMultiplier) * time.Second)
			_, err := sub.out.Write(data)
			if err == nil {
				return err
			}
		}
	}
	return err
}
