package advertisementhandler

type Broadcaster interface {
	SendData(data []byte, kind string, action int) error
}
