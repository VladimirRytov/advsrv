package front

import (
	"bytes"
	"net/url"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/encodedecoder"
	"github.com/VladimirRytov/advsrv/internal/front/requests"

	"github.com/gofiber/fiber/v2"
)

func (r *Reciever) ExtraCharges(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	ExtraCharges, err := r.requestGate.ExtraChargesGetRequest(c.Context(), authHeader)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&ExtraCharges)
}

func (r *Reciever) ExtraChargeByName(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	name, err := url.QueryUnescape(c.Params("name"))
	if err != nil {
		return err
	}

	ExtraCharge, err := r.requestGate.ExtraChargeGetRequest(c.Context(), authHeader, name)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&ExtraCharge)
}

func (r *Reciever) NewExtraCharge(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	var ExtraCharge datatransferobjects.ExtraChargeDTO
	b := bytes.NewBuffer(c.Body())

	err = encodedecoder.FromJSON(&ExtraCharge, b)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	}

	err = r.requestGate.ExtraChargePostRequest(c.Context(), authHeader, &ExtraCharge)
	if err != nil {
		return r.handleError(c, err)
	}
	c.SendStatus(fiber.StatusCreated)
	return c.JSON(&ExtraCharge)
}

func (r *Reciever) UpdateExtraCharge(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	var extraCharge datatransferobjects.ExtraChargeDTO
	b := bytes.NewBuffer(c.Body())

	err = encodedecoder.FromJSON(&extraCharge, b)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}
	extraCharge.ChargeName, err = url.QueryUnescape(c.Params("name"))
	if err != nil {
		return err
	}
	err = r.requestGate.ExtraChargePutRequest(c.Context(), authHeader, &extraCharge)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&extraCharge)
}

func (r *Reciever) RemoveExtraCharge(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	name, err := url.QueryUnescape(c.Params("name"))
	if err != nil {
		return err
	}

	err = r.requestGate.ExtraChargeDeleteRequest(c.Context(), authHeader, name)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
