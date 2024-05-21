package advertisementhandler

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/VladimirRytov/advsrv/internal/datastorage/dbmanager"
	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/encodedecoder"
	"github.com/VladimirRytov/advsrv/internal/handlers/broadcaster"
	"github.com/VladimirRytov/advsrv/internal/logging"
	"github.com/VladimirRytov/advsrv/internal/validator"
)

func CreateControllerForTests() (*AdvertisementController, error) {
	CreateLogger()
	param := &datatransferobjects.NetworkDataBaseDSN{
		Adress: "127.0.0.1", DbName: "gorm_test", UserName: "gorm_test", Password: "gorm_test", Port: 5432}
	mar, err := json.Marshal(&param)
	if err != nil {
		return nil, err
	}
	v, err := validator.NewWorker(encodedecoder.NewBase64Encoder())
	if err != nil {
		return nil, err
	}
	sm := &SenderMaker{}
	br := broadcaster.NewSubHandler(v, sm, encodedecoder.NewBase64Encoder())
	manager := dbmanager.NewDatabaseManager()
	advController := NewAdvertisementController(manager, br)
	_, err = advController.InitializeController("Postgres", mar)
	if err != nil {
		return nil, err
	}
	return advController, err
}

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
}
