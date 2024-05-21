package front

import (
	"bytes"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/encodedecoder"
	"github.com/VladimirRytov/advsrv/internal/front/requests"

	"github.com/gofiber/fiber/v2"
)

func (r *Reciever) LineAdvertisements(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	params, err := r.requestGate.CheckAdvertisementQueries(c.Context(), c.Queries())
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}

	if params.Calculate {
		var lineAdvertisement datatransferobjects.LineAdvertisementDTO
		b := bytes.NewBuffer(c.Body())

		err = encodedecoder.FromJSON(&lineAdvertisement, b)
		if err != nil {
			return r.handleError(c, err)
		}
		calculatedLine, err := r.requestGate.CalculateLineAdvertisementCost(c.Context(), authHeader, params, &lineAdvertisement)
		if err != nil {
			return r.handleError(c, err)
		}
		return c.JSON(calculatedLine)
	}

	cli, err := r.requestGate.LineAdvertisementsGetRequest(c.Context(), authHeader, params)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&cli)
}

func (r *Reciever) LineAdvertisement(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	if len(c.Queries()) > 0 {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}
	cli, err := r.requestGate.LineAdvertisementGetRequest(c.Context(), authHeader, id)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&cli)
}

func (r *Reciever) NewLineAdvertisement(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	if len(c.Queries()) > 0 {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}

	var LineAdvertisement datatransferobjects.LineAdvertisementDTO
	b := bytes.NewBuffer(c.Body())

	err = encodedecoder.FromJSON(&LineAdvertisement, b)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	}

	err = r.requestGate.LineAdvertisementPostRequest(c.Context(), authHeader, &LineAdvertisement)
	if err != nil {
		return r.handleError(c, err)
	}
	c.SendStatus(fiber.StatusCreated)
	return c.JSON(&LineAdvertisement)
}

func (r *Reciever) UpdateLineAdvertisement(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	if len(c.Queries()) > 0 {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}

	var lineAdvertisement datatransferobjects.LineAdvertisementDTO
	b := bytes.NewBuffer(c.Body())
	err = encodedecoder.FromJSON(&lineAdvertisement, b)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	}

	lineAdvertisement.ID, err = c.ParamsInt("id")
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}

	err = r.requestGate.LineAdvertisementPutRequest(c.Context(), authHeader, &lineAdvertisement)
	if err != nil {
		return r.handleError(c, err)
	}
	c.SendStatus(fiber.StatusCreated)
	return c.JSON(&lineAdvertisement)
}

func (r *Reciever) RemoveLineAdvertisement(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	if len(c.Queries()) > 0 {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}

	err = r.requestGate.LineAdvertisementDeleteRequest(c.Context(), authHeader, id)
	if err != nil {
		return r.handleError(c, err)
	}
	c.SendStatus(fiber.StatusNoContent)
	return err
}
