package signalhandler

import (
	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/handlers"
)

type UserRequestsHandler interface {
	InitializeControllerFirstAvailable(networkDBs []datatransferobjects.NetworkDataBaseDSN,
		localDBs []datatransferobjects.LocalDSN, users []datatransferobjects.UserDTO) error
	InitializeController(dbName string, params []byte) error
	DatabasesHandler
}

type RequestsHandler interface {
	InitializeControllerFirstAvailable(networkDBs []datatransferobjects.NetworkDataBaseDSN,
		localDBs []datatransferobjects.LocalDSN) (handlers.DataBase, error)
	InitializeController(dbName string, params []byte) (handlers.DataBase, error)
	DatabasesHandler
}

type DatabasesHandler interface {
	ChangeDatabase(dbName string, params []byte) error
	Close() error
}

type ConfigsHandler interface {
	ParseConfig(path string) error
	ListenParams() (datatransferobjects.ListenParams, error)
	FileStorage() (datatransferobjects.FileStorage, error)
	Databases(dbType string) ([]datatransferobjects.NetworkDataBaseDSN, []datatransferobjects.LocalDSN, error)
	UsersDTO() ([]datatransferobjects.UserDTO, error)
	GenerateTemplate() error
	Records() string
	Users() string
	RewriteConfig() error
}

type Front interface {
	Listen(adress string) error
	ShutDown() error
	ListenTLS(adress, certFile, keyFile string) error
}

type FilesHandler interface {
	SetFolder(string) error
	Init() error
	ReloadCache() error
	CleanCache() error
}
