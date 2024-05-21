package orm

import (
	"context"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/logging"

	"gorm.io/gorm/clause"
)

func (ds *DataStorageOrm) NewClient(ctx context.Context, client *datatransferobjects.ClientDTO) (string, error) {
	logging.Logger.Debug("orm: create request. Saving Client to database")
	newClient := convertClientToModel(client)
	err := ds.db.WithContext(ctx).Table("clients").Create(&newClient).Error
	return newClient.Name, handleError(err)
}

func (ds *DataStorageOrm) NewAdvertisementsOrder(ctx context.Context, order *datatransferobjects.OrderDTO) (datatransferobjects.OrderDTO, error) {
	logging.Logger.Debug("orm: create request. Saving Order to database")
	orderModel := convertOrderToModel(order)
	logging.Logger.Debug("orm: create request. Saving Order to database", "orderModel", orderModel)

	err := ds.db.WithContext(ctx).Table("orders").Create(&orderModel).Error
	if err != nil {
		return datatransferobjects.OrderDTO{}, handleError(err)
	}
	return convertOrderToDTO(&orderModel), handleError(err)
}

func (ds *DataStorageOrm) NewLineAdvertisement(ctx context.Context, lineadv *datatransferobjects.LineAdvertisementDTO) (int, error) {
	logging.Logger.Debug("orm: create request. Saving LineAdvertisement to database")
	newLineAdv := convertLineAdvertisementToModel(lineadv)
	err := ds.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Table("advertisement_lines").Create(&newLineAdv).Error
	return newLineAdv.ID, handleError(err)
}

func (ds *DataStorageOrm) NewBlockAdvertisement(ctx context.Context, blockadv *datatransferobjects.BlockAdvertisementDTO) (int, error) {
	logging.Logger.Debug("orm: create request. Saving BlockAdvertisement to database")
	newBlockAdv := convertBlockAdvertisementToModel(blockadv)
	err := ds.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&newBlockAdv).Error
	return newBlockAdv.ID, handleError(err)
}

func (ds *DataStorageOrm) NewTag(ctx context.Context, tag *datatransferobjects.TagDTO) (string, error) {
	logging.Logger.Debug("orm: create request. Saving Tag to database")
	newTag := convertTagToModel(tag)
	err := ds.db.WithContext(ctx).Create(&newTag).Error
	return newTag.Name, handleError(err)
}

func (ds *DataStorageOrm) NewExtraCharge(ctx context.Context, ExtraCharges *datatransferobjects.ExtraChargeDTO) (string, error) {
	logging.Logger.Debug("orm: create request. Saving ExtraCharges to database")
	newExtraCharges := convertExtraChargeToModel(ExtraCharges)
	err := ds.db.WithContext(ctx).Create(&newExtraCharges).Error
	return newExtraCharges.Name, handleError(err)
}

func (ds *DataStorageOrm) NewCostRate(ctx context.Context, costRate *datatransferobjects.CostRateDTO) (string, error) {
	logging.Logger.Debug("orm: create request. Saving CostRate to database")
	newCostRate := convertCostRateToModel(costRate)
	err := ds.db.WithContext(ctx).Create(&newCostRate).Error
	return newCostRate.Name, handleError(err)
}

func (ds *DataStorageOrm) NewUser(ctx context.Context, user *datatransferobjects.UserDTO) (string, error) {
	logging.Logger.Debug("orm: create request. Saving User to database", "user", user)
	newUser := ConvertUserToModel(user)
	err := ds.db.WithContext(ctx).Create(&newUser).Error
	return newUser.Name, handleError(err)
}
