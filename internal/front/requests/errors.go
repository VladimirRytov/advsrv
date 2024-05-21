package requests

import (
	"errors"

	"github.com/VladimirRytov/advsrv/internal/handlers/advertisementhandler"
	"github.com/VladimirRytov/advsrv/internal/handlers/userhandler"
	"github.com/VladimirRytov/advsrv/internal/validator"
)

var (
	ErrQuery      = errQuery()
	ErrClientSide = errClientSide()
	ErrNotFound   = advertisementhandler.ErrNotFound
	ErrDatabase   = advertisementhandler.ErrDatabase

	ErrWrongLoginPassword = userhandler.ErrWrongData
	ErrNeedBasicMethod    = errWrongAuthMethod()
	ErrNeedBearerMethod   = errWrongAuthMethod()

	ErrAuthorizeParams = errWrongAuthParams()
	ErrValidate        = validator.ErrValidate
)

func errClientSide() error {
	return errors.Join(advertisementhandler.ErrConvertation, advertisementhandler.ErrDuplicate, advertisementhandler.ErrViolatesForeignKey, ErrQuery)
}

func errQuery() error {
	return errors.New("неверные параметры запроса")
}

func errWrongAuthMethod() error {
	return errors.New("неверный метод аутентификации")
}

func errWrongAuthParams() error {
	return errors.New("некорректный параметр Authorize")
}
