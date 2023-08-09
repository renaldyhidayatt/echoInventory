package route

import (
	"echoinventory/handler/api"
	"echoinventory/pb"

	"github.com/labstack/echo/v4"
)

func NewCustomerRoute(client pb.CustomerServiceClient, router *echo.Echo) {
	customerHandler := api.NewHandlerCustomer(client)

	routerCustomer := router.Group("/api/customer")

	routerCustomer.GET("/hello", customerHandler.HandlerHello)
	routerCustomer.POST("/create", customerHandler.CreateCustomer)
	routerCustomer.PUT("/update/:id", customerHandler.UpdateCustomer)
	routerCustomer.DELETE("/delete/:id", customerHandler.DeleteCustomer)
	routerCustomer.GET("/", customerHandler.GetCustomers)
	routerCustomer.GET("/:id", customerHandler.GetCustomer)
}
