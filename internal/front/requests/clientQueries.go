package requests

import (
	"context"
	"errors"
	"strconv"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func (rh *Requesting) CheckClientQueries(ctx context.Context, params map[string]string) (datatransferobjects.ClientParams, error) {
	var clientParams datatransferobjects.ClientParams
	for k := range params {
		switch k {
		case "nested":
			parsed, err := strconv.ParseBool(params[k])
			if err != nil {
				return clientParams, errors.New("неверный параметр атрибута nested")
			}
			clientParams.Nested = parsed
		default:
			return clientParams, errors.New("неверный запрос")

		}
	}
	return clientParams, nil
}
