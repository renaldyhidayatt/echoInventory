package api

import (
	"echoinventory/pb"
	"echoinventory/schemas"

	"github.com/labstack/echo/v4"
)

type handleCustomer struct {
	client pb.CustomerServiceClient
}

func NewHandlerCustomer(client pb.CustomerServiceClient) *handleCustomer {
	return &handleCustomer{client: client}
}

func (h *handleCustomer) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello Customer")
}

func (h *handleCustomer) CreateCustomer(c echo.Context) error {
	var body schemas.SchemaCustomer

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}
	ctx := c.Request().Context()

	data := &pb.CreateCustomerInput{
		Name:    body.Name,
		Alamat:  body.Alamat,
		Email:   body.Email,
		Telepon: body.Telepon,
	}

	res, err := h.client.CreateCustomer(ctx, data)

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

func (h *handleCustomer) GetCustomers(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.GetCustomers(ctx, &pb.CustomersRequest{})

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

func (h *handleCustomer) GetCustomer(c echo.Context) error {
	var body schemas.SchemaCustomer

	id := c.Param("id")

	body.ID = id
	ctx := c.Request().Context()

	data := &pb.CustomerRequest{
		Id: id,
	}

	res, err := h.client.GetCustomer(ctx, data)

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

func (h *handleCustomer) UpdateCustomer(c echo.Context) error {
	var body schemas.SchemaCustomer
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

	data := &pb.UpdateCustomerInput{
		Id:      body.ID,
		Name:    body.Name,
		Alamat:  body.Alamat,
		Email:   body.Email,
		Telepon: body.Telepon,
	}

	res, err := h.client.UpdateCustomer(ctx, data)

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

func (h *handleCustomer) DeleteCustomer(c echo.Context) error {
	var body schemas.SchemaCustomer
	id := c.Param("id")

	body.ID = id

	data := &pb.CustomerRequest{
		Id: id,
	}
	ctx := c.Request().Context()

	res, err := h.client.DeleteCustomer(ctx, data)

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
