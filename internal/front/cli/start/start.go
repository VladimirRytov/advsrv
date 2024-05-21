package start

import (
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/VladimirRytov/advsrv/internal/authorizator"
	"github.com/VladimirRytov/advsrv/internal/configparser"
	"github.com/VladimirRytov/advsrv/internal/datastorage/dbmanager"
	"github.com/VladimirRytov/advsrv/internal/encodedecoder"
	"github.com/VladimirRytov/advsrv/internal/filestorage"
	"github.com/VladimirRytov/advsrv/internal/front/converter"
	"github.com/VladimirRytov/advsrv/internal/front/reciever"
	front "github.com/VladimirRytov/advsrv/internal/front/reciever/fiber"
	"github.com/VladimirRytov/advsrv/internal/front/requests"
	"github.com/VladimirRytov/advsrv/internal/front/rpcworker"
	"github.com/VladimirRytov/advsrv/internal/front/sender"
	"github.com/VladimirRytov/advsrv/internal/handlers/advertisementhandler"
	"github.com/VladimirRytov/advsrv/internal/handlers/authorizatorhandler"
	"github.com/VladimirRytov/advsrv/internal/handlers/broadcaster"
	"github.com/VladimirRytov/advsrv/internal/handlers/costcalculationhandler"
	"github.com/VladimirRytov/advsrv/internal/handlers/filehandler"
	"github.com/VladimirRytov/advsrv/internal/handlers/signalhandler"
	"github.com/VladimirRytov/advsrv/internal/handlers/userhandler"
	"github.com/VladimirRytov/advsrv/internal/logging"
	"github.com/VladimirRytov/advsrv/internal/minimizer"
	"github.com/VladimirRytov/advsrv/internal/validator"
	"github.com/spf13/cobra"

	_ "github.com/VladimirRytov/advsrv/internal/datastorage/dummy"
	_ "github.com/VladimirRytov/advsrv/internal/datastorage/sql/orm"
)

var (
	ConfigFilePath string

	advController  *advertisementhandler.AdvertisementController
	userController *userhandler.UserHandler
	web            reciever.WebServer

	Version string
)

var Start = &cobra.Command{Use: "start", Short: "start server", RunE: func(cmd *cobra.Command, args []string) error {
	err := rpcworker.CheckExisted()
	if err != nil {
		return err
	}
	return RunMain()
}}

func RunMain() error {
	logging.CreateLogger("", 14*(24*time.Hour), &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
	//
	authorizator := authorizator.NewAuthrizator()
	base64Encoder := encodedecoder.NewBase64Encoder()

	authorizeHandler := authorizatorhandler.NewAuthrizator(authorizator)
	databaseManager := dbmanager.NewDatabaseManager()

	token, err := validator.NewWorker(base64Encoder)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// worker with JWT Token
	spraySigner, err := validator.NewWorker(base64Encoder)
	if err != nil {
		log.Fatalln(err.Error())
	}

	broadCaster := broadcaster.NewSubHandler(spraySigner, &sender.SenderMaker{}, base64Encoder)
	advController = advertisementhandler.NewAdvertisementController(databaseManager, broadCaster)
	userController = userhandler.NewUserHandler(token, databaseManager, authorizator)
	defer advController.Close()
	defer userController.Close()

	ch := configparser.NewConfigparser(authorizator, base64Encoder)

	fileStorage := filestorage.NewFileStorage()
	fileHandler, err := filehandler.NewFilehandler(fileStorage, &minimizer.Imager{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	reqHandler := requests.NewRequestHandler(token, authorizeHandler, advController, userController, broadCaster, fileHandler, base64Encoder)
	frontConverter := converter.NewFrontConverter(base64Encoder)
	web = front.Create(reqHandler, frontConverter, base64Encoder)
	costCalculator := costcalculationhandler.NewCostRateCalculator(nil)
	reqHandler.SetCostRateCalculator(costCalculator)

	signalHandler := signalhandler.NewSignalHandler(userController, advController, ch, costCalculator, web, fileHandler)
	rpcWorder := rpcworker.NewRpcServer(signalHandler)

	return rpcWorder.Listen(ConfigFilePath)
}
