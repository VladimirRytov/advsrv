package requests

import (
	"context"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func (rh *Requesting) addNetsteadOrdersToClient(ctx context.Context, client *datatransferobjects.ClientDTO) error {
	var err error
	client.Orders, err = rh.advRepo.OrdersByClientName(ctx, client.Name)
	if err != nil {
		return err
	}
	for i := range client.Orders {
		rh.addNetsteadAdvertisementsToOrders(ctx, &client.Orders[i])
	}
	return nil
}

func (rh *Requesting) addNetsteadAdvertisementsToOrders(ctx context.Context, order *datatransferobjects.OrderDTO) error {
	var err error
	order.BlockAdvertisements, err = rh.advRepo.BlockAdvertisementsByOrderID(ctx, order.ID)
	if err != nil {
		return err
	}
	order.LineAdvertisements, err = rh.advRepo.LineAdvertisementsByOrderID(ctx, order.ID)
	return err
}
