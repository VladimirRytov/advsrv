package reciever

type WebServer interface {
	Listen(string) error
	ListenTLS(string, string, string) error
	ShutDown() error
}
