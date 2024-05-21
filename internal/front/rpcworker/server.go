package rpcworker

import (
	"errors"
	"net"
	"net/rpc"
	"os"
	"path/filepath"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

const socketName = "advertisementer.sock"

const (
	_ = iota
	Stop
	Reload
	ReloadCache
	CleanCache
	Ping
	Pong = "pong"

	rpcSendMethod = "RpcServ.Send"
)

type SignalHandler interface {
	Stop() error
	Start(string) error
	Reload(string) error
	ChangeConfig(string) error
	ReloadCache() error
	CleanCache() error
}

type RpcServ struct {
	signalHandler SignalHandler
	serv          *rpc.Server
	listener      net.Listener
}

func NewRpcServer(s SignalHandler) *RpcServ {
	r := new(RpcServ)
	r.serv = rpc.NewServer()
	r.signalHandler = s
	return r
}

func (srv *RpcServ) Send(in, out *datatransferobjects.Message) error {
	switch in.Code {
	case Stop:
		srv.signalHandler.Stop()
		return srv.listener.Close()
	case Reload:
		return srv.signalHandler.Reload(in.Message)
	case ReloadCache:
		return srv.signalHandler.ReloadCache()
	case CleanCache:
		return srv.signalHandler.CleanCache()
	case Ping:
		out.Code = Ping
		out.Message = Pong
		return nil
	}
	return errors.New("unknown method")
}

func (srv *RpcServ) Listen(configPath string) error {

	err := srv.serv.Register(srv)
	if err != nil {
		return err
	}
	srv.listener, err = net.Listen("unix", filepath.Join(os.TempDir(), socketName))
	if err != nil {
		return err
	}
	go srv.signalHandler.Start(configPath)

	srv.serv.Accept(srv.listener)
	return nil
}

func CheckExisted() error {
	_, err := os.Stat(filepath.Join(os.TempDir(), socketName))
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	client, err := rpc.Dial("unix", filepath.Join(os.TempDir(), socketName))
	if err != nil {
		return os.Remove(filepath.Join(os.TempDir(), socketName))
	}
	got := new(datatransferobjects.Message)
	err = client.Call(rpcSendMethod, datatransferobjects.Message{Code: Ping}, got)
	if err != nil {
		return os.Remove(filepath.Join(os.TempDir(), socketName))
	}
	if got.Message != Pong && got.Code != Ping {
		return os.Remove(filepath.Join(os.TempDir(), socketName))
	}
	return errors.New("server already running")
}

func Send(message *datatransferobjects.Message) error {
	client, err := rpc.Dial("unix", filepath.Join(os.TempDir(), socketName))
	if err != nil {
		return err
	}
	return client.Call(rpcSendMethod, message, nil)
}
