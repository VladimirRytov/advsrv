package costcalculationhandler

import (
	"context"
	"time"

	"github.com/VladimirRytov/advsrv/internal/advertisements"
	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/handlers"
	"github.com/VladimirRytov/advsrv/internal/logging"
	"github.com/VladimirRytov/advsrv/internal/mapper"
)

type CostRateCalculator struct {
	storage handlers.CostRateCalculatorRequests
}

func NewCostRateCalculator(storage handlers.CostRateCalculatorRequests) *CostRateCalculator {
	return &CostRateCalculator{storage: storage}
}

func (c *CostRateCalculator) SetAdvRepo(st handlers.CostRateCalculatorRequests) {
	c.storage = st
}

func (c *CostRateCalculator) CalculateBlockAdvertisementCost(ctx context.Context, adv datatransferobjects.BlockAdvertisementDTO,
	costRateName string) (datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("costRateCalculator.CalculateBlockAdvertisementCost: start calculating block Advertisement cost")
	cont, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()
	costRateAdv, err := c.costRateAdv(cont, costRateName)
	if err != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, err
	}

	blockAdv, err := mapper.DtoToAdvertisementBlock(&adv)
	if err != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, err
	}

	tags, err := c.collectTags(cont, adv.Tags)
	if err != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, err

	}
	charges, err := c.collectExtraCharges(cont, adv.ExtraCharges)
	if err != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, err
	}
	cost, err := costRateAdv.CalculateBlockCost(blockAdv, tags, charges)
	if err != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, err

	}
	err = blockAdv.SetCost(cost)
	if err != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, err

	}
	return mapper.BlockAdvertisementToDTO(&blockAdv), err

}

func (c *CostRateCalculator) CalculateLineAdvertisementCost(ctx context.Context, adv datatransferobjects.LineAdvertisementDTO,
	costRateName string) (datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("costRateCalculator.CalculateLineAdvertisementCost: start calculating line Advertisement cost")
	cont, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()
	costRateAdv, err := c.costRateAdv(cont, costRateName)
	if err != nil {
		return datatransferobjects.LineAdvertisementDTO{}, err
	}

	lineAdv, err := mapper.DtoToAdvertisementLine(&adv)
	if err != nil {
		return datatransferobjects.LineAdvertisementDTO{}, err
	}

	tags, err := c.collectTags(cont, adv.Tags)
	if err != nil {
		return datatransferobjects.LineAdvertisementDTO{}, err
	}
	charges, err := c.collectExtraCharges(cont, adv.ExtraCharges)
	if err != nil {
		return datatransferobjects.LineAdvertisementDTO{}, err
	}
	cost, err := costRateAdv.CalculateLineCost(lineAdv, tags, charges)
	if err != nil {
		return datatransferobjects.LineAdvertisementDTO{}, err
	}
	err = lineAdv.SetCost(cost)
	if err != nil {
		return datatransferobjects.LineAdvertisementDTO{}, err
	}
	return mapper.LineAdvertisementToDTO(&lineAdv), nil
}

func (c *CostRateCalculator) CalculateOrderCost(ctx context.Context, adv datatransferobjects.OrderDTO, costRateName string) (datatransferobjects.OrderDTO, error) {
	logging.Logger.Debug("costRateCalculator.CalculateBlockAdvertisementCost: start calculating order cost")
	cont, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()
	costRateAdv, err := c.costRateAdv(cont, costRateName)
	if err != nil {
		return datatransferobjects.OrderDTO{}, err
	}
	orderAdv, err := mapper.DtoToOrder(&adv)
	if err != nil {
		return datatransferobjects.OrderDTO{}, err

	}
	if orderAdv.OrderId() != 0 {
		adv.BlockAdvertisements, err = c.collectBlockAdvertisements(cont, orderAdv.OrderId())
		if err != nil {
			return datatransferobjects.OrderDTO{}, err
		}
		adv.LineAdvertisements, err = c.collectLineAdvertisements(cont, orderAdv.OrderId())
		if err != nil {
			return datatransferobjects.OrderDTO{}, err
		}
	}
	blocksAdv, err := c.dtoToBlockAdvertisements(adv.BlockAdvertisements)
	if err != nil {
		return datatransferobjects.OrderDTO{}, err
	}
	linesAdv, err := c.dtoToLineAdvertisements(adv.LineAdvertisements)
	if err != nil {
		return datatransferobjects.OrderDTO{}, err
	}

	cost := costRateAdv.CalculateOrderCost(orderAdv, blocksAdv, linesAdv)
	err = orderAdv.SetOrderCost(cost)
	if err != nil {
		return datatransferobjects.OrderDTO{}, err
	}
	return mapper.OrderToDTO(&orderAdv), err
}

func (c *CostRateCalculator) collectTags(ctx context.Context, tags []string) ([]advertisements.Tag, error) {
	tagArr := make([]advertisements.Tag, 0, len(tags))
	for _, v := range tags {
		tag, err := c.storage.TagByName(ctx, v)
		if err != nil {
			return tagArr, err
		}
		tagAdv, err := mapper.DtoToTag(&tag)
		if err != nil {
			return tagArr, err
		}
		tagArr = append(tagArr, tagAdv)
	}
	return tagArr, nil
}

func (c *CostRateCalculator) collectExtraCharges(ctx context.Context, extraCharges []string) ([]advertisements.ExtraCharge, error) {
	chargeArr := make([]advertisements.ExtraCharge, 0, len(extraCharges))
	for _, v := range extraCharges {
		charge, err := c.storage.ExtraChargeByName(ctx, v)
		if err != nil {
			return chargeArr, err
		}

		chargeAdv, err := mapper.DtoToExtraCharge(&charge)
		if err != nil {
			return chargeArr, err
		}
		chargeArr = append(chargeArr, chargeAdv)
	}
	return chargeArr, nil
}

func (c *CostRateCalculator) collectBlockAdvertisements(ctx context.Context, orderID int) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	blocksDto, err := c.storage.BlockAdvertisementsByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	return blocksDto, nil
}

func (c *CostRateCalculator) dtoToBlockAdvertisements(blocksDto []datatransferobjects.BlockAdvertisementDTO) ([]advertisements.AdvertisementBlock, error) {
	blocksAdv := make([]advertisements.AdvertisementBlock, 0, len(blocksDto))
	for i := range blocksDto {
		blockAdv, err := mapper.DtoToAdvertisementBlock(&blocksDto[i])
		if err != nil {
			return nil, err
		}
		blocksAdv = append(blocksAdv, blockAdv)
	}
	return blocksAdv, nil
}

func (c *CostRateCalculator) collectLineAdvertisements(ctx context.Context, orderID int) ([]datatransferobjects.LineAdvertisementDTO, error) {
	lineDto, err := c.storage.LineAdvertisementsByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	return lineDto, nil
}

func (c *CostRateCalculator) dtoToLineAdvertisements(lineDto []datatransferobjects.LineAdvertisementDTO) ([]advertisements.AdvertisementLine, error) {
	linesAdv := make([]advertisements.AdvertisementLine, 0, len(lineDto))
	for i := range lineDto {
		lineAdv, err := mapper.DtoToAdvertisementLine(&lineDto[i])
		if err != nil {
			return nil, err
		}
		linesAdv = append(linesAdv, lineAdv)
	}
	return linesAdv, nil
}

func (c *CostRateCalculator) costRateAdv(ctx context.Context, costRateName string) (advertisements.CostRate, error) {
	costRate, err := c.storage.CostRateByName(ctx, costRateName)
	if err != nil {
		return advertisements.CostRate{}, err
	}
	costRateAdv, err := mapper.DtoToCostRate(&costRate)
	if err != nil {
		return advertisements.CostRate{}, err
	}
	return costRateAdv, nil
}
