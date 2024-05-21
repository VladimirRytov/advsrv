package requests

import (
	"context"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func (rh *Requesting) ClientsGetRequest(ctx context.Context, token string, params datatransferobjects.ClientParams) ([]datatransferobjects.ClientDTO, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return nil, err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return nil, err
	}

	clients, err := rh.advRepo.AllClients(ctx)
	if err != nil {
		return nil, err
	}
	if params.Nested {
		for i := range clients {
			err := rh.addNetsteadOrdersToClient(ctx, &clients[i])
			if err != nil {
				return nil, err
			}
		}
	}
	return clients, nil
}

func (rh *Requesting) ClientGetRequest(ctx context.Context, token string, params datatransferobjects.ClientParams, name string) (datatransferobjects.ClientDTO, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return datatransferobjects.ClientDTO{}, err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return datatransferobjects.ClientDTO{}, err
	}

	client, err := rh.advRepo.ClientByName(ctx, name)
	if err != nil {
		return datatransferobjects.ClientDTO{}, err
	}
	if params.Nested {
		err := rh.addNetsteadOrdersToClient(ctx, &client)
		if err != nil {
			return datatransferobjects.ClientDTO{}, err
		}
	}
	return client, err
}

func (rh *Requesting) ClientPostRequest(ctx context.Context, token string, params datatransferobjects.ClientParams, client *datatransferobjects.ClientDTO) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanCreateRecords(&user)
	if err != nil {
		return err
	}

	err = rh.advRepo.NewClient(ctx, client)
	if err != nil {
		return err
	}
	if params.Nested {
		for i := range client.Orders {
			client.Orders[i].ClientName = client.Name
			client.Orders[i].ID = 0
			err := rh.advRepo.NewAdvertisementsOrder(ctx, &client.Orders[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (rh *Requesting) ClientPutRequest(ctx context.Context, token string, client *datatransferobjects.ClientDTO) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanWriteRecords(&user)
	if err != nil {
		return err
	}
	return rh.advRepo.UpdateClient(ctx, client)
}

func (rh *Requesting) ClientDeleteRequest(ctx context.Context, token string, name string) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanDeleteRecords(&user)
	if err != nil {
		return err
	}
	return rh.advRepo.RemoveClientByName(ctx, name)
}
