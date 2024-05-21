package authorizatorhandler

import "errors"

type Authorizator struct {
	auth PermHandler
}

var ErrAuthrization = errAuthorization()

func errAuthorization() error {
	return errors.New("недостаточно прав")
}
func NewAuthrizator(prmHandler PermHandler) *Authorizator {
	return &Authorizator{auth: prmHandler}
}
