package route

import (
	"echoinventory/handler/api"
	"echoinventory/pb"

	"github.com/labstack/echo/v4"
)

func NewProductKeluarRoute(client pb.ProductKeluarServiceClient, router *echo.Echo) {
	productHandler := api.NewHandleProductKeluar(client)

	routerProduct := router.Group("/api/productkeluar")

	routerProduct.GET("/hello", productHandler.HandlerHello)
	routerProduct.POST("/create", productHandler.CreateProductKeluar)
	routerProduct.PUT("/update/:id", productHandler.UpdateProductKeluar)
	routerProduct.DELETE("/delete/:id", productHandler.DeleteProductKeluar)
	routerProduct.GET("/", productHandler.GetProductKeluars)
	routerProduct.GET("/:id", productHandler.GetProductKeluar)
}
