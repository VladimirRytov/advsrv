package configparser

import (
	"log/slog"
	"os"
	"testing"

	"github.com/VladimirRytov/advsrv/internal/authorizator"
	"github.com/VladimirRytov/advsrv/internal/encodedecoder"
	"github.com/VladimirRytov/advsrv/internal/logging"
)

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
}

func TestGenerateTemplate(t *testing.T) {
	cp := NewConfigparser(authorizator.NewAuthrizator(), encodedecoder.NewBase64Encoder())
	err := cp.GenerateTemplate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestAdvConnectionParams(t *testing.T) {
	cp := NewConfigparser(authorizator.NewAuthrizator(), encodedecoder.NewBase64Encoder())
	err := cp.ParseConfig("config.yaml")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserConnectionParams(t *testing.T) {
	cp := NewConfigparser(authorizator.NewAuthrizator(), encodedecoder.NewBase64Encoder())
	err := cp.ParseConfig("config.yaml")
	if err != nil {
		t.Fatal(err)
	}
}

func TestListenParams(t *testing.T) {
	cp := NewConfigparser(authorizator.NewAuthrizator(), encodedecoder.NewBase64Encoder())
	err := cp.ParseConfig("config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	params, err := cp.ListenParams()
	if err != nil {
		t.Fatal(err, params)
	}
}

func TestFileStorage(t *testing.T) {
	cp := NewConfigparser(authorizator.NewAuthrizator(), encodedecoder.NewBase64Encoder())
	err := cp.ParseConfig("config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	params, err := cp.FileStorage()
	if err != nil {
		t.Fatal(err, params)
	}
	dir, err := os.ReadDir(params.Path)
	if err != nil {
		t.Fatal(err, dir)
	}
}

func TestGetUsers(t *testing.T) {
	cp := NewConfigparser(authorizator.NewAuthrizator(), encodedecoder.NewBase64Encoder())
	err := cp.ParseConfig("config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	users, err := cp.users()
	if err == nil {
		t.Fatal(err, users)
	}
}

func TestGetUsersDTO(t *testing.T) {
	CreateLogger()
	cp := NewConfigparser(authorizator.NewAuthrizator(), encodedecoder.NewBase64Encoder())
	err := cp.ParseConfig("config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	users, err := cp.UsersDTO()
	if err == nil {
		t.Fatal(err, users)
	}
}

func TestGetRecordsDatabases(t *testing.T) {
	cp := NewConfigparser(authorizator.NewAuthrizator(), encodedecoder.NewBase64Encoder())
	err := cp.ParseConfig("config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	n, l, err := cp.Databases(Users)
	if err == nil {
		t.Fatal(err, n, l)
	}
}
