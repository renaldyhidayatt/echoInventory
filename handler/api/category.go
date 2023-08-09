package api

import (
	"echoinventory/pb"
	"echoinventory/schemas"

	"github.com/labstack/echo/v4"
)

type handleCategory struct {
	client pb.CategoryServiceClient
}

func NewHandlerCategory(client pb.CategoryServiceClient) *handleCategory {
	return &handleCategory{client: client}
}

func (h *handleCategory) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello Category")
}

func (h *handleCategory) CreateCategory(c echo.Context) error {
	var body schemas.SchemaCategory

	if err := c.Bind(&body); err != nil {
		c.JSON(400, schemas.SchemaResponses{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}
	ctx := c.Request().Context()

	data := &pb.CreateCategoryInput{
		Name: body.Name,
	}

	res, err := h.client.CreateCategory(ctx, data)

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

func (h *handleCategory) GetCategories(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := h.client.GetCategories(ctx, &pb.CategoriesRequest{})

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

func (h *handleCategory) GetCategory(c echo.Context) error {
	var body schemas.SchemaCategory

	id := c.Param("id")

	body.ID = id
	ctx := c.Request().Context()

	data := &pb.CategoryRequest{
		Id: id,
	}

	res, err := h.client.GetCategory(ctx, data)

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

func (h *handleCategory) UpdateCategory(c echo.Context) error {
	var body schemas.SchemaCategory
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
	// body.Name = body.Name

	data := &pb.UpdateCategoryInput{
		Id:   body.ID,
		Name: body.Name,
	}

	res, err := h.client.UpdateCategory(ctx, data)

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

func (h *handleCategory) DeleteCategory(c echo.Context) error {
	var body schemas.SchemaCategory
	id := c.Param("id")

	body.ID = id

	data := &pb.CategoryRequest{
		Id: id,
	}
	ctx := c.Request().Context()

	res, err := h.client.DeleteCategory(ctx, data)

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
