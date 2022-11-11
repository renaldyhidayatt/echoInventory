package handler

import (
	"echoinventory/entity"
	"echoinventory/schemas"

	"github.com/labstack/echo/v4"
)

type handlerSupplier struct {
	supplier entity.EntitySupplier
}

func NewHandlerSupplier(supplier entity.EntitySupplier) *handlerSupplier {
	return &handlerSupplier{
		supplier: supplier,
	}
}

func (h *handlerSupplier) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello World")
}

func (h *handlerSupplier) HandlerResults(c echo.Context) error {
	res, err := h.supplier.EntityResults()

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

func (h *handlerSupplier) HandlerCreate(c echo.Context) error {
	var body schemas.SchemaSupplier

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	_, error := h.supplier.EntityCreate(&body)

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

func (h *handlerSupplier) HandlerResult(c echo.Context) error {
	var body schemas.SchemaSupplier
	id := c.Param("id")

	body.ID = id

	res, err := h.supplier.EntityResult(&body)

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

func (h *handlerSupplier) HandlerUpdate(c echo.Context) error {
	var body schemas.SchemaSupplier

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	_, error := h.supplier.EntityUpdate(&body)

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

func (h *handlerSupplier) HandlerDelete(c echo.Context) error {
	var body schemas.SchemaSupplier
	id := c.Param("id")

	body.ID = id

	_, error := h.supplier.EntityDelete(&body)

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
