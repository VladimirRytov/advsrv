package sender

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

type SenderMaker struct{}

func (sm *SenderMaker) NewSender(uri string) (io.Writer, error) {
	return &Requester{URL: uri}, nil
}

type Requester struct {
	URL string
}

func (r *Requester) Write(data []byte) (int, error) {
	req := bytes.NewReader(data)
	resp, err := http.Post(r.URL, "asd", req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		return 0, errors.New("got wrong status code")
	}

	return len(data), err
}

func (r *Requester) SendTo() string {
	return r.URL
}
