package broadcaster

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

type Broadcaster struct {
	privateKey  *rsa.PrivateKey
	b64         B64Enc
	sender      SendMaker
	validator   Validator
	buf         chan []byte
	subscribers map[string]Subscriber
}

func NewSubHandler(val Validator, sm SendMaker, b64 B64Enc) *Broadcaster {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	br := &Broadcaster{
		privateKey:  key,
		validator:   val,
		sender:      sm,
		b64:         b64,
		buf:         make(chan []byte, 100),
		subscribers: make(map[string]Subscriber),
	}
	go br.spray()
	return br
}

func (sh *Broadcaster) spray() {
	for data := range sh.buf {
		for k := range sh.subscribers {
			if sh.subscribers[k].Closed() {
				delete(sh.subscribers, k)
				continue
			}
			sh.subscribers[k].AppendData(data)
		}
	}
}

func (sh *Broadcaster) NewPassiveSub(addr, name string) (string, error) {
	if v, ok := sh.subscribers[addr+name]; ok && !v.Closed() {
		return "", ErrSubscriberExist
	}
	s, err := sh.sender.NewSender(addr)
	if err != nil {
		return "", err
	}
	hashed := sha256.Sum256([]byte(addr + name))
	strHash := hex.EncodeToString(hashed[:])
	sh.subscribers[strHash] = NewpassiveSubscriber(100, 5, addr, s)
	return strHash, nil
}

func (sh *Broadcaster) NewActiveSub(ctx context.Context, addr string) (<-chan []byte, error) {
	if v, ok := sh.subscribers[addr]; ok && !v.Closed() {
		return nil, ErrSubscriberExist
	}
	hashed := sha256.Sum256([]byte(addr))
	strHash := hex.EncodeToString(hashed[:])
	sh.subscribers[strHash] = NewActiveSubscriber(ctx, 100, addr)
	return sh.subscribers[strHash].Buffer(), nil
}

func (sh *Broadcaster) RemoveSub(addr string) {
	if _, ok := sh.subscribers[addr]; !ok {
		return
	}
	sh.subscribers[addr].Close()
	delete(sh.subscribers, addr)
}

func (sh *Broadcaster) ListSubs() (int, []datatransferobjects.SubscribeParams) {
	subs := make([]datatransferobjects.SubscribeParams, 0, len(sh.subscribers))
	for k, v := range sh.subscribers {
		subs = append(subs, datatransferobjects.SubscribeParams{
			UserID: k,
			URL:    v.Adress(),
		})
	}
	return len(sh.subscribers), subs
}

func (sh *Broadcaster) Ping(addr string) error {
	if _, ok := sh.subscribers[addr]; !ok {
		return ErrSubscriberNotExist
	}
	p := SendWrapper{
		Type:   "Check",
		Action: 999,
		Entry:  []byte("Pong"),
	}
	raw, err := json.Marshal(p)
	if err != nil {
		return err
	}
	sh.subscribers[addr].AppendData(raw)
	return nil
}

func (sh *Broadcaster) SubInfo(id string) (datatransferobjects.SubscribeParams, error) {
	if v, ok := sh.subscribers[id]; ok {
		return datatransferobjects.SubscribeParams{
			UserID: id,
			URL:    v.Adress(),
		}, nil
	}
	return datatransferobjects.SubscribeParams{}, ErrSubscriberNotExist
}
