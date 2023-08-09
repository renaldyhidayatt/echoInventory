package route

import (
	"echoinventory/handler/api"
	"echoinventory/pb"

	"github.com/labstack/echo/v4"
)

func NewUserRoute(client pb.UserServiceClient, router *echo.Echo) {
	userHandler := api.NewHandlerUser(client)

	routerUser := router.Group("/api/user")

	routerUser.GET("/hello", userHandler.HandlerHello)
	routerUser.POST("/create", userHandler.CreateUser)
	routerUser.PUT("/update/:id", userHandler.UpdateUser)
	routerUser.DELETE("/delete/:id", userHandler.DeleteUser)
	routerUser.GET("/", userHandler.GetUsers)
	routerUser.GET("/:id", userHandler.GetUser)

}
