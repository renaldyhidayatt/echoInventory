package handler

import (
	"echoinventory/entity"
	"echoinventory/schemas"

	"github.com/labstack/echo/v4"
)

type handlerCategory struct {
	category entity.EntityCategory
}

func NewHandlerCategory(category entity.EntityCategory) *handlerCategory {
	return &handlerCategory{
		category: category,
	}
}

func (h *handlerCategory) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello World")
}

func (h *handlerCategory) HandlerResults(c echo.Context) error {
	res, err := h.category.EntityResults()

	if err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})

		return err
	}

	return c.JSON(200, schemas.SchemaResponses{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *handlerCategory) HandlerCreate(c echo.Context) error {
	var body schemas.SchemaCategory

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	_, error := h.category.EntityCreate(&body)

	if error != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Error",
			Data:       nil,
		})

		return error
	}

	return c.JSON(200, schemas.SchemaResponses{
		StatusCode: 200,
		Message:    "Success",
		Data:       body,
	})
}

func (h *handlerCategory) HandlerResult(c echo.Context) error {
	var body schemas.SchemaCategory
	id := c.Param("id")

	body.ID = id

	res, err := h.category.EntityResult(&body)

	if err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})

		return err
	}

	return c.JSON(200, schemas.SchemaResponses{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *handlerCategory) HandlerUpdate(c echo.Context) error {
	var body schemas.SchemaCategory

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	_, error := h.category.EntityUpdate(&body)

	if error != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Error",
			Data:       nil,
		})

		return error
	}

	return c.JSON(200, schemas.SchemaResponses{
		StatusCode: 200,
		Message:    "Success",
		Data:       body,
	})
}

func (h *handlerCategory) HandlerDelete(c echo.Context) error {
	var body schemas.SchemaCategory
	id := c.Param("id")

	body.ID = id

	_, error := h.category.EntityDelete(&body)

	if error != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Error",
			Data:       nil,
		})

		return error
	}

	return c.JSON(200, schemas.SchemaResponses{
		StatusCode: 200,
		Message:    "Success",
		Data:       body,
	})
}
