package handler

import (
	"echoinventory/entity"
	"echoinventory/schemas"

	"github.com/labstack/echo/v4"
)

type handlerProductMasuk struct {
	productmasuk entity.EntityProductMasuk
}

func NewHandlerProductMasuk(productmasuk entity.EntityProductMasuk) *handlerProductMasuk {
	return &handlerProductMasuk{
		productmasuk: productmasuk,
	}
}

func (h *handlerProductMasuk) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello World")
}

func (h *handlerProductMasuk) HandlerResults(c echo.Context) error {
	res, err := h.productmasuk.EntityResults()

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

func (h *handlerProductMasuk) HandlerCreate(c echo.Context) error {
	var body schemas.SchemaProductMasuk

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	_, error := h.productmasuk.EntityCreate(&body)

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

func (h *handlerProductMasuk) HandlerResult(c echo.Context) error {
	var body schemas.SchemaProductMasuk
	id := c.Param("id")

	body.ID = id

	res, err := h.productmasuk.EntityResult(&body)

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

func (h *handlerProductMasuk) HandlerUpdate(c echo.Context) error {
	var body schemas.SchemaProductMasuk

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	_, error := h.productmasuk.EntityUpdate(&body)

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

func (h *handlerProductMasuk) HandlerDelete(c echo.Context) error {
	var body schemas.SchemaProductMasuk
	id := c.Param("id")

	body.ID = id

	_, error := h.productmasuk.EntityDelete(&body)

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
