package route

import (
	"echoinventory/handler/api"
	"echoinventory/pb"

	"github.com/labstack/echo/v4"
)

func NewSupplierRoute(client pb.SupplierServiceClient, router *echo.Echo) {
	supplierHandler := api.NewHandlerSupplier(client)

	routerSale := router.Group("/api/sale")

	routerSale.GET("/hello", supplierHandler.HandlerHello)
	routerSale.POST("/create", supplierHandler.CreateSupplier)
	routerSale.PUT("/update/:id", supplierHandler.UpdateSupplier)
	routerSale.DELETE("/delete/:id", supplierHandler.DeleteSupplier)
	routerSale.GET("/", supplierHandler.GetSuppliers)
	routerSale.GET("/:id", supplierHandler.GetSupplier)
}
