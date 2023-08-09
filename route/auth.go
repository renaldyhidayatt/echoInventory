package route

import (
	"echoinventory/handler/api"
	"echoinventory/pb"

	"github.com/labstack/echo/v4"
)

func NewAuthRoute(client pb.AuthServiceClient, router *echo.Echo) {

	authHandler := api.NewHandlerAuth(client)

	routerAuth := router.Group("/api/auth")

	routerAuth.GET("/hello", authHandler.HandlerHello)
	routerAuth.POST("/login", authHandler.LoginUser)
	routerAuth.POST("/register", authHandler.RegisterUser)
}
