package authorizatorhandler

import (
	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func (au *Authorizator) CanReadUsers(user *datatransferobjects.UserToken) error {
	if au.auth.CanReadUsers(user.Perm) {
		return nil
	}
	return ErrAuthrization
}

func (au *Authorizator) CanWriteUsers(user *datatransferobjects.UserToken) error {
	if au.auth.CanWriteUsers(user.Perm) {
		return nil
	}
	return ErrAuthrization
}

func (au *Authorizator) CanCreateUsers(user *datatransferobjects.UserToken) error {
	if au.auth.CanCreateUsers(user.Perm) {
		return nil
	}
	return ErrAuthrization
}

func (au *Authorizator) CanDeleteUsers(user *datatransferobjects.UserToken) error {
	if au.auth.CanRemoveUsers(user.Perm) {
		return nil
	}
	return ErrAuthrization
}
