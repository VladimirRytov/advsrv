package advertisementhandler

import (
	"context"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/logging"
	"github.com/VladimirRytov/advsrv/internal/mapper"
)

func (a *AdvertisementController) ClientByName(ctx context.Context, name string) (datatransferobjects.ClientDTO, error) {
	logging.Logger.Debug("advertisementhandler: get request. Get Client by Name", "name", name)

	client, err := a.entities.ClientByName(ctx, name)
	if err != nil {
		return datatransferobjects.ClientDTO{}, err
	}
	clientModel, err := mapper.DtoToClient(&client)
	if err != nil {
		return datatransferobjects.ClientDTO{}, err
	}
	return mapper.ClientToDTO(&clientModel), nil
}

func (a *AdvertisementController) OrderByID(ctx context.Context, id int) (datatransferobjects.OrderDTO, error) {
	logging.Logger.Debug("advertisementhandler: get request. Get Order by id", "id", id)

	order, err := a.entities.OrderByID(ctx, id)
	if err != nil {
		return datatransferobjects.OrderDTO{}, err
	}

	orderModel, err := mapper.DtoToOrder(&order)
	if err != nil {
		return datatransferobjects.OrderDTO{}, err
	}

	return mapper.OrderToDTO(&orderModel), nil
}

func (a *AdvertisementController) LineAdvertisementByID(ctx context.Context, id int) (datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("advertisementhandler: get request. get LineAdvertisement by id", "LineAdvertisement", id)

	line, err := a.entities.LineAdvertisementByID(ctx, id)
	if err != nil {
		return datatransferobjects.LineAdvertisementDTO{}, err
	}

	lineModel, err := mapper.DtoToAdvertisementLine(&line)
	if err != nil {
		return datatransferobjects.LineAdvertisementDTO{}, err
	}

	return mapper.LineAdvertisementToDTO(&lineModel), nil
}

func (a *AdvertisementController) BlockAdvertisementByID(ctx context.Context, id int) (datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("advertisementhandler: get request. get BlockAdvertisement by id", "BlockAdvertisement", id)

	block, err := a.entities.BlockAdvertisementByID(ctx, id)
	if err != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, err
	}

	blockModel, err := mapper.DtoToAdvertisementBlock(&block)
	if err != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, err
	}

	return mapper.BlockAdvertisementToDTO(&blockModel), nil
}

func (a *AdvertisementController) TagByName(ctx context.Context, name string) (datatransferobjects.TagDTO, error) {
	logging.Logger.Debug("advertisementhandler: get request. get Tag by name", "Tag", name)

	tag, err := a.entities.TagByName(ctx, name)
	if err != nil {
		return datatransferobjects.TagDTO{}, err
	}

	tagModel, err := mapper.DtoToTag(&tag)
	if err != nil {
		return datatransferobjects.TagDTO{}, err
	}
	return mapper.TagToDTO(&tagModel), nil
}

func (a *AdvertisementController) ExtraChargeByName(ctx context.Context, name string) (datatransferobjects.ExtraChargeDTO, error) {
	logging.Logger.Debug("advertisementhandler: get request. get ExtraCharge by name", "ExtraCharge", name)

	extraCharge, err := a.entities.ExtraChargeByName(ctx, name)
	if err != nil {
		return datatransferobjects.ExtraChargeDTO{}, err
	}

	extraChargeModel, err := mapper.DtoToExtraCharge(&extraCharge)
	if err != nil {
		return datatransferobjects.ExtraChargeDTO{}, err
	}

	return mapper.ExtraChargeToDTO(&extraChargeModel), nil
}

func (a *AdvertisementController) CostRateByName(ctx context.Context, name string) (datatransferobjects.CostRateDTO, error) {
	logging.Logger.Debug("advertisementhandler.CostRateByName: get request. get CostRate by name", "CostRate", name)

	costRate, err := a.entities.CostRateByName(ctx, name)
	if err != nil {
		return datatransferobjects.CostRateDTO{}, err
	}

	costRateModel, err := mapper.DtoToCostRate(&costRate)
	if err != nil {
		return datatransferobjects.CostRateDTO{}, err
	}
	return mapper.CostRateToDTO(&costRateModel), nil
}
