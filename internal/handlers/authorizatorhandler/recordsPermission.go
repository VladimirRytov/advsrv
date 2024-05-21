package authorizatorhandler

import (
	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func (au *Authorizator) CanReadRecords(user *datatransferobjects.UserToken) error {
	if au.auth.CanReadRecords(user.Perm) {
		return nil
	}
	return ErrAuthrization
}

func (au *Authorizator) CanWriteRecords(user *datatransferobjects.UserToken) error {
	if au.auth.CanWriteRecords(user.Perm) {
		return nil
	}
	return ErrAuthrization
}

func (au *Authorizator) CanCreateRecords(user *datatransferobjects.UserToken) error {
	if au.auth.CanCreateRecords(user.Perm) {
		return nil
	}
	return ErrAuthrization
}

func (au *Authorizator) CanDeleteRecords(user *datatransferobjects.UserToken) error {
	if au.auth.CanRemoveRecords(user.Perm) {
		return nil
	}
	return ErrAuthrization
}
