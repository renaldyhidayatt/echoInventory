package api

import (
	"echoinventory/pb"
	"echoinventory/schemas"

	"github.com/labstack/echo/v4"
)

type handleSupplier struct {
	client pb.SupplierServiceClient
}

func NewHandlerSupplier(client pb.SupplierServiceClient) *handleSupplier {
	return &handleSupplier{client: client}
}

func (h *handleSupplier) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello Supplier")
}

func (h *handleSupplier) CreateSupplier(c echo.Context) error {
	var body schemas.SchemaSupplier

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}
	ctx := c.Request().Context()

	data := &pb.CreateSupplierInput{
		Name:    body.Name,
		Alamat:  body.Alamat,
		Email:   body.Email,
		Telepon: body.Telepon,
	}

	res, err := h.client.CreateSupplier(ctx, data)

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

func (h *handleSupplier) GetSuppliers(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.GetSuppliers(ctx, &pb.SuppliersRequest{})

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

func (h *handleSupplier) GetSupplier(c echo.Context) error {
	var body schemas.SchemaSupplier

	id := c.Param("id")

	body.ID = id
	ctx := c.Request().Context()

	data := &pb.SupplierRequest{
		Id: id,
	}

	res, err := h.client.GetSupplier(ctx, data)

	if err != nil {
		return c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	return c.JSON(200, schemas.SchemaResponses{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *handleSupplier) UpdateSupplier(c echo.Context) error {
	var body schemas.SchemaSupplier
	id := c.Param("id")

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}
	ctx := c.Request().Context()

	body.ID = id

	data := &pb.UpdateSupplierInput{
		Id:      body.ID,
		Name:    body.Name,
		Alamat:  body.Alamat,
		Email:   body.Email,
		Telepon: body.Telepon,
	}

	res, err := h.client.UpdateSupplier(ctx, data)

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

func (h *handleSupplier) DeleteSupplier(c echo.Context) error {
	var body schemas.SchemaSupplier
	id := c.Param("id")

	body.ID = id

	data := &pb.SupplierRequest{
		Id: id,
	}
	ctx := c.Request().Context()

	res, err := h.client.DeleteSupplier(ctx, data)

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
