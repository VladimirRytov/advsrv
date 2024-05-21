package userhandler

import (
	"context"
	"encoding/json"
	"errors"
	"slices"
	"time"

	"github.com/VladimirRytov/advsrv/internal/datastorage"
	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/handlers"
	"github.com/VladimirRytov/advsrv/internal/logging"
	"github.com/VladimirRytov/advsrv/internal/mapper"
)

var (
	ErrNotFound     = datastorage.ErrNotFound
	ErrWrondRequest = errWrondRequest()
	ErrWrongData    = errWrongData()
)

func errWrongData() error {
	return errors.New("неверное имя пользователя или пароль")
}

func errWrondRequest() error {
	return errors.New("запрос невозможно выполнить")
}

type UserHandler struct {
	htpasswd bool
	auth     Authorizator

	connected bool
	userRepo  handlers.UserRepo
	connector handlers.DataBaseConnector
	validator Validator
}

func NewUserHandler(validator Validator, manager handlers.DataBaseConnector, auth Authorizator) *UserHandler {
	db, _ := manager.ConnectToDatabase("Dummy", nil)
	return &UserHandler{connector: manager, validator: validator, userRepo: db, auth: auth}
}

func (uh *UserHandler) InitializeControllerFirstAvailable(networkDBs []datatransferobjects.NetworkDataBaseDSN,
	localDBs []datatransferobjects.LocalDSN, users []datatransferobjects.UserDTO) error {
	for i := range networkDBs {
		data, err := json.Marshal(networkDBs[i])
		if err != nil {
			return err
		}
		err = uh.InitializeController(networkDBs[i].Type, data)
		if err != nil {
			continue
		}
		return uh.makeUsers(users)
	}

	for i := range localDBs {
		data, err := json.Marshal(localDBs[i])
		if err != nil {
			return err
		}
		err = uh.InitializeController(networkDBs[i].Type, data)
		if err != nil {
			continue
		}
		return uh.makeUsers(users)
	}
	return errors.New("couldn't connect to all listed databases")
}

func (uh *UserHandler) makeUsers(users []datatransferobjects.UserDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	err := uh.userRepo.RemoveAllUsers(ctx)
	if err != nil {
		return err
	}
	for i := range users {
		_, err := uh.userRepo.NewUser(ctx, &users[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (uh *UserHandler) InitializeController(dbName string, params []byte) error {
	db, err := uh.connector.ConnectToDatabase(dbName, params)
	if err != nil {
		return err
	}
	db.InitializeDatabaseMode(true)
	uh.userRepo = db
	uh.connected = true
	return nil
}

func (uh *UserHandler) Close() error {
	if uh.connected {
		logging.Logger.Info("advertisementhandler: Close connection")
		uh.connected = false
		return uh.userRepo.Close()
	}
	return nil
}

func (uh *UserHandler) ChangeDatabase(dbName string, params []byte) error {
	db, err := uh.connector.ConnectToDatabase(dbName, params)
	if err != nil {
		return err
	}
	db.InitializeDatabaseMode(true)

	uh.userRepo.Close()
	uh.userRepo = db
	uh.connected = true
	return nil
}

func (a *UserHandler) Authenticate(ctx context.Context, user *datatransferobjects.UserDTO) ([]byte, error) {
	userModel, err := mapper.DtoToUser(user, true)
	if err != nil {
		return nil, err
	}

	u, err := a.userRepo.UserByName(ctx, userModel.Name())
	if err != nil {
		return nil, err
	}
	existUser, err := mapper.DtoToUser(&u, false)
	if err != nil {
		return nil, err
	}
	if slices.Equal(existUser.Password(), userModel.Password()) {
		return a.validator.SignDate(userModel.Name(), existUser.Permissions())
	}
	return nil, ErrWrongData
}

func (a *UserHandler) ExtendToken(token []byte) ([]byte, error) {
	userToken, err := a.validator.FetchPayload(token)
	if err != nil {
		return nil, err
	}
	return a.validator.SignDate(userToken.Log, userToken.Perm)
}

func (uh *UserHandler) createAdmin(ctx context.Context, db handlers.UserRepo) error {
	admin := datatransferobjects.UserDTO{
		Name:        "admin",
		Password:    []byte("admin"),
		Permissions: uh.auth.AdminPermissions(),
	}
	adminAdv, err := mapper.DtoToUser(&admin, true)
	if err != nil {
		return err
	}
	user := mapper.UserToDTO(&adminAdv, false)
	_, err = db.NewUser(ctx, &user)
	return err
}
