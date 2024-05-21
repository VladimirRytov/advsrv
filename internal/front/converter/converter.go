package converter

import (
	"bytes"

	"github.com/VladimirRytov/advsrv/internal/encodedecoder"
)

func ConvertToJsonWithSubQuery(readable bool, entry any) ([]byte, error) {
	var (
		err error
		b   bytes.Buffer
	)
	err = encodedecoder.ToJSON(&b, entry, readable)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), err
}
