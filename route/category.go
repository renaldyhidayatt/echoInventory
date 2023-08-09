package route

import (
	"echoinventory/handler/api"
	"echoinventory/pb"

	"github.com/labstack/echo/v4"
)

func NewCategoryRoute(client pb.CategoryServiceClient, router *echo.Echo) {
	categoryHandler := api.NewHandlerCategory(client)

	routerCategory := router.Group("/api/category")

	routerCategory.GET("/hello", categoryHandler.HandlerHello)
	routerCategory.POST("/create", categoryHandler.CreateCategory)
	routerCategory.PUT("/update/:id", categoryHandler.UpdateCategory)
	routerCategory.DELETE("/delete/:id", categoryHandler.DeleteCategory)
	routerCategory.GET("/", categoryHandler.GetCategories)
	routerCategory.GET("/:id", categoryHandler.GetCategory)

}
