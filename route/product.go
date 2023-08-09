package route

import (
	"echoinventory/handler/api"
	"echoinventory/pb"

	"github.com/labstack/echo/v4"
)

func NewProductRoute(client pb.ProductServiceClient, router *echo.Echo) {
	productHandler := api.NewHandlerProduct(client)

	routerProduct := router.Group("/api/product")

	routerProduct.GET("/hello", productHandler.HandlerHello)
	routerProduct.POST("/create", productHandler.CreateProduct)
	routerProduct.PUT("/update/:id", productHandler.UpdateProduct)
	routerProduct.DELETE("/delete/:id", productHandler.DeleteProduct)
	routerProduct.GET("/", productHandler.GetProducts)
	routerProduct.GET("/:id", productHandler.GetProduct)
}
