package handlers

import (
	"github.com/VladimirRytov/advsrv/internal/logging"
)

type DbManagerHandler struct {
	DataBaseConnector
}

func NewDatabaseController(dbm DataBaseConnector) *DbManagerHandler {
	logging.Logger.Debug("dbManagerHandler: Initialize Database Controller")
	return &DbManagerHandler{DataBaseConnector: dbm}
}

func (dbm *DbManagerHandler) AvailableLocalDatabases() {
	logging.Logger.Debug("dbManagerHandler: Requesting available local databases")
}

func (dbm *DbManagerHandler) AvailableNetworkDatabases() {
	logging.Logger.Debug("dbManagerHandler: Requesting available network databases")
}

func (dbm *DbManagerHandler) DefaultPort(dbName string) {
	logging.Logger.Debug("dbManagerHandler: Requesting default network port for database", "database", dbName)
}
