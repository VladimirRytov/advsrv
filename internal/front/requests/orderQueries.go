package requests

import (
	"context"
	"errors"
	"strconv"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func (rh *Requesting) CheckGetSeveralOrdersQueries(ctx context.Context, params map[string]string) (datatransferobjects.OrderParams, error) {
	var orderParams datatransferobjects.OrderParams

	for k := range params {
		switch k {
		case "calculatecost":
			parsed, err := strconv.ParseBool(params[k])
			if err != nil {
				return orderParams, errors.New("неверное значение параметра nested")
			}
			costRate, ok := params["costrate"]
			if !ok {
				return orderParams, errors.New("сostrate должен быть указан")
			}
			orderParams.Calculate = parsed
			orderParams.Costrate = costRate
			return orderParams, nil

		case "nested":
			parsed, err := strconv.ParseBool(params[k])
			if err != nil {
				return orderParams, errors.New("неверное значение параметра nested")
			}
			orderParams.Nested = parsed
		case "client":
			orderParams.Client = params[k]
		default:
			return orderParams, errors.New("неверный запрос")

		}
	}
	return orderParams, nil
}

func (rh *Requesting) CheckOrderQueries(ctx context.Context, params map[string]string) (datatransferobjects.OrderParams, error) {
	var orderParams datatransferobjects.OrderParams
	for k := range params {
		switch k {
		case "nested":
			parsed, err := strconv.ParseBool(params[k])
			if err != nil {
				return orderParams, errors.New("неверное значение атрибута nested")
			}
			orderParams.Nested = parsed
		default:
			return orderParams, errors.New("неверный запрос")

		}
	}
	return orderParams, nil
}
