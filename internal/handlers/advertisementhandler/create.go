package advertisementhandler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/logging"
	"github.com/VladimirRytov/advsrv/internal/mapper"
)

func (a *AdvertisementController) NewClient(ctx context.Context, client *datatransferobjects.ClientDTO) error {
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending client DTO to database", "Client", client)
	clientAdv, err := mapper.DtoToClient(client)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: create request. Client is not created", "err", err)
		return err
	}

	*client = mapper.ClientToDTO(&clientAdv)

	_, err = a.entities.NewClient(ctx, client)
	if err != nil {
		logging.Logger.Error("advertisementhandler: create request. Client is not created")
		return err
	}
	raw, _ := json.Marshal(client)
	a.broadCaster.SendData(raw, ClientType, CreateAction)
	return nil
}

func (a *AdvertisementController) NewAdvertisementsOrder(ctx context.Context, order *datatransferobjects.OrderDTO) error {
	var (
		bloks []datatransferobjects.BlockAdvertisementDTO = make([]datatransferobjects.BlockAdvertisementDTO, 0, len(order.BlockAdvertisements))
		lines []datatransferobjects.LineAdvertisementDTO  = make([]datatransferobjects.LineAdvertisementDTO, 0, len(order.LineAdvertisements))
	)
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending Order DTO to database", "Order", order)
	orderAdv, err := mapper.DtoToOrder(order)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: create request. Order is not created", "err", err)
		return err
	}

	for i := range order.LineAdvertisements {
		line, err := mapper.DtoToAdvertisementLine(&order.LineAdvertisements[i])
		if err != nil {
			logging.Logger.Warn("advertisementhandler.NewAdvertisementsOrder: create request. LineAdvertisement is not created", "err", err)
			return errors.Join(err, fmt.Errorf("порядковый номер строчного объявления: %d", i+1))
		}
		lines = append(lines, mapper.LineAdvertisementToDTO(&line))
	}
	for i := range order.BlockAdvertisements {
		block, err := mapper.DtoToAdvertisementBlock(&order.BlockAdvertisements[i])
		if err != nil {
			logging.Logger.Warn("advertisementhandler.NewAdvertisementsOrder: create request. Block is not created", "err", err)
			return errors.Join(err, fmt.Errorf("порядковый номер блочного объявления: %d", i+1))
		}
		bloks = append(bloks, mapper.BlockAdvertisementToDTO(&block))
	}
	*order = mapper.OrderToDTO(&orderAdv)
	order.BlockAdvertisements = bloks
	order.LineAdvertisements = lines

	logging.Logger.Debug("advertisementhandler: create request. Checking and sending Order DTO to database", "Order", order)
	newOrder, err := a.entities.NewAdvertisementsOrder(ctx, order)
	if err != nil {
		logging.Logger.Error("advertisementhandler: create request. Order is not created")
		return err
	}

	*order = newOrder

	raw, _ := json.Marshal(order)
	a.broadCaster.SendData(raw, OrderType, CreateAction)
	return nil
}

func (a *AdvertisementController) NewLineAdvertisement(ctx context.Context, line *datatransferobjects.LineAdvertisementDTO) error {
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending LineAdvertisement DTO to database", "LineAdvertisement", line)
	lineAdv, err := mapper.DtoToAdvertisementLine(line)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: create request. LineAdvertisement is not created", "err", err)
		return err
	}
	*line = mapper.LineAdvertisementToDTO(&lineAdv)

	id, err := a.entities.NewLineAdvertisement(ctx, line)
	if err != nil {
		logging.Logger.Error("advertisementhandler: create request. LineAdvertisement is not created")
		return err
	}
	line.ID = id

	raw, _ := json.Marshal(line)
	a.broadCaster.SendData(raw, LineAdvertisementType, CreateAction)
	return nil
}

func (a *AdvertisementController) NewBlockAdvertisement(ctx context.Context, block *datatransferobjects.BlockAdvertisementDTO) error {
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending BlockAdvertisement DTO to database", "LineAdvertisement", block)
	blockAdv, err := mapper.DtoToAdvertisementBlock(block)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: create request. BlockAdvertisement is not created", "err", err)
		return err
	}
	*block = mapper.BlockAdvertisementToDTO(&blockAdv)

	id, err := a.entities.NewBlockAdvertisement(ctx, block)
	if err != nil {
		logging.Logger.Error("advertisementhandler: create request. BlockAdvertisement is not created")
		return err
	}
	block.ID = id

	raw, _ := json.Marshal(block)
	a.broadCaster.SendData(raw, BlockAdvertisementType, CreateAction)
	return nil
}

func (a *AdvertisementController) NewExtraCharge(ctx context.Context, extraCharge *datatransferobjects.ExtraChargeDTO) error {
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending Tag DTO to database", "ExtraCharge", extraCharge)

	chargeAdv, err := mapper.DtoToExtraCharge(extraCharge)
	if err != nil {
		logging.Logger.Error("newExtraCharge: create request. ExtraCharge is not created", "err", err)
		return err
	}
	*extraCharge = mapper.ExtraChargeToDTO(&chargeAdv)

	_, err = a.entities.NewExtraCharge(ctx, extraCharge)
	if err != nil {
		logging.Logger.Error("advertisementhandler: create request. ExtraCharge is not created")
		return err
	}

	raw, _ := json.Marshal(extraCharge)
	a.broadCaster.SendData(raw, ExtraChargeType, CreateAction)
	return nil
}

func (a *AdvertisementController) NewTag(ctx context.Context, tag *datatransferobjects.TagDTO) error {
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending ExtraCharge DTO to database", "Tag", tag)
	tagAdv, err := mapper.DtoToTag(tag)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: create request. ExtraCharge is not created", "err", err)
		return err
	}
	*tag = mapper.TagToDTO(&tagAdv)

	_, err = a.entities.NewTag(ctx, tag)
	if err != nil {
		logging.Logger.Error("advertisementhandler: create request. Tag is not created")
		return err
	}

	raw, _ := json.Marshal(tag)
	a.broadCaster.SendData(raw, TagType, CreateAction)
	return nil
}

func (a *AdvertisementController) NewCostRate(ctx context.Context, costRate *datatransferobjects.CostRateDTO) error {
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending CostRate DTO to database", "CostRate", costRate)
	costRateAdv, err := mapper.DtoToCostRate(costRate)
	if err != nil {
		logging.Logger.Warn("advertisementhandler.NewCostRate: create request. CostRate is not created", "err", err)
		return err
	}
	*costRate = mapper.CostRateToDTO(&costRateAdv)

	_, err = a.entities.NewCostRate(ctx, costRate)
	if err != nil {
		logging.Logger.Error("advertisementhandler.NewCostRate: create request. CostRate is not created", "error", err)
		return err
	}

	raw, _ := json.Marshal(costRate)
	a.broadCaster.SendData(raw, CostRateType, CreateAction)
	return nil
}
