package broadcaster

import (
	"context"
)

type ActiveSubscriber struct {
	sendingData bool
	closed      bool

	close   chan struct{}
	cancel  context.Context
	buffer  chan []byte
	buffcap int
	addr    string
}

func NewActiveSubscriber(ctx context.Context, buffLen int, addr string) *ActiveSubscriber {
	sub := &ActiveSubscriber{
		buffer:  make(chan []byte, buffLen),
		close:   make(chan struct{}),
		buffcap: buffLen,
		addr:    addr,
		cancel:  ctx,
	}
	sub.startSendingData()
	return sub
}
func (sub *ActiveSubscriber) AppendData(data []byte) {
	if sub.closed {
		return
	}
	if len(sub.buffer) == sub.buffcap {
		sub.Close()
		return
	}
	sub.buffer <- data
}

func (sub *ActiveSubscriber) Adress() string {
	return sub.addr
}

func (sub *ActiveSubscriber) Buffer() chan []byte {
	return sub.buffer
}

func (sub *ActiveSubscriber) Close() {
	if !sub.closed {
		sub.closed = true
		sub.close <- struct{}{}
	}
}

func (sub *ActiveSubscriber) Closed() bool {
	return sub.closed
}

func (sub *ActiveSubscriber) startSendingData() {
	if sub.sendingData {
		return
	}
	sub.sendingData = true
	go func() {
		<-sub.cancel.Done()
		sub.closed = true
		close(sub.buffer)
		for range sub.buffer {
			if len(sub.buffer) == 0 {
				return
			}
		}
	}()
}
