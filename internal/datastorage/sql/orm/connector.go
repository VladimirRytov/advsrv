package orm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/VladimirRytov/advsrv/internal/datastorage/dbmanager"
	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/handlers"
	"github.com/VladimirRytov/advsrv/internal/logging"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DataStorageOrm struct {
	dbName string
	params []byte
	db     *gorm.DB
}

func ConnectToSqlite(p []byte) (handlers.DataBase, error) {
	logging.Logger.Info("orm: Start Connecting to Sqlite")
	var err error
	reader := bytes.NewReader(p)
	param := &datatransferobjects.LocalDSN{}
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&param)
	if err != nil {
		logging.Logger.Error("orm: Decoding failed", "error", err)
		return nil, err
	}
	db, err := gorm.Open(sqlite.Open(filepath.Join(param.Path, param.Name+"?_foreign_keys=1")), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		logging.Logger.Warn("orm: Can`t open database file", "error", err)
		return nil, err
	}
	return &DataStorageOrm{db: db}, nil
}
func ConnectToPostgres(p []byte) (handlers.DataBase, error) {
	logging.Logger.Info("orm: Start Connecting to Postgresql")
	var err error
	dsnParam := &datatransferobjects.NetworkDataBaseDSN{}
	reader := bytes.NewReader(p)
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&dsnParam)
	if err != nil {
		logging.Logger.Error("orm: Decoding failed", "error", err)
		return nil, err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dsnParam.Adress, dsnParam.UserName, dsnParam.Password, dsnParam.DbName, dsnParam.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		logging.Logger.Warn("orm: Connection failed", "error", err)
		return nil, errors.Unwrap(err)
	}
	return &DataStorageOrm{db: db}, nil
}

func ConnectToSqlServer(p []byte) (handlers.DataBase, error) {
	logging.Logger.Info("orm: Start Connecting to Sql Server")

	var err error
	dsnParam := &datatransferobjects.NetworkDataBaseDSN{}
	reader := bytes.NewReader(p)
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&dsnParam)
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		dsnParam.UserName, dsnParam.Password, dsnParam.Adress, dsnParam.Port, dsnParam.DbName)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		logging.Logger.Warn("orm: Connection failed", "error", err)
		return nil, err
	}
	return &DataStorageOrm{db: db}, nil
}

func ConnectToMysql(p []byte) (handlers.DataBase, error) {
	logging.Logger.Info("orm: Start Connecting to MySql")
	var err error
	dsnParam := &datatransferobjects.NetworkDataBaseDSN{}
	reader := bytes.NewReader(p)
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&dsnParam)
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		dsnParam.UserName, dsnParam.Password, dsnParam.Adress, dsnParam.Port, dsnParam.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		logging.Logger.Warn("orm: Connection failed", "error", err)
		return nil, err
	}

	return &DataStorageOrm{db: db}, nil
}

func (ds *DataStorageOrm) Close() error {
	logging.Logger.Info("orm: Closing connection")
	sql, _ := ds.db.DB()
	err := sql.Close()
	if err != nil {
		return err
	}
	return nil
}

func (ds *DataStorageOrm) InitializeDatabaseMode(userGate bool) error {
	if userGate {
		return ds.db.AutoMigrate(&User{})
	}
	return ds.db.AutoMigrate(&Client{}, &Order{}, &ExtraCharge{}, &AdvertisementBlock{},
		&AdvertisementLine{}, &ReleaseDates{}, &CostRate{})
}

func (ds *DataStorageOrm) SetDbName(name string) {
	ds.dbName = name
}

func (ds *DataStorageOrm) SetDbParams(params []byte) {
	ds.params = params
}

func (ds *DataStorageOrm) DbName() string {
	return ds.dbName
}

func (ds *DataStorageOrm) DbParams() []byte {
	return ds.params
}

func init() {
	dbmanager.Register("Postgres", 5432, dbmanager.NetworkDB, ConnectToPostgres)
	dbmanager.Register("MySql", 3306, dbmanager.NetworkDB, ConnectToMysql)
	dbmanager.Register("Sql Server", 1433, dbmanager.NetworkDB, ConnectToSqlServer)
	dbmanager.Register("Sqlite", 0, dbmanager.LocalDB, ConnectToSqlite)
}
