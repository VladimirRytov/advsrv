package userhandler

import "github.com/VladimirRytov/advsrv/internal/datatransferobjects"

type Authorizator interface {
	AdminPermissions() int32
}

type Validator interface {
	Validate([]byte) error
	SignDate(name string, perm int32) ([]byte, error)
	FetchPayload([]byte) (datatransferobjects.UserToken, error)
}
