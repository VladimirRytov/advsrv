package requests

import (
	"context"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func (rh *Requesting) OrderGetRequest(ctx context.Context, token string, params datatransferobjects.OrderParams, id int) (datatransferobjects.OrderDTO, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return datatransferobjects.OrderDTO{}, err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return datatransferobjects.OrderDTO{}, err
	}

	order, err := rh.advRepo.OrderByID(ctx, id)
	if err != nil {
		return order, err
	}
	switch {
	case params.Nested:
		err := rh.addNetsteadAdvertisementsToOrders(ctx, &order)
		if err != nil {
			return order, err
		}
	}
	return order, err
}

func (rh *Requesting) OrdersGetRequest(ctx context.Context, token string, params datatransferobjects.OrderParams) ([]datatransferobjects.OrderDTO, error) {
	var (
		err    error
		orders []datatransferobjects.OrderDTO
	)

	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return nil, err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return nil, err
	}

	if len(params.Client) > 0 {
		orders, err = rh.advRepo.OrdersByClientName(ctx, params.Client)
	} else {
		orders, err = rh.advRepo.AllOrders(ctx)
	}
	if err != nil {
		return orders, err
	}

	if params.Nested {
		for i := range orders {
			err := rh.addNetsteadAdvertisementsToOrders(ctx, &orders[i])
			if err != nil {
				return orders, err
			}
		}
	}
	return orders, err
}

func (rh *Requesting) OrderPostRequest(ctx context.Context, token string, params datatransferobjects.OrderParams, order *datatransferobjects.OrderDTO) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanCreateRecords(&user)
	if err != nil {
		return err
	}
	order.ID = 0
	if params.Nested {
		return rh.advRepo.NewAdvertisementsOrder(ctx, order)
	}

	order.BlockAdvertisements = nil
	order.LineAdvertisements = nil
	return rh.advRepo.NewAdvertisementsOrder(ctx, order)
}

func (rh *Requesting) OrderPutRequest(ctx context.Context, token string, order *datatransferobjects.OrderDTO) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanWriteRecords(&user)
	if err != nil {
		return err
	}
	order.BlockAdvertisements = nil
	order.LineAdvertisements = nil
	return rh.advRepo.UpdateOrder(ctx, order)
}

func (rh *Requesting) OrderDeleteRequest(ctx context.Context, token string, id int) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanDeleteRecords(&user)
	if err != nil {
		return err
	}
	return rh.advRepo.RemoveOrderByID(ctx, id)
}
