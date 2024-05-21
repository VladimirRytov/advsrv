package advertisementhandler

import (
	"context"
	"time"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/logging"
	"github.com/VladimirRytov/advsrv/internal/mapper"
)

func (a *AdvertisementController) AllClients(ctx context.Context) ([]datatransferobjects.ClientDTO, error) {
	logging.Logger.Debug("advertisementhandler: search request. Search All Clients")

	recievedClientsDto, err := a.entities.AllClients(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllClients: cannot get clients", "error", err)
		return nil, err
	}
	clients := make([]datatransferobjects.ClientDTO, 0, len(recievedClientsDto))
	for i := range recievedClientsDto {
		clientModel, err := mapper.DtoToClient(&recievedClientsDto[i])
		if err != nil {
			logging.Logger.Error("client dont pass the filter", "clientName", recievedClientsDto[i].Name, "error", err)
			continue
		}
		clients = append(clients, mapper.ClientToDTO(&clientModel))
	}
	return clients, nil
}

func (a *AdvertisementController) AllOrders(ctx context.Context) ([]datatransferobjects.OrderDTO, error) {
	logging.Logger.Debug("advertisementhandler: search request. Search All AllOrders")

	orders, err := a.entities.AllOrders(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllOrders: cannot get orders", "error", err)
		return nil, err
	}
	convertedOrders := make([]datatransferobjects.OrderDTO, 0, len(orders))
	for i := range orders {
		orderModel, err := mapper.DtoToOrder(&orders[i])
		if err != nil {
			logging.Logger.Error("order dont pass the filter", "orderID", orders[i].ID, "error", err)
			continue
		}
		convertedOrders = append(convertedOrders, mapper.OrderToDTO(&orderModel))
	}
	return convertedOrders, nil
}

func (a *AdvertisementController) AllBlockAdvertisements(ctx context.Context) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("advertisementhandler: search request. Search All AllBlockAdvertisements")

	blockAdvs, err := a.entities.AllBlockAdvertisements(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllBlockAdvertisements: cannot get all BlockAdvertisements", "error", err)
		return nil, err
	}
	convertedBlocks := make([]datatransferobjects.BlockAdvertisementDTO, 0, len(blockAdvs))
	for i := range blockAdvs {
		blockModel, err := mapper.DtoToAdvertisementBlock(&blockAdvs[i])
		if err != nil {
			logging.Logger.Error("blockAdvertisement dont pass the filter", "id", blockAdvs[i].ID, "error", err)
			continue
		}
		convertedBlocks = append(convertedBlocks, mapper.BlockAdvertisementToDTO(&blockModel))
	}
	return convertedBlocks, nil
}

func (a *AdvertisementController) AllLineAdvertisements(ctx context.Context) ([]datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("advertisementhandler: search request. Search All AllLineAdvertisements")

	lineAdvs, err := a.entities.AllLineAdvertisements(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllLineAdvertisements: cannot get all LineAdvertisements", "error", err)
		return nil, err
	}

	convertedLines := make([]datatransferobjects.LineAdvertisementDTO, 0, len(lineAdvs))
	for i := range lineAdvs {
		lineModel, err := mapper.DtoToAdvertisementLine(&lineAdvs[i])
		if err != nil {
			logging.Logger.Error("lineAdvertisement dont pass the filter", "id", lineAdvs[i].ID, "error", err)
			continue
		}
		convertedLines = append(convertedLines, mapper.LineAdvertisementToDTO(&lineModel))
	}
	return convertedLines, nil
}

func (a *AdvertisementController) OrdersByClientName(ctx context.Context, name string) ([]datatransferobjects.OrderDTO, error) {
	logging.Logger.Debug("advertisementhandler: search request. Search Order by Client name")

	orders, err := a.entities.OrdersByClientName(ctx, name)
	if err != nil {
		logging.Logger.Error("advertisementhandler.OrdersByClientName: cannot get OrdersByClientName", "error", err)
		return nil, err
	}
	convertedOrders := make([]datatransferobjects.OrderDTO, 0, len(orders))
	for i := range orders {
		orderModel, err := mapper.DtoToOrder(&orders[i])
		if err != nil {
			logging.Logger.Error("order dont pass the filter", "orderID", orders[i].ID, "error", err)
			continue
		}
		convertedOrders = append(convertedOrders, mapper.OrderToDTO(&orderModel))
	}
	return convertedOrders, nil
}

func (a *AdvertisementController) BlockAdvertisementsByOrderID(ctx context.Context, id int) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("advertisementhandler: search request. Search BlockAdvertisements by Order id", "id", id)

	blockAdvs, err := a.entities.BlockAdvertisementsByOrderID(ctx, id)
	if err != nil {
		logging.Logger.Error("advertisementhandler.BlockAdvertisementsByOrderID: cannot get BlockAdvertisementsByOrderID", "error", err)
		return nil, err
	}
	convertedBlocks := make([]datatransferobjects.BlockAdvertisementDTO, 0, len(blockAdvs))

	for i := range blockAdvs {
		blockModel, err := mapper.DtoToAdvertisementBlock(&blockAdvs[i])
		if err != nil {
			logging.Logger.Error("blockAdvertisement dont pass the filter", "id", blockAdvs[i].ID, "error", err)
			continue
		}
		convertedBlocks = append(convertedBlocks, mapper.BlockAdvertisementToDTO(&blockModel))
	}
	return convertedBlocks, nil
}

func (a *AdvertisementController) LineAdvertisementsByOrderID(ctx context.Context, id int) ([]datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("advertisementhandler: search request. Search LineAdvertisements by Order id", "id", id)

	lineAdvs, err := a.entities.LineAdvertisementsByOrderID(ctx, id)
	if err != nil {
		logging.Logger.Error("advertisementhandler.LineAdvertisementsByOrderID: cannot get LineAdvertisementsByOrderID", "error", err)
		return nil, err
	}
	convertedLines := make([]datatransferobjects.LineAdvertisementDTO, 0, len(lineAdvs))
	for i := range lineAdvs {
		lineModel, err := mapper.DtoToAdvertisementLine(&lineAdvs[i])
		if err != nil {
			logging.Logger.Error("lineAdvertisement dont pass the filter", "id", lineAdvs[i].ID, "error", err)
			continue
		}
		convertedLines = append(convertedLines, mapper.LineAdvertisementToDTO(&lineModel))
	}
	return convertedLines, nil
}

func (a *AdvertisementController) BlockAdvertisementsBetweenReleaseDates(ctx context.Context, from, to time.Time) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("advertisementhandler: search request. Search BlockAdvertisements Between release dates", "from", from, "to", to)

	blockAdvs, err := a.entities.BlockAdvertisementBetweenReleaseDates(ctx, from, to)
	if err != nil {
		logging.Logger.Error("advertisementhandler.BlockAdvertisementsBetweenReleaseDates: cannot get BlockAdvertisementsBetweenReleaseDates", "error", err)
		return nil, err
	}
	convertedBlocks := make([]datatransferobjects.BlockAdvertisementDTO, 0, len(blockAdvs))
	for i := range blockAdvs {
		blockModel, err := mapper.DtoToAdvertisementBlock(&blockAdvs[i])
		if err != nil {
			logging.Logger.Error("blockAdvertisement dont pass the filter", "id", blockAdvs[i].ID, "error", err)
			continue
		}
		convertedBlocks = append(convertedBlocks, mapper.BlockAdvertisementToDTO(&blockModel))
	}
	return convertedBlocks, nil
}

func (a *AdvertisementController) LineAdvertisementsBetweenReleaseDates(ctx context.Context, from, to time.Time) ([]datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("advertisementhandler: search request. Search LineAdvertisements Between release dates", "from", from, "to", to)

	lineAdvs, err := a.entities.LineAdvertisementBetweenReleaseDates(ctx, from, to)
	if err != nil {
		logging.Logger.Error("advertisementhandler.LineAdvertisementsBetweenReleaseDates: cannot get LineAdvertisementsBetweenReleaseDates", "error", err)
		return nil, err
	}
	convertedLines := make([]datatransferobjects.LineAdvertisementDTO, 0, len(lineAdvs))
	for i := range lineAdvs {
		lineModel, err := mapper.DtoToAdvertisementLine(&lineAdvs[i])
		if err != nil {
			logging.Logger.Error("lineAdvertisement dont pass the filter", "id", lineAdvs[i].ID, "error", err)
			continue
		}
		convertedLines = append(convertedLines, mapper.LineAdvertisementToDTO(&lineModel))
	}
	return convertedLines, nil
}

func (a *AdvertisementController) BlockAdvertisementsActualReleaseDate(ctx context.Context) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("advertisementhandler: search request. Search BlockAdvertisements with actual release dates", "from", time.Now())

	blockAdvs, err := a.entities.BlockAdvertisementFromReleaseDates(ctx, time.Now())
	if err != nil {
		logging.Logger.Error("advertisementhandler.BlockAdvertisementsActualReleaseDate: cannot get BlockAdvertisementsActualReleaseDate", "error", err)
		return nil, err
	}
	convertedBlocks := make([]datatransferobjects.BlockAdvertisementDTO, 0, len(blockAdvs))
	for i := range blockAdvs {
		blockModel, err := mapper.DtoToAdvertisementBlock(&blockAdvs[i])
		if err != nil {
			logging.Logger.Error("blockAdvertisement dont pass the filter", "id", blockAdvs[i].ID, "error", err)
			continue
		}
		convertedBlocks = append(convertedBlocks, mapper.BlockAdvertisementToDTO(&blockModel))
	}
	return convertedBlocks, nil
}

func (a *AdvertisementController) LineAdvertisementsActualReleaseDate(ctx context.Context) ([]datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("advertisementhandler: search request. Search LineAdvertisements with actual release dates", "from", time.Now())

	lineAdvs, err := a.entities.LineAdvertisementFromReleaseDates(ctx, time.Now())
	if err != nil {
		logging.Logger.Error("advertisementhandler.LineAdvertisementsActualReleaseDate: cannot get LineAdvertisementsActualReleaseDate", "error", err)
		return nil, err
	}

	convertedLines := make([]datatransferobjects.LineAdvertisementDTO, 0, len(lineAdvs))
	for i := range lineAdvs {
		lineModel, err := mapper.DtoToAdvertisementLine(&lineAdvs[i])
		if err != nil {
			logging.Logger.Error("lineAdvertisement dont pass the filter", "id", lineAdvs[i].ID, "error", err)
			continue
		}
		convertedLines = append(convertedLines, mapper.LineAdvertisementToDTO(&lineModel))
	}
	return convertedLines, nil
}

func (a *AdvertisementController) BlockAdvertisementsFromReleaseDate(ctx context.Context, from time.Time) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("advertisementhandler: search request. Search BlockAdvertisements from release dates", "from", from)

	blockAdvs, err := a.entities.BlockAdvertisementFromReleaseDates(ctx, from)
	if err != nil {
		logging.Logger.Error("advertisementhandler.BlockAdvertisementsFromReleaseDate: cannot get BlockAdvertisementsFromReleaseDate", "error", err)
		return nil, err
	}
	convertedBlocks := make([]datatransferobjects.BlockAdvertisementDTO, 0, len(blockAdvs))
	for i := range blockAdvs {
		blockModel, err := mapper.DtoToAdvertisementBlock(&blockAdvs[i])
		if err != nil {
			logging.Logger.Error("blockAdvertisement dont pass the filter", "id", blockAdvs[i].ID, "error", err)
			continue
		}
		convertedBlocks = append(convertedBlocks, mapper.BlockAdvertisementToDTO(&blockModel))
	}
	return convertedBlocks, nil
}

func (a *AdvertisementController) LineAdvertisementsFromReleaseDate(ctx context.Context, from time.Time) ([]datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("advertisementhandler: search request. Search LineAdvertisements from release dates", "from", from)

	lineAdvs, err := a.entities.LineAdvertisementFromReleaseDates(ctx, from)
	if err != nil {
		logging.Logger.Error("advertisementhandler.LineAdvertisementsFromReleaseDate: cannot get LineAdvertisementsFromReleaseDate", "error", err)
		return nil, err
	}
	convertedLines := make([]datatransferobjects.LineAdvertisementDTO, 0, len(lineAdvs))
	for i := range lineAdvs {
		lineModel, err := mapper.DtoToAdvertisementLine(&lineAdvs[i])
		if err != nil {
			logging.Logger.Error("lineAdvertisement dont pass the filter", "id", lineAdvs[i].ID, "error", err)
			continue
		}
		convertedLines = append(convertedLines, mapper.LineAdvertisementToDTO(&lineModel))
	}
	return convertedLines, nil
}

func (a *AdvertisementController) AllTags(ctx context.Context) ([]datatransferobjects.TagDTO, error) {
	logging.Logger.Debug("advertisementhandler: search request. Search all Tags")

	tags, err := a.entities.AllTags(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllTags: cannot get AllTags", "error", err)
		return nil, err
	}
	convertedTags := make([]datatransferobjects.TagDTO, 0, len(tags))
	for i := range tags {
		tagModel, err := mapper.DtoToTag(&tags[i])
		if err != nil {
			logging.Logger.Error("tag dont pass the filter", "name", tags[i].TagName, "error", err)
			continue
		}
		convertedTags = append(convertedTags, mapper.TagToDTO(&tagModel))
	}
	return convertedTags, nil
}

func (a *AdvertisementController) AllExtraCharges(ctx context.Context) ([]datatransferobjects.ExtraChargeDTO, error) {
	logging.Logger.Debug("advertisementhandler: search request. Search all ExtraCharges")

	extraCharges, err := a.entities.AllExtraCharges(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllExtraCharges: cannot get AllExtraCharges", "error", err)
		return nil, err
	}
	convertedExtraCharges := make([]datatransferobjects.ExtraChargeDTO, 0, len(extraCharges))
	for i := range extraCharges {
		extraChargeModel, err := mapper.DtoToExtraCharge(&extraCharges[i])
		if err != nil {
			logging.Logger.Error("extraCharge dont pass the filter", "name", extraCharges[i].ChargeName, "error", err)
			continue
		}
		convertedExtraCharges = append(convertedExtraCharges, mapper.ExtraChargeToDTO(&extraChargeModel))
	}
	return convertedExtraCharges, nil
}

func (a *AdvertisementController) AllCostRates(ctx context.Context) ([]datatransferobjects.CostRateDTO, error) {
	logging.Logger.Debug("advertisementhandler: search request. Search all CostRates")

	costRates, err := a.entities.AllCostRates(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllCostRates: cannot get AllCostRates", "error", err)
		return nil, err
	}
	convertedCostRates := make([]datatransferobjects.CostRateDTO, 0, len(costRates))
	for i := range costRates {
		costRateModel, err := mapper.DtoToCostRate(&costRates[i])
		if err != nil {
			logging.Logger.Error("costRate dont pass the filter", "name", costRates[i].Name, "error", err)
			continue
		}
		convertedCostRates = append(convertedCostRates, mapper.CostRateToDTO(&costRateModel))
	}
	return convertedCostRates, nil
}
