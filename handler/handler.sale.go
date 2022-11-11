package handler

import (
	"echoinventory/entity"
	"echoinventory/schemas"

	"github.com/labstack/echo/v4"
)

type handlerSale struct {
	sale entity.EntitySale
}

func NewHandlerSale(sale entity.EntitySale) *handlerSale {
	return &handlerSale{
		sale: sale,
	}
}

func (h *handlerSale) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello World")
}

func (h *handlerSale) HandlerResults(c echo.Context) error {
	res, err := h.sale.EntityResults()

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

func (h *handlerSale) HandlerCreate(c echo.Context) error {
	var body schemas.SchemaSale

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	_, error := h.sale.EntityCreate(&body)

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

func (h *handlerSale) HandlerResult(c echo.Context) error {
	var body schemas.SchemaSale
	id := c.Param("id")

	body.ID = id

	res, err := h.sale.EntityResult(&body)

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

func (h *handlerSale) HandlerUpdate(c echo.Context) error {
	var body schemas.SchemaSale

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	_, error := h.sale.EntityUpdate(&body)

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

func (h *handlerSale) HandlerDelete(c echo.Context) error {
	var body schemas.SchemaSale
	id := c.Param("id")

	body.ID = id

	_, error := h.sale.EntityDelete(&body)

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
