package requests

import (
	"context"
	"strconv"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func (rh *Requesting) CheckFileQueries(ctx context.Context, params map[string]string) (datatransferobjects.FileParams, error) {
	var (
		fileParam datatransferobjects.FileParams
		err       error
	)
	val, ok := params["miniatures"]
	if ok {
		fileParam.Miniatures, err = strconv.ParseBool(val)
		if err != nil {
			return fileParam, ErrQuery
		}
	}
	return fileParam, nil
}
