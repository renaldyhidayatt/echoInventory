package route

import (
	"echoinventory/handler/api"
	"echoinventory/pb"

	"github.com/labstack/echo/v4"
)

func NewProductMasukRoute(client pb.ProductMasukServiceClient, router *echo.Echo) {
	productHandler := api.NewHandleProductMasuk(client)

	routerProduct := router.Group("/api/productmasuk")

	routerProduct.GET("/hello", productHandler.HandlerHello)
	routerProduct.POST("/create", productHandler.CreateProductMasuk)
	routerProduct.PUT("/update/:id", productHandler.UpdateProductMasuk)
	routerProduct.DELETE("/delete/:id", productHandler.DeleteProductMasuk)
	routerProduct.GET("/", productHandler.GetProductMasuks)
	routerProduct.GET("/:id", productHandler.GetProductMasuk)
}
