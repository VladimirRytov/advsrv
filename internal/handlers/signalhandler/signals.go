package signalhandler

import (
	"log"
	"path/filepath"
	"strconv"

	"github.com/VladimirRytov/advsrv/internal/handlers"
	"github.com/VladimirRytov/advsrv/internal/logging"
)

type SignalHandler struct {
	userRepo     UserRequestsHandler
	advRepo      RequestsHandler
	configPath   string
	configs      ConfigsHandler
	costRateCalc handlers.CostRateCalculator
	front        Front
	files        FilesHandler
}

func NewSignalHandler(userRepo UserRequestsHandler, advRepo RequestsHandler,
	configs ConfigsHandler, calc handlers.CostRateCalculator,
	front Front, fh FilesHandler) *SignalHandler {
	return &SignalHandler{
		userRepo:     userRepo,
		advRepo:      advRepo,
		configs:      configs,
		costRateCalc: calc,
		front:        front,
		files:        fh}
}

func (sh *SignalHandler) Start(filePath string) error {
	if len(filePath) != 0 {
		globalPath, err := filepath.Abs(filePath)
		if err != nil {
			return err
		}
		sh.configPath = globalPath
	}
	err := sh.configs.ParseConfig(sh.configPath)
	if err != nil {
		return sh.handleError(err)
	}

	err = sh.connectToRecordsDatabase()
	if err != nil {
		return sh.handleError(err)
	}
	err = sh.connectToUsersDatabase()
	if err != nil {
		return sh.handleError(err)
	}

	err = sh.initFilesHandler()
	if err != nil {
		return sh.handleError(err)
	}
	go sh.startWeb()
	sh.configs.RewriteConfig()
	return sh.handleError(err)
}

func (sh *SignalHandler) connectToRecordsDatabase() error {
	advNetworkDb, advLocalDb, err := sh.configs.Databases(sh.configs.Records())
	if err != nil {
		return err
	}

	db, err := sh.advRepo.InitializeControllerFirstAvailable(advNetworkDb, advLocalDb)
	if err != nil {
		return err
	}
	sh.costRateCalc.SetAdvRepo(db)
	return nil
}

func (sh *SignalHandler) connectToUsersDatabase() error {
	userNetworkDb, userLocalDb, err := sh.configs.Databases(sh.configs.Users())
	if err != nil {
		return err
	}
	users, err := sh.configs.UsersDTO()
	if err != nil {
		return err
	}
	return sh.userRepo.InitializeControllerFirstAvailable(userNetworkDb, userLocalDb, users)
}

func (sh *SignalHandler) startWeb() error {
	listenParams, err := sh.configs.ListenParams()
	if err != nil {
		return sh.handleError(err)
	}
	if listenParams.Tls != nil {
		err = sh.front.ListenTLS(listenParams.Adress+":"+strconv.Itoa(int(listenParams.Port)), listenParams.Tls.PathToCertificate, listenParams.Tls.PathToKey)
		if err != nil {
			logging.Logger.Error("cannot start server with current params", "error", err)
			return sh.handleError(err)
		}
	} else {
		err = sh.front.Listen(listenParams.Adress + ":" + strconv.Itoa(int(listenParams.Port)))
		if err != nil {
			logging.Logger.Error("cannot start server with current params", "error", err)
			return sh.handleError(err)
		}
	}
	return nil
}

func (sh *SignalHandler) Stop() error {
	sh.front.ShutDown()
	sh.advRepo.Close()
	sh.userRepo.Close()
	return nil
}

func (sh *SignalHandler) Reload(filePath string) error {
	err := sh.Stop()
	if err != nil {
		return err
	}
	return sh.Start(filePath)
}

func (sh *SignalHandler) ChangeConfig(filePath string) error {
	err := sh.Stop()
	if err != nil {
		return err
	}
	return sh.Start(filePath)
}

func (sh *SignalHandler) initFilesHandler() error {
	files, err := sh.configs.FileStorage()
	if err != nil {
		return err
	}
	err = sh.files.SetFolder(files.Path)
	if err != nil {
		return err
	}
	return sh.files.Init()
}

func (sh *SignalHandler) ReloadCache() error {
	return sh.files.ReloadCache()
}

func (sh *SignalHandler) CleanCache() error {
	return sh.files.CleanCache()
}

func (sh *SignalHandler) handleError(err error) error {
	if err != nil {
		log.Fatalln(err.Error())
	}
	return nil
}
