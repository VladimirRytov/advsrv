package front

import (
	"bytes"
	"net/url"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/encodedecoder"

	"github.com/gofiber/fiber/v2"
)

func (r *Reciever) CostRates(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	CostRates, err := r.requestGate.CostRatesGetRequest(c.Context(), authHeader)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&CostRates)
}

func (r *Reciever) CostRateByName(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	parsedName, err := url.QueryUnescape(c.Params("name"))
	if err != nil {
		return err
	}

	CostRate, err := r.requestGate.CostRateGetRequest(c.Context(), authHeader, parsedName)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&CostRate)
}

func (r *Reciever) NewCostRate(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	var CostRate datatransferobjects.CostRateDTO
	b := bytes.NewBuffer(c.Body())

	err = encodedecoder.FromJSON(&CostRate, b)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	}

	err = r.requestGate.CostRatePostRequest(c.Context(), authHeader, &CostRate)
	if err != nil {
		return r.handleError(c, err)
	}
	c.SendStatus(fiber.StatusCreated)
	return c.JSON(&CostRate)
}

func (r *Reciever) UpdateCostRate(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	var CostRate datatransferobjects.CostRateDTO
	b := bytes.NewBuffer(c.Body())

	err = encodedecoder.FromJSON(&CostRate, b)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	}
	CostRate.Name, err = url.QueryUnescape(c.Params("name"))
	if err != nil {
		return err
	}

	err = r.requestGate.CostRatePutRequest(c.Context(), authHeader, &CostRate)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&CostRate)
}

func (r *Reciever) RemoveCostRate(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	parsedName, err := url.QueryUnescape(c.Params("name"))
	if err != nil {
		return err
	}

	err = r.requestGate.CostRateDeleteRequest(c.Context(), authHeader, parsedName)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
