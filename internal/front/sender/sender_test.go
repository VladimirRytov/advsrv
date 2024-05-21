package sender

import (
	"slices"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
)

func reply(c *fiber.Ctx) error {
	if slices.Equal(c.Body(), []byte("Hello")) {
		c.SendStatus(fiber.StatusCreated)
		return c.Send([]byte("ok"))
	}
	c.SendStatus(fiber.ErrBadRequest.Code)
	return fiber.ErrBadRequest
}
func TestRegistration(t *testing.T) {
	a := fiber.New()
	a.Post("/auth/subscribe", reply)
	go a.Listen("127.0.0.1:4567")
	time.Sleep(2 * time.Second)
	sender := Requester{URL: "http://127.0.0.1:4567/auth/subscribe"}
	_, err := sender.Write([]byte("Hello"))
	if err != nil {
		t.Fatal(err)
	}
}
