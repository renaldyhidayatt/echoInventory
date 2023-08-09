package api

import (
	"echoinventory/pb"
	"echoinventory/schemas"

	"github.com/labstack/echo/v4"
)

type handleUser struct {
	client pb.UserServiceClient
}

func NewHandlerUser(client pb.UserServiceClient) *handleUser {
	return &handleUser{client: client}
}

func (h *handleUser) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello User")
}

func (h *handleUser) CreateUser(c echo.Context) error {
	var body schemas.SchemaUser

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}
	ctx := c.Request().Context()

	data := &pb.CreateUserInput{
		Firstname: body.FirstName,
		Lastname:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
		Role:      body.Role,
	}

	res, err := h.client.CreateUser(ctx, data)

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

func (h *handleUser) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.GetUsers(ctx, &pb.UsersRequest{})
	if err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
		return err
	}

	var users []schemas.SchemaUser

	for _, pbUser := range res.Users {
		user := schemas.SchemaUser{
			ID:        pbUser.Id,
			FirstName: pbUser.Firstname,
			LastName:  pbUser.Lastname,
			Email:     pbUser.Email,
			Role:      pbUser.Role,
		}
		users = append(users, user)
	}

	return c.JSON(200, schemas.SchemaResponses{
		StatusCode: 200,
		Message:    "Success",
		Data:       users,
	})
}

func (h *handleUser) GetUser(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()

	data := &pb.UserRequest{
		Id: id,
	}

	res, err := h.client.GetUser(ctx, data)

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

func (h *handleUser) UpdateUser(c echo.Context) error {
	var body schemas.SchemaUser
	id := c.Param("id")

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}
	ctx := c.Request().Context()

	data := &pb.UpdateUserInput{
		Id:        id,
		Firstname: body.FirstName,
		Lastname:  body.LastName,
		Email:     body.Email,
		Role:      body.Role,
	}

	res, err := h.client.UpdateUser(ctx, data)

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

func (h *handleUser) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	data := &pb.UserRequest{
		Id: id,
	}
	ctx := c.Request().Context()

	res, err := h.client.DeleteUser(ctx, data)

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
