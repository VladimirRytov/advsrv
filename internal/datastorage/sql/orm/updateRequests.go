package orm

import (
	"context"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/logging"

	"gorm.io/gorm"
)

func (ds *DataStorageOrm) UpdateClient(ctx context.Context, client *datatransferobjects.ClientDTO) error {
	logging.Logger.Debug("orm: update request. Updating Client")
	err := ds.db.WithContext(ctx).Select("*").Updates(convertClientToModel(client)).Error
	return handleError(err)
}

func (ds *DataStorageOrm) UpdateOrder(ctx context.Context, order *datatransferobjects.OrderDTO) error {
	logging.Logger.Debug("orm: update request. Updating Order")
	err := ds.db.WithContext(ctx).Select("*").Updates(convertOrderToModel(order)).Error
	return handleError(err)
}

func (ds *DataStorageOrm) UpdateLineAdvertisement(ctx context.Context, lineadv *datatransferobjects.LineAdvertisementDTO) error {
	logging.Logger.Debug("orm: update request. Updating LineAdvertisement")
	lineModel := convertLineAdvertisementToModel(lineadv)
	err := ds.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Updates(&lineModel).Error
		if err != nil {
			return err
		}

		err = tx.WithContext(ctx).Model(&lineModel).Association("Tags").Replace(lineModel.Tags)
		if err != nil {
			return err
		}

		err = tx.WithContext(ctx).Model(&lineModel).Association("ExtraCharges").Replace(lineModel.ExtraCharges)
		if err != nil {
			return err
		}

		err = tx.WithContext(ctx).Model(&lineModel).Association("ReleaseDates").Replace(lineModel.ReleaseDates)
		if err != nil {
			return err
		}
		return nil
	})
	return handleError(err)
}

func (ds *DataStorageOrm) UpdateBlockAdvertisement(ctx context.Context, blockadv *datatransferobjects.BlockAdvertisementDTO) error {
	logging.Logger.Debug("orm: update request. Updating BlockAdvertisement")
	blockModel := convertBlockAdvertisementToModel(blockadv)
	err := ds.db.Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Select("*").Updates(&blockModel).Error
		if err != nil {
			return err
		}

		err = tx.WithContext(ctx).Model(&blockModel).Association("Tags").Replace(blockModel.Tags)
		if err != nil {
			return err
		}

		err = tx.WithContext(ctx).Model(&blockModel).Association("ExtraCharges").Replace(blockModel.ExtraCharges)
		if err != nil {
			return err
		}

		err = tx.WithContext(ctx).Model(&blockModel).Association("ReleaseDates").Replace(blockModel.ReleaseDates)
		if err != nil {
			return err
		}
		return nil
	})
	return handleError(err)
}

func (ds *DataStorageOrm) UpdateExtraCharge(ctx context.Context, extraCharge *datatransferobjects.ExtraChargeDTO) error {
	logging.Logger.Debug("orm: update request. Updating ExtraCharge")
	err := ds.db.WithContext(ctx).Select("*").Updates(convertExtraChargeToModel(extraCharge)).Error
	return handleError(err)
}

func (ds *DataStorageOrm) UpdateTag(ctx context.Context, tag *datatransferobjects.TagDTO) error {
	logging.Logger.Debug("orm: update request. Updating Tag")
	err := ds.db.WithContext(ctx).Updates(convertTagToModel(tag)).Error
	return handleError(err)
}

func (ds *DataStorageOrm) UpdateCostRate(ctx context.Context, costRate *datatransferobjects.CostRateDTO) error {
	logging.Logger.Debug("orm: update request. Updating CostRate")
	err := ds.db.WithContext(ctx).Updates(convertCostRateToModel(costRate)).Error
	return handleError(err)
}

func (ds *DataStorageOrm) UpdateUser(ctx context.Context, user *datatransferobjects.UserDTO) error {
	logging.Logger.Debug("orm: update request. Updating User")
	err := ds.db.WithContext(ctx).Updates(ConvertUserToModel(user)).Error
	return handleError(err)
}
