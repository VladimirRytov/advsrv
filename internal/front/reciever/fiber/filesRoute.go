package front

import (
	"net/url"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/front/converter"

	"github.com/gofiber/fiber/v2"
)

func (r *Reciever) FileByName(c *fiber.Ctx) error {
	token, err := r.fetchTokenAuthorization(c)
	if err != nil {
		token, err = r.fetchHeaderAuthorization(c)
		if err != nil {
			c.SendStatus(fiber.StatusForbidden)
			return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
		}
	}

	parsedName, err := url.QueryUnescape(c.Params("name"))
	if err != nil {
		return err
	}
	decoded, err := r.b64.FromBase64URLString(parsedName)
	if err == nil {
		parsedName = string(decoded)
	}
	params := c.Queries()
	format := params["format"]
	switch format {
	case "json":
		file, err := r.requestGate.FileGetFormatedRequest(c.Context(), token, parsedName, params["size"])
		if err != nil {
			return r.handleError(c, err)
		}
		return c.JSON(r.frontConv.FileToFront(&file))
	}
	fp, err := r.requestGate.FileGetRequest(c.Context(), token, parsedName)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.SendFile(fp)
}

func (r *Reciever) Files(c *fiber.Ctx) error {
	token, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}

	files, err := r.requestGate.FilesGetRequest(c.Context(), token, c.Queries())
	if err != nil {
		return r.handleError(c, err)
	}
	encFiles := make([]converter.FileFront, 0, len(files))
	for i := range files {
		encFiles = append(encFiles, r.frontConv.FileToFront(&files[i]))
	}
	return c.JSON(encFiles)
}

func (r *Reciever) UploadFile(c *fiber.Ctx) error {
	token, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}
	form, err := c.MultipartForm()
	if err != nil {
		return r.handleError(c, err)
	}
	encFiles := make([]converter.FileFront, 0)
	for _, v := range form.File {
		for i := range v {
			f, err := v[i].Open()
			if err != nil {
				return r.handleError(c, err)
			}
			fileName, err := r.requestGate.FilePostRequest(c.Context(), token, c.Queries(), v[i].Filename, f)
			if err != nil {
				return r.handleError(c, err)
			}
			encFiles = append(encFiles, r.frontConv.FileToFront(&datatransferobjects.File{Name: fileName}))
		}
	}
	c.SendStatus(fiber.StatusCreated)
	return c.JSON(encFiles)
}

func (r *Reciever) RemoveFile(c *fiber.Ctx) error {
	token, err := r.fetchHeaderAuthorization(c)
	if err != nil {
		c.SendStatus(fiber.StatusForbidden)
		return c.JSON(r.frontConv.NewResponceMessage(fiber.StatusForbidden, err.Error()))
	}
	parsedName, err := url.QueryUnescape(c.Params("name"))
	if err != nil {
		return err
	}
	err = r.requestGate.FileDeleteRequest(c.Context(), token, parsedName)
	if err != nil {
		return r.handleError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
