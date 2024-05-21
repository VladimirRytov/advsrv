package front

import (
	"bytes"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/encodedecoder"
	"github.com/VladimirRytov/advsrv/internal/front/requests"

	"github.com/gofiber/fiber/v2"
)

func (r *Reciever) Orders(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	orderparams, err := r.requestGate.CheckGetSeveralOrdersQueries(c.Context(), c.Queries())
	if orderparams.Calculate {
		var order datatransferobjects.OrderDTO
		b := bytes.NewBuffer(c.Body())

		err = encodedecoder.FromJSON(&order, b)
		if err != nil {
			c.SendStatus(fiber.StatusBadRequest)
			return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
		}
		calculatedOrder, err := r.requestGate.CalculateOrderCost(c.Context(), authHeader, orderparams, &order)
		if err != nil {
			return r.handleError(c, err)
		}
		return c.JSON(calculatedOrder)
	}

	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}
	orders, err := r.requestGate.OrdersGetRequest(c.Context(), authHeader, orderparams)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&orders)
}

func (r *Reciever) Order(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	orderparams, err := r.requestGate.CheckOrderQueries(c.Context(), c.Queries())
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}

	order, err := r.requestGate.OrderGetRequest(c.Context(), authHeader, orderparams, id)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&order)
}

func (r *Reciever) NewOrder(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	orderparams, err := r.requestGate.CheckOrderQueries(c.Context(), c.Queries())
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}

	var order datatransferobjects.OrderDTO
	b := bytes.NewBuffer(c.Body())

	err = encodedecoder.FromJSON(&order, b)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	}

	err = r.requestGate.OrderPostRequest(c.Context(), authHeader, orderparams, &order)
	if err != nil {
		return r.handleError(c, err)
	}

	c.SendStatus(fiber.StatusCreated)
	return c.JSON(&order)
}

func (r *Reciever) UpdateOrder(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	if len(c.Queries()) > 0 {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}

	var order datatransferobjects.OrderDTO
	b := bytes.NewBuffer(c.Body())
	err = encodedecoder.FromJSON(&order, b)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}
	order.ID, err = c.ParamsInt("id")
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, requests.ErrQuery.Error()))
	}
	err = r.requestGate.OrderPutRequest(c.Context(), authHeader, &order)
	if err != nil {
		return r.handleError(c, err)
	}

	return c.JSON(&order)
}

func (r *Reciever) DeleteOrder(c *fiber.Ctx) error {
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

	err = r.requestGate.OrderDeleteRequest(c.Context(), authHeader, id)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
