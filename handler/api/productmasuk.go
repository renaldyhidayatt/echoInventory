package api

import (
	"echoinventory/pb"
	"echoinventory/schemas"

	"github.com/labstack/echo/v4"
)

type handleProductMasuk struct {
	client pb.ProductMasukServiceClient
}

func NewHandleProductMasuk(productmasuk pb.ProductMasukServiceClient) *handleProductMasuk {
	return &handleProductMasuk{
		client: productmasuk,
	}
}

func (h *handleProductMasuk) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello Product Masuk")
}

func (h *handleProductMasuk) CreateProductMasuk(c echo.Context) error {
	var body schemas.SchemaProductMasuk

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	ctx := c.Request().Context()

	data := &pb.CreateProductMasukInput{
		Name:       body.Name,
		Qty:        body.Qty,
		ProductId:  body.ProductID,
		SupplierId: body.SupplierID,
	}

	res, err := h.client.CreateProductMasuk(ctx, data)

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

func (h *handleProductMasuk) GetProductMasuks(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.GetProductMasuks(ctx, &pb.ProductMasuksRequest{})

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

func (h *handleProductMasuk) GetProductMasuk(c echo.Context) error {
	var body schemas.SchemaProductMasuk

	id := c.Param("id")

	body.ID = id
	ctx := c.Request().Context()

	data := &pb.ProductMasukRequest{
		Id: id,
	}

	res, err := h.client.GetProductMasuk(ctx, data)

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

func (h *handleProductMasuk) UpdateProductMasuk(c echo.Context) error {
	var body schemas.SchemaProductMasuk
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

	data := &pb.UpdateProductMasukInput{
		Id:         body.ID,
		Name:       body.Name,
		Qty:        body.Qty,
		ProductId:  body.ProductID,
		SupplierId: body.SupplierID,
	}

	res, err := h.client.UpdateProductMasuk(ctx, data)

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

func (h *handleProductMasuk) DeleteProductMasuk(c echo.Context) error {
	var body schemas.SchemaProductMasuk
	id := c.Param("id")

	body.ID = id

	data := &pb.ProductMasukRequest{
		Id: id,
	}
	ctx := c.Request().Context()

	res, err := h.client.DeleteProductMasuk(ctx, data)

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
