package advertisementhandler

import (
	"context"
	"encoding/json"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/logging"
	"github.com/VladimirRytov/advsrv/internal/mapper"
)

func (a *AdvertisementController) UpdateClient(ctx context.Context, client *datatransferobjects.ClientDTO) error {
	logging.Logger.Debug("advertisementhandler: update request. Updating Client", "Client", client)
	clientAdv, err := mapper.DtoToClient(client)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: update request. Client entity is not created", "err", err)
		return err
	}
	*client = mapper.ClientToDTO(&clientAdv)

	err = a.entities.UpdateClient(ctx, client)
	if err != nil {
		return err
	}

	raw, _ := json.Marshal(client)
	a.broadCaster.SendData(raw, ClientType, CreateAction)
	return nil
}

func (a *AdvertisementController) UpdateOrder(ctx context.Context, order *datatransferobjects.OrderDTO) error {
	logging.Logger.Debug("advertisementhandler: update request. Updating Order", "Order", order)
	orderAdv, err := mapper.DtoToOrder(order)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: update request. Order entity is not created", "err", err)
		return err
	}
	*order = mapper.OrderToDTO(&orderAdv)

	err = a.entities.UpdateOrder(ctx, order)
	if err != nil {
		return err
	}

	raw, _ := json.Marshal(order)
	a.broadCaster.SendData(raw, OrderType, CreateAction)
	return nil
}

func (a *AdvertisementController) UpdateLineAdvertisement(ctx context.Context, line *datatransferobjects.LineAdvertisementDTO) error {
	logging.Logger.Debug("advertisementhandler: update request. Updating LineAdvertisement", "LineAdvertisement", line)
	lineAdv, err := mapper.DtoToAdvertisementLine(line)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: update request. LineAdvertisement entity is not created", "err", err)
		return err
	}
	*line = mapper.LineAdvertisementToDTO(&lineAdv)

	err = a.entities.UpdateLineAdvertisement(ctx, line)
	if err != nil {
		return err
	}

	raw, _ := json.Marshal(line)
	a.broadCaster.SendData(raw, LineAdvertisementType, CreateAction)
	return nil
}

func (a *AdvertisementController) UpdateBlockAdvertisement(ctx context.Context, block *datatransferobjects.BlockAdvertisementDTO) error {
	logging.Logger.Debug("advertisementhandler: update request. Updating BlockAdvertisement", "BlockAdvertisement", block)
	blockAdv, err := mapper.DtoToAdvertisementBlock(block)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: update request. BlockAdvertisement entity is not created", "err", err)
		return err
	}
	*block = mapper.BlockAdvertisementToDTO(&blockAdv)

	err = a.entities.UpdateBlockAdvertisement(ctx, block)
	if err != nil {
		return err
	}

	raw, _ := json.Marshal(block)
	a.broadCaster.SendData(raw, BlockAdvertisementType, CreateAction)
	return nil
}

func (a *AdvertisementController) UpdateExtraCharge(ctx context.Context, extraCharge *datatransferobjects.ExtraChargeDTO) error {
	logging.Logger.Debug("advertisementhandler: update request. Updating ExtraCharge", "ExtraCharge", extraCharge)
	chargeAdv, err := mapper.DtoToExtraCharge(extraCharge)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: update request. ExtraCharge entity is not created", "err", err)
		return err
	}
	*extraCharge = mapper.ExtraChargeToDTO(&chargeAdv)

	err = a.entities.UpdateExtraCharge(ctx, extraCharge)
	if err != nil {
		return err
	}

	raw, _ := json.Marshal(extraCharge)
	a.broadCaster.SendData(raw, ExtraChargeType, CreateAction)
	return nil
}

func (a *AdvertisementController) UpdateTag(ctx context.Context, tag *datatransferobjects.TagDTO) error {
	logging.Logger.Debug("advertisementhandler: update request. Updating Tag", "Tag", tag)
	tagAdv, err := mapper.DtoToTag(tag)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: update request. Tag entity is not created", "err", err)
		return err
	}
	*tag = mapper.TagToDTO(&tagAdv)

	err = a.entities.UpdateTag(ctx, tag)
	if err != nil {
		return err
	}

	raw, _ := json.Marshal(tag)
	a.broadCaster.SendData(raw, TagType, CreateAction)
	return nil
}

func (a *AdvertisementController) UpdateCostRate(ctx context.Context, costRate *datatransferobjects.CostRateDTO) error {
	logging.Logger.Debug("advertisementhandler: update request. Updating CostRate", "name", costRate)
	costRateAdv, err := mapper.DtoToCostRate(costRate)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: update request. CostRate entity is not created", "err", err)
		return err
	}
	*costRate = mapper.CostRateToDTO(&costRateAdv)

	err = a.entities.UpdateCostRate(ctx, costRate)
	if err != nil {
		return err
	}

	raw, _ := json.Marshal(costRate)
	a.broadCaster.SendData(raw, CostRateType, CreateAction)
	return nil
}
