package route

import (
	"echoinventory/handler/api"
	"echoinventory/pb"

	"github.com/labstack/echo/v4"
)

func NewSaleRoute(client pb.SaleServiceClient, router *echo.Echo) {
	saleHandler := api.NewHandlerSale(client)

	routerSale := router.Group("/api/sale")

	routerSale.GET("/hello", saleHandler.HandlerHello)
	routerSale.POST("/create", saleHandler.CreateSale)
	routerSale.PUT("/update/:id", saleHandler.UpdateSale)
	routerSale.DELETE("/delete/:id", saleHandler.DeleteSale)
	routerSale.GET("/", saleHandler.GetSales)
	routerSale.GET("/:id", saleHandler.GetSale)
}
