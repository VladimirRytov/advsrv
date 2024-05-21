package front

import (
	"errors"

	"github.com/VladimirRytov/advsrv/internal/front/requests"

	"github.com/gofiber/fiber/v2"
)

func (r *Reciever) handleError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(requests.ErrNeedBasicMethod, err):
		return r.needBearerAuth(c)
	case errors.Is(requests.ErrValidate, err):
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	case errors.Is(requests.ErrNotFound, err):
		c.SendStatus(fiber.StatusNotFound)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusNotFound, err.Error()))
	case errors.Is(requests.ErrClientSide, err):
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	default:
		c.SendStatus(fiber.ErrBadGateway.Code)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadGateway, err.Error()))
	}
}
