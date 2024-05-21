package front

import (
	"github.com/VladimirRytov/advsrv/internal/front/converter"

	"github.com/gofiber/fiber/v2"
)

func (r *Reciever) AuthUser(c *fiber.Ctx) error {
	auth, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		return r.needBasicAuth(c)
	}
	token, err := r.requestGate.AuthenticateUser(c.Context(), auth)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(converter.Token{AccessToken: string(token)})

}

func (r *Reciever) needBasicAuth(c *fiber.Ctx) error {
	c.Append("WWW-Authenticate", "Basic")
	return c.SendStatus(fiber.StatusUnauthorized)
}

func (r *Reciever) needBearerAuth(c *fiber.Ctx) error {
	c.Append("WWW-Authenticate", "Bearer")
	return c.SendStatus(fiber.StatusUnauthorized)
}

func (r *Reciever) fetchHeaderAuthorization(c *fiber.Ctx) (string, error) {
	head := c.GetReqHeaders()
	v, ok := head["Authorization"]
	if !ok || len(v) == 0 {
		return "", fiber.ErrForbidden
	}
	return v[0], nil
}

func (r *Reciever) fetchTokenAuthorization(c *fiber.Ctx) (string, error) {
	query := c.Queries()
	token, ok := query["token"]
	if !ok {
		return "", fiber.ErrForbidden
	}
	return token, nil
}
