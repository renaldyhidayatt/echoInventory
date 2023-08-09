package handler

import (
	"echoinventory/entity"
	"echoinventory/pkg"
	"echoinventory/schemas"

	"github.com/labstack/echo/v4"
)

type handlerAuth struct {
	auth entity.EntityAuth
}

func NewHandlerAuth(auth entity.EntityAuth) *handlerAuth {
	return &handlerAuth{auth: auth}
}

func (h *handlerAuth) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello Auth")
}

func (h *handlerAuth) HandlerRegister(c echo.Context) error {
	var body schemas.SchemaUser

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	_, error := h.auth.EntityRegister(&body)

	if error != nil {
		return c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Error",
			Data:       nil,
		})
	}

	return c.JSON(200, schemas.SchemaResponses{
		StatusCode: 200,
		Message:    "Success",
		Data:       body,
	})
}

func (h *handlerAuth) HandlerLogin(c echo.Context) error {
	var body schemas.SchemaUser

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	res, err := h.auth.EntityLogin(&body)

	if err != nil {
		return c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Error",
			Data:       nil,
		})
	}

	jwt, err := pkg.Sign(&schemas.SchemaJWT{
		ID:    res.ID,
		Email: res.Email,
	})

	if err != nil {
		return c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Error jwt",
			Data:       nil,
		})
	}

	return c.JSON(200, schemas.SchemaResponses{
		StatusCode: 200,
		Message:    "Success",
		Data:       jwt,
	})
}
