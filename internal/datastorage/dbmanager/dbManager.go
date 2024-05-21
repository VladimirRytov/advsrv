package dbmanager

import (
	"errors"
	"slices"
	"sync"

	"github.com/VladimirRytov/advsrv/internal/handlers"
	"github.com/VladimirRytov/advsrv/internal/logging"
)

const (
	_ = iota
	LocalDB
	NetworkDB
)

var (
	mu        sync.Mutex
	databases []AvailableDataBases
)

type AvailableDataBases struct {
	dbType             int
	name               string
	defaultNetworkPort uint
	connection         func([]byte) (handlers.DataBase, error)
}

func Register(name string, port uint, dbType int, connFunc func([]byte) (handlers.DataBase, error)) {
	mu.Lock()
	defer mu.Unlock()
	databases = append(databases, AvailableDataBases{
		name:               name,
		defaultNetworkPort: port,
		dbType:             dbType,
		connection:         connFunc,
	})
}

type DataBaseManager struct {
	networkAvailableDatabases map[string]uint
	localAvailableDatabases   []string
}

func NewDatabaseManager() *DataBaseManager {
	logging.Logger.Info("db manager: Initializing database manager")
	dbm := &DataBaseManager{}
	dbm.networkAvailableDatabases = make(map[string]uint, len(databases))
	dbm.appendDBToList(databases)
	return dbm
}

func (dbm *DataBaseManager) appendDBToList(dbs []AvailableDataBases) {
	for _, v := range dbs {
		switch v.dbType {
		case LocalDB:
			dbm.localAvailableDatabases = append(dbm.localAvailableDatabases, v.name)
		case NetworkDB:
			dbm.networkAvailableDatabases[v.name] = v.defaultNetworkPort
		}
	}
}
func (dbm *DataBaseManager) AvailableLocalDatabases() []string {
	return dbm.localAvailableDatabases
}

func (dbm *DataBaseManager) AvailableNetworkDatabases() []string {
	listNetworkDatabases := make([]string, 0, len(dbm.networkAvailableDatabases))
	for k := range dbm.networkAvailableDatabases {
		if k == "Dummy" {
			continue
		}
		listNetworkDatabases = append(listNetworkDatabases, k)
	}
	slices.Sort(listNetworkDatabases)
	return listNetworkDatabases
}

func (dbm *DataBaseManager) DefaultPort(dnname string) uint {
	return dbm.networkAvailableDatabases[dnname]
}

func (dbm *DataBaseManager) ConnectToDatabase(dbName string, params []byte) (handlers.DataBase, error) {
	logging.Logger.Info("db manager: Search database for connecting")
	for _, v := range databases {
		if v.name == dbName {
			db, err := v.connection(params)
			if err != nil {
				return nil, err
			}
			db.SetDbName(dbName)
			db.SetDbParams(params)
			logging.Logger.Info("db manager: Connection to the database was successful")
			return db, nil
		}
	}
	logging.Logger.Error("db manager: Database not found", "dbName", dbName)
	return nil, errors.New("error: this database not supported")
}
