package configparser

import (
	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/mapper"
)

func (c *ConfigHandler) userToDTO(user User) (datatransferobjects.UserDTO, error) {
	userPermissions, err := c.parseUserPermission(user.Permissions)
	if err != nil {
		return datatransferobjects.UserDTO{}, err
	}
	userDto := datatransferobjects.UserDTO{
		Name:        user.Name,
		Password:    []byte(user.Password),
		Permissions: userPermissions,
	}
	userAdv, err := mapper.DtoToUser(&userDto, !user.Salted)
	if err != nil {
		return datatransferobjects.UserDTO{}, err
	}
	return mapper.UserToDTO(&userAdv, false), nil
}
