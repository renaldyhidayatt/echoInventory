package api

import (
	"echoinventory/pb"
	"echoinventory/schemas"

	"github.com/labstack/echo/v4"
)

type handleAuth struct {
	client pb.AuthServiceClient
}

func NewHandlerAuth(client pb.AuthServiceClient) *handleAuth {
	return &handleAuth{client: client}
}

func (h *handleAuth) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello Auth")
}

func (h *handleAuth) RegisterUser(c echo.Context) error {
	var body schemas.SchemaUser

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}
	data := &pb.SignUpUserInput{
		Firstname: body.FirstName,
		Lastname:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
		Role:      body.Role,
	}

	ctx := c.Request().Context()

	res, err := h.client.RegisterUser(ctx, data)

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

func (h *handleAuth) LoginUser(c echo.Context) error {
	var body schemas.SchemaUser

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	data := &pb.SignInUserInput{
		Email:    body.Email,
		Password: body.Password,
	}

	ctx := c.Request().Context()

	res, err := h.client.LoginUser(ctx, data)

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
