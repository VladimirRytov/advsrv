package requests

import (
	"context"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func (rh *Requesting) UserGetRequest(ctx context.Context, token string, name string) (datatransferobjects.UserDTO, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return datatransferobjects.UserDTO{}, err
	}
	err = rh.authorizator.CanReadUsers(&user)
	if err != nil {
		return datatransferobjects.UserDTO{}, err
	}
	return rh.userRepo.UserByName(ctx, name)
}

func (rh *Requesting) UsersGetRequest(ctx context.Context, token string) ([]datatransferobjects.UserDTO, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return nil, err
	}
	err = rh.authorizator.CanReadUsers(&user)
	if err != nil {
		return nil, err
	}
	return rh.userRepo.UserList(ctx)
}

func (rh *Requesting) UserPostRequest(ctx context.Context, token string, user *datatransferobjects.UserDTO) (datatransferobjects.UserDTO, error) {
	userToken, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return datatransferobjects.UserDTO{}, err
	}
	err = rh.authorizator.CanCreateUsers(&userToken)
	if err != nil {
		return datatransferobjects.UserDTO{}, err
	}
	return rh.userRepo.NewUser(ctx, user)
}

func (rh *Requesting) UserPutRequest(ctx context.Context, token string, user *datatransferobjects.UserDTO) (datatransferobjects.UserDTO, error) {
	userToken, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return datatransferobjects.UserDTO{}, err
	}
	err = rh.authorizator.CanWriteUsers(&userToken)
	if err != nil {
		return datatransferobjects.UserDTO{}, err
	}

	return rh.userRepo.UpdateUser(ctx, user)
}

func (rh *Requesting) UserDeleteRequest(ctx context.Context, token string, name string) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanDeleteUsers(&user)
	if err != nil {
		return err
	}
	return rh.userRepo.RemoveUser(ctx, name)
}
