package handler

import (
	"echoinventory/entity"
	"echoinventory/schemas"

	"github.com/labstack/echo/v4"
)

type handlerProductKeluar struct {
	productkeluar entity.EntityProductKeluar
}

func NewHandlerProductKeluar(productkeluar entity.EntityProductKeluar) *handlerProductKeluar {
	return &handlerProductKeluar{
		productkeluar: productkeluar,
	}
}

func (h *handlerProductKeluar) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello World")
}

func (h *handlerProductKeluar) HandlerResults(c echo.Context) error {
	res, err := h.productkeluar.EntityResults()

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

func (h *handlerProductKeluar) HandlerCreate(c echo.Context) error {
	var body schemas.SchemaProductKeluar

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	_, error := h.productkeluar.EntityCreate(&body)

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

func (h *handlerProductKeluar) HandlerResult(c echo.Context) error {
	var body schemas.SchemaProductKeluar
	id := c.Param("id")

	body.ID = id

	res, err := h.productkeluar.EntityResult(&body)

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

func (h *handlerProductKeluar) HandlerUpdate(c echo.Context) error {
	var body schemas.SchemaProductKeluar

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	_, error := h.productkeluar.EntityUpdate(&body)

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

func (h *handlerProductKeluar) HandlerDelete(c echo.Context) error {
	var body schemas.SchemaProductKeluar
	id := c.Param("id")

	body.ID = id

	_, error := h.productkeluar.EntityDelete(&body)

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
