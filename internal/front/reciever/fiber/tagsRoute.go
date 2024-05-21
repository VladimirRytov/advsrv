package front

import (
	"bytes"
	"net/url"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/encodedecoder"

	"github.com/gofiber/fiber/v2"
)

func (r *Reciever) Tags(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	tags, err := r.requestGate.TagsGetRequest(c.Context(), authHeader)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&tags)
}

func (r *Reciever) TagByName(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	name, err := url.QueryUnescape(c.Params("name"))
	if err != nil {
		return err
	}

	tag, err := r.requestGate.TagGetRequest(c.Context(), authHeader, name)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&tag)
}

func (r *Reciever) NewTag(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	var tag datatransferobjects.TagDTO
	b := bytes.NewBuffer(c.Body())

	err = encodedecoder.FromJSON(&tag, b)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	}

	err = r.requestGate.TagPostRequest(c.Context(), authHeader, &tag)
	if err != nil {
		return r.handleError(c, err)
	}
	c.SendStatus(fiber.StatusCreated)
	return c.JSON(&tag)
}

func (r *Reciever) UpdateTag(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	var tag datatransferobjects.TagDTO
	b := bytes.NewBuffer(c.Body())

	err = encodedecoder.FromJSON(&tag, b)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusBadRequest, err.Error()))
	}
	tag.TagName, err = url.QueryUnescape(c.Params("name"))
	if err != nil {
		return err
	}
	err = r.requestGate.TagPutRequest(c.Context(), authHeader, &tag)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.JSON(&tag)
}

func (r *Reciever) RemoveTag(c *fiber.Ctx) error {
	authHeader, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	parsedName, err := url.QueryUnescape(c.Params("name"))
	if err != nil {
		return err
	}

	err = r.requestGate.TagDeleteRequest(c.Context(), authHeader, parsedName)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
