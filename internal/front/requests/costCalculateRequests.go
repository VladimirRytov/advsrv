package requests

import (
	"context"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func (rh *Requesting) CalculateOrderCost(ctx context.Context, token string, params datatransferobjects.OrderParams,
	order *datatransferobjects.OrderDTO) (datatransferobjects.OrderDTO, error) {

	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return datatransferobjects.OrderDTO{}, err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return datatransferobjects.OrderDTO{}, err
	}
	return rh.costCalculator.CalculateOrderCost(ctx, *order, params.Costrate)
}

func (rh *Requesting) CalculateBlockAdvertisementCost(ctx context.Context, token string, params datatransferobjects.AdvertisementParams,
	blockAdv *datatransferobjects.BlockAdvertisementDTO) (datatransferobjects.BlockAdvertisementDTO, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, err
	}
	return rh.costCalculator.CalculateBlockAdvertisementCost(ctx, *blockAdv, params.Costrate)

}

func (rh *Requesting) CalculateLineAdvertisementCost(ctx context.Context, token string, params datatransferobjects.AdvertisementParams,
	lineAdv *datatransferobjects.LineAdvertisementDTO) (datatransferobjects.LineAdvertisementDTO, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return datatransferobjects.LineAdvertisementDTO{}, err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return datatransferobjects.LineAdvertisementDTO{}, err
	}
	return rh.costCalculator.CalculateLineAdvertisementCost(ctx, *lineAdv, params.Costrate)

}
