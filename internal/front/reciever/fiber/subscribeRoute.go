package front

import (
	"bytes"
	"context"
	"net/url"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/encodedecoder"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func (r *Reciever) Subscriber(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}
	id, err := r.requestGate.SubscribeGetRequest(c.Context(), authHeader, c.Params("id"), c.QueryBool("ping"))
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&id)
}

func (r *Reciever) Subscribers(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}
	subs, err := r.requestGate.SubscribersGetRequest(c.Context(), authHeader)
	if err != nil {
		return r.handleError(c, err)
	}

	return c.JSON(&subs)
}

func (r *Reciever) Subscribe(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}
	var sub datatransferobjects.SubscribeParams
	b := bytes.NewBuffer(c.Body())

	err = encodedecoder.FromJSON(&sub, b)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	}
	id, err := r.requestGate.SubscribePostRequest(c.Context(), authHeader, sub)
	if err != nil {
		return r.handleError(c, err)
	}
	sub.UserID = id
	c.SendStatus(fiber.StatusCreated)
	return c.JSON(&sub)
}

func (r *Reciever) WebSock(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		header, err := r.fetchHeaderAuthorization(c)
		if err != nil {
			c.SendStatus(fiber.StatusForbidden)
			return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
		}
		_, err = r.requestGate.CheckToken(c.Context(), header)
		if err != nil {
			c.SendStatus(fiber.StatusForbidden)
			return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
		}
		c.Locals("allowed", true)
		return c.Next()
	}

	return fiber.ErrUpgradeRequired
}

func (r *Reciever) ActiveSubscribe(c *websocket.Conn) {
	defer c.Close()
	c.RemoteAddr()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	b, err := r.requestGate.SubscribeActiveRequest(ctx, c.Headers("Authorization"), c.RemoteAddr().String())
	if err != nil {
		return
	}

	go readMessage(c)
	for v := range b {
		w, err := c.NextWriter(websocket.TextMessage)
		if err != nil {
			w.Close()
			break
		}
		_, err = w.Write(v)
		if err != nil {

			break
		}
		w.Close()
	}
}

func readMessage(c *websocket.Conn) {
	for {
		mt, _, err := c.NextReader()
		if err != nil {
			c.WriteMessage(websocket.CloseMessage, []byte{})
			c.Close()
			return
		}
		switch mt {
		case websocket.CloseNormalClosure:
			c.Close()
			return
		}
	}
}

func (r *Reciever) CancelSubscribe(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}
	name, err := url.QueryUnescape(c.Params("name"))
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	}
	err = r.requestGate.SubscribeDeleteRequest(c.Context(), authHeader, name)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
