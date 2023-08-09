package api

import (
	"echoinventory/pb"
	"echoinventory/schemas"

	"github.com/labstack/echo/v4"
)

type handleProductKeluar struct {
	client pb.ProductKeluarServiceClient
}

func NewHandleProductKeluar(productkeluar pb.ProductKeluarServiceClient) *handleProductKeluar {
	return &handleProductKeluar{
		client: productkeluar,
	}
}

func (h *handleProductKeluar) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello Product Keluar")
}

func (h *handleProductKeluar) CreateProductKeluar(c echo.Context) error {
	var body schemas.SchemaProductKeluar

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	ctx := c.Request().Context()

	data := &pb.CreateProductKeluarInput{
		Qty:        body.Qty,
		ProductId:  body.ProductID,
		CategoryId: body.CategoryID,
	}

	res, err := h.client.CreateProductKeluar(ctx, data)

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

func (h *handleProductKeluar) GetProductKeluars(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := h.client.GetProductKeluars(ctx, &pb.ProductKeluarsRequest{})

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

func (h *handleProductKeluar) GetProductKeluar(c echo.Context) error {
	productKeluarID := c.Param("id")

	ctx := c.Request().Context()

	data := &pb.ProductKeluarRequest{
		Id: productKeluarID,
	}

	res, err := h.client.GetProductKeluar(ctx, data)

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

func (h *handleProductKeluar) UpdateProductKeluar(c echo.Context) error {
	var body schemas.SchemaProductKeluar

	productKeluarID := c.Param("id")

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Error",
			Data:       nil,
		})
	}

	ctx := c.Request().Context()

	data := &pb.UpdateProductKeluarInput{
		Id:         productKeluarID,
		Qty:        body.Qty,
		ProductId:  body.ProductID,
		CategoryId: body.CategoryID,
	}

	res, err := h.client.UpdateProductKeluar(ctx, data)

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

func (h *handleProductKeluar) DeleteProductKeluar(c echo.Context) error {
	productKeluarID := c.Param("id")

	ctx := c.Request().Context()

	data := &pb.ProductKeluarRequest{
		Id: productKeluarID,
	}

	res, err := h.client.DeleteProductKeluar(ctx, data)

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
