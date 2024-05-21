package front

import (
	"bytes"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/encodedecoder"
	"github.com/VladimirRytov/advsrv/internal/front/requests"

	"github.com/gofiber/fiber/v2"
)

func (r *Reciever) BlockAdvertisements(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	params, err := r.requestGate.CheckAdvertisementQueries(c.Context(), c.Queries())
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	}

	if params.Calculate {
		var blockAdvertisement datatransferobjects.BlockAdvertisementDTO
		b := bytes.NewBuffer(c.Body())

		err = encodedecoder.FromJSON(&blockAdvertisement, b)
		if err != nil {
			return r.handleError(c, err)
		}
		calculatedBlock, err := r.requestGate.CalculateBlockAdvertisementCost(c.Context(), authHeader, params, &blockAdvertisement)
		if err != nil {
			return r.handleError(c, err)
		}
		return c.JSON(calculatedBlock)
	}

	cli, err := r.requestGate.BlockAdvertisementsGetRequest(c.Context(), authHeader, params)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&cli)
}

func (r *Reciever) BlockAdvertisement(c *fiber.Ctx) error {
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
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	}

	cli, err := r.requestGate.BlockAdvertisementGetRequest(c.Context(), authHeader, id)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&cli)
}

func (r *Reciever) NewBlockAdvertisement(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	if len(c.Queries()) > 0 {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}

	var BlockAdvertisement datatransferobjects.BlockAdvertisementDTO
	b := bytes.NewBuffer(c.Body())

	err = encodedecoder.FromJSON(&BlockAdvertisement, b)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	}

	err = r.requestGate.BlockAdvertisementPostRequest(c.Context(), authHeader, &BlockAdvertisement)
	if err != nil {
		return r.handleError(c, err)
	}

	c.SendStatus(fiber.StatusCreated)
	return c.JSON(&BlockAdvertisement)
}

func (r *Reciever) UpdateBlockAdvertisement(c *fiber.Ctx) error {
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

	var BlockAdvertisement datatransferobjects.BlockAdvertisementDTO
	b := bytes.NewBuffer(c.Body())
	err = encodedecoder.FromJSON(&BlockAdvertisement, b)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	}

	BlockAdvertisement.ID = id
	err = r.requestGate.BlockAdvertisementPutRequest(c.Context(), authHeader, &BlockAdvertisement)
	if err != nil {
		return r.handleError(c, err)
	}
	c.SendStatus(fiber.StatusCreated)
	return c.JSON(&BlockAdvertisement)
}

func (r *Reciever) RemoveBlockAdvertisement(c *fiber.Ctx) error {
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

	err = r.requestGate.BlockAdvertisementDeleteRequest(c.Context(), authHeader, id)
	if err != nil {
		return r.handleError(c, err)
	}
	c.SendStatus(fiber.StatusNoContent)
	return err
}
