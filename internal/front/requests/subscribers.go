package requests

import (
	"context"
	"errors"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func (rh *Requesting) SubscribeGetRequest(ctx context.Context, token string, id string, ping bool) (datatransferobjects.SubscribeParams, error) {
	tok, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return datatransferobjects.SubscribeParams{}, err
	}
	err = rh.authorizator.CanReadRecords(&tok)
	if err != nil {
		return datatransferobjects.SubscribeParams{}, err
	}
	if ping {
		err = rh.broadCaster.Ping(id)
		if err != nil {
			if errors.Is(rh.broadCaster.ErrSubscriberNotExist(), err) {
				return datatransferobjects.SubscribeParams{}, ErrNotFound
			}
			return datatransferobjects.SubscribeParams{}, err
		}
	}
	sub, err := rh.broadCaster.SubInfo(id)
	if err != nil {
		if errors.Is(rh.broadCaster.ErrSubscriberNotExist(), err) {
			return datatransferobjects.SubscribeParams{}, ErrNotFound
		}
	}
	return sub, err
}

func (rh *Requesting) SubscribersGetRequest(ctx context.Context, token string) ([]datatransferobjects.SubscribeParams, error) {
	tok, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return nil, err
	}
	err = rh.authorizator.CanReadUsers(&tok)
	if err != nil {
		return nil, err
	}
	_, subs := rh.broadCaster.ListSubs()
	return subs, err
}

func (rh *Requesting) SubscribePostRequest(ctx context.Context, token string, params datatransferobjects.SubscribeParams) (string, error) {
	tok, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return "", err
	}
	err = rh.authorizator.CanReadRecords(&tok)
	if err != nil {
		return "", err
	}
	id, err := rh.broadCaster.NewPassiveSub(params.URL, tok.Log)
	if err != nil {
		return "", err
	}
	return id, err
}

func (rh *Requesting) SubscribeActiveRequest(ctx context.Context, token string, addr string) (<-chan []byte, error) {
	tok, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return nil, err
	}
	err = rh.authorizator.CanReadRecords(&tok)
	if err != nil {
		return nil, err
	}
	return rh.broadCaster.NewActiveSub(ctx, addr)
}

func (rh *Requesting) SubscribeDeleteRequest(ctx context.Context, token string, name string) error {
	tok, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	if tok.Log != name {
		return ErrClientSide
	}
	rh.broadCaster.RemoveSub(name)
	return nil
}
