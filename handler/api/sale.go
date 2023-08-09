package api

import (
	"echoinventory/pb"
	"echoinventory/schemas"

	"github.com/labstack/echo/v4"
)

type handleSale struct {
	client pb.SaleServiceClient
}

func NewHandlerSale(client pb.SaleServiceClient) *handleSale {
	return &handleSale{client: client}
}

func (h *handleSale) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello Sale")
}

func (h *handleSale) CreateSale(c echo.Context) error {
	var body schemas.SchemaSale

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}
	ctx := c.Request().Context()

	data := &pb.CreateSaleInput{
		Name:    body.Name,
		Alamat:  body.Alamat,
		Email:   body.Email,
		Telepon: body.Telepon,
	}

	res, err := h.client.CreateSale(ctx, data)

	if err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Error",
			Data:       nil,
		})
	}

	return c.JSON(200, schemas.SchemaResponses{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *handleSale) GetSales(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.GetSales(ctx, &pb.SalesRequest{})

	if err != nil {
		return c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Error",
			Data:       nil,
		})
	}

	return c.JSON(200, schemas.SchemaResponses{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *handleSale) GetSale(c echo.Context) error {
	var body schemas.SchemaSale

	id := c.Param("id")

	body.ID = id
	ctx := c.Request().Context()

	data := &pb.SaleRequest{
		Id: id,
	}

	res, err := h.client.GetSale(ctx, data)

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

func (h *handleSale) UpdateSale(c echo.Context) error {
	var body schemas.SchemaSale
	id := c.Param("id")

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}
	ctx := c.Request().Context()

	body.ID = id

	data := &pb.UpdateSaleInput{
		Id:      body.ID,
		Name:    body.Name,
		Alamat:  body.Alamat,
		Email:   body.Email,
		Telepon: body.Telepon,
	}

	res, err := h.client.UpdateSale(ctx, data)

	if err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Error",
			Data:       nil,
		})
	}

	return c.JSON(200, schemas.SchemaResponses{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *handleSale) DeleteSale(c echo.Context) error {
	var body schemas.SchemaSale
	id := c.Param("id")

	body.ID = id

	data := &pb.SaleRequest{
		Id: id,
	}
	ctx := c.Request().Context()

	res, err := h.client.DeleteSale(ctx, data)

	if err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Error",
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
