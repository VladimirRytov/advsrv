package orm

import (
	"context"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/logging"

	"gorm.io/gorm/clause"
)

func (ds *DataStorageOrm) ClientByName(ctx context.Context, name string) (datatransferobjects.ClientDTO, error) {
	logging.Logger.Debug("orm: get request. Getting Client by name")
	var client Client
	err := ds.db.WithContext(ctx).First(&client, "name = ?", name).Error
	return convertClientToDTO(&client), handleError(err)
}

func (ds *DataStorageOrm) OrderByID(ctx context.Context, id int) (datatransferobjects.OrderDTO, error) {
	logging.Logger.Debug("orm: get request. Getting Order by id")
	var order Order
	err := ds.db.WithContext(ctx).First(&order, id).Error
	return convertOrderToDTO(&order), handleError(err)
}

func (ds *DataStorageOrm) LineAdvertisementByID(ctx context.Context, id int) (datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("orm: get request. Getting LineAdvertisement by id")
	var lineAdv AdvertisementLine
	err := ds.db.WithContext(ctx).Preload(clause.Associations).First(&lineAdv, id).Error
	return convertLineAdvertisementToDTO(&lineAdv), handleError(err)
}

func (ds *DataStorageOrm) BlockAdvertisementByID(ctx context.Context, id int) (datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("orm: get request. Getting BlockAdvertisement by id")
	var blockAdv AdvertisementBlock
	err := ds.db.WithContext(ctx).Preload(clause.Associations).First(&blockAdv, id).Error
	return convertBlockAdvertisementToDTO(&blockAdv), handleError(err)
}

func (ds *DataStorageOrm) TagByName(ctx context.Context, tagName string) (datatransferobjects.TagDTO, error) {
	logging.Logger.Debug("orm: get request. Getting Tag by name")
	var tag Tag
	err := ds.db.WithContext(ctx).Where("name = ?", tagName).First(&tag).Error
	return convertTagToDTO(&tag), handleError(err)
}

func (ds *DataStorageOrm) ExtraChargeByName(ctx context.Context, chargeName string) (datatransferobjects.ExtraChargeDTO, error) {
	logging.Logger.Debug("orm: get request. Getting ExtraCharge by name")
	var extraCharge ExtraCharge
	err := ds.db.WithContext(ctx).Where("name = ?", chargeName).First(&extraCharge).Error
	return convertExtraChargeToDTO(&extraCharge), handleError(err)
}

func (ds *DataStorageOrm) CostRateByName(ctx context.Context, name string) (datatransferobjects.CostRateDTO, error) {
	logging.Logger.Debug("orm: get request. Getting CostRate by name")
	var costRate CostRate
	err := ds.db.WithContext(ctx).Where("name = ?", name).First(&costRate).Error
	return convertCostRateToDto(&costRate), handleError(err)
}

func (ds *DataStorageOrm) UserByName(ctx context.Context, name string) (datatransferobjects.UserDTO, error) {
	logging.Logger.Debug("orm: get request. Getting CostRate by name")
	var user User
	err := ds.db.WithContext(ctx).Where("name = ?", name).First(&user).Error
	return ConvertUserToDto(&user), handleError(err)
}

func (ds *DataStorageOrm) UserByNameHidePassword(ctx context.Context, name string) (datatransferobjects.UserDTO, error) {
	logging.Logger.Debug("orm: get request. Getting CostRate by name")
	var user User
	err := ds.db.WithContext(ctx).Select("name", "permissions").Where("name = ?", name).First(&user).Error
	user.Password = []byte("*******")
	return ConvertUserToDto(&user), handleError(err)
}
