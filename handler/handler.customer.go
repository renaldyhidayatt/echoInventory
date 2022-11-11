package handler

import (
	"echoinventory/entity"
	"echoinventory/schemas"

	"github.com/labstack/echo/v4"
)

type handlerCustomer struct {
	customer entity.EntityCustomer
}

func NewHandlerCustomer(customer entity.EntityCustomer) *handlerCustomer {
	return &handlerCustomer{
		customer: customer,
	}
}

func (h *handlerCustomer) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello World")
}

func (h *handlerCustomer) HandlerResults(c echo.Context) error {
	res, err := h.customer.EntityResults()

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

func (h *handlerCustomer) HandlerCreate(c echo.Context) error {
	var body schemas.SchemaCustomer

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	_, error := h.customer.EntityCreate(&body)

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

func (h *handlerCustomer) HandlerResult(c echo.Context) error {
	var body schemas.SchemaCustomer
	id := c.Param("id")

	body.ID = id

	res, err := h.customer.EntityResult(&body)

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

func (h *handlerCustomer) HandlerUpdate(c echo.Context) error {
	var body schemas.SchemaCustomer

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	_, error := h.customer.EntityUpdate(&body)

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

func (h *handlerCustomer) HandlerDelete(c echo.Context) error {
	var body schemas.SchemaCustomer
	id := c.Param("id")

	body.ID = id

	_, error := h.customer.EntityDelete(&body)

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
