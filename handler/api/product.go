package api

import (
	"echoinventory/pb"
	"echoinventory/schemas"

	"github.com/labstack/echo/v4"
)

type handleProduct struct {
	client pb.ProductServiceClient
}

func NewHandlerProduct(product pb.ProductServiceClient) *handleProduct {
	return &handleProduct{client: product}
}

func (h *handleProduct) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello Product")
}

func (h *handleProduct) CreateProduct(c echo.Context) error {
	var body schemas.SchemaProduct

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	ctx := c.Request().Context()

	data := &pb.CreateProductInput{
		Name:       body.Name,
		Image:      body.Image,
		Qty:        body.Qty,
		CategoryId: body.CategoryID,
	}

	res, err := h.client.CreateProduct(ctx, data)

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

func (h *handleProduct) GetProducts(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := h.client.GetProducts(ctx, &pb.ProductsRequest{})

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

func (h *handleProduct) GetProduct(c echo.Context) error {
	productID := c.Param("id")

	ctx := c.Request().Context()

	data := &pb.ProductRequest{
		Id: productID,
	}

	res, err := h.client.GetProduct(ctx, data)

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

func (h *handleProduct) UpdateProduct(c echo.Context) error {

	var body schemas.SchemaProduct

	productID := c.Param("id")

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Error",
			Data:       nil,
		})
	}

	ctx := c.Request().Context()

	data := &pb.UpdateProductInput{
		Id:         productID,
		Name:       body.Name,
		Image:      body.Image,
		Qty:        body.Qty,
		CategoryId: body.CategoryID,
	}

	res, err := h.client.UpdateProduct(ctx, data)

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

func (h *handleProduct) DeleteProduct(c echo.Context) error {
	productID := c.Param("id")

	ctx := c.Request().Context()

	data := &pb.ProductRequest{
		Id: productID,
	}

	res, err := h.client.DeleteProduct(ctx, data)

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
