package broadcaster

import (
	"io"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

type Subscriber interface {
	AppendData([]byte)
	Adress() string
	Close()
	Buffer() chan []byte
	Closed() bool
}

type B64Enc interface {
	ToBase64([]byte) []byte
	FromBase64([]byte) ([]byte, error)
}

type SubscribeCreator interface {
	NewPassiveSub(int, int, string, io.Writer) Subscriber
}

type Validator interface {
	SignDate(string, int32) ([]byte, error)
	FetchPayload(token []byte) (datatransferobjects.UserToken, error)
	Validate([]byte) error
}

type SendMaker interface {
	NewSender(URL string) (io.Writer, error)
}

type SendWrapper struct {
	Type   string `json:"type"`
	Action int    `json:"action"`
	Entry  []byte `json:"entry"`
}
