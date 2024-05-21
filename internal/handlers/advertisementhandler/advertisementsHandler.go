package advertisementhandler

import (
	"encoding/json"
	"errors"

	"github.com/VladimirRytov/advsrv/internal/datastorage"
	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/handlers"
	"github.com/VladimirRytov/advsrv/internal/logging"
	"github.com/VladimirRytov/advsrv/internal/mapper"
)

const (
	CancelStr  = "операция была отменена пользователем"
	TimeoutStr = "превышено время ожидания"

	ClientType             = "Client"
	OrderType              = "Order"
	BlockAdvertisementType = "BlockAdv"
	LineAdvertisementType  = "LineAdv"
	TagType                = "Tag"
	ExtraChargeType        = "ExtraCharge"
	CostRateType           = "CostRate"

	DeleteAction = 1
	CreateAction = 2
)

var (
	ErrConvertation       = mapper.ErrConvertation
	ErrDuplicate          = datastorage.ErrDuplicate
	ErrNotFound           = datastorage.ErrNotFound
	ErrViolatesForeignKey = datastorage.ErrViolatesForeignKey

	ErrDatabase = datastorage.ErrDatabase
)

type AdvertisementController struct {
	status      bool
	broadCaster Broadcaster
	connector   handlers.DataBaseConnector
	entities    handlers.AdvertisementRepo
}

func NewAdvertisementController(connector handlers.DataBaseConnector, bc Broadcaster) *AdvertisementController {
	logging.Logger.Info("advertisementhandler: Initialize Advertisement controller")
	db, _ := connector.ConnectToDatabase("Dummy", nil)
	return &AdvertisementController{connector: connector, broadCaster: bc, entities: db}
}

func (ac *AdvertisementController) InitializeController(dbName string, params []byte) (handlers.DataBase, error) {
	logging.Logger.Info("advertisementhandler: Initialize Advertisement controller")
	db, err := ac.connector.ConnectToDatabase(dbName, params)
	if err != nil {
		return nil, err
	}
	db.InitializeDatabaseMode(false)
	ac.entities = db
	return db, nil
}
func (ac *AdvertisementController) InitializeControllerFirstAvailable(networkDBs []datatransferobjects.NetworkDataBaseDSN,
	localDBs []datatransferobjects.LocalDSN) (handlers.DataBase, error) {
	for i := range networkDBs {
		data, err := json.Marshal(networkDBs[i])
		if err != nil {
			return nil, err
		}

		db, err := ac.InitializeController(networkDBs[i].Type, data)
		if err != nil {
			continue
		}
		return db, nil
	}

	for i := range localDBs {
		data, err := json.Marshal(localDBs[i])
		if err != nil {
			return nil, err
		}
		db, err := ac.InitializeController(localDBs[i].Type, data)
		if err != nil {
			continue
		}
		return db, nil
	}
	return nil, errors.New("couldn't connect to all listed databases")

}

func (ac *AdvertisementController) ChangeDatabase(dbName string, params []byte) error {
	logging.Logger.Info("advertisementhandler: Initialize Advertisement controller")
	db, err := ac.connector.ConnectToDatabase(dbName, params)
	if err != nil {
		return err
	}
	db.InitializeDatabaseMode(false)
	ac.entities.Close()
	ac.entities = db
	ac.status = true
	return nil
}

func (ac *AdvertisementController) Close() error {
	if ac.status {
		logging.Logger.Info("advertisementhandler: Close connection")
		ac.status = false
		return ac.entities.Close()
	}
	return nil
}

func (ac *AdvertisementController) ConnectionInfo() (string, []byte) {
	return ac.entities.DbName(), ac.entities.DbParams()
}
