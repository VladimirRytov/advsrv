package requests

import (
	"context"
	"time"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func (rh *Requesting) LineAdvertisementsGetRequest(ctx context.Context, token string, params datatransferobjects.AdvertisementParams) ([]datatransferobjects.LineAdvertisementDTO, error) {
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
		return rh.advRepo.LineAdvertisementsActualReleaseDate(ctx)

	case params.FromDate != time.Time{} && params.ToDate != time.Time{}:
		return rh.advRepo.LineAdvertisementsBetweenReleaseDates(ctx, params.FromDate, params.ToDate)

	case params.FromDate != time.Time{}:
		return rh.advRepo.LineAdvertisementsFromReleaseDate(ctx, params.FromDate)

	case params.ToDate != time.Time{}:
		return rh.advRepo.LineAdvertisementsBetweenReleaseDates(ctx, params.FromDate, params.ToDate)
	case params.OrderID > 0:
		return rh.advRepo.LineAdvertisementsByOrderID(ctx, params.OrderID)
	default:
		return rh.advRepo.AllLineAdvertisements(ctx)
	}
}

func (rh *Requesting) LineAdvertisementGetRequest(ctx context.Context, token string, id int) (datatransferobjects.LineAdvertisementDTO, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return datatransferobjects.LineAdvertisementDTO{}, err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return datatransferobjects.LineAdvertisementDTO{}, err
	}
	return rh.advRepo.LineAdvertisementByID(ctx, id)
}

func (rh *Requesting) LineAdvertisementPostRequest(ctx context.Context, token string, LineAdvertisement *datatransferobjects.LineAdvertisementDTO) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanCreateRecords(&user)
	if err != nil {
		return err
	}
	LineAdvertisement.ID = 0
	return rh.advRepo.NewLineAdvertisement(ctx, LineAdvertisement)
}

func (rh *Requesting) LineAdvertisementPutRequest(ctx context.Context, token string, LineAdvertisement *datatransferobjects.LineAdvertisementDTO) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanWriteRecords(&user)
	if err != nil {
		return err
	}
	return rh.advRepo.UpdateLineAdvertisement(ctx, LineAdvertisement)
}

func (rh *Requesting) LineAdvertisementDeleteRequest(ctx context.Context, token string, id int) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanDeleteRecords(&user)
	if err != nil {
		return err
	}
	return rh.advRepo.RemoveLineAdvertisementByID(ctx, id)
}
