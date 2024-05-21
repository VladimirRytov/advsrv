package requests

import (
	"context"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func (rh *Requesting) CostRatesGetRequest(ctx context.Context, token string) ([]datatransferobjects.CostRateDTO, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return nil, err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return nil, err
	}

	return rh.advRepo.AllCostRates(ctx)
}

func (rh *Requesting) CostRateGetRequest(ctx context.Context, token string, name string) (datatransferobjects.CostRateDTO, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return datatransferobjects.CostRateDTO{}, err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return datatransferobjects.CostRateDTO{}, err
	}

	return rh.advRepo.CostRateByName(ctx, name)
}

func (rh *Requesting) CostRatePostRequest(ctx context.Context, token string, costRate *datatransferobjects.CostRateDTO) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanCreateRecords(&user)
	if err != nil {
		return err
	}

	return rh.advRepo.NewCostRate(ctx, costRate)
}

func (rh *Requesting) CostRatePutRequest(ctx context.Context, token string, costRate *datatransferobjects.CostRateDTO) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanWriteRecords(&user)
	if err != nil {
		return err
	}

	return rh.advRepo.UpdateCostRate(ctx, costRate)
}

func (rh *Requesting) CostRateDeleteRequest(ctx context.Context, token string, name string) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanDeleteRecords(&user)
	if err != nil {
		return err
	}

	return rh.advRepo.RemoveCostRateByName(ctx, name)
}
