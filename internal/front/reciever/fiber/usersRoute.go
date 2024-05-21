package front

import (
	"bytes"
	"net/url"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/encodedecoder"

	"github.com/gofiber/fiber/v2"
)

func (r *Reciever) UserByName(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	name, err := url.QueryUnescape(c.Params("name"))
	if err != nil {
		return err
	}

	user, err := r.requestGate.UserGetRequest(c.Context(), authHeader, name)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&user)
}

func (r *Reciever) Users(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	users, err := r.requestGate.UsersGetRequest(c.Context(), authHeader)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&users)
}

func (r *Reciever) NewUser(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	var user datatransferobjects.UserDTO
	b := bytes.NewBuffer(c.Body())

	err = encodedecoder.FromJSON(&user, b)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	}

	newUser, err := r.requestGate.UserPostRequest(c.Context(), authHeader, &user)
	if err != nil {
		return r.handleError(c, err)
	}
	c.SendStatus(fiber.StatusCreated)
	return c.JSON(&newUser)
}

func (r *Reciever) UpdateUser(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	var user datatransferobjects.UserDTO
	b := bytes.NewBuffer(c.Body())

	err = encodedecoder.FromJSON(&user, b)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	}

	user.Name, err = url.QueryUnescape(c.Params("name"))
	if err != nil {
		return err
	}

	newUser, err := r.requestGate.UserPutRequest(c.Context(), authHeader, &user)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&newUser)
}

func (r *Reciever) DeleteUser(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	name, err := url.QueryUnescape(c.Params("name"))
	if err != nil {
		return err
	}

	err = r.requestGate.UserDeleteRequest(c.Context(), authHeader, name)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)

}
