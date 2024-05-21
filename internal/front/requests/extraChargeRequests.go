package requests

import (
	"context"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func (rh *Requesting) ExtraChargesGetRequest(ctx context.Context, token string) ([]datatransferobjects.ExtraChargeDTO, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return nil, err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return nil, err
	}
	return rh.advRepo.AllExtraCharges(ctx)
}

func (rh *Requesting) ExtraChargeGetRequest(ctx context.Context, token string, name string) (datatransferobjects.ExtraChargeDTO, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return datatransferobjects.ExtraChargeDTO{}, err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return datatransferobjects.ExtraChargeDTO{}, err
	}

	return rh.advRepo.ExtraChargeByName(ctx, name)
}

func (rh *Requesting) ExtraChargePostRequest(ctx context.Context, token string, ExtraCharge *datatransferobjects.ExtraChargeDTO) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanCreateRecords(&user)
	if err != nil {
		return err
	}

	return rh.advRepo.NewExtraCharge(ctx, ExtraCharge)
}

func (rh *Requesting) ExtraChargePutRequest(ctx context.Context, token string, ExtraCharge *datatransferobjects.ExtraChargeDTO) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanWriteRecords(&user)
	if err != nil {
		return err
	}

	return rh.advRepo.UpdateExtraCharge(ctx, ExtraCharge)
}

func (rh *Requesting) ExtraChargeDeleteRequest(ctx context.Context, token string, name string) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanDeleteRecords(&user)
	if err != nil {
		return err
	}

	return rh.advRepo.RemoveExtraChargeByName(ctx, name)
}
