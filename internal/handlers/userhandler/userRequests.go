package userhandler

import (
	"context"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/mapper"
)

func (a *UserHandler) NewUser(ctx context.Context, user *datatransferobjects.UserDTO) (datatransferobjects.UserDTO, error) {
	userAdv, err := mapper.DtoToUser(user, true)
	if err != nil {
		return datatransferobjects.UserDTO{}, err
	}
	userDto := mapper.UserToDTO(&userAdv, true)
	_, err = a.userRepo.NewUser(ctx, &userDto)
	return userDto, err
}

func (a *UserHandler) UserList(ctx context.Context) ([]datatransferobjects.UserDTO, error) {
	users, err := a.userRepo.AllUsersHidePassword(ctx)
	if err != nil {
		return nil, err
	}
	convertedUsers := make([]datatransferobjects.UserDTO, 0, len(users))
	for i := range users {
		userAdv, err := mapper.DtoToUser(&users[i], false)
		if err != nil {
			continue
		}
		convertedUsers = append(convertedUsers, mapper.UserToDTO(&userAdv, true))
	}
	return convertedUsers, nil
}

func (a *UserHandler) UserByName(ctx context.Context, name string) (datatransferobjects.UserDTO, error) {
	user, err := a.userRepo.UserByNameHidePassword(ctx, name)
	if err != nil {
		return datatransferobjects.UserDTO{}, err
	}
	userAdv, err := mapper.DtoToUser(&user, false)
	if err != nil {
		return datatransferobjects.UserDTO{}, err
	}

	return mapper.UserToDTO(&userAdv, true), nil
}

func (a *UserHandler) UpdateUser(ctx context.Context, user *datatransferobjects.UserDTO) (datatransferobjects.UserDTO, error) {
	userAdv, err := mapper.DtoToUser(user, true)
	if err != nil {
		return datatransferobjects.UserDTO{}, err
	}
	userDto := mapper.UserToDTO(&userAdv, true)
	return userDto, a.userRepo.UpdateUser(ctx, &userDto)
}

func (a *UserHandler) RemoveUser(ctx context.Context, name string) error {
	if name == "admin" {
		return ErrWrondRequest
	}
	return a.userRepo.RemoveUser(ctx, name)
}
