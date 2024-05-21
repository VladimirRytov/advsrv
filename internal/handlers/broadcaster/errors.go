package broadcaster

import "errors"

var (
	ErrSubscriberExist    = errSubscriberExist()
	ErrSubscriberNotExist = errSubscriberNotExist()
)

func errSubscriberExist() error {
	return errors.New("подписка уже существует")
}

func errSubscriberNotExist() error {
	return errors.New("подписчика не существует")
}

func (sh *Broadcaster) ErrSubscriberExist() error {
	return ErrSubscriberExist
}

func (sh *Broadcaster) ErrSubscriberNotExist() error {
	return ErrSubscriberNotExist
}
