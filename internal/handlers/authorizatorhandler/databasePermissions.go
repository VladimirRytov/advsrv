package authorizatorhandler

import (
	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func (au *Authorizator) CanReadDatabase(user *datatransferobjects.UserToken) error {
	if au.auth.CanReadDatabases(user.Perm) {
		return nil
	}
	return ErrAuthrization
}

func (au *Authorizator) CanWriteDatabase(user *datatransferobjects.UserToken) error {
	if au.auth.CanWriteDatabases(user.Perm) {
		return nil
	}
	return ErrAuthrization
}

func (au *Authorizator) CanCreateDatabase(user *datatransferobjects.UserToken) error {
	if au.auth.CanCreateDatabases(user.Perm) {
		return nil
	}
	return ErrAuthrization
}

func (au *Authorizator) CanDeleteDatabase(user *datatransferobjects.UserToken) error {
	if au.auth.CanRemoveDatabases(user.Perm) {
		return nil
	}
	return ErrAuthrization
}
