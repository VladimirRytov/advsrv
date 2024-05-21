package dummy

import (
	"context"
	"time"

	"github.com/VladimirRytov/advsrv/internal/datastorage"
	"github.com/VladimirRytov/advsrv/internal/datastorage/dbmanager"
	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/handlers"
)

type DummyDB struct{}

func ConnectToDummy(p []byte) (handlers.DataBase, error) {
	return &DummyDB{}, nil
}

func (dd *DummyDB) InitializeDatabaseMode(userGate bool) error {
	return nil
}

func (dd *DummyDB) DbName() string {
	return ""
}

func (dd *DummyDB) RemoveAllUsers(ctx context.Context) error {
	return nil
}

func (dd *DummyDB) SetDbName(name string) {
}

func (dd *DummyDB) SetDbParams(params []byte) {
}
func (dd *DummyDB) DbParams() []byte {
	return nil
}

func (dd *DummyDB) NewClient(ctx context.Context, client *datatransferobjects.ClientDTO) (string, error) {
	return "", datastorage.ErrDatabase
}

func (dd *DummyDB) NewAdvertisementsOrder(ctx context.Context, order *datatransferobjects.OrderDTO) (datatransferobjects.OrderDTO, error) {
	return datatransferobjects.OrderDTO{}, datastorage.ErrDatabase
}

func (dd *DummyDB) NewLineAdvertisement(ctx context.Context, lineadv *datatransferobjects.LineAdvertisementDTO) (int, error) {
	return 0, datastorage.ErrDatabase
}

func (dd *DummyDB) NewBlockAdvertisement(ctx context.Context, blockadv *datatransferobjects.BlockAdvertisementDTO) (int, error) {
	return 0, datastorage.ErrDatabase
}

func (dd *DummyDB) NewTag(ctx context.Context, tag *datatransferobjects.TagDTO) (string, error) {
	return "", datastorage.ErrDatabase
}

func (dd *DummyDB) NewExtraCharge(ctx context.Context, ExtraCharges *datatransferobjects.ExtraChargeDTO) (string, error) {
	return "", datastorage.ErrDatabase
}

func (dd *DummyDB) NewCostRate(ctx context.Context, costRate *datatransferobjects.CostRateDTO) (string, error) {
	return "", datastorage.ErrDatabase
}

func (dd *DummyDB) NewUser(ctx context.Context, user *datatransferobjects.UserDTO) (string, error) {
	return "", datastorage.ErrDatabase
}

func (dd *DummyDB) ClientByName(ctx context.Context, name string) (datatransferobjects.ClientDTO, error) {
	return datatransferobjects.ClientDTO{}, datastorage.ErrDatabase
}

func (dd *DummyDB) OrderByID(ctx context.Context, id int) (datatransferobjects.OrderDTO, error) {
	return datatransferobjects.OrderDTO{}, datastorage.ErrDatabase
}

func (dd *DummyDB) LineAdvertisementByID(ctx context.Context, id int) (datatransferobjects.LineAdvertisementDTO, error) {
	return datatransferobjects.LineAdvertisementDTO{}, datastorage.ErrDatabase
}

func (dd *DummyDB) BlockAdvertisementByID(ctx context.Context, id int) (datatransferobjects.BlockAdvertisementDTO, error) {
	return datatransferobjects.BlockAdvertisementDTO{}, datastorage.ErrDatabase
}

func (dd *DummyDB) TagByName(ctx context.Context, tagName string) (datatransferobjects.TagDTO, error) {
	return datatransferobjects.TagDTO{}, datastorage.ErrDatabase
}

func (dd *DummyDB) ExtraChargeByName(ctx context.Context, chargeName string) (datatransferobjects.ExtraChargeDTO, error) {
	return datatransferobjects.ExtraChargeDTO{}, datastorage.ErrDatabase
}

func (dd *DummyDB) CostRateByName(ctx context.Context, name string) (datatransferobjects.CostRateDTO, error) {
	return datatransferobjects.CostRateDTO{}, datastorage.ErrDatabase
}

func (dd *DummyDB) UserByName(ctx context.Context, name string) (datatransferobjects.UserDTO, error) {
	return datatransferobjects.UserDTO{}, datastorage.ErrDatabase
}

func (dd *DummyDB) UserByNameHidePassword(ctx context.Context, name string) (datatransferobjects.UserDTO, error) {
	return datatransferobjects.UserDTO{}, datastorage.ErrDatabase
}

func (dd *DummyDB) RemoveClientByName(ctx context.Context, name string) error {
	return datastorage.ErrDatabase
}

func (dd *DummyDB) RemoveOrderByID(ctx context.Context, id int) error {
	return datastorage.ErrDatabase
}

func (dd *DummyDB) RemoveLineAdvertisementByID(ctx context.Context, id int) error {
	return datastorage.ErrDatabase
}

func (dd *DummyDB) RemoveBlockAdvertisementByID(ctx context.Context, id int) error {
	return datastorage.ErrDatabase
}

func (dd *DummyDB) RemoveTagByName(ctx context.Context, name string) error {
	return datastorage.ErrDatabase
}

func (dd *DummyDB) RemoveExtraChargeByName(ctx context.Context, name string) error {
	return datastorage.ErrDatabase
}

func (dd *DummyDB) RemoveCostRateByName(ctx context.Context, name string) error {
	return datastorage.ErrDatabase
}

func (dd *DummyDB) RemoveUser(ctx context.Context, name string) error {
	return datastorage.ErrDatabase
}

func (dd *DummyDB) AllClients(ctx context.Context) ([]datatransferobjects.ClientDTO, error) {
	return nil, datastorage.ErrDatabase
}

func (dd *DummyDB) AllOrders(ctx context.Context) ([]datatransferobjects.OrderDTO, error) {
	return nil, datastorage.ErrDatabase
}

func (dd *DummyDB) OrdersByClientName(ctx context.Context, name string) ([]datatransferobjects.OrderDTO, error) {
	return nil, datastorage.ErrDatabase
}

func (dd *DummyDB) AllLineAdvertisements(ctx context.Context) ([]datatransferobjects.LineAdvertisementDTO, error) {
	return nil, datastorage.ErrDatabase
}

func (dd *DummyDB) LineAdvertisementsByOrderID(ctx context.Context, id int) ([]datatransferobjects.LineAdvertisementDTO, error) {
	return nil, datastorage.ErrDatabase
}

func (dd *DummyDB) AllBlockAdvertisements(ctx context.Context) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	return nil, datastorage.ErrDatabase
}

func (dd *DummyDB) BlockAdvertisementsByOrderID(ctx context.Context, id int) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	return nil, datastorage.ErrDatabase
}

func (dd *DummyDB) BlockAdvertisementBetweenReleaseDates(ctx context.Context, from, to time.Time) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	return nil, datastorage.ErrDatabase
}

func (dd *DummyDB) BlockAdvertisementFromReleaseDates(ctx context.Context, from time.Time) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	return nil, datastorage.ErrDatabase
}

func (dd *DummyDB) LineAdvertisementBetweenReleaseDates(ctx context.Context, from, to time.Time) ([]datatransferobjects.LineAdvertisementDTO, error) {
	return nil, datastorage.ErrDatabase
}

func (dd *DummyDB) LineAdvertisementFromReleaseDates(ctx context.Context, from time.Time) ([]datatransferobjects.LineAdvertisementDTO, error) {
	return nil, datastorage.ErrDatabase
}

func (dd *DummyDB) AllTags(ctx context.Context) ([]datatransferobjects.TagDTO, error) {
	return nil, datastorage.ErrDatabase
}

func (dd *DummyDB) AllExtraCharges(ctx context.Context) ([]datatransferobjects.ExtraChargeDTO, error) {
	return nil, datastorage.ErrDatabase
}

func (dd *DummyDB) AllCostRates(ctx context.Context) ([]datatransferobjects.CostRateDTO, error) {
	return nil, datastorage.ErrDatabase
}

func (dd *DummyDB) AllUsers(ctx context.Context) ([]datatransferobjects.UserDTO, error) {
	return nil, datastorage.ErrDatabase
}

func (dd *DummyDB) AllUsersHidePassword(ctx context.Context) ([]datatransferobjects.UserDTO, error) {
	return nil, datastorage.ErrDatabase
}

func (dd *DummyDB) UpdateClient(ctx context.Context, client *datatransferobjects.ClientDTO) error {
	return datastorage.ErrDatabase
}

func (dd *DummyDB) UpdateOrder(ctx context.Context, order *datatransferobjects.OrderDTO) error {
	return datastorage.ErrDatabase
}

func (dd *DummyDB) UpdateLineAdvertisement(ctx context.Context, lineadv *datatransferobjects.LineAdvertisementDTO) error {
	return datastorage.ErrDatabase
}

func (dd *DummyDB) UpdateBlockAdvertisement(ctx context.Context, blockadv *datatransferobjects.BlockAdvertisementDTO) error {
	return datastorage.ErrDatabase
}

func (dd *DummyDB) UpdateExtraCharge(ctx context.Context, extraCharge *datatransferobjects.ExtraChargeDTO) error {
	return datastorage.ErrDatabase
}

func (dd *DummyDB) UpdateTag(ctx context.Context, tag *datatransferobjects.TagDTO) error {
	return datastorage.ErrDatabase
}

func (dd *DummyDB) UpdateCostRate(ctx context.Context, costRate *datatransferobjects.CostRateDTO) error {
	return datastorage.ErrDatabase
}

func (dd *DummyDB) UpdateUser(ctx context.Context, user *datatransferobjects.UserDTO) error {
	return datastorage.ErrDatabase
}

func (dd *DummyDB) Close() error {
	return nil
}

func init() {
	dbmanager.Register("Dummy", 0000, dbmanager.NetworkDB, ConnectToDummy)
}
