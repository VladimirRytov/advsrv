package requests

import (
	"context"
	"time"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func (rh *Requesting) BlockAdvertisementsGetRequest(ctx context.Context, token string, params datatransferobjects.AdvertisementParams) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return nil, err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return nil, err
	}

	switch {
	case params.Actual:
		return rh.advRepo.BlockAdvertisementsActualReleaseDate(ctx)
	case params.FromDate != time.Time{} && params.ToDate != time.Time{}:
		return rh.advRepo.BlockAdvertisementsBetweenReleaseDates(ctx, params.FromDate, params.ToDate)
	case params.FromDate != time.Time{}:
		return rh.advRepo.BlockAdvertisementsFromReleaseDate(ctx, params.FromDate)
	case params.ToDate != time.Time{}:
		return rh.advRepo.BlockAdvertisementsBetweenReleaseDates(ctx, params.FromDate, params.ToDate)
	case params.OrderID > 0:
		return rh.advRepo.BlockAdvertisementsByOrderID(ctx, params.OrderID)
	default:
		return rh.advRepo.AllBlockAdvertisements(ctx)
	}
}

func (rh *Requesting) BlockAdvertisementGetRequest(ctx context.Context, token string, id int) (datatransferobjects.BlockAdvertisementDTO, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, err
	}
	return rh.advRepo.BlockAdvertisementByID(ctx, id)
}

func (rh *Requesting) BlockAdvertisementPostRequest(ctx context.Context, token string, blockAdvertisement *datatransferobjects.BlockAdvertisementDTO) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanCreateRecords(&user)
	if err != nil {
		return err
	}
	blockAdvertisement.ID = 0
	return rh.advRepo.NewBlockAdvertisement(ctx, blockAdvertisement)
}

func (rh *Requesting) BlockAdvertisementPutRequest(ctx context.Context, token string, blockAdvertisement *datatransferobjects.BlockAdvertisementDTO) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanWriteRecords(&user)
	if err != nil {
		return err
	}
	return rh.advRepo.UpdateBlockAdvertisement(ctx, blockAdvertisement)
}

func (rh *Requesting) BlockAdvertisementDeleteRequest(ctx context.Context, token string, id int) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanDeleteRecords(&user)
	if err != nil {
		return err
	}
	return rh.advRepo.RemoveBlockAdvertisementByID(ctx, id)
}
