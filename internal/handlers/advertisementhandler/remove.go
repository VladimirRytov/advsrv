package advertisementhandler

import (
	"context"
	"encoding/binary"

	"github.com/VladimirRytov/advsrv/internal/logging"
)

func (a *AdvertisementController) RemoveClientByName(ctx context.Context, name string) error {
	logging.Logger.Debug("advertisementhandler: remove request. Removing Client by name", "Client", name)

	err := a.entities.RemoveClientByName(ctx, name)
	if err != nil {
		return err
	}
	a.broadCaster.SendData([]byte(name), ClientType, DeleteAction)
	return nil
}

func (a *AdvertisementController) RemoveOrderByID(ctx context.Context, id int) error {
	logging.Logger.Debug("advertisementhandler: remove request. Removing Order by id", "Order", id)
	err := a.entities.RemoveOrderByID(ctx, id)
	if err != nil {
		return err
	}
	b := make([]byte, binary.MaxVarintLen64)
	binary.BigEndian.PutUint64(b, uint64(id))
	a.broadCaster.SendData(b, OrderType, DeleteAction)

	return nil
}

func (a *AdvertisementController) RemoveLineAdvertisementByID(ctx context.Context, id int) error {
	logging.Logger.Debug("advertisementhandler: remove request. Removing LineAdvertisement by id", "LineAdvertisement", id)
	err := a.entities.RemoveLineAdvertisementByID(ctx, id)
	if err != nil {
		return err
	}

	b := make([]byte, binary.MaxVarintLen64)
	binary.BigEndian.PutUint64(b, uint64(id))
	a.broadCaster.SendData(b, LineAdvertisementType, DeleteAction)
	return nil
}

func (a *AdvertisementController) RemoveBlockAdvertisementByID(ctx context.Context, id int) error {
	logging.Logger.Debug("advertisementhandler: remove request. Removing BlockAdvertisement by id", "BlockAdvertisement", id)

	err := a.entities.RemoveBlockAdvertisementByID(ctx, id)
	if err != nil {
		return err
	}

	b := make([]byte, binary.MaxVarintLen64)
	binary.BigEndian.PutUint64(b, uint64(id))
	a.broadCaster.SendData(b, BlockAdvertisementType, DeleteAction)
	return nil
}

func (a *AdvertisementController) RemoveTagByName(ctx context.Context, name string) error {
	logging.Logger.Debug("advertisementhandler: remove request. Removing Tag by name", "Tag", name)

	err := a.entities.RemoveTagByName(ctx, name)
	if err != nil {
		return err
	}

	a.broadCaster.SendData([]byte(name), TagType, DeleteAction)
	return nil
}

func (a *AdvertisementController) RemoveExtraChargeByName(ctx context.Context, name string) error {
	logging.Logger.Debug("advertisementhandler: remove request. Removing ExtraCharge by name", "ExtraCharge", name)

	err := a.entities.RemoveExtraChargeByName(ctx, name)
	if err != nil {
		return err
	}

	a.broadCaster.SendData([]byte(name), ExtraChargeType, DeleteAction)
	return nil
}

func (a *AdvertisementController) RemoveCostRateByName(ctx context.Context, name string) error {
	logging.Logger.Debug("advertisementhandler: remove request. Removing CostRate by name", "CostRate", name)

	err := a.entities.RemoveCostRateByName(ctx, name)
	if err != nil {
		return err
	}

	a.broadCaster.SendData([]byte(name), CostRateType, DeleteAction)
	return nil
}
