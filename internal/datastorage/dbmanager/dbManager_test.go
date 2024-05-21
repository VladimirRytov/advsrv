package dbmanager

import (
	"log/slog"
	"os"
	"reflect"
	"testing"

	"github.com/VladimirRytov/advsrv/internal/handlers"
	"github.com/VladimirRytov/advsrv/internal/logging"
)

var connectedDbName string

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
}

func fillDBList(dbm *DataBaseManager) {
	CreateLogger()
	localDB := []string{"sqlite"}
	networkDB := map[string]uint{"Postgres": 5432, "Mysql": 3306}
	dbm.localAvailableDatabases = append(dbm.localAvailableDatabases, localDB...)
	for k, v := range networkDB {
		dbm.networkAvailableDatabases[k] = v
	}
}
func fillRegisterDB(dbm *DataBaseManager) {
	CreateLogger()
	localDB := []string{"sqlite"}
	newworkDB := []string{"Postgres", "Mysql"}
	fakeFunc := func(b []byte) (handlers.DataBase, error) {
		connectedDbName = string(b)
		return nil, nil
	}
	for _, v := range localDB {
		Register(v, 0, LocalDB, fakeFunc)
	}
	for _, v := range newworkDB {
		Register(v, 0, NetworkDB, fakeFunc)
	}
}

func TestNewDatabaseManager(t *testing.T) {
	CreateLogger()
	want := &DataBaseManager{}
	got := NewDatabaseManager()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestAvailableDatabases(t *testing.T) {
	CreateLogger()
	wantNetwork := []string{"Postgres", "Mysql"}
	wantLocal := []string{"sqlite"}
	db := NewDatabaseManager()
	fillDBList(db)
	got := db.AvailableLocalDatabases()
	if !reflect.DeepEqual(wantLocal, got) {
		t.Errorf("want %v, got %v", wantLocal, got)
	}
	got = db.AvailableNetworkDatabases()
	if !reflect.DeepEqual(wantNetwork, got) {
		t.Errorf("want %v, got %v", wantNetwork, got)
	}
}

func TestConnectToDataBase(t *testing.T) {
	CreateLogger()
	dbm := NewDatabaseManager()
	fillRegisterDB(dbm)
	localDB := []string{"sqlite"}
	networkDB := []string{"Postgres", "Mysql"}
	for _, v := range localDB {
		dbm.ConnectToDatabase(v, []byte(v))
		if connectedDbName != v {
			t.Fatal("Connect to wrong DB")
		}
	}
	for _, v := range networkDB {
		dbm.ConnectToDatabase(v, []byte(v))
		if connectedDbName != v {
			t.Fatal("Connect to wrong DB")
		}
	}
	_, err := dbm.ConnectToDatabase("asdasd", []byte("asd"))
	if err.Error() != "error: this database not supported" {
		t.Fatalf("want error \"error: this database not supported\",got %v", err)
	}

}
