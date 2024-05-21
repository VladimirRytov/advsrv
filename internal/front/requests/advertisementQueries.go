package requests

import (
	"context"
	"errors"
	"regexp"
	"strconv"
	"time"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

var betweebRegex = regexp.MustCompile(`(\d{2}\.\d{2}\.\d{4})-(\d{2}\.\d{2}\.\d{4})$`)

func (rh *Requesting) CheckAdvertisementQueries(ctx context.Context, params map[string]string) (datatransferobjects.AdvertisementParams, error) {
	var err error
	var advertisementParams datatransferobjects.AdvertisementParams

	for k := range params {
		switch k {
		case "calculatecost":
			parsed, err := strconv.ParseBool(params[k])

			if err != nil {
				return advertisementParams, errors.New("неверное значение параметра nested")
			}
			costRate, ok := params["costrate"]
			if !ok {
				return advertisementParams, errors.New("сostrate должен быть указан")
			}
			advertisementParams.Calculate = parsed
			advertisementParams.Costrate = costRate
			return advertisementParams, nil

		case "fromdate":
			advertisementParams.FromDate, err = time.Parse("02.01.2006", params[k])
			if err != nil {
				return advertisementParams, ErrQuery
			}
		case "between":
			dates := betweebRegex.FindStringSubmatch(params[k])
			if len(dates) < 3 {
				return advertisementParams, ErrQuery
			}

			advertisementParams.FromDate, err = time.Parse("02.01.2006", dates[1])
			if err != nil {
				return advertisementParams, ErrQuery
			}

			advertisementParams.ToDate, err = time.Parse("02.01.2006", dates[2])
			if err != nil {
				return advertisementParams, ErrQuery
			}
		case "actual":
			advertisementParams.Actual, err = strconv.ParseBool(params[k])
			if err != nil {
				return advertisementParams, ErrQuery
			}
		case "orderid":
			advertisementParams.OrderID, err = strconv.Atoi(params[k])
			if err != nil {
				return advertisementParams, ErrQuery
			}
		default:
			return advertisementParams, ErrQuery
		}
	}
	return advertisementParams, nil
}
