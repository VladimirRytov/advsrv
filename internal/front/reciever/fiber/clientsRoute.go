package front

import (
	"bytes"
	"net/url"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/encodedecoder"
	"github.com/VladimirRytov/advsrv/internal/front/requests"

	"github.com/gofiber/fiber/v2"
)

func (r *Reciever) Clients(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	params, err := r.requestGate.CheckClientQueries(c.Context(), c.Queries())
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}

	cli, err := r.requestGate.ClientsGetRequest(c.Context(), authHeader, params)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&cli)
}

func (r *Reciever) Client(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	params, err := r.requestGate.CheckClientQueries(c.Context(), c.Queries())
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}
	clientName, err := url.QueryUnescape(c.Params("name"))
	if err != nil {
		return err
	}

	cli, err := r.requestGate.ClientGetRequest(c.Context(), authHeader, params, clientName)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&cli)
}

func (r *Reciever) NewClient(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	params, err := r.requestGate.CheckClientQueries(c.Context(), c.Queries())
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}

	var client datatransferobjects.ClientDTO
	b := bytes.NewBuffer(c.Body())

	err = encodedecoder.FromJSON(&client, b)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	}
	err = r.requestGate.ClientPostRequest(c.Context(), authHeader, params, &client)
	if err != nil {
		return r.handleError(c, err)
	}
	c.SendStatus(fiber.StatusCreated)
	return c.JSON(&client)
}

func (r *Reciever) UpdateClient(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	if len(c.Queries()) > 0 {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}

	var client datatransferobjects.ClientDTO
	b := bytes.NewBuffer(c.Body())

	err = encodedecoder.FromJSON(&client, b)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	}

	client.Name, err = url.QueryUnescape(c.Params("name"))
	if err != nil {
		return err
	}
	err = r.requestGate.ClientPutRequest(c.Context(), authHeader, &client)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&client)
}

func (r *Reciever) RemoveClient(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	if len(c.Queries()) > 0 {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}

	name, err := url.QueryUnescape(c.Params("name"))
	if err != nil {
		return err
	}

	err = r.requestGate.ClientDeleteRequest(c.Context(), authHeader, name)
	if err != nil {
		return r.handleError(c, err)
	}
	c.SendStatus(fiber.StatusNoContent)
	return err
}
